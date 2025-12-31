package menu

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// sys_menu 菜单

// SysMenuDao 数据转换
func SysMenuDao(item *SysMenuObject) *dao.SysMenu {
	daoItem := &dao.SysMenu{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 名称
	}
	if item != nil && item.Code != nil {
		daoItem.Code = null.StringFrom(item.GetCode()) // 编码
	}
	if item != nil && item.Type != nil {
		daoItem.Type = item.Type // 类型:0 目录/1 菜单/2 按钮
	}
	if item != nil && item.Sort != nil {
		daoItem.Sort = item.Sort // 排序
	}
	if item != nil && item.ParentId != nil {
		daoItem.ParentId = item.ParentId // 上级
	}
	if item != nil && item.Path != nil {
		daoItem.Path = null.StringFrom(item.GetPath()) // 地址
	}
	if item != nil && item.Icon != nil {
		daoItem.Icon = null.StringFrom(item.GetIcon()) // 图标
	}
	if item != nil && item.Component != nil {
		daoItem.Component = null.StringFrom(item.GetComponent()) // 组件路径
	}
	if item != nil && item.ComponentName != nil {
		daoItem.ComponentName = null.StringFrom(item.GetComponentName()) // 组件名称
	}
	if item != nil && item.Hide != nil {
		daoItem.Hide = item.Hide // 隐藏:0 否/1 是
	}
	if item != nil && item.Link != nil {
		daoItem.Link = null.StringFrom(item.GetLink()) // 外部地址
	}
	if item != nil && item.Cache != nil {
		daoItem.Cache = item.Cache // 缓存:0否/1 是
	}
	if item != nil && item.Remark != nil {
		daoItem.Remark = null.StringFrom(item.GetRemark()) // 备注
	}
	if item != nil && item.Active != nil {
		daoItem.Active = null.StringFrom(item.GetActive()) // 激活地址
	}
	if item != nil && item.FullScreen != nil {
		daoItem.FullScreen = item.FullScreen // 全屏:1 开/0 关
	}
	if item != nil && item.Redirect != nil {
		daoItem.Redirect = null.StringFrom(item.GetRedirect()) // 重定向
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 状态:0正常/1停用
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

// SysMenuProto 数据绑定
func SysMenuProto(item dao.SysMenu) *SysMenuObject {
	res := &SysMenuObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.Code.IsValid() {
		res.Code = item.Code.Ptr()
	}
	if item.Type != nil {
		res.Type = item.Type
	}
	if item.Sort != nil {
		res.Sort = item.Sort
	}
	if item.ParentId != nil {
		res.ParentId = item.ParentId
	}
	if item.Path.IsValid() {
		res.Path = item.Path.Ptr()
	}
	if item.Icon.IsValid() {
		res.Icon = item.Icon.Ptr()
	}
	if item.Component.IsValid() {
		res.Component = item.Component.Ptr()
	}
	if item.ComponentName.IsValid() {
		res.ComponentName = item.ComponentName.Ptr()
	}
	if item.Hide != nil {
		res.Hide = item.Hide
	}
	if item.Link.IsValid() {
		res.Link = item.Link.Ptr()
	}
	if item.Cache != nil {
		res.Cache = item.Cache
	}
	if item.Remark.IsValid() {
		res.Remark = item.Remark.Ptr()
	}
	if item.Active.IsValid() {
		res.Active = item.Active.Ptr()
	}
	if item.FullScreen != nil {
		res.FullScreen = item.FullScreen
	}
	if item.Redirect.IsValid() {
		res.Redirect = item.Redirect.Ptr()
	}
	if item.Status != nil {
		res.Status = item.Status
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
