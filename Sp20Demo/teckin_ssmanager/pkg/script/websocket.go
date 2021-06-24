package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"time"
)

type A struct {
	Header Header `json:"header"`
	Play   Playod `json:"playod"`
}
type Header struct {
	From string `json:"from"`
}

var (
	method = "post"
	algorithm = "AWS4-HMAC-SHA256"
	access_key = "AKIAWED5N4XQSN2BHBXT"
	secret_key = "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"
	region     = "ap-northeast-1"
	service    = "apigateway "
	endpoint   = "aws4_request"
	content_type = "application/x-amz-json-1.0"
	host = "l3u1zeammb.execute-api.ap-northeast-1.amazonaws.com"
	end = "https://l3u1zeammb.execute-api.ap-northeast-1.amazonaws.com/production/@connections"
)

type Playod struct {
	Key string `json:"key"`
}

func GetCredential(key, regionName, serviceName, endpoint string) string {
	datastamp := time.Now().Format("20060102")
	kDate := sign("AWS4" + secret_key, datastamp)
	kRegion := sign(kDate, regionName)
	kService := sign(kRegion, serviceName)
	kSigning := sign(kService, endpoint)
	sign := sign(kSigning, signMsg)
	return sign
}

func getCanonicalRequest(amzdate string) string{
	return fmt.Sprintf("content-type:%s\nhost:%s\nx-amz-date:%s\nx-amz-target:%s\n",
		content_type, host, amzdate, )
}

func getStringToSign(amzData string)string{

}

func sign (key, msg string) string{
	k := []byte(key)
	mac := hmac.New(sha256.New, k)
	mac.Write([]byte(msg))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

//websocket连接到aws gateway，并发送数据
func main() {
	var wsurl = "wss://l3u1zeammb.execute-api.ap-northeast-1.amazonaws.com/production"
	var origin = "https://l3u1zeammb.execute-api.ap-northeast-1.amazonaws.com"
	config, _ := websocket.NewConfig(wsurl, origin)



	config.Header.Add("Authorization", "MTo0MTBhZmU0ODE1ODJmMDk3ZTg4OTk3ZjE3ZmVlODgwOQ==")
	config.Header.Add("SignedHeader", "host;x-amz-date")
	config.Header.Add("Content-Type", "application/x-amz-json-1.0")

	w, err := websocket.DialConfig(config)
	defer w.Close()
	fmt.Println(w, err)
	var a A
	a.Header.From = "test/cloud/1"
	a.Play.Key = "kkk"
	d, _ := json.Marshal(a)
	fmt.Println(w.Write(d))
	go func() {
		for {
			msg := make([]byte, 1280)
			n, _ := w.Read(msg)
			fmt.Println(string(msg[:n]))
		}
	}()
	var c = make(chan int)
	<-c
}
