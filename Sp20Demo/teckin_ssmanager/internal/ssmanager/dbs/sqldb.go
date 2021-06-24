package dbs

import (
	"github.com/sirupsen/logrus"
	"teckin_ssmanager/config"
	"teckin_ssmanager/pkg/sql"
)

type MysqlDB struct {
	conn sql.SqlFactory
	Logger *logrus.Entry
}

func InitDB(conf *config.MysqlConfig, log *logrus.Entry) *MysqlDB{
	conn := sql.NewSql(0)
	err := conn.Open(conf)
	if err != nil{
		panic(err)
	}
	return &MysqlDB{conn:conn}
}
