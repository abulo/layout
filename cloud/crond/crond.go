package crond

import (
	"time"

	"cloud/initial"

	"github.com/abulo/ratel/v3/core/task"
	"github.com/abulo/ratel/v3/core/task/driver"
	"github.com/abulo/ratel/v3/stores/redis"
)

func CronJob() func() {
	redisHandler := initial.Core.Store.LoadRedis("redis")
	client, err := redis.OriginalClient(redisHandler)
	if err != nil {
		panic(err)
	}
	driverHandler := driver.NewRedisZSetDriver(client)
	cron := task.NewTaskWithOption(
		"WorkerService",
		driverHandler,
		task.WithHashReplicas(10),
		task.CronOptionSeconds(),
		task.WithNodeUpdateDuration(time.Second*10),
		task.CronOptionLocation(initial.Core.Local),
	)
	cron.AddFunc("SysMenu", "0 */1 * * * *", SysMenu)                   // 刷新菜单模块名称缓存数据
	cron.AddFunc("SysLoggerOperate", "*/2 * * * * *", SysLoggerOperate) // 后台操作日志写入
	// 后的登录日志写入
	// cron.AddFunc("SystemLoginLogQueue", "*/2 * * * * *", SystemLoginLogQueue)
	// cron.AddFunc("CommonQueue", "*/2 * * * * *", CommonQueue)
	cron.Start()
	return func() { cron.Stop() }
}
