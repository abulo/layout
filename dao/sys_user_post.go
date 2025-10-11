package dao

import "github.com/abulo/ratel/v3/stores/null"

// SysUserPost 用户职位 sys_user_post
type SysUserPost struct {
	Id       *int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	UserId   *int64 `gorm:"column:user_id" json:"userId"`                 //bigint 用户
	PostId   *int64 `gorm:"column:post_id" json:"postId"`                 //bigint 部门
	TenantId *int64 `gorm:"column:tenant_id" json:"tenantId"`             //bigint 租户
}

func (SysUserPost) TableName() string {
	return "sys_user_post"
}
