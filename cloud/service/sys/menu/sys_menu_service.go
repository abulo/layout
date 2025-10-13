package menu

import (
	"cloud/code"
	"cloud/module/sys/menu"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// sys_menu 菜单

// SrvSysMenuServiceServer 菜单
type SrvSysMenuServiceServer struct {
	UnimplementedSysMenuServiceServer
	Server *xgrpc.Server
}

// SysMenuCreate 创建数据
func (srv SrvSysMenuServiceServer) SysMenuCreate(ctx context.Context, request *SysMenuCreateRequest) (*SysMenuCreateResponse, error) {
	req := SysMenuDao(request.GetData())
	data, err := menu.SysMenuCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:菜单:sys_menu:SysMenuCreate")
		return &SysMenuCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysMenuCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysMenuUpdate 更新数据
func (srv SrvSysMenuServiceServer) SysMenuUpdate(ctx context.Context, request *SysMenuUpdateRequest) (*SysMenuUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysMenuUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysMenuDao(request.GetData())
	_, err := menu.SysMenuUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:菜单:sys_menu:SysMenuUpdate")
		return &SysMenuUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysMenuUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysMenuDelete 删除数据
func (srv SrvSysMenuServiceServer) SysMenuDelete(ctx context.Context, request *SysMenuDeleteRequest) (*SysMenuDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysMenuDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := menu.SysMenuDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:菜单:sys_menu:SysMenuDelete")
		return &SysMenuDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysMenuDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysMenu 查询单条数据
func (srv SrvSysMenuServiceServer) SysMenu(ctx context.Context, request *SysMenuRequest) (*SysMenuResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysMenuResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := menu.SysMenu(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:菜单:sys_menu:SysMenu")
		return &SysMenuResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysMenuResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysMenuProto(res),
	}, nil
}

// SysMenuParent 列表数据
func (srv SrvSysMenuServiceServer) SysMenuParent(ctx context.Context, request *SysMenuParentRequest) (*SysMenuParentResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.ParentId != nil {
		condition["parentId"] = request.GetParentId()
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
	list, err := menu.SysMenuParent(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:菜单:sys_menu:SysMenuParent")
		return &SysMenuParentResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysMenuObject
	for _, item := range list {
		res = append(res, SysMenuProto(item))
	}
	return &SysMenuParentResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysMenuParentTotal 获取总数
func (srv SrvSysMenuServiceServer) SysMenuParentTotal(ctx context.Context, request *SysMenuParentTotalRequest) (*SysMenuTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.ParentId != nil {
		condition["parentId"] = request.GetParentId()
	}

	// 获取数据集合
	total, err := menu.SysMenuParentTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:菜单:sys_menu:SysMenuParentTotal")
		return &SysMenuTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysMenuTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}

// SysMenuList 列表数据
func (srv SrvSysMenuServiceServer) SysMenuList(ctx context.Context, request *SysMenuListRequest) (*SysMenuListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.ParentId != nil {
		condition["parentId"] = request.GetParentId()
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
	list, err := menu.SysMenuList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:菜单:sys_menu:SysMenuList")
		return &SysMenuListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysMenuObject
	for _, item := range list {
		res = append(res, SysMenuProto(item))
	}
	return &SysMenuListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysMenuListTotal 获取总数
func (srv SrvSysMenuServiceServer) SysMenuListTotal(ctx context.Context, request *SysMenuListTotalRequest) (*SysMenuTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.ParentId != nil {
		condition["parentId"] = request.GetParentId()
	}

	// 获取数据集合
	total, err := menu.SysMenuListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:菜单:sys_menu:SysMenuListTotal")
		return &SysMenuTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysMenuTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
