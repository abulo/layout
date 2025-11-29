package routes

import (
	"cloud/api/sys/user"
	"cloud/api/verify"

	"github.com/abulo/ratel/v3/server/xhertz"
)

func GlobalInitRoute(handle *xhertz.Server) {
	route := handle.Group("/v1")
	{
		// 验证码生成
		route.GET("/verify/generate", verify.Generate)
		// sys_user->用户->用户登录
		route.POST("/sys/user/login", user.Login)
	}
}
