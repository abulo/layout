package dao

import "github.com/abulo/ratel/v3/stores/null"

// SysDict 字典 sys_dict
type SysDict struct {
	Id         *int64        `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	Sort       *int32        `gorm:"column:sort" json:"sort"`                      //int 排序
	Label      *string       `gorm:"column:label" json:"label"`                    //varchar 标签
	Value      *string       `gorm:"column:value" json:"value"`                    //varchar 键值
	DictTypeId *int64        `gorm:"column:dict_type_id" json:"dictTypeId"`        //bigint 字典类型
	ColorType  null.String   `gorm:"column:color_type" json:"colorType"`           //varchar 颜色类型
	CssClass   null.String   `gorm:"column:css_class" json:"cssClass"`             //varchar CSS样式
	Status     *int32        `gorm:"column:status" json:"status"`                  //tinyint 状态:0正常/1停用
	Remark     null.String   `gorm:"column:remark" json:"remark"`                  //varchar 备注
	Creator    null.String   `gorm:"column:creator" json:"creator"`                //varchar 创建人
	CreateTime null.DateTime `gorm:"column:create_time" json:"createTime"`         //datetime 创建时间
	Updater    null.String   `gorm:"column:updater" json:"updater"`                //varchar 更新人
	UpdateTime null.DateTime `gorm:"column:update_time" json:"updateTime"`         //datetime 更新时间
}

func (SysDict) TableName() string {
	return "sys_dict"
}
