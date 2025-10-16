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

// sys_user_tenant 租户用户

// SrvSysUserTenantServiceServer 租户用户
type SrvSysUserTenantServiceServer struct {
	UnimplementedSysUserTenantServiceServer
	Server *xgrpc.Server
}

// SysUserTenantCreate 创建数据
func (srv SrvSysUserTenantServiceServer) SysUserTenantCreate(ctx context.Context, request *SysUserTenantCreateRequest) (*SysUserTenantCreateResponse, error) {
	req := SysUserTenantDao(request.GetData())
	data, err := user.SysUserTenantCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户用户:sys_user_tenant:SysUserTenantCreate")
		return &SysUserTenantCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserTenantCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysUserTenantUpdate 更新数据
func (srv SrvSysUserTenantServiceServer) SysUserTenantUpdate(ctx context.Context, request *SysUserTenantUpdateRequest) (*SysUserTenantUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserTenantUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysUserTenantDao(request.GetData())
	_, err := user.SysUserTenantUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户用户:sys_user_tenant:SysUserTenantUpdate")
		return &SysUserTenantUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserTenantUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysUserTenantDelete 删除数据
func (srv SrvSysUserTenantServiceServer) SysUserTenantDelete(ctx context.Context, request *SysUserTenantDeleteRequest) (*SysUserTenantDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserTenantDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := user.SysUserTenantDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户用户:sys_user_tenant:SysUserTenantDelete")
		return &SysUserTenantDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserTenantDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysUserTenant 查询单条数据
func (srv SrvSysUserTenantServiceServer) SysUserTenant(ctx context.Context, request *SysUserTenantRequest) (*SysUserTenantResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserTenantResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := user.SysUserTenant(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户用户:sys_user_tenant:SysUserTenant")
		return &SysUserTenantResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserTenantResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysUserTenantProto(res),
	}, nil
}

// SysUserTenantBind 查询单条数据
func (srv SrvSysUserTenantServiceServer) SysUserTenantBind(ctx context.Context, request *SysUserTenantBindRequest) (*SysUserTenantBindResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}

	if util.Empty(condition) {
		err := errors.New("condition is empty")
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户用户:sys_user_tenant:SysUserTenantBind")
		return &SysUserTenantBindResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	res, err := user.SysUserTenantBind(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户用户:sys_user_tenant:SysUserTenantBind")
		return &SysUserTenantBindResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserTenantBindResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysUserTenantProto(res),
	}, nil
}

// SysUserTenantList 列表数据
func (srv SrvSysUserTenantServiceServer) SysUserTenantList(ctx context.Context, request *SysUserTenantListRequest) (*SysUserTenantListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
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
	list, err := user.SysUserTenantList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户用户:sys_user_tenant:SysUserTenantList")
		return &SysUserTenantListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysUserTenantObject
	for _, item := range list {
		res = append(res, SysUserTenantProto(item))
	}
	return &SysUserTenantListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysUserTenantListTotal 获取总数
func (srv SrvSysUserTenantServiceServer) SysUserTenantListTotal(ctx context.Context, request *SysUserTenantListTotalRequest) (*SysUserTenantTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}

	// 获取数据集合
	total, err := user.SysUserTenantListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户用户:sys_user_tenant:SysUserTenantListTotal")
		return &SysUserTenantTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserTenantTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
