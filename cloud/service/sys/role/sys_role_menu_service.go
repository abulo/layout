package role

import (
	"cloud/code"
	"cloud/module/sys/role"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// sys_role_menu 角色菜单

// SrvSysRoleMenuServiceServer 角色菜单
type SrvSysRoleMenuServiceServer struct {
	UnimplementedSysRoleMenuServiceServer
	Server *xgrpc.Server
}

// SysRoleMenuCreate 创建数据
func (srv SrvSysRoleMenuServiceServer) SysRoleMenuCreate(ctx context.Context, request *SysRoleMenuCreateRequest) (*SysRoleMenuCreateResponse, error) {
	req := SysRoleMenuDao(request.GetData())
	data, err := role.SysRoleMenuCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:角色菜单:sys_role_menu:SysRoleMenuCreate")
		return &SysRoleMenuCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleMenuCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysRoleMenuUpdate 更新数据
func (srv SrvSysRoleMenuServiceServer) SysRoleMenuUpdate(ctx context.Context, request *SysRoleMenuUpdateRequest) (*SysRoleMenuUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysRoleMenuUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysRoleMenuDao(request.GetData())
	_, err := role.SysRoleMenuUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:角色菜单:sys_role_menu:SysRoleMenuUpdate")
		return &SysRoleMenuUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleMenuUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysRoleMenuDelete 删除数据
func (srv SrvSysRoleMenuServiceServer) SysRoleMenuDelete(ctx context.Context, request *SysRoleMenuDeleteRequest) (*SysRoleMenuDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysRoleMenuDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := role.SysRoleMenuDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:角色菜单:sys_role_menu:SysRoleMenuDelete")
		return &SysRoleMenuDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleMenuDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysRoleMenu 查询单条数据
func (srv SrvSysRoleMenuServiceServer) SysRoleMenu(ctx context.Context, request *SysRoleMenuRequest) (*SysRoleMenuResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysRoleMenuResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := role.SysRoleMenu(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:角色菜单:sys_role_menu:SysRoleMenu")
		return &SysRoleMenuResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleMenuResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysRoleMenuProto(res),
	}, nil
}

// SysRoleMenuItem 查询单条数据
func (srv SrvSysRoleMenuServiceServer) SysRoleMenuItem(ctx context.Context, request *SysRoleMenuItemRequest) (*SysRoleMenuItemResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.RoleId != nil {
		condition["roleId"] = request.GetRoleId()
	}
	if request.MenuId != nil {
		condition["menuId"] = request.GetMenuId()
	}

	if util.Empty(condition) {
		err := errors.New("condition is empty")
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:角色菜单:sys_role_menu:SysRoleMenuItem")
		return &SysRoleMenuItemResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	res, err := role.SysRoleMenuItem(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:角色菜单:sys_role_menu:SysRoleMenuItem")
		return &SysRoleMenuItemResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleMenuItemResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysRoleMenuProto(res),
	}, nil
}

// SysRoleMenuList 列表数据
func (srv SrvSysRoleMenuServiceServer) SysRoleMenuList(ctx context.Context, request *SysRoleMenuListRequest) (*SysRoleMenuListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.RoleId != nil {
		condition["roleId"] = request.GetRoleId()
	}
	if request.MenuId != nil {
		condition["menuId"] = request.GetMenuId()
	}

	paginationRequest := request.GetPagination()
	if paginationRequest != nil {
		// 当前页面
		pageNum := paginationRequest.GetPageNum()
		// 每页多少数据
		pageSize := paginationRequest.GetPageSize()
		if pageNum < 1 {
			pageNum = 1
		}
		if pageSize < 1 {
			pageSize = 10
		}
		// 分页数据
		offset := pageSize * (pageNum - 1)
		pagination := &sql.Pagination{
			Offset: &offset,
			Limit:  &pageSize,
		}
		condition["pagination"] = pagination
	}
	// 获取数据集合
	list, err := role.SysRoleMenuList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:角色菜单:sys_role_menu:SysRoleMenuList")
		return &SysRoleMenuListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysRoleMenuObject
	for _, item := range list {
		res = append(res, SysRoleMenuProto(item))
	}
	return &SysRoleMenuListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysRoleMenuListTotal 获取总数
func (srv SrvSysRoleMenuServiceServer) SysRoleMenuListTotal(ctx context.Context, request *SysRoleMenuListTotalRequest) (*SysRoleMenuTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.RoleId != nil {
		condition["roleId"] = request.GetRoleId()
	}
	if request.MenuId != nil {
		condition["menuId"] = request.GetMenuId()
	}

	// 获取数据集合
	total, err := role.SysRoleMenuListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:角色菜单:sys_role_menu:SysRoleMenuListTotal")
		return &SysRoleMenuTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleMenuTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
