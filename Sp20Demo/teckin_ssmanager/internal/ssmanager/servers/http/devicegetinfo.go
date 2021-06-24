package http

import (
	"github.com/gin-gonic/gin"
	"teckin_ssmanager/internal/ssmanager/model"
	scv "teckin_ssmanager/pkg/strconv"
)

//获取设备信息
func (s *HttpServer) DeviceGetInfo(ctx *gin.Context) {
	var devInfo model.DeviceList
	var err error
	uid := ctx.PostForm("userID")
	devCode := ctx.PostForm("deviceCode")

	userid := scv.S2I(uid)

	_, err =s.uidCheckToken(ctx, userid)
	if err != nil{
		return
	}

	if devCode == ""{  //设备码为空代表获取所有用户关联设备
		devInfo, err = s.service.Dao.Mysql.GetUserAllRelateDeviceInfo(userid)
	} else{
		devInfo, err = s.service.Dao.Mysql.GetUserRelateDeviceInfoByDeviceCode(userid, devCode)
	}


	if err != nil{
		responseFail(ctx, 400, err.Error())
		return
	}
	if len(devInfo.Info) > 0{  //获取设备在线状态
		for i, _ := range devInfo.Info{
			devInfo.Info[i].Switch = 2
		}
	}
	responseSucc(ctx, 200, "获取成功", devInfo)
}
