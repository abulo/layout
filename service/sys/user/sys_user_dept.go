package user

import (
	"cloud/dao"
)

// sys_user_dept 用户部门

// SysUserDeptDao 数据转换
func SysUserDeptDao(item *SysUserDeptObject) *dao.SysUserDept {
	daoItem := &dao.SysUserDept{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.UserId != nil {
		daoItem.UserId = item.UserId // 用户
	}
	if item != nil && item.DeptId != nil {
		daoItem.DeptId = item.DeptId // 部门
	}
	if item != nil && item.TenantId != nil {
		daoItem.TenantId = item.TenantId // 租户
	}

	return daoItem
}

// SysUserDeptProto 数据绑定
func SysUserDeptProto(item dao.SysUserDept) *SysUserDeptObject {
	res := &SysUserDeptObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.UserId != nil {
		res.UserId = item.UserId
	}
	if item.DeptId != nil {
		res.DeptId = item.DeptId
	}
	if item.TenantId != nil {
		res.TenantId = item.TenantId
	}

	return res
}
