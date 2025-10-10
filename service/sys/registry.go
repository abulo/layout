package sys

import (
	"cloud/service/sys/menu"
	"cloud/service/sys/user"

	"github.com/abulo/ratel/v3/server/xgrpc"
)

// Registry 注册服务
func Registry(server *xgrpc.Server) {
	// 用户->sys_user
	user.RegisterSysUserServiceServer(server.Server, &user.SrvSysUserServiceServer{
		Server: server,
	})
	// 菜单->sys_menu
	menu.RegisterSysMenuServiceServer(server.Server, &menu.SrvSysMenuServiceServer{
		Server: server,
	})
}
