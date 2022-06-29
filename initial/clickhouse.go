package initial

import (
	"time"

	"github.com/abulo/ratel/v3/stores/clickhouse"
	"github.com/abulo/ratel/v3/stores/proxy"
	"github.com/abulo/ratel/v3/stores/query"
	"github.com/abulo/ratel/v3/util"
)

// InitClickHouse load clickhouse && returns an clickhouse instance.
func (initial *Initial) InitClickHouse() *Initial {
	configs := initial.Config.Get("clickhouse")
	list := configs.(map[string]interface{})

	links := make(map[string]*query.QueryDb)
	for node, nodeConfig := range list {
		opt := &clickhouse.Config{}
		res := nodeConfig.(map[string]interface{})
		if Username := util.ToString(res["Username"]); Username != "" {
			opt.Username = Username
		}
		if Password := util.ToString(res["Password"]); Password != "" {
			opt.Password = Password
		}
		if Addr := util.ToStringSlice(res["Addr"]); len(Addr) > 0 {
			opt.Addr = Addr
		}
		if Database := util.ToString(res["Database"]); Database != "" {
			opt.Database = Database
		}
		if DialTimeout := util.ToString(res["DialTimeout"]); DialTimeout != "" {
			opt.DialTimeout = DialTimeout
		}
		if OpenStrategy := util.ToString(res["OpenStrategy"]); OpenStrategy != "" {
			opt.OpenStrategy = OpenStrategy
		}
		if Compress := util.ToBool(res["Compress"]); Compress {
			opt.Compress = true
		} else {
			opt.Compress = false
		}
		if MaxExecutionTime := util.ToString(res["MaxExecutionTime"]); MaxExecutionTime != "" {
			opt.MaxExecutionTime = MaxExecutionTime
		}

		opt.DisableDebug = util.ToBool(res["DisableDebug"])
		// # MaxOpenConns 连接池最多同时打开的连接数
		// MaxOpenConns = 128
		// # MaxIdleConns 连接池里最大空闲连接数。必须要比maxOpenConns小
		// MaxIdleConns = 32
		// # MaxLifetime 连接池里面的连接最大存活时长(分钟)
		// MaxLifetime = 10
		// # MaxIdleTime 连接池里面的连接最大空闲时长(分钟)
		// MaxIdleTime = 5

		if MaxLifetime := util.ToInt(res["MaxLifetime"]); MaxLifetime > 0 {
			opt.MaxLifetime = time.Duration(MaxLifetime) * time.Minute
		}
		if MaxIdleTime := util.ToInt(res["MaxIdleTime"]); MaxIdleTime > 0 {
			opt.MaxIdleTime = time.Duration(MaxIdleTime) * time.Minute
		}
		if MaxIdleConns := util.ToInt(res["MaxIdleConns"]); MaxIdleConns > 0 {
			opt.MaxIdleConns = util.ToInt(MaxIdleConns)
		}
		if MaxOpenConns := util.ToInt(res["MaxOpenConns"]); MaxOpenConns > 0 {
			opt.MaxOpenConns = util.ToInt(MaxOpenConns)
		}
		opt.DriverName = "clickhouse"
		opt.DisableDebug = util.ToBool(res["DisableDebug"])
		opt.DisableMetric = util.ToBool(res["DisableMetric"])
		opt.DisableTrace = util.ToBool(res["DisableTrace"])
		conn := clickhouse.NewClient(opt)
		links["clickhouse."+node] = conn
	}

	proxyConfigs := initial.Config.Get("proxyclickhouse")
	proxyRes := proxyConfigs.([]map[string]interface{})
	for _, val := range proxyRes {
		proxyPool := proxy.NewProxySQL()
		if node := util.ToString(val["Node"]); node != "" {
			proxyPool.SetWrite(links[node])
		}
		if Name := util.ToString(val["Name"]); Name != "" {
			initial.Store.StoreSQL(Name, proxyPool)
		}
	}
	return initial
}
