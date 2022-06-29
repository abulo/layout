package initial

import (
	"github.com/abulo/ratel/v3/config"
	"github.com/abulo/ratel/v3/config/toml"
	"github.com/abulo/ratel/v3/util"
	"github.com/fsnotify/fsnotify"
)

func (initial *Initial) GetEnvironment(dir, key string) string {
	envConfig := config.NewWithOptions("go-ratel-evn", config.Readonly, config.EnableCache)
	driver := toml.Driver
	envConfig.AddDriver(driver)
	envConfig.LoadDir(dir, driver.Name())
	return envConfig.String(key)
}

// InitConfig set app config toml files
func (initial *Initial) InitConfig(dirs ...string) *Initial {
	Config := config.NewWithOptions("go-ratel", config.Readonly, config.EnableCache)
	driver := toml.Driver
	Config.AddDriver(driver)
	for _, dir := range dirs {
		isDir, err := util.IsDir(dir)
		if err != nil {
			panic(dir)
		}
		if !isDir {
			panic(dir + "not a directory")
		}
		//load file
		Config.LoadDir(dir, driver.Name())
	}
	Config.OnConfigChange(func(e fsnotify.Event) {
		Config.LoadFiles(e.Name)
	})
	Config.WatchConfig(driver.Name())
	initial.Config = Config
	return initial
}
