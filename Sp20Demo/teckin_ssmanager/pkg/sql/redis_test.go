package sql

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"teckin_ssmanager/config"
	"testing"
)

func TestRedisConn(t *testing.T){
	r := NewRedis(config.Conf, &logrus.Entry{})
	err := r.Open()
	if err == nil{
		fmt.Println("连接成功")
	}
	//r.Set("pfx", "1234", 14)
	err = r.Del("pfx")
	fmt.Println(err)
}
