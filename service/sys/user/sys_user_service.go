package user

import (
	"cloud/code"
	"cloud/module/sys/user"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// sys_user 用户

// SrvSysUserServiceServer 用户
type SrvSysUserServiceServer struct {
	UnimplementedSysUserServiceServer
	Server *xgrpc.Server
}

// SysUserCreate 创建数据
func (srv SrvSysUserServiceServer) SysUserCreate(ctx context.Context, request *SysUserCreateRequest) (*SysUserCreateResponse, error) {
	req := SysUserDao(request.GetData())
	data, err := user.SysUserCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:用户:sys_user:SysUserCreate")
		return &SysUserCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysUserUpdate 更新数据
func (srv SrvSysUserServiceServer) SysUserUpdate(ctx context.Context, request *SysUserUpdateRequest) (*SysUserUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysUserDao(request.GetData())
	_, err := user.SysUserUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:用户:sys_user:SysUserUpdate")
		return &SysUserUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysUserDelete 删除数据
func (srv SrvSysUserServiceServer) SysUserDelete(ctx context.Context, request *SysUserDeleteRequest) (*SysUserDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := user.SysUserDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:用户:sys_user:SysUserDelete")
		return &SysUserDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysUser 查询单条数据
func (srv SrvSysUserServiceServer) SysUser(ctx context.Context, request *SysUserRequest) (*SysUserResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := user.SysUser(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:用户:sys_user:SysUser")
		return &SysUserResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysUserProto(res),
	}, nil
}

// SysUserRecover 恢复数据
func (srv SrvSysUserServiceServer) SysUserRecover(ctx context.Context, request *SysUserRecoverRequest) (*SysUserRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := user.SysUserRecover(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:用户:sys_user:SysUserRecover")
		return &SysUserRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysUserDrop 清理数据
func (srv SrvSysUserServiceServer) SysUserDrop(ctx context.Context, request *SysUserDropRequest) (*SysUserDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := user.SysUserDrop(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:用户:sys_user:SysUserDrop")
		return &SysUserDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysUserLogin 查询单条数据
func (srv SrvSysUserServiceServer) SysUserLogin(ctx context.Context, request *SysUserLoginRequest) (*SysUserLoginResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}

	if util.Empty(condition) {
		err := errors.New("condition is empty")
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户:sys_user:SysUserLogin")
		return &SysUserLoginResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	res, err := user.SysUserLogin(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户:sys_user:SysUserLogin")
		return &SysUserLoginResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserLoginResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysUserProto(res),
	}, nil
}

// SysUserList 列表数据
func (srv SrvSysUserServiceServer) SysUserList(ctx context.Context, request *SysUserListRequest) (*SysUserListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Mobile != nil {
		condition["mobile"] = request.GetMobile()
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
	list, err := user.SysUserList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户:sys_user:SysUserList")
		return &SysUserListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysUserObject
	for _, item := range list {
		res = append(res, SysUserProto(item))
	}
	return &SysUserListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysUserListTotal 获取总数
func (srv SrvSysUserServiceServer) SysUserListTotal(ctx context.Context, request *SysUserListTotalRequest) (*SysUserTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Mobile != nil {
		condition["mobile"] = request.GetMobile()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}

	// 获取数据集合
	total, err := user.SysUserListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户:sys_user:SysUserListTotal")
		return &SysUserTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
