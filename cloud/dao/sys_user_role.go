package dao

// SysUserRole 用户角色 sys_user_role
type SysUserRole struct {
	Id       *int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	UserId   *int64 `gorm:"column:user_id" json:"userId"`                 //bigint 用户
	RoleId   *int64 `gorm:"column:role_id" json:"roleId"`                 //bigint 角色
	TenantId *int64 `gorm:"column:tenant_id" json:"tenantId"`             //bigint 租户
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
