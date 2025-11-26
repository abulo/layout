package role

import (
	"cloud/dao"
	"encoding/json"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// sys_role 角色

// SysRoleDao 数据转换
func SysRoleDao(item *SysRoleObject) *dao.SysRole {
	daoItem := &dao.SysRole{}

	if item != nil && item.Id != nil {
		daoItem.Id = item.Id // 编号
	}
	if item != nil && item.Name != nil {
		daoItem.Name = item.Name // 名称
	}
	if item != nil && item.Code != nil {
		daoItem.Code = item.Code // 编码
	}
	if item != nil && item.Scope != nil {
		daoItem.Scope = null.Int32From(item.GetScope()) // 数据范围:1:全部数据权限/2:自定数据权限/3:本部门数据权限/4:本部门及以下数据权限
	}
	if item != nil && item.ScopeDept != nil {
		var deptIds []int64
		var deptIdsNew []int64
		var deptIdsByte []byte
		if err := json.Unmarshal(item.GetScopeDept(), &deptIds); err == nil {
			for _, deptId := range deptIds {
				if !util.InArray(deptId, deptIdsNew) {
					deptIdsNew = append(deptIdsNew, deptId)
				}
			}
			deptIdsByte, _ = json.Marshal(deptIdsNew)
		}
		daoItem.ScopeDept = null.JSONFrom(deptIdsByte)
	}
	if item != nil && item.Sort != nil {
		daoItem.Sort = item.Sort // 排序
	}
	if item != nil && item.Status != nil {
		daoItem.Status = item.Status // 状态:0正常/1停用
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
	if item != nil && item.MenuIds != nil {
		// 这里需要将数据去重一下
		var menuIds []int64
		var menuIdsNew []int64
		var menuIdsByte []byte
		if err := json.Unmarshal(item.GetMenuIds(), &menuIds); err == nil {
			for _, menuId := range menuIds {
				if !util.InArray(menuId, menuIdsNew) {
					menuIdsNew = append(menuIdsNew, menuId)
				}
			}
			menuIdsByte, _ = json.Marshal(menuIdsNew)
		}
		daoItem.MenuIds = null.JSONFrom(menuIdsByte)
	}

	return daoItem
}

// SysRoleProto 数据绑定
func SysRoleProto(item dao.SysRole) *SysRoleObject {
	res := &SysRoleObject{}
	if item.Id != nil {
		res.Id = item.Id
	}
	if item.Name != nil {
		res.Name = item.Name
	}
	if item.Code != nil {
		res.Code = item.Code
	}
	if item.Scope.IsValid() {
		res.Scope = item.Scope.Ptr()
	}
	if item.ScopeDept.IsValid() {
		res.ScopeDept = *item.ScopeDept.Ptr()
	}
	if item.Sort != nil {
		res.Sort = item.Sort
	}
	if item.Status != nil {
		res.Status = item.Status
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
	if item.MenuIds.IsValid() {
		res.MenuIds = *item.MenuIds.Ptr()
	}

	return res
}
