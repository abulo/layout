package dao

import "github.com/abulo/ratel/v3/stores/null"

// SysTenantPackage 租户套餐 sys_tenant_package
type SysTenantPackage struct {
	Id         *int64        `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	Name       *string       `gorm:"column:name" json:"name"`                      //varchar 名称
	ScopeMenu  null.JSON     `gorm:"column:scope_menu" json:"scopeMenu"`           //json 数据范围
	Sort       *int32        `gorm:"column:sort" json:"sort"`                      //int 排序
	Status     *int32        `gorm:"column:status" json:"status"`                  //tinyint 状态:0正常/1停用
	Remark     *string       `gorm:"column:remark" json:"remark"`                  //varchar 备注
	Creator    null.String   `gorm:"column:creator" json:"creator"`                //varchar 创建人
	CreateTime null.TimeStamp `gorm:"column:create_time" json:"createTime"`         //datetime 创建时间
	Updater    null.String   `gorm:"column:updater" json:"updater"`                //varchar 更新人
	UpdateTime null.TimeStamp `gorm:"column:update_time" json:"updateTime"`         //datetime 更新时间
}

func (SysTenantPackage) TableName() string {
	return "sys_tenant_package"
}
