package dao

import "github.com/abulo/ratel/v3/stores/null"

// SysLoggerOperate 操作日志 sys_logger_operate
type SysLoggerOperate struct {
	Id           *int64        `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	Name         null.String   `gorm:"column:name" json:"name"`                      //varchar 姓名
	Username     *string       `gorm:"column:username" json:"username"`              //varchar 用户名
	UserId       *int64        `gorm:"column:user_id" json:"userId"`                 //bigint 用户编号
	Module       null.String   `gorm:"column:module" json:"module"`                  //varchar 模块名称
	Method       null.String   `gorm:"column:method" json:"method"`                  //varchar 请求方法
	Url          null.String   `gorm:"column:url" json:"url"`                        //varchar 请求地址
	Ip           null.String   `gorm:"column:ip" json:"ip"`                          //varchar IP
	Ua           null.String   `gorm:"column:ua" json:"ua"`                          //varchar UA
	GoMethod     null.String   `gorm:"column:go_method" json:"goMethod"`             //varchar 方法名
	GoMethodArgs null.JSON     `gorm:"column:go_method_args" json:"goMethodArgs"`    //json 方法参数
	StartTime    null.TimeStamp `gorm:"column:start_time" json:"startTime"`           //datetime 开始时间
	Duration     null.Int32    `gorm:"column:duration" json:"duration"`              //int 执行时长
	Channel      null.String   `gorm:"column:channel" json:"channel"`                //varchar 渠道
	Result       null.Int32    `gorm:"column:result" json:"result"`                  //tinyint 结果:0 成功/1 失败
	ResultMsg    null.String   `gorm:"column:result_msg" json:"resultMsg"`           //varchar 结果信息
	Deleted      *int32        `gorm:"column:deleted" json:"deleted"`                //tinyint 删除:0否/1是
	TenantId     *int64        `gorm:"column:tenant_id" json:"tenantId"`             //bigint 租户
	Creator      null.String   `gorm:"column:creator" json:"creator"`                //varchar 创建人
	CreateTime   null.TimeStamp `gorm:"column:create_time" json:"createTime"`         //datetime 创建时间
	Updater      null.String   `gorm:"column:updater" json:"updater"`                //varchar 更新人
	UpdateTime   null.TimeStamp `gorm:"column:update_time" json:"updateTime"`         //datetime 更新时间
}

func (SysLoggerOperate) TableName() string {
	return "sys_logger_operate"
}
