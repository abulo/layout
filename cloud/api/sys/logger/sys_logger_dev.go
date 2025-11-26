package logger

import (
	"context"
	"time"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pagination"
	"cloud/service/sys/logger"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// sys_logger_dev 开发日志
// SysLoggerDevItem 查询单条数据
func SysLoggerDevItem(ctx context.Context, newCtx *app.RequestContext, id int64) (*logger.SysLoggerDevResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:开发日志:sys_logger_dev:SysLoggerDevItem")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := logger.NewSysLoggerDevServiceClient(grpcClient)
	request := &logger.SysLoggerDevRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysLoggerDev(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:开发日志:sys_logger_dev:SysLoggerDevItem")
		return nil, err
	}
	return res, nil
}

// SysLoggerDevDelete 删除数据
func SysLoggerDevDelete(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysLoggerDevItem(ctx, newCtx, id); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:开发日志:sys_logger_dev:SysLoggerDevDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSysLoggerDevServiceClient(grpcClient)
	request := &logger.SysLoggerDevDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysLoggerDevDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:开发日志:sys_logger_dev:SysLoggerDevDelete")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SysLoggerDev 查询单条数据
func SysLoggerDev(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	// 执行服务
	res, err := SysLoggerDevItem(ctx, newCtx, id)
	if err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": logger.SysLoggerDevDao(res.GetData()),
	})
}

// SysLoggerDevList 列表数据
func SysLoggerDevList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:开发日志:sys_logger_dev:SysLoggerDevList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSysLoggerDevServiceClient(grpcClient)
	// 构造查询条件
	request := &logger.SysLoggerDevListRequest{}
	requestTotal := &logger.SysLoggerDevListTotalRequest{}

	if val, ok := newCtx.GetQuery("beginTimestamp"); ok {
		request.BeginTimestamp = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 时间
		requestTotal.BeginTimestamp = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 时间
	}
	if val, ok := newCtx.GetQuery("finishTimestamp"); ok {
		request.FinishTimestamp = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 时间
		requestTotal.FinishTimestamp = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 时间
	}
	if val, ok := newCtx.GetQuery("host"); ok {
		request.Host = proto.String(val)      // 服务名
		requestTotal.Host = proto.String(val) // 服务名
	}
	if val, ok := newCtx.GetQuery("level"); ok {
		request.Level = proto.String(val)      // 等级
		requestTotal.Level = proto.String(val) // 等级
	}

	// 执行服务,获取数据量
	resTotal, err := client.SysLoggerDevListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:开发日志:sys_logger_dev:SysLoggerDevList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var total int64
	paginationRequest := &pagination.PaginationRequest{}
	paginationRequest.PageNum = proto.Int64(cast.ToInt64(newCtx.Query("pageNum")))
	paginationRequest.PageSize = proto.Int64(cast.ToInt64(newCtx.Query("pageSize")))
	request.Pagination = paginationRequest
	if resTotal.GetCode() == code.Success {
		total = resTotal.GetData()
	}
	// 执行服务
	res, err := client.SysLoggerDevList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:开发日志:sys_logger_dev:SysLoggerDevList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SysLoggerDev
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, logger.SysLoggerDevDao(item))
		}
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": utils.H{
			"total":    total,
			"list":     list,
			"pageNum":  paginationRequest.PageNum,
			"pageSize": paginationRequest.PageSize,
		},
	})
}
