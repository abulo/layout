package dept

import (
	"cloud/code"
	"cloud/module/sys/dept"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// sys_dept 部门

// SrvSysDeptServiceServer 部门
type SrvSysDeptServiceServer struct {
	UnimplementedSysDeptServiceServer
	Server *xgrpc.Server
}

// SysDeptCreate 创建数据
func (srv SrvSysDeptServiceServer) SysDeptCreate(ctx context.Context, request *SysDeptCreateRequest) (*SysDeptCreateResponse, error) {
	req := SysDeptDao(request.GetData())
	data, err := dept.SysDeptCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:部门:sys_dept:SysDeptCreate")
		return &SysDeptCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDeptCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysDeptUpdate 更新数据
func (srv SrvSysDeptServiceServer) SysDeptUpdate(ctx context.Context, request *SysDeptUpdateRequest) (*SysDeptUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysDeptUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysDeptDao(request.GetData())
	_, err := dept.SysDeptUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:部门:sys_dept:SysDeptUpdate")
		return &SysDeptUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDeptUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysDeptDelete 删除数据
func (srv SrvSysDeptServiceServer) SysDeptDelete(ctx context.Context, request *SysDeptDeleteRequest) (*SysDeptDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysDeptDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dept.SysDeptDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:部门:sys_dept:SysDeptDelete")
		return &SysDeptDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDeptDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysDept 查询单条数据
func (srv SrvSysDeptServiceServer) SysDept(ctx context.Context, request *SysDeptRequest) (*SysDeptResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysDeptResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := dept.SysDept(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:部门:sys_dept:SysDept")
		return &SysDeptResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDeptResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysDeptProto(res),
	}, nil
}

// SysDeptRecover 恢复数据
func (srv SrvSysDeptServiceServer) SysDeptRecover(ctx context.Context, request *SysDeptRecoverRequest) (*SysDeptRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysDeptRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dept.SysDeptRecover(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:部门:sys_dept:SysDeptRecover")
		return &SysDeptRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDeptRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysDeptDrop 清理数据
func (srv SrvSysDeptServiceServer) SysDeptDrop(ctx context.Context, request *SysDeptDropRequest) (*SysDeptDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysDeptDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dept.SysDeptDrop(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:部门:sys_dept:SysDeptDrop")
		return &SysDeptDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDeptDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysDeptList 列表数据
func (srv SrvSysDeptServiceServer) SysDeptList(ctx context.Context, request *SysDeptListRequest) (*SysDeptListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.ParentId != nil {
		condition["parentId"] = request.GetParentId()
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
	list, err := dept.SysDeptList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:部门:sys_dept:SysDeptList")
		return &SysDeptListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysDeptObject
	for _, item := range list {
		res = append(res, SysDeptProto(item))
	}
	return &SysDeptListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysDeptListTotal 获取总数
func (srv SrvSysDeptServiceServer) SysDeptListTotal(ctx context.Context, request *SysDeptListTotalRequest) (*SysDeptTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.ParentId != nil {
		condition["parentId"] = request.GetParentId()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}

	// 获取数据集合
	total, err := dept.SysDeptListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:部门:sys_dept:SysDeptListTotal")
		return &SysDeptTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDeptTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
