package logger

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// sys_logger_dev 开发日志

// SysLoggerDevDao 数据转换
func SysLoggerDevDao(item *SysLoggerDevObject) *dao.SysLoggerDev {
	daoItem := &dao.SysLoggerDev{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.Host != nil {
		daoItem.Host = null.StringFrom(item.GetHost()) // 服务名
	}
	if item != nil && item.Timestamp != nil {
		daoItem.Timestamp = null.TimeStampFrom(util.GrpcTime(item.Timestamp)) // 时间
	}
	if item != nil && item.File != nil {
		daoItem.File = null.StringFrom(item.GetFile()) //
	}
	if item != nil && item.Func != nil {
		daoItem.Func = null.StringFrom(item.GetFunc()) // 方法名
	}
	if item != nil && item.Message != nil {
		daoItem.Message = null.StringFrom(item.GetMessage()) // 消息
	}
	if item != nil && item.Level != nil {
		daoItem.Level = null.StringFrom(item.GetLevel()) // 等级
	}
	if item != nil && item.Data != nil {
		daoItem.Data = null.JSONFrom(item.GetData()) // 数据
	}

	return daoItem
}

// SysLoggerDevProto 数据绑定
func SysLoggerDevProto(item dao.SysLoggerDev) *SysLoggerDevObject {
	res := &SysLoggerDevObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Host.IsValid() {
		res.Host = item.Host.Ptr()
	}
	if item.Timestamp.IsValid() {
		res.Timestamp = timestamppb.New(*item.Timestamp.Ptr())
	}
	if item.File.IsValid() {
		res.File = item.File.Ptr()
	}
	if item.Func.IsValid() {
		res.Func = item.Func.Ptr()
	}
	if item.Message.IsValid() {
		res.Message = item.Message.Ptr()
	}
	if item.Level.IsValid() {
		res.Level = item.Level.Ptr()
	}
	if item.Data.IsValid() {
		res.Data = *item.Data.Ptr()
	}

	return res
}
