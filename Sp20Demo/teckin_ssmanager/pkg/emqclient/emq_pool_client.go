package emqclient

import (
	"errors"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	clientPool Pool
)

type MqttPoolClient struct {
	Pool       Pool
	PoolConfig *EmqPoolConf
}

func NewMqttPoolClient(poolConfig *EmqPoolConf, factory func() (interface{}, error)) (client *MqttPoolClient, err error) {
	client = &MqttPoolClient{
		PoolConfig: poolConfig,
	}
	client.Pool, err = initPool(poolConfig, factory, 0)
	return client, err
}

func initPool(poolConf *EmqPoolConf, factory func() (interface{}, error), retryCount int) (Pool, error) {
	if retryCount > poolConf.MaxRetryInitPool {
		return nil, errors.New("err to retry init pool of emqx client")
	}
	if clientPool != nil {
		return clientPool, nil
	}

	//close 关闭连接的方法
	closeFunc := func(v interface{}) error {
		v.(*MqttClient).Close(0)
		return nil
	}

	poolConfig := &Config{
		//连接池中拥有的最小连接数
		InitialCap: poolConf.InitialCap,
		//最大空闲连接
		MaxIdle: poolConf.MaxIdle,
		//最大并发存活连接数
		MaxCap:  poolConf.MaxCap,
		Factory: factory,
		Close:   closeFunc,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: poolConf.IdleTimeout * time.Second,
	}
	clientPool, _ = NewChannelPool(poolConfig)
	if clientPool == nil {
		time.Sleep(1000 * time.Millisecond)
		_, _ = initPool(poolConf, factory, retryCount+1)
	}
	return clientPool, nil
}

func (c *MqttPoolClient) getClientFromPool(retryCount int) (interface{}, error) {
	poolClient, err := c.Pool.Get()
	if err != nil {
		retryCount++
		if retryCount > c.PoolConfig.MaxRetryGetClient {
			return nil, errors.New("max retry to get pool client failed")
		}
		poolClient, err = c.getClientFromPool(retryCount)
	}
	return poolClient, nil
}

func (c *MqttPoolClient) SyncPublish(from, topic string, message string) (msg string, err error) {
	poolClient, err := c.getClientFromPool(0)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = c.Pool.Put(poolClient)
	}()
	client := poolClient.(*MqttClient)
	msg, err = client.SyncPublish(from, topic, message)
	return msg, err
}

func (c *MqttPoolClient) Publish(topic string, message string) (status bool, err error) {
	poolClient, err := c.getClientFromPool(0)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = c.Pool.Put(poolClient)
	}()
	client := poolClient.(*MqttClient)
	return client.Publish(topic, message)
}

func (c *MqttPoolClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) (status bool, err error) {
	poolClient, err := c.getClientFromPool(0)
	if err != nil {
		fmt.Println("111111111111111111111")
		return false, err
	}
	defer func() {
		_ = c.Pool.Put(poolClient)
	}()
	client := poolClient.(*MqttClient)
	return client.Subscribe(topic, qos, callback)
}

func (c *MqttPoolClient) UnSubscribe(topic []string) (status bool, err error) {
	poolClient, err := c.getClientFromPool(0)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = c.Pool.Put(poolClient)
	}()
	client := poolClient.(*MqttClient)
	return client.UnSubscribe(topic)
}
