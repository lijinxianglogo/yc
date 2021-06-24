package dbs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"teckin_ssmanager/config"
	"teckin_ssmanager/internal/ssmanager/model"
	"testing"
)

func NewService() *MysqlDB {
	return InitDB(config.Conf, &logrus.Entry{})
}

func TestInsertDeviceInfo(t *testing.T) {
	db := NewService()
	a := model.DeviceInfo{
		Code:"08100100000001",
		Name: "one",
		Model: "SP20",
		Version: "v1.0",
	}
	if err := db.InsertDeviceInfo(a, 1); err != nil{
		fmt.Println(err.Error())
	}
}
