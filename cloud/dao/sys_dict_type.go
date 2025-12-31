package dao

import "github.com/abulo/ratel/v3/stores/null"

// SysDictType 字典类型 sys_dict_type
type SysDictType struct {
	Id         *int64        `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	Name       *string       `gorm:"column:name" json:"name"`                      //varchar 字典名称
	Type       *string       `gorm:"column:type" json:"type"`                      //varchar 字典类型
	Status     *int32        `gorm:"column:status" json:"status"`                  //tinyint 状态:0正常/1停用
	Remark     null.String   `gorm:"column:remark" json:"remark"`                  //varchar 备注
	Creator    null.String   `gorm:"column:creator" json:"creator"`                //varchar 创建人
	CreateTime null.TimeStamp `gorm:"column:create_time" json:"createTime"`         //datetime 创建时间
	Updater    null.String   `gorm:"column:updater" json:"updater"`                //varchar 更新人
	UpdateTime null.TimeStamp `gorm:"column:update_time" json:"updateTime"`         //datetime 更新时间
}

func (SysDictType) TableName() string {
	return "sys_dict_type"
}
