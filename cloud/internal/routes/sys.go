package routes

import (
	"cloud/api/sys/dept"
	"cloud/api/sys/dict"
	"cloud/api/sys/menu"
	"cloud/api/sys/post"
	"cloud/api/sys/tenant"

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

		// sys_dict_type->字典类型->创建
		route.POST("/sys/dict/type", dict.SysDictTypeCreate)
		// sys_dict_type->字典类型->更新
		route.PUT("/sys/dict/type/:id/update", dict.SysDictTypeUpdate)
		// sys_dict_type->字典类型->删除
		route.DELETE("/sys/dict/type/:id/delete", dict.SysDictTypeDelete)
		// sys_dict_type->字典类型->单条数据信息查看
		route.GET("/sys/dict/type/:id/item", dict.SysDictType)
		// sys_dict_type->字典类型->列表
		route.GET("/sys/dict/type", dict.SysDictTypeList)
		// sys_dict_type->字典类型->列表(精简)
		route.GET("/sys/dict/type/simple", dict.SysDictTypeListSimple)

		// sys_dict->字典->创建
		route.POST("/sys/dict", dict.SysDictCreate)
		// sys_dict->字典->更新
		route.PUT("/sys/dict/:id/update", dict.SysDictUpdate)
		// sys_dict->字典->删除
		route.DELETE("/sys/dict/:id/delete", dict.SysDictDelete)
		// sys_dict->字典->单条数据信息查看
		route.GET("/sys/dict/:id/item", dict.SysDict)
		// sys_dict->字典->列表
		route.GET("/sys/dict", dict.SysDictList)
		// sys_dict->字典->列表(精简)
		route.GET("/sys/dict/simple", dict.SysDictListSimple)
		// sys_dict->字典->列表
		route.GET("/dict", dict.DictList)

		// sys_tenant->租户->创建
		route.POST("/sys/tenant", tenant.SysTenantCreate)
		// sys_tenant->租户->更新
		route.PUT("/sys/tenant/:id/update", tenant.SysTenantUpdate)
		// sys_tenant->租户->删除
		route.DELETE("/sys/tenant/:id/delete", tenant.SysTenantDelete)
		// sys_tenant->租户->单条数据信息查看
		route.GET("/sys/tenant/:id/item", tenant.SysTenant)
		// sys_tenant->租户->恢复
		route.PUT("/sys/tenant/:id/recover", tenant.SysTenantRecover)
		// sys_tenant->租户->清理
		route.DELETE("/sys/tenant/:id/drop", tenant.SysTenantDrop)
		// sys_tenant->租户->列表
		route.GET("/sys/tenant", tenant.SysTenantList)
		// sys_tenant->租户->列表(精简)
		route.GET("/sys/tenant/simple", tenant.SysTenantListSimple)

		// sys_tenant_package->租户套餐->创建
		route.POST("/sys/tenant/package", tenant.SysTenantPackageCreate)
		// sys_tenant_package->租户套餐->更新
		route.PUT("/sys/tenant/package/:id/update", tenant.SysTenantPackageUpdate)
		// sys_tenant_package->租户套餐->删除
		route.DELETE("/sys/tenant/package/:id/delete", tenant.SysTenantPackageDelete)
		// sys_tenant_package->租户套餐->单条数据信息查看
		route.GET("/sys/tenant/package/:id/item", tenant.SysTenantPackage)
		// sys_tenant_package->租户套餐->列表
		route.GET("/sys/tenant/package", tenant.SysTenantPackageList)
		// sys_tenant_package->租户套餐->列表(精简)
		route.GET("/sys/tenant/package/simple", tenant.SysTenantPackageListSimple)

		// sys_dept->部门->创建
		route.POST("/sys/dept", dept.SysDeptCreate)
		// sys_dept->部门->更新
		route.PUT("/sys/dept/:id/update", dept.SysDeptUpdate)
		// sys_dept->部门->删除
		route.DELETE("/sys/dept/:id/delete", dept.SysDeptDelete)
		// sys_dept->部门->单条数据信息查看
		route.GET("/sys/dept/:id/item", dept.SysDept)
		// sys_dept->部门->恢复
		route.PUT("/sys/dept/:id/recover", dept.SysDeptRecover)
		// sys_dept->部门->清理
		route.DELETE("/sys/dept/:id/drop", dept.SysDeptDrop)
		// sys_dept->部门->列表
		route.GET("/sys/dept", dept.SysDeptList)
		// sys_dept->部门->列表(精简)
		route.GET("/sys/dept/simple", dept.SysDeptListSimple)

	}
}
