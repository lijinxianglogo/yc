## emqclient

使用该客户端，通过http的方式向emqx发送命令。<br/>
（*注意：例子中所提供的参数为非真模拟参数*）


## 1.首先需要实例化该客户端。

```golang

    url := "ssl://test-mqtt-ap.meross.com:443"
    clientId := "fmware:1801234948558f10effc_mQf36OYX6zWGfFmf"
    userName := "34:24:85:10:e4:fc"
    password := "1054411_" + md5Str("34:22:8f:60:ef:fcefeae5d261129437bc53c")
    config := emqclient.NewConfig().SetUrl(url).SetClientId(clientId).SetUserName(userName).SetPassword(password)
    	mqttClient := emqclient.New(config)
```
传入参数：
- url:mqtt地址 ，如：ssl://test-mqtt-ap.meross.com:443
- clientId: 客户端id
- userName: 用户名
- password: 密码

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
	topic := "/appliance/1801234948725025132134298f10effc/subscribe"
        messageString := string(json.Marshal(message))
        status, _ := mqttClient.Publish(topic, messageString)
        fmt.Println(status)
```
请求参数：
- topic:mqtt消息的topic，主题
- message:mqtt消息内容。

返回参数：

- status：请求状态，成功为true，错误为false
- err：错误信息，无错误信息为nil



