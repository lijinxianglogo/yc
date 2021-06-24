package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"teckin_ssmanager/internal/ssmanager/model"
	"teckin_ssmanager/pkg/encoding"
	"time"
)


//用户注册
func (s *HttpServer)UserRegister(ctx *gin.Context){
	user := model.User{}
	pwd := ctx.PostForm("Password")
	name := ctx.PostForm("Name")
	email := ctx.PostForm("Email")
	verifyCode := ctx.PostForm("VerifyCode")
	if len(pwd) <=0 || len(name) <=0 || len(email) <=0{
		responseFail(ctx, 400, "缺少参数")
		return
	}

	if err:= s.service.Dao.Mysql.CheckUserExist(email);err != nil{
		responseFail(ctx, 400, err.Error())
		return
	}

	//TODO demo版本默认验证码为123456
	if verifyCode != "123456"{
		responseFail(ctx, 400, "验证码错误")
		return
	}

	user.Name = name
	user.Password = encoding.StringMD5(pwd)
	user.Email = email
	user.Key = encoding.StringMD5(email)  //TODO 生成用户唯一识别码，暂用email生成MD5
	user.AddTime = time.Now().Unix()
	uid, err := s.service.Dao.Mysql.CreateUser(user)
	if err != nil{
		responseFail(ctx, 400, "注册失败, 原因:" + err.Error())
		return
	}
	if err := s.service.Dao.InsertUserInfo(uid, user); err != nil{
		fmt.Println("redis:插入用户信息报错：", err)
	}

	responseSucc(ctx, 200, "注册成功", nil)
}
