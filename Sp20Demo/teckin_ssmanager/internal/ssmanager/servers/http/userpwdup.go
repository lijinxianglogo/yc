package http

import (
	"github.com/gin-gonic/gin"
	scv "teckin_ssmanager/pkg/strconv"
)

//用户密码修改
func (s *HttpServer) UserPasswordUpdate(ctx *gin.Context) {
	uid, ok := ctx.Params.Get("userID")
	if !ok {
		responseFail(ctx, 400, "缺少userID参数")
		return
	}
	userid := scv.S2I(uid)
	oldPwd := ctx.PostForm("OldPassword")
	newPwd := ctx.PostForm("NewPassword")

	if oldPwd == newPwd {
		responseFail(ctx, 400, "新旧密码不能相同")
		return
	}
	user, err := s.uidCheckToken(ctx, userid)
	if err != nil {
		return
	}

	if oldPwd != user.Password {
		responseFail(ctx, 400, "旧密码输入错误")
		return
	}

	if err := s.service.Dao.Mysql.UpdateUserPassword(newPwd, userid); err != nil {
		responseFail(ctx, 400, "密码修改错误,原因:"+err.Error())
		return
	}
	responseSucc(ctx, 200, "密码修改成功", nil)
}
