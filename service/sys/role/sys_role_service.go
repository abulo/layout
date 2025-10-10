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

// sys_role 角色

// SrvSysRoleServiceServer 角色
type SrvSysRoleServiceServer struct {
	UnimplementedSysRoleServiceServer
	Server *xgrpc.Server
}

// SysRoleCreate 创建数据
func (srv SrvSysRoleServiceServer) SysRoleCreate(ctx context.Context, request *SysRoleCreateRequest) (*SysRoleCreateResponse, error) {
	req := SysRoleDao(request.GetData())
	data, err := role.SysRoleCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:角色:sys_role:SysRoleCreate")
		return &SysRoleCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysRoleUpdate 更新数据
func (srv SrvSysRoleServiceServer) SysRoleUpdate(ctx context.Context, request *SysRoleUpdateRequest) (*SysRoleUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysRoleUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysRoleDao(request.GetData())
	_, err := role.SysRoleUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:角色:sys_role:SysRoleUpdate")
		return &SysRoleUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysRoleDelete 删除数据
func (srv SrvSysRoleServiceServer) SysRoleDelete(ctx context.Context, request *SysRoleDeleteRequest) (*SysRoleDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysRoleDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := role.SysRoleDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:角色:sys_role:SysRoleDelete")
		return &SysRoleDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysRole 查询单条数据
func (srv SrvSysRoleServiceServer) SysRole(ctx context.Context, request *SysRoleRequest) (*SysRoleResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysRoleResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := role.SysRole(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:角色:sys_role:SysRole")
		return &SysRoleResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysRoleProto(res),
	}, nil
}

// SysRoleRecover 恢复数据
func (srv SrvSysRoleServiceServer) SysRoleRecover(ctx context.Context, request *SysRoleRecoverRequest) (*SysRoleRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysRoleRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := role.SysRoleRecover(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:角色:sys_role:SysRoleRecover")
		return &SysRoleRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysRoleDrop 清理数据
func (srv SrvSysRoleServiceServer) SysRoleDrop(ctx context.Context, request *SysRoleDropRequest) (*SysRoleDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysRoleDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := role.SysRoleDrop(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:角色:sys_role:SysRoleDrop")
		return &SysRoleDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysRoleCode 查询单条数据
func (srv SrvSysRoleServiceServer) SysRoleCode(ctx context.Context, request *SysRoleCodeRequest) (*SysRoleCodeResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Code != nil {
		condition["code"] = request.GetCode()
	}

	if util.Empty(condition) {
		err := errors.New("condition is empty")
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:角色:sys_role:SysRoleCode")
		return &SysRoleCodeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	res, err := role.SysRoleCode(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:角色:sys_role:SysRoleCode")
		return &SysRoleCodeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleCodeResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysRoleProto(res),
	}, nil
}

// SysRoleList 列表数据
func (srv SrvSysRoleServiceServer) SysRoleList(ctx context.Context, request *SysRoleListRequest) (*SysRoleListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Code != nil {
		condition["code"] = request.GetCode()
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
	list, err := role.SysRoleList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:角色:sys_role:SysRoleList")
		return &SysRoleListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysRoleObject
	for _, item := range list {
		res = append(res, SysRoleProto(item))
	}
	return &SysRoleListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysRoleListTotal 获取总数
func (srv SrvSysRoleServiceServer) SysRoleListTotal(ctx context.Context, request *SysRoleListTotalRequest) (*SysRoleTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Code != nil {
		condition["code"] = request.GetCode()
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

	// 获取数据集合
	total, err := role.SysRoleListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:角色:sys_role:SysRoleListTotal")
		return &SysRoleTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysRoleTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
