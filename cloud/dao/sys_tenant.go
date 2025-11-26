package dao

import "github.com/abulo/ratel/v3/stores/null"

// SysTenant 租户 sys_tenant
type SysTenant struct {
	Id            *int64        `gorm:"primaryKey;autoIncrement;column:id" json:"id"`       //bigint 编号,PRI
	Name          *string       `gorm:"column:name" json:"name"`                            //varchar 名称
	UserId        null.Int64    `gorm:"column:user_id" json:"userId"`                       //bigint 用户
	ContactName   null.String   `gorm:"column:contact_name" json:"contactName"`             //varchar 联系人
	ContactMobile null.String   `gorm:"column:contact_mobile" json:"contactMobile"`         //varchar 联系电话
	ExpireDate    null.Date     `gorm:"column:expire_date" json:"expireDate"`               //date 过期时间
	AccountTotal  *int64        `gorm:"column:account_total" json:"accountTotal"`           //bigint 账号数量
	PackageId     *int64        `gorm:"column:package_id" json:"packageId"`                 //bigint 套餐编号
	Status        *int32        `gorm:"column:status" json:"status"`                        //tinyint 状态:0正常/1停用
	Deleted       *int32        `gorm:"column:deleted" json:"deleted"`                      //tinyint 删除:0否/1是
	Creator       null.String   `gorm:"column:creator" json:"creator"`                      //varchar 创建人
	CreateTime    null.DateTime `gorm:"column:create_time" json:"createTime"`               //datetime 创建时间
	Updater       null.String   `gorm:"column:updater" json:"updater"`                      //varchar 更新人
	UpdateTime    null.DateTime `gorm:"column:update_time" json:"updateTime"`               //datetime 更新时间
	Username      *string       `gorm:"column:username;<-:false" json:"username,omitempty"` //varchar 用户名称
	Password      *string       `gorm:"column:password;<-:false" json:"password,omitempty"` //varchar 用户密码
}

func (SysTenant) TableName() string {
	return "sys_tenant"
}
