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

// sys_logger_login 登录日志
// SysLoggerLoginItem 查询单条数据
func SysLoggerLoginItem(ctx context.Context, newCtx *app.RequestContext, id int64) (*logger.SysLoggerLoginResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:登录日志:sys_logger_login:SysLoggerLoginItem")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := logger.NewSysLoggerLoginServiceClient(grpcClient)
	request := &logger.SysLoggerLoginRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysLoggerLogin(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:sys_logger_login:SysLoggerLoginItem")
		return nil, err
	}
	return res, nil
}

// SysLoggerLoginDelete 删除数据
func SysLoggerLoginDelete(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysLoggerLoginItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:登录日志:sys_logger_login:SysLoggerLoginDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSysLoggerLoginServiceClient(grpcClient)
	request := &logger.SysLoggerLoginDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysLoggerLoginDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:sys_logger_login:SysLoggerLoginDelete")
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

// SysLoggerLogin 查询单条数据
func SysLoggerLogin(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	// 执行服务
	res, err := SysLoggerLoginItem(ctx, newCtx, id)
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
		"data": logger.SysLoggerLoginDao(res.GetData()),
	})
}

// SysLoggerLoginRecover 恢复数据
func SysLoggerLoginRecover(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysLoggerLoginItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:登录日志:sys_logger_login:SysLoggerLoginRecover")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSysLoggerLoginServiceClient(grpcClient)
	request := &logger.SysLoggerLoginRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysLoggerLoginRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:sys_logger_login:SysLoggerLoginRecover")
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

// SysLoggerLoginDrop 清理数据
func SysLoggerLoginDrop(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysLoggerLoginItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:登录日志:sys_logger_login:SysLoggerLoginDrop")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := logger.NewSysLoggerLoginServiceClient(grpcClient)
	request := &logger.SysLoggerLoginDropRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysLoggerLoginDrop(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:sys_logger_login:SysLoggerLoginDrop")
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

// SysLoggerLoginList 列表数据
func SysLoggerLoginList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:登录日志:sys_logger_login:SysLoggerLoginList")
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
	client := logger.NewSysLoggerLoginServiceClient(grpcClient)
	// 构造查询条件
	request := &logger.SysLoggerLoginListRequest{}
	requestTotal := &logger.SysLoggerLoginListTotalRequest{}

	if val, ok := newCtx.GetQuery("channel"); ok {
		request.Channel = proto.String(val)      // 渠道
		requestTotal.Channel = proto.String(val) // 渠道
	}
	if val, ok := newCtx.GetQuery("beginLoginTime"); ok {
		request.BeginLoginTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 登录时间
		requestTotal.BeginLoginTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 登录时间
	}
	if val, ok := newCtx.GetQuery("finishLoginTime"); ok {
		request.FinishLoginTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 登录时间
		requestTotal.FinishLoginTime = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 登录时间
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
	resTotal, err := client.SysLoggerLoginListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:sys_logger_login:SysLoggerLoginList")
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
	res, err := client.SysLoggerLoginList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:登录日志:sys_logger_login:SysLoggerLoginList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SysLoggerLogin
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, logger.SysLoggerLoginDao(item))
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
