package user

import (
	"context"
	"encoding/json"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pagination"
	"cloud/service/sys/tenant"
	"cloud/service/sys/user"

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

// sys_user 用户
// SysUserItem 查询单条数据
func SysUserItem(ctx context.Context, newCtx *app.RequestContext, id int64) (*user.SysUserResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:用户:sys_user:SysUserItem")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := user.NewSysUserServiceClient(grpcClient)
	request := &user.SysUserRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysUser(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:用户:sys_user:SysUserItem")
		return nil, err
	}
	return res, nil
}

// SysUserCreate 创建数据
func SysUserCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:用户:sys_user:SysUserCreate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSysUserServiceClient(grpcClient)
	request := &user.SysUserCreateRequest{}
	// 数据绑定
	var reqInfo dao.SysUser
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
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())
	request.Data = user.SysUserProto(reqInfo)
	// 执行服务
	res, err := client.SysUserCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:用户:sys_user:SysUserCreate")
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

// SysUserUpdate 更新数据
func SysUserUpdate(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysUserItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:用户:sys_user:SysUserUpdate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSysUserServiceClient(grpcClient)
	request := &user.SysUserUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SysUser
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
	reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
	reqInfo.Creator = null.StringFromPtr(nil)
	reqInfo.CreateTime = null.DateTimeFromPtr(nil)
	request.Data = user.SysUserProto(reqInfo)
	// 执行服务
	res, err := client.SysUserUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:用户:sys_user:SysUserUpdate")
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

// SysUserDelete 删除数据
func SysUserDelete(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysUserItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:用户:sys_user:SysUserDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSysUserServiceClient(grpcClient)
	request := &user.SysUserDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysUserDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:用户:sys_user:SysUserDelete")
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

// SysUser 查询单条数据
func SysUser(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	// 执行服务
	res, err := SysUserItem(ctx, newCtx, id)
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
		"data": user.SysUserDao(res.GetData()),
	})
}

// SysUserRecover 恢复数据
func SysUserRecover(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysUserItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:用户:sys_user:SysUserRecover")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSysUserServiceClient(grpcClient)
	request := &user.SysUserRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysUserRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:用户:sys_user:SysUserRecover")
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

// SysUserDrop 清理数据
func SysUserDrop(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysUserItem(ctx, newCtx, id); err != nil {
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
		}).Error("Grpc:用户:sys_user:SysUserDrop")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := user.NewSysUserServiceClient(grpcClient)
	request := &user.SysUserDropRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysUserDrop(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:用户:sys_user:SysUserDrop")
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

// SysUserList 列表数据
func SysUserList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:用户:sys_user:SysUserList")
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
	client := user.NewSysUserServiceClient(grpcClient)
	// 构造查询条件
	request := &user.SysUserListRequest{}
	requestTotal := &user.SysUserListTotalRequest{}
	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId"))      // 租户ID
	requestTotal.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                                 // 删除状态
	requestTotal.Deleted = proto.Int32(0)                            // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
			requestTotal.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("deptId"); ok {
		request.DeptId = proto.Int64(cast.ToInt64(val))      // 用户状态（0正常 1停用）
		requestTotal.DeptId = proto.Int64(cast.ToInt64(val)) // 用户状态（0正常 1停用）
	}
	if val, ok := newCtx.GetQuery("username"); ok {
		request.Username = proto.String(val)      // 用户名
		requestTotal.Username = proto.String(val) // 用户名
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 状态:0正常/1停用
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 状态:0正常/1停用
	}
	if val, ok := newCtx.GetQuery("mobile"); ok {
		request.Mobile = proto.String(val)      // 手机号码
		requestTotal.Mobile = proto.String(val) // 手机号码
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val)      // 姓名
		requestTotal.Name = proto.String(val) // 姓名
	}
	request.UserId = tenantItem.UserId.Ptr()      // 用户ID
	requestTotal.UserId = tenantItem.UserId.Ptr() // 用户ID
	request.Scope = userScope.Scope
	requestTotal.Scope = userScope.Scope
	scopeDept, _ := json.Marshal(userScope.ScopeDept)
	request.ScopeDept = scopeDept
	requestTotal.ScopeDept = scopeDept

	// 执行服务,获取数据量
	resTotal, err := client.SysUserListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:用户:sys_user:SysUserList")
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
	res, err := client.SysUserList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:用户:sys_user:SysUserList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SysUser
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, user.SysUserDao(item))
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

// SysUserListSimple 列表精简数据
func SysUserListSimple(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:用户:sys_user:SysUserListSimple")
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
	client := user.NewSysUserServiceClient(grpcClient)
	// 构造查询条件
	request := &user.SysUserListRequest{}
	requestTotal := &user.SysUserListTotalRequest{}

	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId"))      // 租户ID
	requestTotal.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                                 // 删除状态
	requestTotal.Deleted = proto.Int32(0)                            // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
			requestTotal.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("deptId"); ok {
		request.DeptId = proto.Int64(cast.ToInt64(val))      // 用户状态（0正常 1停用）
		requestTotal.DeptId = proto.Int64(cast.ToInt64(val)) // 用户状态（0正常 1停用）
	}
	if val, ok := newCtx.GetQuery("username"); ok {
		request.Username = proto.String(val)      // 用户名
		requestTotal.Username = proto.String(val) // 用户名
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 状态:0正常/1停用
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 状态:0正常/1停用
	}
	if val, ok := newCtx.GetQuery("mobile"); ok {
		request.Mobile = proto.String(val)      // 手机号码
		requestTotal.Mobile = proto.String(val) // 手机号码
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val)      // 姓名
		requestTotal.Name = proto.String(val) // 姓名
	}
	request.UserId = tenantItem.UserId.Ptr()      // 用户ID
	requestTotal.UserId = tenantItem.UserId.Ptr() // 用户ID
	request.Scope = userScope.Scope
	requestTotal.Scope = userScope.Scope
	scopeDept, _ := json.Marshal(userScope.ScopeDept)
	request.ScopeDept = scopeDept
	requestTotal.ScopeDept = scopeDept

	// 执行服务,获取数据量
	resTotal, err := client.SysUserListTotal(ctx, requestTotal)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:用户:sys_user:SysUserListSimple")
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
	res, err := client.SysUserList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:用户:sys_user:SysUserListSimple")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SysUser
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, user.SysUserDao(item))
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
