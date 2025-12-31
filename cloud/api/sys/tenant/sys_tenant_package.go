package tenant

import (
	"context"

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
)

// sys_tenant_package 租户套餐
// SysTenantPackageItem 查询单条数据
func SysTenantPackageItem(ctx context.Context, newCtx *app.RequestContext, id int64) (*tenant.SysTenantPackageResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户套餐:sys_tenant_package:SysTenantPackageItem")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := tenant.NewSysTenantPackageServiceClient(grpcClient)
	request := &tenant.SysTenantPackageRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysTenantPackage(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户套餐:sys_tenant_package:SysTenantPackageItem")
		return nil, err
	}
	return res, nil
}

// SysTenantPackageCreate 创建数据
func SysTenantPackageCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户套餐:sys_tenant_package:SysTenantPackageCreate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSysTenantPackageServiceClient(grpcClient)
	request := &tenant.SysTenantPackageCreateRequest{}
	// 数据绑定
	var reqInfo dao.SysTenantPackage
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
	request.Data = tenant.SysTenantPackageProto(reqInfo)
	// 执行服务
	res, err := client.SysTenantPackageCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户套餐:sys_tenant_package:SysTenantPackageCreate")
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

// SysTenantPackageUpdate 更新数据
func SysTenantPackageUpdate(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysTenantPackageItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:租户套餐:sys_tenant_package:SysTenantPackageUpdate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSysTenantPackageServiceClient(grpcClient)
	request := &tenant.SysTenantPackageUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SysTenantPackage
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
	request.Data = tenant.SysTenantPackageProto(reqInfo)
	// 执行服务
	res, err := client.SysTenantPackageUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户套餐:sys_tenant_package:SysTenantPackageUpdate")
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

// SysTenantPackageDelete 删除数据
func SysTenantPackageDelete(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysTenantPackageItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:租户套餐:sys_tenant_package:SysTenantPackageDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSysTenantPackageServiceClient(grpcClient)
	request := &tenant.SysTenantPackageDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysTenantPackageDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户套餐:sys_tenant_package:SysTenantPackageDelete")
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

// SysTenantPackage 查询单条数据
func SysTenantPackage(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	// 执行服务
	res, err := SysTenantPackageItem(ctx, newCtx, id)
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
		"data": tenant.SysTenantPackageDao(res.GetData()),
	})
}

// SysTenantPackageList 列表数据
func SysTenantPackageList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户套餐:sys_tenant_package:SysTenantPackageList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSysTenantPackageServiceClient(grpcClient)
	// 构造查询条件
	request := &tenant.SysTenantPackageListRequest{}
	requestTotal := &tenant.SysTenantPackageListTotalRequest{}

	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 状态:0正常/1停用
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 状态:0正常/1停用
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val)      // 名称
		requestTotal.Name = proto.String(val) // 名称
	}

	// 执行服务,获取数据量
	resTotal, err := client.SysTenantPackageListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户套餐:sys_tenant_package:SysTenantPackageList")
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
	res, err := client.SysTenantPackageList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户套餐:sys_tenant_package:SysTenantPackageList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SysTenantPackage
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, tenant.SysTenantPackageDao(item))
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

// SysTenantPackageListSimple 列表精简数据
func SysTenantPackageListSimple(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:租户套餐:sys_tenant_package:SysTenantPackageListSimple")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := tenant.NewSysTenantPackageServiceClient(grpcClient)
	// 构造查询条件
	request := &tenant.SysTenantPackageListRequest{}

	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val)) // 状态:0正常/1停用
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val) // 名称
	}
	// 执行服务
	res, err := client.SysTenantPackageList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:租户套餐:sys_tenant_package:SysTenantPackageListSimple")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SysTenantPackage
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, tenant.SysTenantPackageDao(item))
		}
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": list,
	})
}
