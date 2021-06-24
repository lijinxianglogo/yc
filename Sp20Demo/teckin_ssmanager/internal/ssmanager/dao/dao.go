package dao

import (
	"github.com/sirupsen/logrus"
	"teckin_ssmanager/config"
	"teckin_ssmanager/internal/ssmanager/dbs"
	"teckin_ssmanager/pkg/awsiot"
	"teckin_ssmanager/pkg/emqclient"
	"teckin_ssmanager/pkg/kafka"
	"teckin_ssmanager/pkg/sql"
)

type Dao struct {
	Mysql        *dbs.MysqlDB
	Redis        *sql.Redis
	Kafka        *kafka.Kafka
	logger       *logrus.Entry
	EmqxClient   *emqclient.MqttPoolClient
	AwsIotClient *awsiot.MqttClient
}

func New(logger *logrus.Entry) *Dao {
	dao := &Dao{
		logger: logger,
		Mysql:  dbs.InitDB(config.Conf.Mysql, logger),
		Redis:  sql.NewRedis(config.Conf.Redis, logger),
		//Kafka:      kafka.NewKafka(config.Conf),
		AwsIotClient: awsiot.New(config.Conf),
		EmqxClient:   emqclient.CreateMqttClient(),
	}
	_ = dao.Redis.Open()
	return dao
}
