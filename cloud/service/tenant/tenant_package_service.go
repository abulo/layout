package tenant

import (
	"cloud/code"
	"cloud/module/tenant"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// tenant_package 租户套餐

// SrvTenantPackageServiceServer 租户套餐
type SrvTenantPackageServiceServer struct {
	UnimplementedTenantPackageServiceServer
	Server *xgrpc.Server
}

// TenantPackageCreate 创建数据
func (srv SrvTenantPackageServiceServer) TenantPackageCreate(ctx context.Context, request *TenantPackageCreateRequest) (*TenantPackageCreateResponse, error) {
	req := TenantPackageDao(request.GetData())
	data, err := tenant.TenantPackageCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户套餐:tenant_package:TenantPackageCreate")
		return &TenantPackageCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &TenantPackageCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// TenantPackageUpdate 更新数据
func (srv SrvTenantPackageServiceServer) TenantPackageUpdate(ctx context.Context, request *TenantPackageUpdateRequest) (*TenantPackageUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &TenantPackageUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := TenantPackageDao(request.GetData())
	_, err := tenant.TenantPackageUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户套餐:tenant_package:TenantPackageUpdate")
		return &TenantPackageUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &TenantPackageUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// TenantPackageDelete 删除数据
func (srv SrvTenantPackageServiceServer) TenantPackageDelete(ctx context.Context, request *TenantPackageDeleteRequest) (*TenantPackageDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &TenantPackageDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.TenantPackageDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户套餐:tenant_package:TenantPackageDelete")
		return &TenantPackageDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &TenantPackageDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// TenantPackage 查询单条数据
func (srv SrvTenantPackageServiceServer) TenantPackage(ctx context.Context, request *TenantPackageRequest) (*TenantPackageResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &TenantPackageResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := tenant.TenantPackage(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户套餐:tenant_package:TenantPackage")
		return &TenantPackageResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &TenantPackageResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: TenantPackageProto(res),
	}, nil
}

// TenantPackageList 列表数据
func (srv SrvTenantPackageServiceServer) TenantPackageList(ctx context.Context, request *TenantPackageListRequest) (*TenantPackageListResponse, error) {
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
	list, err := tenant.TenantPackageList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户套餐:tenant_package:TenantPackageList")
		return &TenantPackageListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*TenantPackageObject
	for _, item := range list {
		res = append(res, TenantPackageProto(item))
	}
	return &TenantPackageListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// TenantPackageListTotal 获取总数
func (srv SrvTenantPackageServiceServer) TenantPackageListTotal(ctx context.Context, request *TenantPackageListTotalRequest) (*TenantPackageTotalResponse, error) {
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
	total, err := tenant.TenantPackageListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户套餐:tenant_package:TenantPackageListTotal")
		return &TenantPackageTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &TenantPackageTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
