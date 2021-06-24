package mqttModel

type Header struct {
	MessageId string `json:"messageId"`
	Namespace string `json:"namespace"`
	Src       string `json:"src"`
	Method    string `json:"method"`
	Sign      string `json:"sign"`
	From      string `json:"from"`
	Ver       string `json:"ver"`
}

var (
	AppOnlineStatusTopic = "iot/online/status/"
)
