package user

import (
	"cloud/dao"
)

// sys_user_post 用户职位

// SysUserPostDao 数据转换
func SysUserPostDao(item *SysUserPostObject) *dao.SysUserPost {
	daoItem := &dao.SysUserPost{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.UserId != nil {
		daoItem.UserId = item.UserId // 用户
	}
	if item != nil && item.PostId != nil {
		daoItem.PostId = item.PostId // 部门
	}
	if item != nil && item.TenantId != nil {
		daoItem.TenantId = item.TenantId // 租户
	}

	return daoItem
}

// SysUserPostProto 数据绑定
func SysUserPostProto(item dao.SysUserPost) *SysUserPostObject {
	res := &SysUserPostObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.UserId != nil {
		res.UserId = item.UserId
	}
	if item.PostId != nil {
		res.PostId = item.PostId
	}
	if item.TenantId != nil {
		res.TenantId = item.TenantId
	}

	return res
}
