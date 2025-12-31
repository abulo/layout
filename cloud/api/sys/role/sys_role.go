package role

import (
	"context"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pagination"
	"cloud/service/sys/role"

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

// sys_role 角色
// SysRoleItem 查询单条数据
func SysRoleItem(ctx context.Context, newCtx *app.RequestContext, id int64) (*role.SysRoleResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:角色:sys_role:SysRoleItem")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := role.NewSysRoleServiceClient(grpcClient)
	request := &role.SysRoleRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysRole(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:角色:sys_role:SysRoleItem")
		return nil, err
	}
	return res, nil
}

// SysRoleCreate 创建数据
func SysRoleCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:角色:sys_role:SysRoleCreate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := role.NewSysRoleServiceClient(grpcClient)
	request := &role.SysRoleCreateRequest{}
	// 数据绑定
	var reqInfo dao.SysRole
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Id = nil
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId"))
	reqInfo.Creator = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.CreateTime = null.TimeStampFrom(util.Now())
	request.Data = role.SysRoleProto(reqInfo)
	// 执行服务
	res, err := client.SysRoleCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:角色:sys_role:SysRoleCreate")
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

// SysRoleUpdate 更新数据
func SysRoleUpdate(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysRoleItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:角色:sys_role:SysRoleUpdate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := role.NewSysRoleServiceClient(grpcClient)
	request := &role.SysRoleUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SysRole
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Id = nil
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId"))
	reqInfo.Updater = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.UpdateTime = null.TimeStampFrom(util.Now())
	reqInfo.Creator = null.StringFromPtr(nil)
	reqInfo.CreateTime = null.TimeStampFromPtr(nil)
	request.Data = role.SysRoleProto(reqInfo)
	// 执行服务
	res, err := client.SysRoleUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:角色:sys_role:SysRoleUpdate")
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

// SysRoleDelete 删除数据
func SysRoleDelete(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysRoleItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:角色:sys_role:SysRoleDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := role.NewSysRoleServiceClient(grpcClient)
	request := &role.SysRoleDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysRoleDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:角色:sys_role:SysRoleDelete")
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

// SysRole 查询单条数据
func SysRole(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	// 执行服务
	res, err := SysRoleItem(ctx, newCtx, id)
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
		"data": role.SysRoleDao(res.GetData()),
	})
}

// SysRoleRecover 恢复数据
func SysRoleRecover(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysRoleItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:角色:sys_role:SysRoleRecover")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := role.NewSysRoleServiceClient(grpcClient)
	request := &role.SysRoleRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysRoleRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:角色:sys_role:SysRoleRecover")
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

// SysRoleDrop 清理数据
func SysRoleDrop(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysRoleItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:角色:sys_role:SysRoleDrop")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := role.NewSysRoleServiceClient(grpcClient)
	request := &role.SysRoleDropRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysRoleDrop(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:角色:sys_role:SysRoleDrop")
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

// SysRoleList 列表数据
func SysRoleList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:角色:sys_role:SysRoleList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := role.NewSysRoleServiceClient(grpcClient)
	// 构造查询条件
	request := &role.SysRoleListRequest{}
	requestTotal := &role.SysRoleListTotalRequest{}

	if val, ok := newCtx.GetQuery("code"); ok {
		request.Code = proto.String(val)      // 编码
		requestTotal.Code = proto.String(val) // 编码
	}
	if val, ok := newCtx.GetQuery("deleted"); ok {
		request.Deleted = proto.Int32(cast.ToInt32(val))      // 删除:0否/1是
		requestTotal.Deleted = proto.Int32(cast.ToInt32(val)) // 删除:0否/1是
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 状态:0正常/1停用
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 状态:0正常/1停用
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val)      // 名称
		requestTotal.Name = proto.String(val) // 名称
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
	// 执行服务,获取数据量
	resTotal, err := client.SysRoleListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:角色:sys_role:SysRoleList")
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
	res, err := client.SysRoleList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:角色:sys_role:SysRoleList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SysRole
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, role.SysRoleDao(item))
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

// SysRoleListSimple 列表精简数据
func SysRoleListSimple(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:角色:sys_role:SysRoleListSimple")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := role.NewSysRoleServiceClient(grpcClient)
	// 构造查询条件
	request := &role.SysRoleListRequest{}

	if val, ok := newCtx.GetQuery("code"); ok {
		request.Code = proto.String(val) // 编码
	}

	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val)) // 状态:0正常/1停用
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val) // 名称
	}

	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                            // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = proto.Int32(1)
		}
	}
	// 执行服务
	res, err := client.SysRoleList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:角色:sys_role:SysRoleListSimple")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SysRole
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, role.SysRoleDao(item))
		}
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": list,
	})
}
