package http

import (
	"fmt"
	utils "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
	"testing"
)
func init() {
	router := gin.Default()  // 这需要写到init中，启动gin框架
	g := router.Group("user")
	g.PUT("/password/:userID", )
	utils.SetRouter(router)  //把启动的engine 对象传入到test框架中
}

func TestLoginHandler(t *testing.T) {
	// 定义发送POST请求传的内容
	user := map[string]interface{}{
		"OldPassword": "123456",
		"NewPassword": "123458",
	}
	// 把返回response解析到resp中
	resp := Response{}
	// 调用函数发起http请求
	utils.AddHeader("token", "9f801a376f716f942810a3beb5c44dd3")
	err := utils.TestHandlerUnMarshalResp("PUT", "/password/1", "form", user, &resp)
	if err != nil {
		t.Errorf("TestLoginHandler: %v\n", err)
		return
	}
	// 得到返回数据结构体， 至此，完美完成一次post请求测试，
	// 如果需要benchmark 输出性能报告也是可以的
	fmt.Println("result:", resp)
}