package dao

import "github.com/abulo/ratel/v3/stores/null"

// SysLoggerLogin 登录日志 sys_logger_login
type SysLoggerLogin struct {
	Id         *int64        `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	Name       null.String   `gorm:"column:name" json:"name"`                      //varchar 姓名
	Username   *string       `gorm:"column:username" json:"username"`              //varchar 用户名
	UserId     *int64        `gorm:"column:user_id" json:"userId"`                 //bigint 用户编号
	Ua         null.String   `gorm:"column:ua" json:"ua"`                          //varchar UA
	LoginTime  null.DateTime `gorm:"column:login_time" json:"loginTime"`           //datetime 登录时间
	Channel    null.String   `gorm:"column:channel" json:"channel"`                //varchar 渠道
	Ip         null.String   `gorm:"column:ip" json:"ip"`                          //varchar IP
	Deleted    *int32        `gorm:"column:deleted" json:"deleted"`                //tinyint 删除:0否/1是
	TenantId   *int64        `gorm:"column:tenant_id" json:"tenantId"`             //bigint 租户
	Creator    null.String   `gorm:"column:creator" json:"creator"`                //varchar 创建人
	CreateTime null.DateTime `gorm:"column:create_time" json:"createTime"`         //datetime 创建时间
	Updater    null.String   `gorm:"column:updater" json:"updater"`                //varchar 更新人
	UpdateTime null.DateTime `gorm:"column:update_time" json:"updateTime"`         //datetime 更新时间
}

func (SysLoggerLogin) TableName() string {
	return "sys_logger_login"
}
