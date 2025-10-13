package dao

import "github.com/abulo/ratel/v3/stores/null"

// SysUser 用户 sys_user
type SysUser struct {
	Id         *int64        `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	Name       null.String   `gorm:"column:name" json:"name"`                      //varchar 姓名
	Mobile     null.String   `gorm:"column:mobile" json:"mobile"`                  //varchar 手机号码
	Username   *string       `gorm:"column:username" json:"username"`              //varchar 用户名
	Password   *string       `gorm:"column:password" json:"password"`              //varchar 密码
	Status     *int32        `gorm:"column:status" json:"status"`                  //tinyint 状态:0正常/1停用
	Deleted    *int32        `gorm:"column:deleted" json:"deleted"`                //tinyint 删除:0否/1是
	TenantId   *int64        `gorm:"column:tenant_id" json:"tenantId"`             //bigint 租户
	Creator    null.String   `gorm:"column:creator" json:"creator"`                //varchar 创建人
	CreateTime null.DateTime `gorm:"column:create_time" json:"createTime"`         //datetime 创建时间
	Updater    null.String   `gorm:"column:updater" json:"updater"`                //varchar 更新人
	UpdateTime null.DateTime `gorm:"column:update_time" json:"updateTime"`         //datetime 更新时间
}

func (SysUser) TableName() string {
	return "sys_user"
}
