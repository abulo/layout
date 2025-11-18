package user

import (
	"cloud/code"
	"cloud/module/sys/user"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// sys_user_role 用户角色

// SrvSysUserRoleServiceServer 用户角色
type SrvSysUserRoleServiceServer struct {
	UnimplementedSysUserRoleServiceServer
	Server *xgrpc.Server
}

// SysUserRoleCreate 创建数据
func (srv SrvSysUserRoleServiceServer) SysUserRoleCreate(ctx context.Context, request *SysUserRoleCreateRequest) (*SysUserRoleCreateResponse, error) {
	req := SysUserRoleDao(request.GetData())
	data, err := user.SysUserRoleCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:用户角色:sys_user_role:SysUserRoleCreate")
		return &SysUserRoleCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserRoleCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysUserRoleUpdate 更新数据
func (srv SrvSysUserRoleServiceServer) SysUserRoleUpdate(ctx context.Context, request *SysUserRoleUpdateRequest) (*SysUserRoleUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserRoleUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysUserRoleDao(request.GetData())
	_, err := user.SysUserRoleUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:用户角色:sys_user_role:SysUserRoleUpdate")
		return &SysUserRoleUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserRoleUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysUserRoleDelete 删除数据
func (srv SrvSysUserRoleServiceServer) SysUserRoleDelete(ctx context.Context, request *SysUserRoleDeleteRequest) (*SysUserRoleDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserRoleDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := user.SysUserRoleDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:用户角色:sys_user_role:SysUserRoleDelete")
		return &SysUserRoleDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserRoleDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysUserRole 查询单条数据
func (srv SrvSysUserRoleServiceServer) SysUserRole(ctx context.Context, request *SysUserRoleRequest) (*SysUserRoleResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserRoleResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := user.SysUserRole(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:用户角色:sys_user_role:SysUserRole")
		return &SysUserRoleResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserRoleResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysUserRoleProto(res),
	}, nil
}

// SysUserRoleItem 查询单条数据
func (srv SrvSysUserRoleServiceServer) SysUserRoleItem(ctx context.Context, request *SysUserRoleItemRequest) (*SysUserRoleItemResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.RoleId != nil {
		condition["roleId"] = request.GetRoleId()
	}

	if util.Empty(condition) {
		err := errors.New("condition is empty")
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户角色:sys_user_role:SysUserRoleItem")
		return &SysUserRoleItemResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	res, err := user.SysUserRoleItem(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户角色:sys_user_role:SysUserRoleItem")
		return &SysUserRoleItemResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserRoleItemResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysUserRoleProto(res),
	}, nil
}

// SysUserRoleList 列表数据
func (srv SrvSysUserRoleServiceServer) SysUserRoleList(ctx context.Context, request *SysUserRoleListRequest) (*SysUserRoleListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.RoleId != nil {
		condition["roleId"] = request.GetRoleId()
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
	list, err := user.SysUserRoleList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户角色:sys_user_role:SysUserRoleList")
		return &SysUserRoleListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysUserRoleObject
	for _, item := range list {
		res = append(res, SysUserRoleProto(item))
	}
	return &SysUserRoleListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysUserRoleListTotal 获取总数
func (srv SrvSysUserRoleServiceServer) SysUserRoleListTotal(ctx context.Context, request *SysUserRoleListTotalRequest) (*SysUserRoleTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.RoleId != nil {
		condition["roleId"] = request.GetRoleId()
	}

	// 获取数据集合
	total, err := user.SysUserRoleListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户角色:sys_user_role:SysUserRoleListTotal")
		return &SysUserRoleTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserRoleTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
