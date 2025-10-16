package user

import (
	"cloud/dao"
)

// sys_user_tenant 租户用户

// SysUserTenantDao 数据转换
func SysUserTenantDao(item *SysUserTenantObject) *dao.SysUserTenant {
	daoItem := &dao.SysUserTenant{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.UserId != nil {
		daoItem.UserId = item.UserId // 用户
	}
	if item != nil && item.TenantId != nil {
		daoItem.TenantId = item.TenantId // 租户
	}

	return daoItem
}

// SysUserTenantProto 数据绑定
func SysUserTenantProto(item dao.SysUserTenant) *SysUserTenantObject {
	res := &SysUserTenantObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.UserId != nil {
		res.UserId = item.UserId
	}
	if item.TenantId != nil {
		res.TenantId = item.TenantId
	}

	return res
}
