package dict

import (
	"cloud/dao"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// dict 字典

// DictDao 数据转换
func DictDao(item *DictObject) *dao.Dict {
	daoItem := &dao.Dict{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.Sort != nil {
		daoItem.Sort = item.Sort // 排序
	}
	if item != nil && item.Label != nil {
		daoItem.Label = item.Label // 标签
	}
	if item != nil && item.Value != nil {
		daoItem.Value = item.Value // 键值
	}
	if item != nil && item.DictType != nil {
		daoItem.DictType = item.DictType // 字典类型
	}
	if item != nil && item.ColorType != nil {
		daoItem.ColorType = null.StringFrom(item.GetColorType()) // 颜色类型
	}
	if item != nil && item.CssClass != nil {
		daoItem.CssClass = null.StringFrom(item.GetCssClass()) // CSS样式
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 状态:0正常/1停用
	}
	if item != nil && item.Remark != nil {
		daoItem.Remark = item.Remark // 备注
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

// DictProto 数据绑定
func DictProto(item dao.Dict) *DictObject {
	res := &DictObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Sort != nil {
		res.Sort = item.Sort
	}
	if item.Label != nil {
		res.Label = item.Label
	}
	if item.Value != nil {
		res.Value = item.Value
	}
	if item.DictType != nil {
		res.DictType = item.DictType
	}
	if item.ColorType.IsValid() {
		res.ColorType = item.ColorType.Ptr()
	}
	if item.CssClass.IsValid() {
		res.CssClass = item.CssClass.Ptr()
	}
	if item.Status != nil {
		res.Status = item.Status
	}
	if item.Remark != nil {
		res.Remark = item.Remark
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
