package dao

// SysUserTenant 租户用户 sys_user_tenant
type SysUserTenant struct {
	Id       *int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	UserId   *int64 `gorm:"column:user_id" json:"userId"`                 //bigint 用户
	TenantId *int64 `gorm:"column:tenant_id" json:"tenantId"`             //bigint 租户
}

func (SysUserTenant) TableName() string {
	return "sys_user_tenant"
}
