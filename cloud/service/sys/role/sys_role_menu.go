package role

import (
	"cloud/dao"
)

// sys_role_menu 角色菜单

// SysRoleMenuDao 数据转换
func SysRoleMenuDao(item *SysRoleMenuObject) *dao.SysRoleMenu {
	daoItem := &dao.SysRoleMenu{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.RoleId != nil {
		daoItem.RoleId = item.RoleId // 角色
	}
	if item != nil && item.MenuId != nil {
		daoItem.MenuId = item.MenuId // 菜单
	}
	if item != nil && item.TenantId != nil {
		daoItem.TenantId = item.TenantId // 租户
	}

	return daoItem
}

// SysRoleMenuProto 数据绑定
func SysRoleMenuProto(item dao.SysRoleMenu) *SysRoleMenuObject {
	res := &SysRoleMenuObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.RoleId != nil {
		res.RoleId = item.RoleId
	}
	if item.MenuId != nil {
		res.MenuId = item.MenuId
	}
	if item.TenantId != nil {
		res.TenantId = item.TenantId
	}

	return res
}
