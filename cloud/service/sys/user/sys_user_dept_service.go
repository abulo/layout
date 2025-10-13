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

// sys_user_dept 用户部门

// SrvSysUserDeptServiceServer 用户部门
type SrvSysUserDeptServiceServer struct {
	UnimplementedSysUserDeptServiceServer
	Server *xgrpc.Server
}

// SysUserDeptCreate 创建数据
func (srv SrvSysUserDeptServiceServer) SysUserDeptCreate(ctx context.Context, request *SysUserDeptCreateRequest) (*SysUserDeptCreateResponse, error) {
	req := SysUserDeptDao(request.GetData())
	data, err := user.SysUserDeptCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:用户部门:sys_user_dept:SysUserDeptCreate")
		return &SysUserDeptCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserDeptCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysUserDeptUpdate 更新数据
func (srv SrvSysUserDeptServiceServer) SysUserDeptUpdate(ctx context.Context, request *SysUserDeptUpdateRequest) (*SysUserDeptUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserDeptUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysUserDeptDao(request.GetData())
	_, err := user.SysUserDeptUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:用户部门:sys_user_dept:SysUserDeptUpdate")
		return &SysUserDeptUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserDeptUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysUserDeptDelete 删除数据
func (srv SrvSysUserDeptServiceServer) SysUserDeptDelete(ctx context.Context, request *SysUserDeptDeleteRequest) (*SysUserDeptDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserDeptDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := user.SysUserDeptDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:用户部门:sys_user_dept:SysUserDeptDelete")
		return &SysUserDeptDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserDeptDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysUserDept 查询单条数据
func (srv SrvSysUserDeptServiceServer) SysUserDept(ctx context.Context, request *SysUserDeptRequest) (*SysUserDeptResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserDeptResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := user.SysUserDept(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:用户部门:sys_user_dept:SysUserDept")
		return &SysUserDeptResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserDeptResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysUserDeptProto(res),
	}, nil
}

// SysUserDeptItem 查询单条数据
func (srv SrvSysUserDeptServiceServer) SysUserDeptItem(ctx context.Context, request *SysUserDeptItemRequest) (*SysUserDeptItemResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.DeptId != nil {
		condition["deptId"] = request.GetDeptId()
	}

	if util.Empty(condition) {
		err := errors.New("condition is empty")
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户部门:sys_user_dept:SysUserDeptItem")
		return &SysUserDeptItemResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	res, err := user.SysUserDeptItem(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户部门:sys_user_dept:SysUserDeptItem")
		return &SysUserDeptItemResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserDeptItemResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysUserDeptProto(res),
	}, nil
}

// SysUserDeptList 列表数据
func (srv SrvSysUserDeptServiceServer) SysUserDeptList(ctx context.Context, request *SysUserDeptListRequest) (*SysUserDeptListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.DeptId != nil {
		condition["deptId"] = request.GetDeptId()
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
	list, err := user.SysUserDeptList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户部门:sys_user_dept:SysUserDeptList")
		return &SysUserDeptListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysUserDeptObject
	for _, item := range list {
		res = append(res, SysUserDeptProto(item))
	}
	return &SysUserDeptListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysUserDeptListTotal 获取总数
func (srv SrvSysUserDeptServiceServer) SysUserDeptListTotal(ctx context.Context, request *SysUserDeptListTotalRequest) (*SysUserDeptTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.DeptId != nil {
		condition["deptId"] = request.GetDeptId()
	}

	// 获取数据集合
	total, err := user.SysUserDeptListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户部门:sys_user_dept:SysUserDeptListTotal")
		return &SysUserDeptTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserDeptTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
