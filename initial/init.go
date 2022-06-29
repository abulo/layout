package initial

import (
	"time"

	"github.com/abulo/ratel/v3/config"
	"github.com/abulo/ratel/v3/stores/proxy"
	"github.com/abulo/ratel/v3/util"
	"github.com/sarulabs/di"
)

type Initial struct {
	Path       string           // 应用程序执行路径
	Config     *config.Config   // 配置文件
	Store      *proxy.ProxyPool // 数据库链接
	LaunchTime time.Time        // 时间设置
	Container  di.Container     // 依赖注入
}

//系统
var Core *Initial

// Default returns an Initial instance.
func Default() *Initial {
	engine := New()
	return engine
}

func New() *Initial {
	Core = &Initial{
		Store: proxy.NewProxyPool(),
	}
	Core.InitPath(util.GetAppRootPath())
	Core.InitLaunchTime(util.Now())
	return Core
}
