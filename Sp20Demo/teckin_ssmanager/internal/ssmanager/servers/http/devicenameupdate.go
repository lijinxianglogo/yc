package http

import (
	"github.com/gin-gonic/gin"
	scv "teckin_ssmanager/pkg/strconv"
)


func (s *HttpServer) DeviceNameUpdate(ctx *gin.Context) {
	var err error
	uid := ctx.PostForm("userID")
	devName := ctx.PostForm("deviceName")
	devCode := ctx.PostForm("deviceCode")

	if devCode== "" || devName == ""{
		responseFail(ctx, 400, "缺少参数")
		return
	}

	userid := scv.S2I(uid)
	_, err =s.uidCheckToken(ctx, userid)
	if err != nil{
		return
	}

	if err := s.service.Dao.Mysql.UpdateDeviceName(devName, devCode, userid); err != nil{
		responseFail(ctx, 400, err.Error())
		return
	}

	responseSucc(ctx, 200, "修改成功", nil)
}
