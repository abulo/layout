package crond

import (
	"cloud/initial"
	"cloud/module/sys/menu"
	"context"

	"github.com/abulo/ratel/v3/util"
)

func SysMenu() {
	ctx := context.Background()
	condition := make(map[string]any)
	// 获取数据集合
	list, err := menu.SysMenuList(ctx, condition)
	if err != nil {
		return
	}

	redisHandler := initial.Core.Store.LoadRedis("redis")
	key := util.NewReplacer(initial.Core.Config.String("Cache.SysMenu.Module"))
	menuKey := util.NewReplacer(initial.Core.Config.String("Cache.SysMenu.Code"))

	permissionList := make([]any, 0)
	for _, v := range list {
		if v.Code.String != "" {
			permissionList = append(permissionList, v.Code.String)
			redisHandler.SAdd(ctx, menuKey, v.Code.String)
			itemList, err := menu.SysMenuListRecursive(ctx, *v.Id)
			if err != nil {
				continue
			}
			data := make([]string, 0)
			for _, item := range itemList {
				data = append(data, *item.Name)
			}
			redisHandler.HSet(ctx, key, v.Code.String, util.Implode("-", data))
		}
	}
	// 现在需要删除没有用的数据
	if modelKeys, err := redisHandler.HKeys(ctx, key); err == nil {
		for _, v := range modelKeys {
			if !util.InArray(v, permissionList) {
				redisHandler.HDel(ctx, key, v)
			}
		}
	}
	if permissionKeys, err := redisHandler.SMembers(ctx, menuKey); err == nil {
		for _, v := range permissionKeys {
			if !util.InArray(v, permissionList) {
				redisHandler.SRem(ctx, menuKey, v)
			}
		}
	}

}
