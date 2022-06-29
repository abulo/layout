package initial

import (
	"github.com/abulo/ratel/v3/stores/elasticsearch"
	"github.com/abulo/ratel/v3/stores/proxy"
	"github.com/abulo/ratel/v3/util"
)

// InitElasticSearch load elasticsearch && returns an elasticsearch instance.
func (initial *Initial) InitElasticSearch() *Initial {
	configs := initial.Config.Get("elasticsearch")
	list := configs.(map[string]interface{})
	links := make(map[string]*elasticsearch.Client)
	for node, nodeConfig := range list {
		opts := &elasticsearch.Config{}
		res := nodeConfig.(map[string]interface{})
		opts.URL = util.ToStringSlice(res["URL"])
		opts.Username = util.ToString(res["Username"])
		opts.Password = util.ToString(res["Password"])
		opts.DisableMetric = util.ToBool(res["DisableMetric"])
		opts.DisableTrace = util.ToBool(res["DisableTrace"])
		conn := elasticsearch.NewClient(opts)
		links["elasticsearch."+node] = conn
	}
	proxyConfigs := initial.Config.Get("proxyelasticsearch")
	proxyRes := proxyConfigs.([]map[string]interface{})
	for _, val := range proxyRes {
		proxyPool := proxy.NewProxyElasticSearch()
		if node := util.ToStringSlice(val["Node"]); len(node) > 0 {
			for _, v := range node {
				proxyPool.Store(links[v])
			}
		}
		if Name := util.ToString(val["Name"]); Name != "" {
			initial.Store.StoreElasticSearch(Name, proxyPool)
		}
	}
	return initial
}
