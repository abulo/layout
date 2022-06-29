package services

import (
	"layout/internal/backstage"
	"layout/routes"

	"github.com/sarulabs/di"
)

//Services 服务
var BackServices = []di.Def{
	//后台管理会话
	{
		Name:  "internal.backstage",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &backstage.Backstage{}, nil
		},
	},
	{
		Name:  "routes",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &routes.Routes{}, nil
		},
	},
}
