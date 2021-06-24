package model

type ConnHeader struct {
	Auth string `json:"AuthInfo"`
}

type ConnResponseHeader struct {
	Header string `json:"my_header"`
}

type ConnectResponseBody struct {
	Key string `json:"key"`
}

type ConnectResponse struct {
	Code int32 `json:"statusCode"`
	Body string `json:"body"`
	IsBase64Encoded bool `json:"isBase64Encoded"`
	ConnHeader ConnResponseHeader `json:"headers"`
}

type ConnRequest struct {
	Header ConnHeader `json:"headers"`
}
