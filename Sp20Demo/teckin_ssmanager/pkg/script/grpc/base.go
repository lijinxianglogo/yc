package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	pb "teckin_ssmanager/api/ssmanager/grpc"
	"time"
)

const (
	grpcServer = "172.31.46.78:21000"
	//grpcServer = "127.0.0.1:21000"
)

var (
	// grpc options
	grpcKeepAliveTime    = time.Duration(5) * time.Second
	grpcKeepAliveTimeout = time.Duration(5) * time.Second
	grpcBackoffMaxDelay  = time.Duration(5) * time.Second
	grpcMaxSendMsgSize   = 1 << 24
	grpcMaxCallMsgSize   = 1 << 24
)

func GetLambdaClient() (pb.LambdaClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(grpcServer,
		[]grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                grpcKeepAliveTime,
				Timeout:             grpcKeepAliveTimeout,
				PermitWithoutStream: true,
			}),
		}...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return pb.NewLambdaClient(conn), conn
}

func GetUserInfo(c pb.LambdaClient, uid string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	uinfo, err := c.StreamGetUserInfo(ctx, &pb.CheckUserRequest{Uid: uid})
	if err != nil {
		return "", err
	}
	return uinfo.Password, nil
}

func GetMqttCheck(c pb.LambdaClient, body string) (int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.StreamMqttCheck(ctx, &pb.MqttCheckRequest{Msg: body})
	if err != nil{
		return 400, err
	}
	return resp.Code, nil
}
