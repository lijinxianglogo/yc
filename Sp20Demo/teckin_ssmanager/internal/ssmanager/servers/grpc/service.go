package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"strings"
	pb "teckin_ssmanager/api/ssmanager/grpc"
	"teckin_ssmanager/config"
	"teckin_ssmanager/internal/ssmanager/model"
	"teckin_ssmanager/internal/ssmanager/service"
	"teckin_ssmanager/pkg/encoding"
	scv "teckin_ssmanager/pkg/strconv"
)

type GrpcServer struct {
	pb.LambdaServer
	srv *service.Service
}

func New(srv *service.Service) *GrpcServer {
	return &GrpcServer{
		srv: srv,
	}
}

func (gsrv *GrpcServer) Dial(conf *config.GrpcServerConfig) error {
	lis, err := net.Listen("tcp", conf.Addr)
	if err != nil {
		fmt.Println("failed to listen: ", err)
	}
	s := grpc.NewServer()
	pb.RegisterLambdaServer(s, &GrpcServer{srv: gsrv.srv})
	go func(){
		if err := s.Serve(lis); err != nil {
			fmt.Println("failed to serve: " + err.Error())
		}
	}()

	return err
}

func (gsrv *GrpcServer) StreamGetUserInfo(ctx context.Context, in *pb.CheckUserRequest) (*pb.CheckUserResponse, error) {
	var user model.User
	var err error
	uid := scv.S2I64(in.Uid)

	user, err = gsrv.srv.Dao.GetUserInfoByUid(uid)
	if err == nil{
		//lambda校验返回规则：base64(md5(uid+key))
		return &pb.CheckUserResponse{Email: user.Email, Password: encoding.StringMD5(in.Uid + user.Key)}, nil
	}
	user, err = gsrv.srv.Dao.Mysql.GetUserInfoByUserID(scv.S2I(in.Uid))
	if err == nil{
		return &pb.CheckUserResponse{Email: user.Email, Password: encoding.StringMD5(in.Uid + user.Key)}, nil
	}
	return nil, err
}

func (gsrv *GrpcServer) StreamMqttCheck(ctx context.Context, in *pb.MqttCheckRequest) (*pb.OperateResponse, error){
	var mqttTopic = []string{"test/device", "test/cloud"}
	var mqttBody model.MyqqBody
	var mqtt model.MqttCheck
	var resp = new(pb.OperateResponse)
	var check bool
	fmt.Println("get info::: ",in.Msg)
	//TODO demo版只为了检测redis取出信息时间，不做其他用
	_, _ = gsrv.srv.Dao.GetUserInfoByUid(2)
	err := json.Unmarshal([]byte(in.Msg), &mqttBody)
	if err != nil{
		resp.Code = 300
		return resp, err
	}
	mqttStr := encoding.DecodeBase64(mqttBody.Body)
	err = json.Unmarshal([]byte(mqttStr), &mqtt)
	if err != nil{
		resp.Code = 300
		return resp, err
	}
	for _, topic := range mqttTopic {
		if strings.HasPrefix(mqtt.Header.Topic, topic){
			check = true
			go gsrv.srv.Dao.AwsIotClient.Publish(mqtt.Header.Topic, []byte(mqttStr))
		}
	}
	fmt.Println("get mqtt from:", mqtt.Header.Topic)
	if check{
		resp.Code = 200
	} else{
		resp.Code = 400
	}
	return resp, nil
}