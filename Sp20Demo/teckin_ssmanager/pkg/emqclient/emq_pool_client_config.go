package emqclient

import "time"

type EmqPoolConf struct {
	InitialCap        int           //连接池中拥有的最小连接数
	MaxCap            int           //最大并发存活连接数
	MaxIdle           int           //最大空闲连接
	IdleTimeout       time.Duration //连接最大空闲时间，超过该事件则将失效
	MaxRetryInitPool  int           //初始化连接池超时时间，必须大于客户端处理超时时间
	MaxRetryGetClient int           //获取客户端最大重试次数
}

func NewPoolConf() *EmqPoolConf {
	c := &EmqPoolConf{}
	return c
}

/*
连接池中拥有的最小连接数,必填项
*/
func (c *EmqPoolConf) SetInitialCap(initialCap int) *EmqPoolConf {
	c.InitialCap = initialCap
	return c
}

/*
最大并发存活连接数,必填项
*/
func (c *EmqPoolConf) SetMaxCap(maxCap int) *EmqPoolConf {
	c.MaxCap = maxCap
	return c
}

/*
连接最大空闲时间，超过该事件则将失效,必填项
*/
func (c *EmqPoolConf) SetIdleTimeout(idleTimeout time.Duration) *EmqPoolConf {
	c.IdleTimeout = idleTimeout
	return c
}

/*
最大空闲连接,必填项
*/
func (c *EmqPoolConf) SetMaxIdle(maxIdle int) *EmqPoolConf {
	c.MaxIdle = maxIdle
	return c
}

/*
初始化连接池超时时间，必须大于客户端处理超时时间,必填项
*/
func (c *EmqPoolConf) SetMaxRetryInitPool(maxRetryInitPool int) *EmqPoolConf {
	c.MaxRetryInitPool = maxRetryInitPool
	return c
}

/*
获取客户端最大重试次数,必填项
*/
func (c *EmqPoolConf) SetMaxRetryGetClient(maxRetryGetClient int) *EmqPoolConf {
	c.MaxRetryGetClient = maxRetryGetClient
	return c
}
