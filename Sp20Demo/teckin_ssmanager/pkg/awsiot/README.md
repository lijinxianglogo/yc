## awsiot

使用该客户端，通过http的方式向awsIOTcore发送命令<br/>
 （*注意：例子中所提供的参数为非真模拟参数*）

## 1.首先需要实例化该客户端。

```golang
     config :=  awsiotclient.NewConfig().
                  SetEndpoint("https://search-test-meross-o4rdxkrmjg6g2ubu2vhvj265za.ap-northeast-1.es.amazonaws.com").
                  SetAkId("AKIAQN5ENAWXLISW2EVU").
                  SetSecretKey("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").
                  SetRegion("ap-northeast-1")
    mqttClient := awsiotclient.New(config)
```
配置参数：
- AkId:具有相关权限的秘钥id
- SecretKey:具有相关权限的秘钥
- Region:服务所在区域
- Endpoint:AWS IoT 的终端节点

返回参数：
- client：实例化的客户端

## 2.开始调用,推送mqtt消息。Publish方法
```golang
message := make(map[string]interface{})
	header := make(map[string]interface{})
	payload := make(map[string]interface{})
	toggle := make(map[string]interface{})
	time := time.Now().Unix()

	header["messageId"] = "65461bf0765eda3e54dee3f7"
	header["timestamp"] = time
	header["namespace"] = "Appliance.Control.Toggle"
	header["method"] = "SET"
	sign := md5Str(header["messageId"].(string) + "efeae5d261165447bc53c" + strconv.FormatInt(time,10))
	header["sign"] = sign
	header["from"] = "TopicAppSubsAck"
	toggle["onoff"] = 0
	toggle["lmTime"] = time
	payload["toggle"] = toggle
	message["header"] = header
	message["payload"] = payload
	topic := "sdk/test/Python"
    messageString := string(json.Marshal(message))
err := mqttClient.Publish(topic, messageString)
```
请求参数：
- topic:mqtt消息的topic，主题
- message:mqtt消息内容。

返回参数：
- err：错误信息，无错误信息为nil




