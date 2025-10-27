package routes

import (
	"cloud/api/sys/menu"
	"cloud/api/sys/post"

	"github.com/abulo/ratel/v3/server/xhertz"
)

func SysInitRoute(handle *xhertz.Server) {
	route := handle.Group("/v1")
	{
		// sys_post->职位->创建
		route.POST("/sys/post", post.SysPostCreate)
		// sys_post->职位->更新
		route.PUT("/sys/post/:id/update", post.SysPostUpdate)
		// sys_post->职位->删除
		route.DELETE("/sys/post/:id/delete", post.SysPostDelete)
		// sys_post->职位->单条数据信息查看
		route.GET("/sys/post/:id/item", post.SysPost)
		// sys_post->职位->恢复
		route.PUT("/sys/post/:id/recover", post.SysPostRecover)
		// sys_post->职位->清理
		route.DELETE("/sys/post/:id/drop", post.SysPostDrop)
		// sys_post->职位->列表
		route.GET("/sys/post", post.SysPostList)
		// sys_post->职位->列表(精简)
		route.GET("/sys/post/simple", post.SysPostListSimple)

		// sys_menu->菜单->创建
		route.POST("/sys/menu", menu.SysMenuCreate)
		// sys_menu->菜单->更新
		route.PUT("/sys/menu/:id/update", menu.SysMenuUpdate)
		// sys_menu->菜单->删除
		route.DELETE("/sys/menu/:id/delete", menu.SysMenuDelete)
		// sys_menu->菜单->单条数据信息查看
		route.GET("/sys/menu/:id/item", menu.SysMenu)
		// sys_menu->菜单->列表
		route.GET("/sys/menu", menu.SysMenuList)
		// sys_menu->菜单->列表(精简)
		route.GET("/sys/menu/simple", menu.SysMenuListSimple)
	}
}
