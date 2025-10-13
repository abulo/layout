package dao

import "github.com/abulo/ratel/v3/stores/null"

// SysPost 职位 sys_post
type SysPost struct {
	Id         *int64        `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	Name       *string       `gorm:"column:name" json:"name"`                      //varchar 名称
	Sort       *int32        `gorm:"column:sort" json:"sort"`                      //int 排序
	Status     *int32        `gorm:"column:status" json:"status"`                  //tinyint 状态:0正常/1停用
	Deleted    *int32        `gorm:"column:deleted" json:"deleted"`                //tinyint 删除:0否/1是
	TenantId   *int64        `gorm:"column:tenant_id" json:"tenantId"`             //bigint 租户
	Creator    null.String   `gorm:"column:creator" json:"creator"`                //varchar 创建人
	CreateTime null.DateTime `gorm:"column:create_time" json:"createTime"`         //datetime 创建时间
	Updater    null.String   `gorm:"column:updater" json:"updater"`                //varchar 更新人
	UpdateTime null.DateTime `gorm:"column:update_time" json:"updateTime"`         //datetime 更新时间
}

func (SysPost) TableName() string {
	return "sys_post"
}
