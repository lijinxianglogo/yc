//用于校验用户信息的lambda函数
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	_ "net/http/pprof"
	"teckin_ssmanager/pkg/encoding"
	"teckin_ssmanager/pkg/script/grpc"
	"teckin_ssmanager/pkg/sql"
)





type dao struct {
	redis *sql.Redis
}

type statement struct {
	Action   string `json:"Action"`
	Effect   string `json:"Effect"`
	Resource string `json:"Resource"`
}

type publishStatement struct {
	statement
}

type connectStatement struct {
	statement
}

type subStatement struct {
	statement
}

type recStatement struct {
	statement
}

type policyDocument struct {
	Statement []interface{} `json:"Statement"`
	Version   string        `json:"Version"`
}

type authResponse struct {
	IsAuthenticated          bool             `json:"isAuthenticated"`
	PrincipalId              string           `json:"principalId"`
	PolicyDocuments          []policyDocument `json:"policyDocuments"`
	DisconnectAfterInSeconds int              `json:"disconnectAfterInSeconds"`
	RefreshAfterInSeconds    int              `json:"refreshAfterInSeconds"`
}

type tls struct {
	ServerName string `json:"serverName"`
}

type http struct {
	Headers     map[string]string `json:"headers"`
	QueryString string            `json:"queryString"`
}

type mqtt struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	ClientId string `json:"clientId"`
}

type protocolData struct {
	Tls  tls  `json:"tls"`
	Http http `json:"http"`
	Mqtt mqtt `json:"mqtt"`
}

type connectionMetadata struct {
	Id string `json:"id"`
}

type event struct {
	Token              string             `json:"token"`
	SignatureVerified  bool               `json:"signatureVerified"`
	Protocols          []string           `json:"protocols"`
	ProtocolData       protocolData       `json:"protocolData"`
	ConnectionMetadata connectionMetadata `json:"connectionMetadata"`
}

func Handle(ctx context.Context, e event) (string, error) {
	fmt.Println("event info: ", e)
	cli, conn := grpc.GetLambdaClient()
	defer conn.Close()
	password, err := grpc.GetUserInfo(cli, e.ProtocolData.Mqtt.UserName)
	if err != nil{
		password = "0"
	}

	pwd := encoding.DecodeBase64(e.ProtocolData.Mqtt.Password)
	switch pwd {
	case password:
		data, err := json.Marshal(generateAuthResponse(2, "Allow"))
		return string(data), err
	default:
		authResp := generateAuthResponse(2, "Deny")
		data, err := json.Marshal(authResp)
		return string(data), err
	}
}



func generateAuthResponse(token int, effect string) authResponse {
	var authResp authResponse
	authResp.IsAuthenticated = true
	authResp.PrincipalId = encoding.HexTimeNow()

	var policyDoc policyDocument
	policyDoc.Version = "2012-10-17"
	var pubStat publishStatement
	var connStat connectStatement
	var subStat subStatement
	var recStat recStatement

	connStat.Action = "iot:Connect"
	connStat.Effect = effect
	connStat.Resource = "*"
	pubStat.Action = "iot:Publish"
	pubStat.Effect = effect
	pubStat.Resource = "*"
	subStat.Action = "iot:Subscribe"
	subStat.Effect = effect
	subStat.Resource = "*"
	recStat.Action = "iot:Receive"
	recStat.Effect = effect
	recStat.Resource = "*"

	policyDoc.Statement = append(policyDoc.Statement, pubStat)
	policyDoc.Statement = append(policyDoc.Statement, connStat)
	policyDoc.Statement = append(policyDoc.Statement, subStat)
	policyDoc.Statement = append(policyDoc.Statement, recStat)

	authResp.PolicyDocuments = append(authResp.PolicyDocuments, policyDoc)
	authResp.DisconnectAfterInSeconds = 3600
	authResp.RefreshAfterInSeconds = 300
	return authResp
}

func main() {
	lambda.Start(Handle)

}
