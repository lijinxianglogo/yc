package http

import (
	"github.com/gin-gonic/gin"
	scv "teckin_ssmanager/pkg/strconv"
)

//用户信息修改
func (s *HttpServer) UserInfoUpdate(ctx *gin.Context) {
	userID, ok := ctx.Params.Get("userID")
	if !ok {
		responseFail(ctx, 400, "缺少userID参数")
		return
	}
	userid := scv.S2I(userID)
	name := ctx.PostForm("Name")
	_, err := s.uidCheckToken(ctx, userid)
	if err != nil {
		return
	}

	if err := s.service.Dao.Mysql.UpdateUserPassword(name, userid); err != nil {
		responseFail(ctx, 400, "密码修改错误,原因:"+err.Error())
		return
	}
	responseSucc(ctx, 200, "密码修改成功", nil)
}
