package service

import (
	"github.com/sirupsen/logrus"
	"teckin_ssmanager/internal/ssmanager/dao"
	"time"
)

type Service struct {
	Dao *dao.Dao
	Logger *logrus.Entry
}

func New() *Service{
	logger := logrus.WithFields(logrus.Fields{
		"time": time.Now(),
	})
	return &Service{
		Logger: logger,
		Dao: dao.New(logger),
	}
}