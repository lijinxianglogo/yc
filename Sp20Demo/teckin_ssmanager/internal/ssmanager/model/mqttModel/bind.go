package mqttModel

//设备绑定 Appliance.Control.Bind
type ApplianceControlBind struct {
	Header  Header                      `json:"header"`
	Payload ApplianceControlBindPayload `json:"payload"`
}

type ApplianceControlBindPayload struct {
	Bind Bind `json:"bind"`
}

type Bind struct {
	Uid      int      `json:"uid"`
	Time     TimeInfo `json:"time"`
	BindTime int      `json:"lmTime"`
}

type TimeInfo struct {
	TimeZone  string `json:"timezone"`
	TimeStamp int    `json:"timestamp"`
}

//设备绑定的ACK消息
type ApplianceControlBindAck struct {
	Header  Header         `json:"header"`
	Payload BindAckPayload `json:"payload"`
}

type BindAckPayload struct {
	Result Result `json:"result"`
}

type Result struct {
	Status int    `json:"status"`
	Uuid   string `json:"uuid"`
}
