package mqttModel

//设备开关 ApplianceControlToggle
type ApplianceControlToggle struct {
	Header  Header                        `json:"header"`
	Payload ApplianceControlTogglePayload `json:"payload"`
}

type ApplianceControlTogglePayload struct {
	Toggle []Toggle `json:"toggle"`
}

type Toggle struct {
	Channel int   `json:"channel"`
	Onoff   int   `json:"onoff"`
	LmTime  int64 `json:"lmTime"`
}
