package tenant

import (
	"cloud/code"
	"cloud/module/sys/tenant"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// sys_tenant 租户

// SrvSysTenantServiceServer 租户
type SrvSysTenantServiceServer struct {
	UnimplementedSysTenantServiceServer
	Server *xgrpc.Server
}

// SysTenantCreate 创建数据
func (srv SrvSysTenantServiceServer) SysTenantCreate(ctx context.Context, request *SysTenantCreateRequest) (*SysTenantCreateResponse, error) {
	req := SysTenantDao(request.GetData())
	data, err := tenant.SysTenantCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户:sys_tenant:SysTenantCreate")
		return &SysTenantCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysTenantCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysTenantUpdate 更新数据
func (srv SrvSysTenantServiceServer) SysTenantUpdate(ctx context.Context, request *SysTenantUpdateRequest) (*SysTenantUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysTenantUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysTenantDao(request.GetData())
	_, err := tenant.SysTenantUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:租户:sys_tenant:SysTenantUpdate")
		return &SysTenantUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysTenantUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysTenantDelete 删除数据
func (srv SrvSysTenantServiceServer) SysTenantDelete(ctx context.Context, request *SysTenantDeleteRequest) (*SysTenantDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysTenantDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.SysTenantDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户:sys_tenant:SysTenantDelete")
		return &SysTenantDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysTenantDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysTenant 查询单条数据
func (srv SrvSysTenantServiceServer) SysTenant(ctx context.Context, request *SysTenantRequest) (*SysTenantResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysTenantResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := tenant.SysTenant(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户:sys_tenant:SysTenant")
		return &SysTenantResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysTenantResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysTenantProto(res),
	}, nil
}

// SysTenantRecover 恢复数据
func (srv SrvSysTenantServiceServer) SysTenantRecover(ctx context.Context, request *SysTenantRecoverRequest) (*SysTenantRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysTenantRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.SysTenantRecover(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户:sys_tenant:SysTenantRecover")
		return &SysTenantRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysTenantRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysTenantDrop 清理数据
func (srv SrvSysTenantServiceServer) SysTenantDrop(ctx context.Context, request *SysTenantDropRequest) (*SysTenantDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysTenantDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := tenant.SysTenantDrop(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:租户:sys_tenant:SysTenantDrop")
		return &SysTenantDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysTenantDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysTenantList 列表数据
func (srv SrvSysTenantServiceServer) SysTenantList(ctx context.Context, request *SysTenantListRequest) (*SysTenantListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}
	if request.BeginExpireDate != nil {
		condition["beginExpireDate"] = util.Date("Y-m-d", util.GrpcTime(request.GetBeginExpireDate()))
	}
	if request.FinishExpireDate != nil {
		condition["finishExpireDate"] = util.Date("Y-m-d", util.GrpcTime(request.GetFinishExpireDate()))
	}
	if request.PackageId != nil {
		condition["packageId"] = request.GetPackageId()
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
	list, err := tenant.SysTenantList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户:sys_tenant:SysTenantList")
		return &SysTenantListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysTenantObject
	for _, item := range list {
		res = append(res, SysTenantProto(item))
	}
	return &SysTenantListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysTenantListTotal 获取总数
func (srv SrvSysTenantServiceServer) SysTenantListTotal(ctx context.Context, request *SysTenantListTotalRequest) (*SysTenantTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}
	if request.BeginExpireDate != nil {
		condition["beginExpireDate"] = util.Date("Y-m-d", util.GrpcTime(request.GetBeginExpireDate()))
	}
	if request.FinishExpireDate != nil {
		condition["finishExpireDate"] = util.Date("Y-m-d", util.GrpcTime(request.GetFinishExpireDate()))
	}
	if request.PackageId != nil {
		condition["packageId"] = request.GetPackageId()
	}

	// 获取数据集合
	total, err := tenant.SysTenantListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:租户:sys_tenant:SysTenantListTotal")
		return &SysTenantTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysTenantTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
