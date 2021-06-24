package emqclient

import config2 "teckin_ssmanager/config"

func CreateMqttClient() *MqttPoolClient {
	var err error
	EmqxClient, err := initEmqxPoolClient()
	if err != nil {
		panic(err)
	}
	return EmqxClient
}

/**
  获取mqtt客户端
*/
func initEmqxClient() *MqttClient {
	//获取mqtt连接配置
	//初始化默认参数 以及创建连接客户端
	//url := "ssl://a1on17nd3yssr4-ats.iot.ap-northeast-1.amazonaws.com:8883"
	url := config2.Conf.Awsiot.Endpoint

	/*****************     production      ********************/
	clientId := "loyal_test"
	username := ""
	password := ""
	/*****************     production      ********************/

	/*****************     test      ********************/
	//本地无法请求测试环境，该代码为模拟app发送消息的代码，后期需删除
	//clientId := "app:devTest" + uuid.NewV5(uuid.NewV4(), "emqx_client").String()
	//username := "415"
	//password := "ba82f4312aefa5689c4bb287f4de7892"
	/*****************     test      ********************/

	config := NewConfig().
		SetUrl(url).
		SetClientId(clientId).
		SetUserName(username).
		SetPassword(password).
		SetWaitTimeout(5)
	client, _ := NewMqttClient(config)
	return client
}

/**
  获取mqtt的连接池客户端
*/
func initEmqxPoolClient() (*MqttPoolClient, error) {
	//连接池配置
	poolConfig := NewPoolConf().
		SetIdleTimeout(5).
		SetInitialCap(1).
		SetMaxCap(10).
		SetMaxIdle(10).
		SetMaxRetryGetClient(5).
		SetMaxRetryInitPool(10)

	poolClient, err := NewMqttPoolClient(poolConfig, func() (interface{}, error) {
		client := initEmqxClient()
		return client, nil
	})
	return poolClient, err
}
