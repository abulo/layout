package dao

import "github.com/abulo/ratel/v3/stores/null"

// SysDept 部门 sys_dept
type SysDept struct {
	Id         *int64        `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	Name       *string       `gorm:"column:name" json:"name"`                      //varchar 名称
	ParentId   *int64        `gorm:"column:parent_id" json:"parentId"`             //bigint 上级
	Sort       *int32        `gorm:"column:sort" json:"sort"`                      //int 排序
	UserId     null.Int64    `gorm:"column:user_id" json:"userId"`                 //bigint 负责人
	Phone      null.String   `gorm:"column:phone" json:"phone"`                    //varchar 联系电话
	Email      null.String   `gorm:"column:email" json:"email"`                    //varchar 邮件
	Status     *int32        `gorm:"column:status" json:"status"`                  //tinyint 状态:0正常/1停用
	Deleted    *int32        `gorm:"column:deleted" json:"deleted"`                //tinyint 删除:0否/1是
	TenantId   *int64        `gorm:"column:tenant_id" json:"tenantId"`             //bigint 租户
	Creator    null.String   `gorm:"column:creator" json:"creator"`                //varchar 创建人
	CreateTime null.TimeStamp `gorm:"column:create_time" json:"createTime"`         //datetime 创建时间
	Updater    null.String   `gorm:"column:updater" json:"updater"`                //varchar 更新人
	UpdateTime null.TimeStamp `gorm:"column:update_time" json:"updateTime"`         //datetime 更新时间
}

func (SysDept) TableName() string {
	return "sys_dept"
}
