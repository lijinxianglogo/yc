package awsiot

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"teckin_ssmanager/config"
	"testing"
	"time"
)

func TestMqttClient(t *testing.T) {
	sendMsg()

}
func sendMsg() {
	t := time.Now()
	tNano := strconv.Itoa(int(t.UnixNano()/1e6))
	fmt.Println("time start:", t.Format("15:04:05.")+tNano[len(tNano)-3:])
	akId := "AKIAWED5N4XQVXOQCAPN"
	secretKey := "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"
	region := "ap-northeast-1"
	endpoint := "a1on17nd3yssr4-ats.iot.ap-northeast-1.amazonaws.com"
	mqttClient := New(&config.Config{Awsiot: &config.AwsIotConfig{
		AkId:      akId,
		SecretKey: secretKey,
		Region:    region,
		Endpoint:  endpoint,
	}})

	message := make(map[string]interface{})
	header := make(map[string]interface{})
	payload := make(map[string]interface{})
	toggle := make(map[string]interface{})
	time := time.Now().Unix()

	header["messageId"] = "7048c18d41871bf0eda3e54dee3f7"
	header["timestamp"] = time
	header["namespace"] = "Appliance.Control.Toggle"
	header["method"] = "SET"
	sign := md5Str(header["messageId"].(string) + "efeae5d2611291aa845745408e7bc53c" + strconv.FormatInt(time, 10))
	header["sign"] = sign
	header["from"] = "TopicAppSubsAck"
	toggle["onoff"] = 0
	toggle["lmTime"] = time
	payload["toggle"] = toggle
	message["header"] = header
	message["payload"] = payload
	topic := "sdk/test/Python"
	messageByte, _ := json.Marshal(message)
	messageString := string(messageByte)
	err := mqttClient.Publish(topic, []byte(messageString))
	fmt.Println(err)
}
func md5Str(str string) string {
	h := md5.New()
	_, _ = h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 输出加密结果
}
