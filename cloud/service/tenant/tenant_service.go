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

// tenant 租户

// SrvTenantServiceServer 租户
type SrvTenantServiceServer struct {
	UnimplementedTenantServiceServer
	Server *xgrpc.Server
}

// TenantCreate 创建数据
func (srv SrvTenantServiceServer) TenantCreate(ctx context.Context, request *TenantCreateRequest) (*TenantCreateResponse, error) {
	req := TenantDao(request.GetData())
	data, err := tenant.TenantCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户:tenant:TenantCreate")
		return &TenantCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &TenantCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// TenantUpdate 更新数据
func (srv SrvTenantServiceServer) TenantUpdate(ctx context.Context, request *TenantUpdateRequest) (*TenantUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &TenantUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := TenantDao(request.GetData())
	_, err := tenant.TenantUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户:tenant:TenantUpdate")
		return &TenantUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &TenantUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// TenantDelete 删除数据
func (srv SrvTenantServiceServer) TenantDelete(ctx context.Context, request *TenantDeleteRequest) (*TenantDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &TenantDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.TenantDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户:tenant:TenantDelete")
		return &TenantDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &TenantDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// Tenant 查询单条数据
func (srv SrvTenantServiceServer) Tenant(ctx context.Context, request *TenantRequest) (*TenantResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &TenantResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := tenant.Tenant(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户:tenant:Tenant")
		return &TenantResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &TenantResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: TenantProto(res),
	}, nil
}

// TenantRecover 恢复数据
func (srv SrvTenantServiceServer) TenantRecover(ctx context.Context, request *TenantRecoverRequest) (*TenantRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &TenantRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.TenantRecover(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户:tenant:TenantRecover")
		return &TenantRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &TenantRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// TenantDrop 清理数据
func (srv SrvTenantServiceServer) TenantDrop(ctx context.Context, request *TenantDropRequest) (*TenantDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &TenantDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.TenantDrop(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户:tenant:TenantDrop")
		return &TenantDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &TenantDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// TenantList 列表数据
func (srv SrvTenantServiceServer) TenantList(ctx context.Context, request *TenantListRequest) (*TenantListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.PackageId != nil {
		condition["packageId"] = request.GetPackageId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}
	if request.ExpireDate != nil {
		condition["expireDate"] = request.GetExpireDate()
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
	list, err := tenant.TenantList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户:tenant:TenantList")
		return &TenantListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*TenantObject
	for _, item := range list {
		res = append(res, TenantProto(item))
	}
	return &TenantListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// TenantListTotal 获取总数
func (srv SrvTenantServiceServer) TenantListTotal(ctx context.Context, request *TenantListTotalRequest) (*TenantTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.PackageId != nil {
		condition["packageId"] = request.GetPackageId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}
	if request.ExpireDate != nil {
		condition["expireDate"] = request.GetExpireDate()
	}

	// 获取数据集合
	total, err := tenant.TenantListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户:tenant:TenantListTotal")
		return &TenantTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &TenantTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
