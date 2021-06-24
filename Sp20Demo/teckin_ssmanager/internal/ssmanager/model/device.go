package model

type DeviceInfo struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Model   string `json:"model"`
	Version string `json:"version"`
	Switch  int32  `json:"switch"`
}

type DeviceList struct {
	Info []DeviceInfo `json:"info"`
}
