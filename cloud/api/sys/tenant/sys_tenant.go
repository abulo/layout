package tenant

import (
	"context"
	"time"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pagination"
	"cloud/service/sys/tenant"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// sys_tenant 租户
// SysTenantItem 查询单条数据
func SysTenantItem(ctx context.Context, newCtx *app.RequestContext, id int64) (*tenant.SysTenantResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:sys_tenant:SysTenantItem")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := tenant.NewSysTenantServiceClient(grpcClient)
	request := &tenant.SysTenantRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysTenant(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:sys_tenant:SysTenantItem")
		return nil, err
	}
	return res, nil
}

// SysTenantCreate 创建数据
func SysTenantCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:sys_tenant:SysTenantCreate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSysTenantServiceClient(grpcClient)
	request := &tenant.SysTenantCreateRequest{}
	// 数据绑定
	var reqInfo dao.SysTenant
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Id = nil
	reqInfo.Creator = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.CreateTime = null.TimeStampFrom(util.Now())
	request.Data = tenant.SysTenantProto(reqInfo)
	// 执行服务
	res, err := client.SysTenantCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:sys_tenant:SysTenantCreate")
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

// SysTenantUpdate 更新数据
func SysTenantUpdate(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysTenantItem(ctx, newCtx, id); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:sys_tenant:SysTenantUpdate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSysTenantServiceClient(grpcClient)
	request := &tenant.SysTenantUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SysTenant
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Id = nil
	reqInfo.Updater = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.UpdateTime = null.TimeStampFrom(util.Now())
	reqInfo.Creator = null.StringFromPtr(nil)
	reqInfo.CreateTime = null.TimeStampFromPtr(nil)
	request.Data = tenant.SysTenantProto(reqInfo)
	// 执行服务
	res, err := client.SysTenantUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:sys_tenant:SysTenantUpdate")
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

// SysTenantDelete 删除数据
func SysTenantDelete(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysTenantItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:租户:sys_tenant:SysTenantDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSysTenantServiceClient(grpcClient)
	request := &tenant.SysTenantDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysTenantDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:sys_tenant:SysTenantDelete")
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

// SysTenant 查询单条数据
func SysTenant(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	// 执行服务
	res, err := SysTenantItem(ctx, newCtx, id)
	if err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	tenantItem := tenant.SysTenantDao(res.GetData())
	tenantItem.Password = nil
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": tenantItem,
	})
}

// SysTenantRecover 恢复数据
func SysTenantRecover(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysTenantItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:租户:sys_tenant:SysTenantRecover")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSysTenantServiceClient(grpcClient)
	request := &tenant.SysTenantRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysTenantRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:sys_tenant:SysTenantRecover")
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

// SysTenantDrop 清理数据
func SysTenantDrop(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysTenantItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:租户:sys_tenant:SysTenantDrop")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSysTenantServiceClient(grpcClient)
	request := &tenant.SysTenantDropRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysTenantDrop(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:sys_tenant:SysTenantDrop")
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

// SysTenantList 列表数据
func SysTenantList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:sys_tenant:SysTenantList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSysTenantServiceClient(grpcClient)
	// 构造查询条件
	request := &tenant.SysTenantListRequest{}
	requestTotal := &tenant.SysTenantListTotalRequest{}

	request.Deleted = proto.Int32(0)      // 删除状态
	requestTotal.Deleted = proto.Int32(0) // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = proto.Int32(1)
			requestTotal.Deleted = proto.Int32(1)
		}
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 状态:0正常/1停用
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 状态:0正常/1停用
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val)      // 名称
		requestTotal.Name = proto.String(val) // 名称
	}
	if val, ok := newCtx.GetQuery("beginExpireDate"); ok {
		request.BeginExpireDate = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 过期时间
		requestTotal.BeginExpireDate = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 过期时间
	}
	if val, ok := newCtx.GetQuery("finishExpireDate"); ok {
		request.FinishExpireDate = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 过期时间
		requestTotal.FinishExpireDate = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 过期时间
	}
	if val, ok := newCtx.GetQuery("packageId"); ok {
		request.PackageId = proto.Int64(cast.ToInt64(val))      // 套餐编号
		requestTotal.PackageId = proto.Int64(cast.ToInt64(val)) // 套餐编号
	}

	// 执行服务,获取数据量
	resTotal, err := client.SysTenantListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:sys_tenant:SysTenantList")
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
	res, err := client.SysTenantList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:sys_tenant:SysTenantList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SysTenant
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			tenantItem := tenant.SysTenantDao(item)
			tenantItem.Password = nil
			list = append(list, tenantItem)
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

// SysTenantListSimple 列表精简数据
func SysTenantListSimple(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户:sys_tenant:SysTenantListSimple")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSysTenantServiceClient(grpcClient)
	// 构造查询条件
	request := &tenant.SysTenantListRequest{}
	requestTotal := &tenant.SysTenantListTotalRequest{}

	request.Deleted = proto.Int32(0)      // 删除状态
	requestTotal.Deleted = proto.Int32(0) // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = proto.Int32(1)
			requestTotal.Deleted = proto.Int32(1)
		}
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 状态:0正常/1停用
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 状态:0正常/1停用
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val)      // 名称
		requestTotal.Name = proto.String(val) // 名称
	}
	if val, ok := newCtx.GetQuery("beginExpireDate"); ok {
		request.BeginExpireDate = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 过期时间
		requestTotal.BeginExpireDate = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 过期时间
	}
	if val, ok := newCtx.GetQuery("finishExpireDate"); ok {
		request.FinishExpireDate = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local))      // 过期时间
		requestTotal.FinishExpireDate = timestamppb.New(cast.ToTimeInDefaultLocation(val, time.Local)) // 过期时间
	}
	if val, ok := newCtx.GetQuery("packageId"); ok {
		request.PackageId = proto.Int64(cast.ToInt64(val))      // 套餐编号
		requestTotal.PackageId = proto.Int64(cast.ToInt64(val)) // 套餐编号
	}

	// 执行服务,获取数据量
	resTotal, err := client.SysTenantListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:sys_tenant:SysTenantListSimple")
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
	res, err := client.SysTenantList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户:sys_tenant:SysTenantListSimple")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SysTenant
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			tenantItem := tenant.SysTenantDao(item)
			tenantItem.Password = nil
			list = append(list, tenantItem)
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
