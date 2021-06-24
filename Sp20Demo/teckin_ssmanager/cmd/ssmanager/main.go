package main

import (
	"os"
	"os/signal"
	"syscall"
	"teckin_ssmanager/config"
	"teckin_ssmanager/internal/ssmanager/servers/grpc"
	"teckin_ssmanager/internal/ssmanager/servers/http"
	"teckin_ssmanager/internal/ssmanager/service"
	"teckin_ssmanager/internal/ssmanager/task"
)

func main() {
	if err := config.Init(); err != nil {
		//panic(err)
		config.DefaultConfig()
	}
	srv := service.New()
	httpsrv := http.New(srv, config.Conf.Http)
	grpcsrv := grpc.New(srv)
	grpcsrv.Dial(config.Conf.Grpc)
	go task.NewTask(srv)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			httpsrv.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
