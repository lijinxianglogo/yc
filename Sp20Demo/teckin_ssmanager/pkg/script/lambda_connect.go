package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"strings"
	"teckin_ssmanager/pkg/encoding"
	"teckin_ssmanager/pkg/script/grpc"
	"teckin_ssmanager/pkg/script/model"
)



func connectHandle(ctx context.Context, data model.ConnRequest) (model.ConnectResponse, error) {
	fmt.Println("request context:", ctx)
	fmt.Println("request data:", data)
	var resp model.ConnectResponse
	auth := data.Header.Auth
	if strings.HasPrefix(auth, "Basic "){
		auth = strings.Split(auth, "Basic ")[1]
	}
	info := encoding.DecodeBase64(auth)
	spliInfo := strings.Split(info, ":")
	if len(spliInfo) != 2{ //TODO info格式暂定为Uid:MD5(Uid+key)
		resp.Code = 300
		return resp, nil
	}
	cli, conn := grpc.GetLambdaClient()
	defer conn.Close()
	password, err := grpc.GetUserInfo(cli, spliInfo[0])
	if err != nil{
		resp.Code = 500
		return resp, err
	}
	switch spliInfo[1] {
	case password:
		resp.Code = 200
	default:
		resp.Code = 401
	}
	return resp, nil
}


func main() {
	lambda.Start(connectHandle)
}
