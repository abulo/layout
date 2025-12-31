package tenant

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// sys_tenant 租户

// SysTenantDao 数据转换
func SysTenantDao(item *SysTenantObject) *dao.SysTenant {
	daoItem := &dao.SysTenant{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 名称
	}
	if item != nil && item.UserId != nil {
		daoItem.UserId = null.Int64From(item.GetUserId()) // 用户
	}
	if item != nil && item.ContactName != nil {
		daoItem.ContactName = null.StringFrom(item.GetContactName()) // 联系人
	}
	if item != nil && item.ContactMobile != nil {
		daoItem.ContactMobile = null.StringFrom(item.GetContactMobile()) // 联系电话
	}
	if item != nil && item.ExpireDate != nil {
		daoItem.ExpireDate = null.DateFrom(util.GrpcTime(item.ExpireDate)) // 过期时间
	}
	if item != nil && item.AccountTotal != nil {
		daoItem.AccountTotal = item.AccountTotal // 账号数量
	}
	if item != nil && item.PackageId != nil {
		daoItem.PackageId = item.PackageId // 套餐编号
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 状态:0正常/1停用
	}
	if item != nil && item.Deleted != nil {
		daoItem.Deleted = item.Deleted // 删除:0否/1是
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
	if item != nil && item.Username != nil {
		daoItem.Username = item.Username
	}
	if item != nil && item.Password != nil {
		daoItem.Password = item.Password
	}

	return daoItem
}

// SysTenantProto 数据绑定
func SysTenantProto(item dao.SysTenant) *SysTenantObject {
	res := &SysTenantObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.UserId.IsValid() {
		res.UserId = item.UserId.Ptr()
	}
	if item.ContactName.IsValid() {
		res.ContactName = item.ContactName.Ptr()
	}
	if item.ContactMobile.IsValid() {
		res.ContactMobile = item.ContactMobile.Ptr()
	}
	if item.ExpireDate.IsValid() {
		res.ExpireDate = timestamppb.New(*item.ExpireDate.Ptr())
	}
	if item.AccountTotal != nil {
		res.AccountTotal = item.AccountTotal
	}
	if item.PackageId != nil {
		res.PackageId = item.PackageId
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Deleted != nil {
		res.Deleted = item.Deleted
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
	if item.Username != nil {
		res.Username = item.Username
	}
	if item.Password != nil {
		res.Password = item.Password
	}
	return res
}
