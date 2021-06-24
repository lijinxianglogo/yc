package emqclient

//import (
//	"crypto/md5"
//	"encoding/hex"
//	"encoding/json"
//	"fmt"
//	"strconv"
//	"testing"
//	"time"
//)
//
//func TestMqttClient(t *testing.T) {
//	//sendPublish()
//
//}
//func sendPublish() {
//	url := "ssl://test-mqtt-ap.meross.com:443"
//	clientId := "fmware:1801234948725025132134298f10eff1_Yz9hqq7F6s000015"
//	userName := "34:29:8f:10:ef:fc"
//	password := "10500311_" + md5Str(userName+"efeae5d2611291aa845745408e7bc53c")
//	config := NewConfig().SetUrl(url).SetClientId(clientId).SetUserName(userName).SetPassword(password)
//	mqttClient := InitCloudClient(config)
//
//	message := make(map[string]interface{})
//	header := make(map[string]interface{})
//	payload := make(map[string]interface{})
//	toggle := make(map[string]interface{})
//	time := time.Now().Unix()
//
//	header["messageId"] = "7048c18d41871bf0eda3e54dee3f7"
//	header["timestamp"] = time
//	header["namespace"] = "Appliance.Control.Toggle"
//	header["method"] = "SET"
//	sign := md5Str(header["messageId"].(string) + "efeae5d2611291aa845745408e7bc53c" + strconv.FormatInt(time, 10))
//	header["sign"] = sign
//	header["from"] = "TopicAppSubsAck"
//	toggle["onoff"] = 0
//	toggle["lmTime"] = time
//	payload["toggle"] = toggle
//	message["header"] = header
//	message["payload"] = payload
//	topic := "/appliance/1801234948725025132134298f10effc/subscribe"
//	messageByte, _ := json.Marshal(message)
//	messageString := string(messageByte)
//	fmt.Println(messageString)
//	status, err := mqttClient.Publish(topic, messageString)
//	fmt.Println(status, err)
//}
//func md5Str(str string) string {
//	h := md5.New()
//	_, _ = h.Write([]byte(str)) // 需要加密的字符串为 123456
//	cipherStr := h.Sum(nil)
//	return hex.EncodeToString(cipherStr) // 输出加密结果
//}
