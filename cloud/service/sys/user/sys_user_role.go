package user

import (
	"cloud/dao"
)

// sys_user_role 用户角色

// SysUserRoleDao 数据转换
func SysUserRoleDao(item *SysUserRoleObject) *dao.SysUserRole {
	daoItem := &dao.SysUserRole{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.UserId != nil {
		daoItem.UserId = item.UserId // 用户
	}
	if item != nil && item.RoleId != nil {
		daoItem.RoleId = item.RoleId // 角色
	}
	if item != nil && item.TenantId != nil {
		daoItem.TenantId = item.TenantId // 租户
	}

	return daoItem
}

// SysUserRoleProto 数据绑定
func SysUserRoleProto(item dao.SysUserRole) *SysUserRoleObject {
	res := &SysUserRoleObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.UserId != nil {
		res.UserId = item.UserId
	}
	if item.RoleId != nil {
		res.RoleId = item.RoleId
	}
	if item.TenantId != nil {
		res.TenantId = item.TenantId
	}

	return res
}
