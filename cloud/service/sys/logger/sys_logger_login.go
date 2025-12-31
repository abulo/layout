package logger

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// sys_logger_login 登录日志

// SysLoggerLoginDao 数据转换
func SysLoggerLoginDao(item *SysLoggerLoginObject) *dao.SysLoggerLogin {
	daoItem := &dao.SysLoggerLogin{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = null.StringFrom(item.GetName()) // 姓名
	}
	if item != nil && item.Username != nil {
		daoItem.Username = item.Username // 用户名
	}
	if item != nil && item.UserId != nil {
		daoItem.UserId = item.UserId // 用户编号
	}
	if item != nil && item.Ua != nil {
		daoItem.Ua = null.StringFrom(item.GetUa()) // UA
	}
	if item != nil && item.LoginTime != nil {
		daoItem.LoginTime = null.TimeStampFrom(util.GrpcTime(item.LoginTime)) // 登录时间
	}
	if item != nil && item.Channel != nil {
		daoItem.Channel = null.StringFrom(item.GetChannel()) // 渠道
	}
	if item != nil && item.Ip != nil {
		daoItem.Ip = null.StringFrom(item.GetIp()) // IP
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 删除:0否/1是
	}
	if item != nil && item.TenantId != nil {
		daoItem.TenantId = item.TenantId // 租户
	}
	if item != nil && item.Creator != nil {
		daoItem.Creator = null.StringFrom(item.GetCreator()) // 创建人
	}
	if item != nil && item.CreateTime != nil {
		daoItem.CreateTime = null.TimeStampFrom(util.GrpcTime(item.CreateTime)) // 创建时间
	}
	if item != nil && item.Updater != nil {
		daoItem.Updater = null.StringFrom(item.GetUpdater()) // 更新人
	}
	if item != nil && item.UpdateTime != nil {
		daoItem.UpdateTime = null.TimeStampFrom(util.GrpcTime(item.UpdateTime)) // 更新时间
	}

	return daoItem
}

// SysLoggerLoginProto 数据绑定
func SysLoggerLoginProto(item dao.SysLoggerLogin) *SysLoggerLoginObject {
	res := &SysLoggerLoginObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name.IsValid() {
		res.Name = item.Name.Ptr()
	}
	if item.Username != nil {
		res.Username = item.Username
	}
	if item.UserId != nil {
		res.UserId = item.UserId
	}
	if item.Ua.IsValid() {
		res.Ua = item.Ua.Ptr()
	}
	if item.LoginTime.IsValid() {
		res.LoginTime = timestamppb.New(*item.LoginTime.Ptr())
	}
	if item.Channel.IsValid() {
		res.Channel = item.Channel.Ptr()
	}
	if item.Ip.IsValid() {
		res.Ip = item.Ip.Ptr()
	}
	if item.Deleted != nil {
		res.Deleted = item.Deleted
	}
	if item.TenantId != nil {
		res.TenantId = item.TenantId
	}
	if item.Creator.IsValid() {
		res.Creator = item.Creator.Ptr()
	}
	if item.CreateTime.IsValid() {
		res.CreateTime = timestamppb.New(*item.CreateTime.Ptr())
	}
	if item.Updater.IsValid() {
		res.Updater = item.Updater.Ptr()
	}
	if item.UpdateTime.IsValid() {
		res.UpdateTime = timestamppb.New(*item.UpdateTime.Ptr())
	}

	return res
}
