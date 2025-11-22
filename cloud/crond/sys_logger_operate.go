package crond

import (
	"cloud/dao"
	"cloud/initial"
	"cloud/module/sys/logger"
	"context"
	"encoding/json"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
)

func SysLoggerOperate() {
	ctx := context.Background()
	redisHandler := initial.Core.Store.LoadRedis("redis")
	key := util.NewReplacer(initial.Core.Config.String("Cache.SysLoggerOperate.Queue"))
	systemMenuModuleKey := util.NewReplacer(initial.Core.Config.String("Cache.SysMenu.Module"))
	//获取列队长度
	currentLen, err := redisHandler.LLen(ctx, key)
	if err != nil {
		return
	}
	length := cast.ToInt64(300)
	if currentLen > length {
		currentLen = length
	}

	for i := 0; i < cast.ToInt(currentLen); i++ {
		data, err := redisHandler.RPop(ctx, key)
		if err != nil {
			continue
		}
		var res dao.SysLoggerOperate
		err = json.Unmarshal([]byte(data), &res)
		if err != nil {
			continue
		}
		method := util.Explode("/", cast.ToString(res.GoMethod.Ptr()))
		handlerName := method[len(method)-1]
		// 获取菜单模块名称
		moduleName, err := redisHandler.HGet(ctx, systemMenuModuleKey, handlerName)
		if err != nil {
			continue
		}
		if moduleName == "" {
			continue
		}
		res.Module = null.StringFrom(moduleName)
		logger.SysLoggerOperateCreate(ctx, res)
	}
}
