package logger

import (
	"context"
	"encoding/json"
	"time"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pagination"
	"cloud/service/sys/logger"
	"cloud/service/sys/tenant"
	"cloud/service/sys/user"

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

// sys_logger_operate 操作日志
// SysLoggerOperateItem 查询单条数据
func SysLoggerOperateItem(ctx context.Context, newCtx *app.RequestContext, id int64) (*logger.SysLoggerOperateResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:操作日志:sys_logger_operate:SysLoggerOperateItem")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := logger.NewSysLoggerOperateServiceClient(grpcClient)
	request := &logger.SysLoggerOperateRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysLoggerOperate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:sys_logger_operate:SysLoggerOperateItem")
		return nil, err
	}
	return res, nil
}

// SysLoggerOperateDelete 删除数据
func SysLoggerOperateDelete(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysLoggerOperateItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:操作日志:sys_logger_operate:SysLoggerOperateDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSysLoggerOperateServiceClient(grpcClient)
	request := &logger.SysLoggerOperateDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysLoggerOperateDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:sys_logger_operate:SysLoggerOperateDelete")
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

// SysLoggerOperate 查询单条数据
func SysLoggerOperate(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	// 执行服务
	res, err := SysLoggerOperateItem(ctx, newCtx, id)
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
		"data": logger.SysLoggerOperateDao(res.GetData()),
	})
}

// SysLoggerOperateRecover 恢复数据
func SysLoggerOperateRecover(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysLoggerOperateItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:操作日志:sys_logger_operate:SysLoggerOperateRecover")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSysLoggerOperateServiceClient(grpcClient)
	request := &logger.SysLoggerOperateRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysLoggerOperateRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:sys_logger_operate:SysLoggerOperateRecover")
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

// SysLoggerOperateDrop 清理数据
func SysLoggerOperateDrop(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysLoggerOperateItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:操作日志:sys_logger_operate:SysLoggerOperateDrop")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSysLoggerOperateServiceClient(grpcClient)
	request := &logger.SysLoggerOperateDropRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysLoggerOperateDrop(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:sys_logger_operate:SysLoggerOperateDrop")
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

// SysLoggerOperateList 列表数据
func SysLoggerOperateList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:操作日志:sys_logger_operate:SysLoggerOperateList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}

	clientTenant := tenant.NewSysTenantServiceClient(grpcClient)
	requestTenant := &tenant.SysTenantRequest{}
	requestTenant.Id = newCtx.GetInt64("tenantId")
	// 执行服务
	resTenant, err := clientTenant.SysTenant(ctx, requestTenant)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestTenant,
			"err": err,
		}).Error("GrpcCall:租户:sys_tenant:SysTenant")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	tenantItem := tenant.SysTenantDao(resTenant.GetData())

	clientUser := user.NewSysUserServiceClient(grpcClient)
	scopeReq := &user.SysUserScopeRequest{}
	scopeReq.UserId = *tenantItem.UserId.Ptr()
	scopeReq.TenantId = newCtx.GetInt64("tenantId")
	userRes, err := clientUser.SysUserScope(ctx, scopeReq)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": scopeReq,
			"err": err,
		}).Error("GrpcCall:部门:sys_dept:SysDeptList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	userScope := user.SysUserScopeDao(userRes.GetData())
	if len(userScope.ScopeDept) < 1 {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.DeptError,
			"msg":  code.StatusText(code.DeptError),
		})
		return
	}

	//链接服务
	client := logger.NewSysLoggerOperateServiceClient(grpcClient)
	// 构造查询条件
	request := &logger.SysLoggerOperateListRequest{}
	requestTotal := &logger.SysLoggerOperateListTotalRequest{}

	if val, ok := newCtx.GetQuery("result"); ok {
		request.Result = proto.Int32(cast.ToInt32(val))      // 结果:0 成功/1 失败
		requestTotal.Result = proto.Int32(cast.ToInt32(val)) // 结果:0 成功/1 失败
	}
	if val, ok := newCtx.GetQuery("channel"); ok {
		request.Channel = proto.String(val)      // 渠道
		requestTotal.Channel = proto.String(val) // 渠道
	}
	if val, ok := newCtx.GetQuery("method"); ok {
		request.Method = proto.String(val)      // 请求方法
		requestTotal.Method = proto.String(val) // 请求方法
	}
	if val, ok := newCtx.GetQuery("beginStartTime"); ok {
		request.BeginStartTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 开始时间
		requestTotal.BeginStartTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 开始时间
	}
	if val, ok := newCtx.GetQuery("finishStartTime"); ok {
		request.FinishStartTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 结束时间
		requestTotal.FinishStartTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 结束时间
	}
	if val, ok := newCtx.GetQuery("username"); ok {
		request.Username = proto.String(val)      // 用户名
		requestTotal.Username = proto.String(val) // 用户名
	}

	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId"))      // 租户ID
	requestTotal.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                                 // 删除状态
	requestTotal.Deleted = proto.Int32(0)                            // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = proto.Int32(1)
			requestTotal.Deleted = proto.Int32(1)
		}
	}
	request.UserId = tenantItem.UserId.Ptr()      // 用户ID
	requestTotal.UserId = tenantItem.UserId.Ptr() // 用户ID
	request.Scope = userScope.Scope
	requestTotal.Scope = userScope.Scope
	scopeDept, _ := json.Marshal(userScope.ScopeDept)
	request.ScopeDept = scopeDept
	requestTotal.ScopeDept = scopeDept
	// 执行服务,获取数据量
	resTotal, err := client.SysLoggerOperateListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:sys_logger_operate:SysLoggerOperateList")
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
	res, err := client.SysLoggerOperateList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:操作日志:sys_logger_operate:SysLoggerOperateList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SysLoggerOperate
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, logger.SysLoggerOperateDao(item))
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
