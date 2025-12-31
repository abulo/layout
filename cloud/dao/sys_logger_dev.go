package dao

import "github.com/abulo/ratel/v3/stores/null"

// SysLoggerDev 开发日志 sys_logger_dev
type SysLoggerDev struct {
	Id        *int64        `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	Host      null.String   `gorm:"column:host" json:"host"`                      //varchar 服务名
	Timestamp null.TimeStamp `gorm:"column:timestamp" json:"timestamp"`            //datetime 时间
	File      null.String   `gorm:"column:file" json:"file"`                      //varchar
	Func      null.String   `gorm:"column:func" json:"func"`                      //varchar 方法名
	Message   null.String   `gorm:"column:message" json:"message"`                //varchar 消息
	Level     null.String   `gorm:"column:level" json:"level"`                    //varchar 等级
	Data      null.JSON     `gorm:"column:data" json:"data"`                      //json 数据
}

func (SysLoggerDev) TableName() string {
	return "sys_logger_dev"
}
