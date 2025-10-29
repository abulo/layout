package logger

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// sys_logger_operate 操作日志

// SysLoggerOperateDao 数据转换
func SysLoggerOperateDao(item *SysLoggerOperateObject) *dao.SysLoggerOperate {
	daoItem := &dao.SysLoggerOperate{}

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
	if item != nil && item.Module != nil {
		daoItem.Module = null.StringFrom(item.GetModule()) // 模块名称
	}
	if item != nil && item.Method != nil {
		daoItem.Method = null.StringFrom(item.GetMethod()) // 请求方法
	}
	if item != nil && item.Url != nil {
		daoItem.Url = null.StringFrom(item.GetUrl()) // 请求地址
	}
	if item != nil && item.Ip != nil {
		daoItem.Ip = null.StringFrom(item.GetIp()) // IP
	}
	if item != nil && item.Ua != nil {
		daoItem.Ua = null.StringFrom(item.GetUa()) // UA
	}
	if item != nil && item.GoMethod != nil {
		daoItem.GoMethod = null.StringFrom(item.GetGoMethod()) // 方法名
	}
	if item != nil && item.GoMethodArgs != nil {
		daoItem.GoMethodArgs = null.JSONFrom(item.GetGoMethodArgs()) // 方法参数
	}
	if item != nil && item.StartTime != nil {
		daoItem.StartTime = null.DateTimeFrom(util.GrpcTime(item.StartTime)) // 开始时间
	}
	if item != nil && item.Duration != nil {
		daoItem.Duration = null.Int32From(item.GetDuration()) // 执行时长
	}
	if item != nil && item.Channel != nil {
		daoItem.Channel = null.StringFrom(item.GetChannel()) // 渠道
	}
	if item != nil && item.Result != nil {
		daoItem.Result = null.Int32From(item.GetResult()) // 结果:0 成功/1 失败
	}
	if item != nil && item.ResultMsg != nil {
		daoItem.ResultMsg = null.StringFrom(item.GetResultMsg()) // 结果信息
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
		daoItem.CreateTime = null.DateTimeFrom(util.GrpcTime(item.CreateTime)) // 创建时间
	}
	if item != nil && item.Updater != nil {
		daoItem.Updater = null.StringFrom(item.GetUpdater()) // 更新人
	}
	if item != nil && item.UpdateTime != nil {
		daoItem.UpdateTime = null.DateTimeFrom(util.GrpcTime(item.UpdateTime)) // 更新时间
	}

	return daoItem
}

// SysLoggerOperateProto 数据绑定
func SysLoggerOperateProto(item dao.SysLoggerOperate) *SysLoggerOperateObject {
	res := &SysLoggerOperateObject{}
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
	if item.Module.IsValid() {
		res.Module = item.Module.Ptr()
	}
	if item.Method.IsValid() {
		res.Method = item.Method.Ptr()
	}
	if item.Url.IsValid() {
		res.Url = item.Url.Ptr()
	}
	if item.Ip.IsValid() {
		res.Ip = item.Ip.Ptr()
	}
	if item.Ua.IsValid() {
		res.Ua = item.Ua.Ptr()
	}
	if item.GoMethod.IsValid() {
		res.GoMethod = item.GoMethod.Ptr()
	}
	if item.GoMethodArgs.IsValid() {
		res.GoMethodArgs = *item.GoMethodArgs.Ptr()
	}
	if item.StartTime.IsValid() {
		res.StartTime = timestamppb.New(*item.StartTime.Ptr())
	}
	if item.Duration.IsValid() {
		res.Duration = item.Duration.Ptr()
	}
	if item.Channel.IsValid() {
		res.Channel = item.Channel.Ptr()
	}
	if item.Result.IsValid() {
		res.Result = item.Result.Ptr()
	}
	if item.ResultMsg.IsValid() {
		res.ResultMsg = item.ResultMsg.Ptr()
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
