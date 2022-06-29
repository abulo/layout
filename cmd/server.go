package main

import (
	"io/ioutil"
	"layout/cmd/service"
	"layout/initial"
	"layout/services"
	"os"

	"github.com/abulo/ratel/v3/logger"
	"github.com/abulo/ratel/v3/logger/mongo"
	"github.com/abulo/ratel/v3/util"
	"github.com/sarulabs/di"
	"github.com/sirupsen/logrus"
)

func init() {
	//全局设置
	global := initial.Default()
	//获取配置文件目录
	configPath := global.GetEnvironment(global.Path+"/config/env", "configDir")
	if util.Empty(configPath) {
		panic("configPath is empty")
	}
	//配置加载 toml 文件
	dirs := make([]string, 0)
	dirs = append(dirs, global.Path+"/config/"+configPath)
	global.InitConfig(dirs...)
	global.InitMongoDB()
	global.InitRedis()
	global.InitMysql()
	global.InitElasticSearch()
	global.InitClickHouse()
	global.InitTrace()
}

// main 主入口
func main() {
	mgClient := initial.Core.Store.LoadMongoDB("mongodb")
	//日志
	loggerHook := mongo.DefaultWithURL(mgClient)
	defer loggerHook.Flush()
	logger.Logger.AddHook(loggerHook)
	logger.Logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Logger.SetReportCaller(true)
	// 如想不想在终端查看日志输出可以打开注释,生产环境中必须打开
	if !initial.Core.Config.Bool("DisableDebug", true) {
		logger.Logger.SetOutput(os.Stdout)
	} else {
		logger.Logger.SetOutput(ioutil.Discard)
	}

	//依赖注入
	builder, err := di.NewBuilder()
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
	def := make([]di.Def, 0)
	def = append(def, services.BackServices...)

	err = builder.Add(def...)
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
	initial.Core.Container = builder.Build()
	//依赖注入
	defer initial.Core.Container.Delete()

	eng := service.NewEngine()
	//加载计划任务
	// eng.Schedule(eng.CrontabWork())
	// 注册函数
	// eng.RegisterHooks(hooks.Stage_AfterLoadConfig, eng.BeforeInit)
	// eng.Startup(
	// 	eng.AdminServer,
	// 	eng.ApiServer,
	// )
	// if err := eng.Startup(
	// 	eng.AdminServer,
	// 	eng.ApiServer,
	// ); err != nil {
	// 	logger.Logger.WithFields(logrus.Fields{
	// 		"err": err,
	// 	}).Panic("startup")
	// }
	// err = eng.Run()
	// logger.Logger.Info(eng.Run())

	if err := eng.Run(); err != nil {
		logger.Logger.Panic(err.Error())
	}
}
