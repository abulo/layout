package logger

import (
	"cloud/code"
	"cloud/module/sys/logger"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// sys_logger_login 登录日志

// SrvSysLoggerLoginServiceServer 登录日志
type SrvSysLoggerLoginServiceServer struct {
	UnimplementedSysLoggerLoginServiceServer
	Server *xgrpc.Server
}

// SysLoggerLoginCreate 创建数据
func (srv SrvSysLoggerLoginServiceServer) SysLoggerLoginCreate(ctx context.Context, request *SysLoggerLoginCreateRequest) (*SysLoggerLoginCreateResponse, error) {
	req := SysLoggerLoginDao(request.GetData())
	data, err := logger.SysLoggerLoginCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:登录日志:sys_logger_login:SysLoggerLoginCreate")
		return &SysLoggerLoginCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerLoginCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysLoggerLoginUpdate 更新数据
func (srv SrvSysLoggerLoginServiceServer) SysLoggerLoginUpdate(ctx context.Context, request *SysLoggerLoginUpdateRequest) (*SysLoggerLoginUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysLoggerLoginUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysLoggerLoginDao(request.GetData())
	_, err := logger.SysLoggerLoginUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:登录日志:sys_logger_login:SysLoggerLoginUpdate")
		return &SysLoggerLoginUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerLoginUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysLoggerLoginDelete 删除数据
func (srv SrvSysLoggerLoginServiceServer) SysLoggerLoginDelete(ctx context.Context, request *SysLoggerLoginDeleteRequest) (*SysLoggerLoginDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysLoggerLoginDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := logger.SysLoggerLoginDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:登录日志:sys_logger_login:SysLoggerLoginDelete")
		return &SysLoggerLoginDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerLoginDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysLoggerLogin 查询单条数据
func (srv SrvSysLoggerLoginServiceServer) SysLoggerLogin(ctx context.Context, request *SysLoggerLoginRequest) (*SysLoggerLoginResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysLoggerLoginResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := logger.SysLoggerLogin(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:登录日志:sys_logger_login:SysLoggerLogin")
		return &SysLoggerLoginResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerLoginResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysLoggerLoginProto(res),
	}, nil
}

// SysLoggerLoginRecover 恢复数据
func (srv SrvSysLoggerLoginServiceServer) SysLoggerLoginRecover(ctx context.Context, request *SysLoggerLoginRecoverRequest) (*SysLoggerLoginRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysLoggerLoginRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := logger.SysLoggerLoginRecover(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:登录日志:sys_logger_login:SysLoggerLoginRecover")
		return &SysLoggerLoginRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerLoginRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysLoggerLoginDrop 清理数据
func (srv SrvSysLoggerLoginServiceServer) SysLoggerLoginDrop(ctx context.Context, request *SysLoggerLoginDropRequest) (*SysLoggerLoginDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysLoggerLoginDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := logger.SysLoggerLoginDrop(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:登录日志:sys_logger_login:SysLoggerLoginDrop")
		return &SysLoggerLoginDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerLoginDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysLoggerLoginList 列表数据
func (srv SrvSysLoggerLoginServiceServer) SysLoggerLoginList(ctx context.Context, request *SysLoggerLoginListRequest) (*SysLoggerLoginListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Channel != nil {
		condition["channel"] = request.GetChannel()
	}
	if request.LoginTime != nil {
		condition["loginTime"] = request.GetLoginTime()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.Username != nil {
		condition["username"] = request.GetUsername()
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
	list, err := logger.SysLoggerLoginList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:登录日志:sys_logger_login:SysLoggerLoginList")
		return &SysLoggerLoginListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysLoggerLoginObject
	for _, item := range list {
		res = append(res, SysLoggerLoginProto(item))
	}
	return &SysLoggerLoginListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysLoggerLoginListTotal 获取总数
func (srv SrvSysLoggerLoginServiceServer) SysLoggerLoginListTotal(ctx context.Context, request *SysLoggerLoginListTotalRequest) (*SysLoggerLoginTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Channel != nil {
		condition["channel"] = request.GetChannel()
	}
	if request.LoginTime != nil {
		condition["loginTime"] = request.GetLoginTime()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}

	// 获取数据集合
	total, err := logger.SysLoggerLoginListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:登录日志:sys_logger_login:SysLoggerLoginListTotal")
		return &SysLoggerLoginTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerLoginTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
