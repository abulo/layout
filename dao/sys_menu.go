package dao

import "github.com/abulo/ratel/v3/stores/null"

// SysMenu 菜单 sys_menu
type SysMenu struct {
	Id            *int64        `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	Name          *string       `gorm:"column:name" json:"name"`                      //varchar 名称
	Code          null.String   `gorm:"column:code" json:"code"`                      //varchar 编码
	Type          *int32        `gorm:"column:type" json:"type"`                      //tinyint 类型:0 目录/1 菜单/2 按钮
	Sort          *int32        `gorm:"column:sort" json:"sort"`                      //int 排序
	ParentId      *int64        `gorm:"column:parent_id" json:"parentId"`             //bigint 上级
	Path          null.String   `gorm:"column:path" json:"path"`                      //varchar 地址
	Icon          null.String   `gorm:"column:icon" json:"icon"`                      //varchar 图标
	Component     null.String   `gorm:"column:component" json:"component"`            //varchar 组件路径
	ComponentName null.String   `gorm:"column:component_name" json:"componentName"`   //varchar 组件名称
	Hide          *int32        `gorm:"column:hide" json:"hide"`                      //tinyint 隐藏:0 否/1 是
	Link          null.String   `gorm:"column:link" json:"link"`                      //varchar 外部地址
	Cache         *int32        `gorm:"column:cache" json:"cache"`                    //tinyint 缓存:0否/1 是
	Affix         *int32        `gorm:"column:affix" json:"affix"`                    //tinyint 固定:0否/1 是
	Active        null.String   `gorm:"column:active" json:"active"`                  //varchar 激活地址
	Full          *int32        `gorm:"column:full" json:"full"`                      //tinyint
	Redirect      null.String   `gorm:"column:redirect" json:"redirect"`              //varchar 重定向
	Status        *int32        `gorm:"column:status" json:"status"`                  //tinyint 状态:0正常/1停用
	Creator       null.String   `gorm:"column:creator" json:"creator"`                //varchar 创建人
	CreateTime    null.DateTime `gorm:"column:create_time" json:"createTime"`         //datetime 创建时间
	Updater       null.String   `gorm:"column:updater" json:"updater"`                //varchar 更新人
	UpdateTime    null.DateTime `gorm:"column:update_time" json:"updateTime"`         //datetime 更新时间
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
