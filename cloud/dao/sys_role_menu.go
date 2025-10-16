package dao

import "github.com/abulo/ratel/v3/stores/null"

// SysRoleMenu 角色菜单 sys_role_menu
type SysRoleMenu struct {
	Id       *int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	RoleId   *int64 `gorm:"column:role_id" json:"roleId"`                 //bigint 角色
	MenuId   *int64 `gorm:"column:menu_id" json:"menuId"`                 //bigint 菜单
	TenantId *int64 `gorm:"column:tenant_id" json:"tenantId"`             //bigint 租户
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
