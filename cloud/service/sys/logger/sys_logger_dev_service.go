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

// sys_logger_dev 开发日志

// SrvSysLoggerDevServiceServer 开发日志
type SrvSysLoggerDevServiceServer struct {
	UnimplementedSysLoggerDevServiceServer
	Server *xgrpc.Server
}

// SysLoggerDevCreate 创建数据
func (srv SrvSysLoggerDevServiceServer) SysLoggerDevCreate(ctx context.Context, request *SysLoggerDevCreateRequest) (*SysLoggerDevCreateResponse, error) {
	req := SysLoggerDevDao(request.GetData())
	data, err := logger.SysLoggerDevCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:开发日志:sys_logger_dev:SysLoggerDevCreate")
		return &SysLoggerDevCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerDevCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysLoggerDevUpdate 更新数据
func (srv SrvSysLoggerDevServiceServer) SysLoggerDevUpdate(ctx context.Context, request *SysLoggerDevUpdateRequest) (*SysLoggerDevUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysLoggerDevUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysLoggerDevDao(request.GetData())
	_, err := logger.SysLoggerDevUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:开发日志:sys_logger_dev:SysLoggerDevUpdate")
		return &SysLoggerDevUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerDevUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysLoggerDevDelete 删除数据
func (srv SrvSysLoggerDevServiceServer) SysLoggerDevDelete(ctx context.Context, request *SysLoggerDevDeleteRequest) (*SysLoggerDevDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysLoggerDevDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := logger.SysLoggerDevDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:开发日志:sys_logger_dev:SysLoggerDevDelete")
		return &SysLoggerDevDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerDevDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysLoggerDev 查询单条数据
func (srv SrvSysLoggerDevServiceServer) SysLoggerDev(ctx context.Context, request *SysLoggerDevRequest) (*SysLoggerDevResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysLoggerDevResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := logger.SysLoggerDev(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:开发日志:sys_logger_dev:SysLoggerDev")
		return &SysLoggerDevResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerDevResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysLoggerDevProto(res),
	}, nil
}

// SysLoggerDevList 列表数据
func (srv SrvSysLoggerDevServiceServer) SysLoggerDevList(ctx context.Context, request *SysLoggerDevListRequest) (*SysLoggerDevListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Timestamp != nil {
		condition["timestamp"] = request.GetTimestamp()
	}
	if request.Host != nil {
		condition["host"] = request.GetHost()
	}
	if request.Level != nil {
		condition["level"] = request.GetLevel()
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
	list, err := logger.SysLoggerDevList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:开发日志:sys_logger_dev:SysLoggerDevList")
		return &SysLoggerDevListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysLoggerDevObject
	for _, item := range list {
		res = append(res, SysLoggerDevProto(item))
	}
	return &SysLoggerDevListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysLoggerDevListTotal 获取总数
func (srv SrvSysLoggerDevServiceServer) SysLoggerDevListTotal(ctx context.Context, request *SysLoggerDevListTotalRequest) (*SysLoggerDevTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Timestamp != nil {
		condition["timestamp"] = request.GetTimestamp()
	}
	if request.Host != nil {
		condition["host"] = request.GetHost()
	}
	if request.Level != nil {
		condition["level"] = request.GetLevel()
	}

	// 获取数据集合
	total, err := logger.SysLoggerDevListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:开发日志:sys_logger_dev:SysLoggerDevListTotal")
		return &SysLoggerDevTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerDevTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
