package http

import (
	"github.com/gin-gonic/gin"
	"teckin_ssmanager/internal/ssmanager/model"
	"teckin_ssmanager/pkg/encoding"
	"time"
)

//用户登录
func (s *HttpServer)UserLogin(ctx *gin.Context){
	var loginResp model.UserLoginResponse
	pwd := ctx.PostForm("Password")
	email := ctx.PostForm("Email")
	if len(pwd) <=0 || len(email) <=0{
		responseFail(ctx, 400, "缺少参数")
		return
	}
	pwd = encoding.StringMD5(pwd)
	user, err := s.service.Dao.Mysql.GetUserInfoByEmailAndPwd(email, pwd)
	if err != nil{
		responseFail(ctx, 400, err.Error())
		return
	}
	token := encoding.StringMD5(email + time.Now().String())
	if err := s.service.Dao.SetUserToken(email, token); err != nil{
		responseFail(ctx, 400, err.Error())
		return
	}
	loginResp.ID = user.ID
	loginResp.Email = email
	loginResp.Name = user.Name
	loginResp.Key = user.Key
	loginResp.Token = token
	responseSucc(ctx, 200, "登录成功", loginResp)
}
