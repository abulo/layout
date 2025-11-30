package crond

import (
	"cloud/dao"
	"cloud/initial"
	"cloud/module/sys/logger"
	"context"
	"encoding/json"

	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
)

func SysLoggerLogin() {
	ctx := context.Background()
	redisHandler := initial.Core.Store.LoadRedis("redis")
	key := util.NewReplacer(initial.Core.Config.String("Cache.SysLoggerLogin.Queue"))
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
		var res dao.SysLoggerLogin
		err = json.Unmarshal([]byte(data), &res)
		if err != nil {
			continue
		}
		logger.SysLoggerLoginCreate(ctx, res)
	}
}
