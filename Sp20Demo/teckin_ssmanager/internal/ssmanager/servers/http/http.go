package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"teckin_ssmanager/config"
	"teckin_ssmanager/internal/ssmanager/service"
	"time"
)

type HttpServer struct {
	httpsrv *http.Server
	service *service.Service
}

func New(srv *service.Service, conf *config.HttpServerConfig) *HttpServer {
	engine := gin.Default()

	go func() {
		//TODO 未加https逻辑判断 需要加加密钥及证书
		//TODO 端口应由conf传入参数，现直接写，以后改
		if err := engine.Run(conf.Addr); err != nil {
			panic(err)
		}
	}()
	server := &http.Server{
		Addr:           conf.Addr,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s := &HttpServer{
		httpsrv: server,
		service: srv,
	}
	s.InitRouter()
	return s

}

func (s *HttpServer) InitRouter() {
	engine := s.httpsrv.Handler.(*gin.Engine)
	ugroup := engine.Group("/user")
	dgroup := engine.Group("/device")
	//登录及注册不需要校验token
	ugroup.POST("/login", s.UserLogin)
	ugroup.POST("/register", s.UserRegister)
	ugroup.Use(headerTokenVerify)
	ugroup.POST("/signout", s.UserSignOut)
	ugroup.POST("/info/:userID", s.UserInfoUpdate)
	ugroup.POST("/password/:userID", s.UserPasswordUpdate)
	dgroup.POST("/getinfo", s.DeviceGetInfo)
	dgroup.POST("/updatename", s.DeviceNameUpdate)
	dgroup.POST("/add", s.DeviceAdd)
}

func (s *HttpServer) Close() {
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err := s.httpsrv.Shutdown(cxt)
	if err != nil{
		fmt.Println("err", err)
	}
	fmt.Println("------exited--------", time.Since(now))
}
