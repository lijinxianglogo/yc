package task

import "teckin_ssmanager/internal/ssmanager/service"

func NewTask(srv *service.Service) {
	go InitBind(srv)
	//go InitTest(srv)
}
