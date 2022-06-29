package initial

import (
	"github.com/abulo/ratel/v3/stores/proxy"
	"github.com/abulo/ratel/v3/stores/redis"
	"github.com/abulo/ratel/v3/util"
)

// InitRedis load redis && returns an redis instance.
func (initial *Initial) InitRedis() *Initial {
	configs := initial.Config.Get("redis")
	list := configs.(map[string]interface{})
	links := make(map[string]*redis.Client)
	for node, nodeConfig := range list {
		opt := &redis.Config{}
		res := nodeConfig.(map[string]interface{})
		if KeyPrefix := util.ToString(res["KeyPrefix"]); KeyPrefix != "" {
			opt.KeyPrefix = KeyPrefix
		}
		if Password := util.ToString(res["Password"]); Password != "" {
			opt.Password = Password
		}
		if Database := util.ToInt(res["Database"]); Database > 0 {
			opt.Database = util.ToInt(Database)
		}
		if PoolSize := util.ToInt(res["PoolSize"]); PoolSize > 0 {
			opt.PoolSize = util.ToInt(PoolSize)
		}
		opt.Type = util.ToBool(res["Type"])
		if Hosts := util.ToStringSlice(res["Hosts"]); len(Hosts) > 0 {
			opt.Hosts = Hosts
		}
		opt.DisableMetric = util.ToBool(res["DisableMetric"])
		opt.DisableTrace = util.ToBool(res["DisableTrace"])
		conn := redis.New(opt)
		links["redis."+node] = conn
	}
	proxyConfigs := initial.Config.Get("proxyredis")
	proxyRes := proxyConfigs.([]map[string]interface{})
	for _, val := range proxyRes {
		proxyPool := proxy.NewProxyRedis()
		if node := util.ToStringSlice(val["Node"]); len(node) > 0 {
			for _, v := range node {
				proxyPool.Store(links[v])
			}
		}
		if Name := util.ToString(val["Name"]); Name != "" {
			initial.Store.StoreRedis(Name, proxyPool)
		}
	}
	return initial
}
