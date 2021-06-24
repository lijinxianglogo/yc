package task

import (
	"fmt"
	"strconv"
	"teckin_ssmanager/internal/ssmanager/service"
)

func InitTest(srv *service.Service) {
	fmt.Println("kaishi")
	for i := 0; i < 100; i++ {
		str := strconv.Itoa(i)
		_, _ = srv.Dao.EmqxClient.Publish("test121", str)
		fmt.Println("publish==========" + str)
	}
}
