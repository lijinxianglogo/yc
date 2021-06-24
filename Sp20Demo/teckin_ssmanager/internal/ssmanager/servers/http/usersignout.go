package http

import (
	"github.com/gin-gonic/gin"
)

//用户登出
func (s *HttpServer) UserSignOut(ctx *gin.Context) {
	email := ctx.PostForm("Email")
	if len(email) <= 0 {
		responseFail(ctx, 400, "参数不能为空")
		return
	}
	if err := s.tokenCheck(ctx, email); err != nil {
		responseFail(ctx, 400, err.Error())
		return
	}
	if err := s.service.Dao.DelUserToken(email); err != nil {
		responseFail(ctx, 400, "登出失败, 原因:"+err.Error())
		return
	}
	responseSucc(ctx, 200, "登出成功", nil)
}
