package emqclient

import "time"

type EmqConfig struct {
	url         string        // mqtt地址 ，如：ssl://test-mqtt-ap.meross.com:443
	clientId    string        // 客户端id
	userName    string        // 用户名
	password    string        // 密码
	waitTimeout time.Duration // 等待mqtt消息回复的时间
}

func NewConfig() *EmqConfig {
	c := &EmqConfig{
		url:         "",
		clientId:    "",
		userName:    "",
		password:    "",
		waitTimeout: 0,
	}
	return c
}

/*
设置服务地址,必填项
*/
func (c *EmqConfig) SetUrl(url string) *EmqConfig {
	c.url = url
	return c
}

/*
设置clientid，必填项
*/
func (c *EmqConfig) SetClientId(clientId string) *EmqConfig {
	c.clientId = clientId
	return c
}

/*
设置用户名，必填项
*/
func (c *EmqConfig) SetUserName(userName string) *EmqConfig {
	c.userName = userName
	return c
}

/*
设置密码，必填项
*/
func (c *EmqConfig) SetPassword(password string) *EmqConfig {
	c.password = password
	return c
}

/*
等待mqtt消息回复的时间，非必填项
*/
func (c *EmqConfig) SetWaitTimeout(waitTimeout time.Duration) *EmqConfig {
	c.waitTimeout = waitTimeout
	return c
}
