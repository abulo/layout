package tenant

import (
	"cloud/code"
	"cloud/module/sys/tenant"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// sys_tenant_package 租户套餐

// SrvSysTenantPackageServiceServer 租户套餐
type SrvSysTenantPackageServiceServer struct {
	UnimplementedSysTenantPackageServiceServer
	Server *xgrpc.Server
}

// SysTenantPackageCreate 创建数据
func (srv SrvSysTenantPackageServiceServer) SysTenantPackageCreate(ctx context.Context, request *SysTenantPackageCreateRequest) (*SysTenantPackageCreateResponse, error) {
	req := SysTenantPackageDao(request.GetData())
	data, err := tenant.SysTenantPackageCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户套餐:sys_tenant_package:SysTenantPackageCreate")
		return &SysTenantPackageCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysTenantPackageCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysTenantPackageUpdate 更新数据
func (srv SrvSysTenantPackageServiceServer) SysTenantPackageUpdate(ctx context.Context, request *SysTenantPackageUpdateRequest) (*SysTenantPackageUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysTenantPackageUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysTenantPackageDao(request.GetData())
	_, err := tenant.SysTenantPackageUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户套餐:sys_tenant_package:SysTenantPackageUpdate")
		return &SysTenantPackageUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysTenantPackageUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysTenantPackageDelete 删除数据
func (srv SrvSysTenantPackageServiceServer) SysTenantPackageDelete(ctx context.Context, request *SysTenantPackageDeleteRequest) (*SysTenantPackageDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysTenantPackageDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.SysTenantPackageDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户套餐:sys_tenant_package:SysTenantPackageDelete")
		return &SysTenantPackageDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysTenantPackageDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysTenantPackage 查询单条数据
func (srv SrvSysTenantPackageServiceServer) SysTenantPackage(ctx context.Context, request *SysTenantPackageRequest) (*SysTenantPackageResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysTenantPackageResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := tenant.SysTenantPackage(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户套餐:sys_tenant_package:SysTenantPackage")
		return &SysTenantPackageResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysTenantPackageResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysTenantPackageProto(res),
	}, nil
}

// SysTenantPackageList 列表数据
func (srv SrvSysTenantPackageServiceServer) SysTenantPackageList(ctx context.Context, request *SysTenantPackageListRequest) (*SysTenantPackageListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
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
	list, err := tenant.SysTenantPackageList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户套餐:sys_tenant_package:SysTenantPackageList")
		return &SysTenantPackageListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysTenantPackageObject
	for _, item := range list {
		res = append(res, SysTenantPackageProto(item))
	}
	return &SysTenantPackageListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysTenantPackageListTotal 获取总数
func (srv SrvSysTenantPackageServiceServer) SysTenantPackageListTotal(ctx context.Context, request *SysTenantPackageListTotalRequest) (*SysTenantPackageTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}

	// 获取数据集合
	total, err := tenant.SysTenantPackageListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户套餐:sys_tenant_package:SysTenantPackageListTotal")
		return &SysTenantPackageTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysTenantPackageTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
