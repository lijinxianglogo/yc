//用于校验topic信息
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"teckin_ssmanager/pkg/script/grpc"
	"teckin_ssmanager/pkg/script/model"
)


func mqttSendHandle(ctx context.Context, data interface{}) (model.ConnectResponse, error) {
	var resp model.ConnectResponse
	var respBody model.ConnectResponseBody
	respBody.Key = "111"
	rd, _ := json.Marshal(respBody)
	resp.Body = string(rd)
	fmt.Println("request data:", data)
	cli, conn := grpc.GetLambdaClient()
	defer conn.Close()
	info, _ := json.Marshal(data)
	code, err := grpc.GetMqttCheck(cli, string(info))
	if err != nil{
		fmt.Println("check err:", err.Error())
	}
	resp.Code = code

	d, _ := json.Marshal(resp)
	fmt.Println("response:", string(d))
	//fmt.Println("response:", resp)
	return resp, err
}

func main() {
	//data := model2.MqttCheck{model2.MqttHeader{From: "test/cloud/14"}}
	//resp, _ := mqttSendHandle(context.Background(), data)
	//fmt.Printf("%v", resp)

	lambda.Start(mqttSendHandle)
}
