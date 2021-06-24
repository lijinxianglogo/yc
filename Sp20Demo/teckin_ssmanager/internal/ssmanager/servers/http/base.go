package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"teckin_ssmanager/internal/ssmanager/model"
	"teckin_ssmanager/pkg/encoding"
	"time"
)

type Header struct {
	Token     string `json:"token"`                        //用户登录后token，没有登录则为空字符串
	Vendor    string `json:"vendor" binding:"required"`    //APP类型，teckin,teckinBeta
	Version   string `json:"version" binding:"required"`   //版本号，如1.1.36
	Apptype   string `json:"apptype" binding:"required"`   //app系统平台，iOS、 Android
	TimeStamp string `json:"TimeStamp" binding:"required"` //当前UNIX时间戳（秒级）
	Nonce     string `json:"nonce" binding:"required"`     //随机流水号（防止重复提交）
	I18n      string `json:"i18n"`                         //多语言，如zh_CN，简体中文
	Sign      string `json:"sign" binding:"required"`      //请求签名，全大写
}

type Response struct {
	Code      int
	Msg       string
	Data      interface{}
	MessageId string
	TimeStamp int64
}

func responseSucc(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(200, Response{
		Code:      code,
		Msg:       msg,
		Data:      data,
		MessageId: encoding.HexTimeNow(),
		TimeStamp: time.Now().Unix(),
	})
}

func responseFail(ctx *gin.Context, code int, msg string) {
	ctx.JSON(200, Response{
		Code:      code,
		Msg:       msg,
		MessageId: encoding.HexTimeNow(),
		TimeStamp: time.Now().Unix(),
	})
}

func (s *HttpServer)tokenCheck(ctx *gin.Context, email string) error {
	token := ctx.GetHeader("token")
	token2, err := s.service.Dao.GetUserToken(email)
	if err != nil {
		return err
	}
	if token != token2 {
		return errors.New("token错误")
	}
	return nil
}

func (s *HttpServer)uidCheckToken(ctx *gin.Context, userid int) (user model.User, err error) {
	user, err = s.service.Dao.Mysql.GetUserInfoByUserID(userid)
	if err != nil {
		responseFail(ctx, 400, err.Error())
		return
	}
	if err = s.tokenCheck(ctx, user.Email); err != nil {
		responseFail(ctx, 400, err.Error())
		return
	}
	return user, nil
}

func Aa(){

}