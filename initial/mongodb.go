package initial

import (
	"time"

	"github.com/abulo/ratel/v3/stores/mongodb"
	"github.com/abulo/ratel/v3/stores/proxy"
	"github.com/abulo/ratel/v3/util"
)

// InitMongoDB load mongodb && returns an mongodb instance.
func (initial *Initial) InitMongoDB() *Initial {
	configs := initial.Config.Get("mongodb")
	list := configs.(map[string]interface{})
	links := make(map[string]*mongodb.MongoDB)
	for node, nodeConfig := range list {
		opt := &mongodb.Config{}
		res := nodeConfig.(map[string]interface{})
		if URI := util.ToString(res["URI"]); URI != "" {
			opt.URI = URI
		}
		if MaxConnIdleTime := util.ToInt64(res["MaxConnIdleTime"]); MaxConnIdleTime > 0 {
			opt.MaxConnIdleTime = util.ToDuration(MaxConnIdleTime) * time.Minute
		}
		if MaxPoolSize := util.ToInt64(res["MaxPoolSize"]); MaxPoolSize > 0 {
			opt.MaxPoolSize = util.ToUint64(MaxPoolSize)
		}
		if MinPoolSize := util.ToInt64(res["MinPoolSize"]); MinPoolSize > 0 {
			opt.MinPoolSize = util.ToUint64(MinPoolSize)
		}

		opt.DisableMetric = util.ToBool(res["DisableMetric"])
		opt.DisableTrace = util.ToBool(res["DisableTrace"])
		conn := mongodb.NewClient(opt)
		links["mongodb."+node] = conn
	}
	proxyConfigs := initial.Config.Get("proxymongodb")
	proxyRes := proxyConfigs.([]map[string]interface{})
	for _, val := range proxyRes {
		proxyPool := proxy.NewProxyMongoDB()
		if node := util.ToStringSlice(val["Node"]); len(node) > 0 {
			for _, v := range node {
				proxyPool.Store(links[v])
			}
		}
		if Name := util.ToString(val["Name"]); Name != "" {
			initial.Store.StoreMongoDB(Name, proxyPool)
		}
	}
	return initial
}
