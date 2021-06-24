package model

type MyqqBody struct {
	Body string `json:"body"`
}

type MqttCheck struct {
	Header MqttHeader `json:"header"`
}

type MqttHeader struct {
	Topic string `json:"topic"`
}
