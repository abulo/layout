package service

import (
	"github.com/abulo/ratel/v3/app"
	"github.com/abulo/ratel/v3/logger"
	"github.com/sirupsen/logrus"
)

type Engine struct {
	app.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	//加载计划任务
	// eng.Schedule(eng.CrontabWork())
	// 注册函数
	// eng.RegisterHooks(hooks.Stage_AfterLoadConfig, eng.BeforeInit)
	if err := eng.Startup(
		eng.AdminServer,
		// eng.ApiServer,
	); err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Panic("startup")
	}
	return eng
}
