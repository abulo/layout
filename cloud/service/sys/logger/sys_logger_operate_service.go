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

// sys_logger_operate 操作日志

// SrvSysLoggerOperateServiceServer 操作日志
type SrvSysLoggerOperateServiceServer struct {
	UnimplementedSysLoggerOperateServiceServer
	Server *xgrpc.Server
}

// SysLoggerOperateCreate 创建数据
func (srv SrvSysLoggerOperateServiceServer) SysLoggerOperateCreate(ctx context.Context, request *SysLoggerOperateCreateRequest) (*SysLoggerOperateCreateResponse, error) {
	req := SysLoggerOperateDao(request.GetData())
	data, err := logger.SysLoggerOperateCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:操作日志:sys_logger_operate:SysLoggerOperateCreate")
		return &SysLoggerOperateCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerOperateCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysLoggerOperateUpdate 更新数据
func (srv SrvSysLoggerOperateServiceServer) SysLoggerOperateUpdate(ctx context.Context, request *SysLoggerOperateUpdateRequest) (*SysLoggerOperateUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysLoggerOperateUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysLoggerOperateDao(request.GetData())
	_, err := logger.SysLoggerOperateUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:操作日志:sys_logger_operate:SysLoggerOperateUpdate")
		return &SysLoggerOperateUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerOperateUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysLoggerOperateDelete 删除数据
func (srv SrvSysLoggerOperateServiceServer) SysLoggerOperateDelete(ctx context.Context, request *SysLoggerOperateDeleteRequest) (*SysLoggerOperateDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysLoggerOperateDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := logger.SysLoggerOperateDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:操作日志:sys_logger_operate:SysLoggerOperateDelete")
		return &SysLoggerOperateDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerOperateDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysLoggerOperate 查询单条数据
func (srv SrvSysLoggerOperateServiceServer) SysLoggerOperate(ctx context.Context, request *SysLoggerOperateRequest) (*SysLoggerOperateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysLoggerOperateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := logger.SysLoggerOperate(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:操作日志:sys_logger_operate:SysLoggerOperate")
		return &SysLoggerOperateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerOperateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysLoggerOperateProto(res),
	}, nil
}

// SysLoggerOperateRecover 恢复数据
func (srv SrvSysLoggerOperateServiceServer) SysLoggerOperateRecover(ctx context.Context, request *SysLoggerOperateRecoverRequest) (*SysLoggerOperateRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysLoggerOperateRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := logger.SysLoggerOperateRecover(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:操作日志:sys_logger_operate:SysLoggerOperateRecover")
		return &SysLoggerOperateRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerOperateRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysLoggerOperateDrop 清理数据
func (srv SrvSysLoggerOperateServiceServer) SysLoggerOperateDrop(ctx context.Context, request *SysLoggerOperateDropRequest) (*SysLoggerOperateDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysLoggerOperateDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := logger.SysLoggerOperateDrop(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:操作日志:sys_logger_operate:SysLoggerOperateDrop")
		return &SysLoggerOperateDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerOperateDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysLoggerOperateList 列表数据
func (srv SrvSysLoggerOperateServiceServer) SysLoggerOperateList(ctx context.Context, request *SysLoggerOperateListRequest) (*SysLoggerOperateListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Result != nil {
		condition["result"] = request.GetResult()
	}
	if request.Channel != nil {
		condition["channel"] = request.GetChannel()
	}
	if request.Method != nil {
		condition["method"] = request.GetMethod()
	}
	if request.StartTime != nil {
		condition["startTime"] = request.GetStartTime()
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
	list, err := logger.SysLoggerOperateList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:操作日志:sys_logger_operate:SysLoggerOperateList")
		return &SysLoggerOperateListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysLoggerOperateObject
	for _, item := range list {
		res = append(res, SysLoggerOperateProto(item))
	}
	return &SysLoggerOperateListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysLoggerOperateListTotal 获取总数
func (srv SrvSysLoggerOperateServiceServer) SysLoggerOperateListTotal(ctx context.Context, request *SysLoggerOperateListTotalRequest) (*SysLoggerOperateTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Result != nil {
		condition["result"] = request.GetResult()
	}
	if request.Channel != nil {
		condition["channel"] = request.GetChannel()
	}
	if request.Method != nil {
		condition["method"] = request.GetMethod()
	}
	if request.StartTime != nil {
		condition["startTime"] = request.GetStartTime()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.Username != nil {
		condition["username"] = request.GetUsername()
	}

	// 获取数据集合
	total, err := logger.SysLoggerOperateListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:操作日志:sys_logger_operate:SysLoggerOperateListTotal")
		return &SysLoggerOperateTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysLoggerOperateTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
