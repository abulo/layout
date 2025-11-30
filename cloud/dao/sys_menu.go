package dao

import (
	"context"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

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
	Remark        null.String   `gorm:"column:remark" json:"remark"`                  //varchar 备注
	Active        null.String   `gorm:"column:active" json:"active"`                  //varchar 激活地址
	Full          *int32        `gorm:"column:full" json:"full"`                      //tinyint 全屏:1 开/0 关
	Redirect      null.String   `gorm:"column:redirect" json:"redirect"`              //varchar 重定向
	Status        *int32        `gorm:"column:status" json:"status"`                  //tinyint 状态:0正常/1停用
	Creator       null.String   `gorm:"column:creator" json:"creator"`                //varchar 创建人
	CreateTime    null.DateTime `gorm:"column:create_time" json:"createTime"`         //datetime 创建时间
	Updater       null.String   `gorm:"column:updater" json:"updater"`                //varchar 更新人
	UpdateTime    null.DateTime `gorm:"column:update_time" json:"updateTime"`         //datetime 更新时间
}

// SysMenuTreeMetaRes 菜单元信息
type SysMenuTreeMeta struct {
	Icon        string `json:"icon"`                 // 菜单图标
	Title       string `json:"title"`                // 菜单标题
	IsLink      string `json:"isLink,omitempty"`     // 是否外链
	IsHide      bool   `json:"isHide"`               // 是否隐藏
	IsFull      bool   `json:"isFull"`               // 是否全屏
	IsAffix     bool   `json:"isAffix"`              // 是否固定
	IsKeepAlive bool   `json:"isKeepAlive"`          // 是否缓存
	ActiveMenu  string `json:"activeMenu,omitempty"` // 激活菜单
}

// SysMenuTreeRes 菜单数据
type SysMenuTree struct {
	Id        int64            `json:"-"`                   // 菜单ID
	ParentId  int64            `json:"-"`                   // 父菜单ID
	Type      int32            `json:"-"`                   // 菜单类型
	Code      string           `json:"-"`                   // 菜单编码
	Path      string           `json:"path"`                // 路由地址
	Name      string           `json:"name"`                // 路由名称
	Component string           `json:"component,omitempty"` // 组件路径
	Redirect  string           `json:"redirect,omitempty"`  // 重定向地址
	Meta      *SysMenuTreeMeta `json:"meta"`                // 菜单元信息
	Children  []*SysMenuTree   `json:"children,omitempty"`  // 子菜单
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

func (r *SysMenu) AfterDelete(tx *gorm.DB) (err error) {
	id := cast.ToInt64(r.Id)
	var data []SysRoleMenu
	ctx := context.Background()
	result := tx.WithContext(ctx).Model(&SysRoleMenu{}).Where("menu_id = ?", id).Find(&data).Delete(&data)
	return result.Error
}
