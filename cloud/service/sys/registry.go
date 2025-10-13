package sys

import (
	"cloud/service/sys/dept"
	"cloud/service/sys/menu"
	"cloud/service/sys/post"
	"cloud/service/sys/role"
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
	// 角色->sys_role
	role.RegisterSysRoleServiceServer(server.Server, &role.SrvSysRoleServiceServer{
		Server: server,
	})
	// 职位->sys_post
	post.RegisterSysPostServiceServer(server.Server, &post.SrvSysPostServiceServer{
		Server: server,
	})
	// 部门->sys_dept
	dept.RegisterSysDeptServiceServer(server.Server, &dept.SrvSysDeptServiceServer{
		Server: server,
	})
	// 用户部门->sys_user_dept
	user.RegisterSysUserDeptServiceServer(server.Server, &user.SrvSysUserDeptServiceServer{
		Server: server,
	})
	// 角色菜单->sys_role_menu
	role.RegisterSysRoleMenuServiceServer(server.Server, &role.SrvSysRoleMenuServiceServer{
		Server: server,
	})
	// 用户职位->sys_user_post
	user.RegisterSysUserPostServiceServer(server.Server, &user.SrvSysUserPostServiceServer{
		Server: server,
	})
}
