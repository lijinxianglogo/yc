package http

import (
	"github.com/gin-gonic/gin"
	"teckin_ssmanager/internal/ssmanager/model"
	scv "teckin_ssmanager/pkg/strconv"
)


func (s *HttpServer) DeviceAdd(ctx *gin.Context) {
	var err error
	var devInfo model.DeviceInfo
	uid := ctx.PostForm("userID")
	devInfo.Code = ctx.PostForm("code")
	devInfo.Name = ctx.PostForm("name")
	devInfo.Model = ctx.PostForm("model")
	devInfo.Version = ctx.PostForm("version")

	if uid == "" || devInfo.Code == ""{
		responseFail(ctx, 400, "缺少参数")
		return
	}

	userid := scv.S2I(uid)
	_, err =s.uidCheckToken(ctx, userid)
	if err != nil{
		return
	}

	if err := s.service.Dao.Mysql.InsertDeviceInfo(devInfo, userid); err != nil{
		responseFail(ctx, 400, err.Error())
		return
	}

	responseSucc(ctx, 200, "添加成功", nil)
}
