package dao

// SysUserDept 用户部门 sys_user_dept
type SysUserDept struct {
	Id       *int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	UserId   *int64 `gorm:"column:user_id" json:"userId"`                 //bigint 用户
	DeptId   *int64 `gorm:"column:dept_id" json:"deptId"`                 //bigint 部门
	TenantId *int64 `gorm:"column:tenant_id" json:"tenantId"`             //bigint 租户
}

func (SysUserDept) TableName() string {
	return "sys_user_dept"
}
