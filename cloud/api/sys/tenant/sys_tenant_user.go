package tenant

import (
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/pagination"
	"cloud/service/sys/tenant"
	"cloud/service/sys/user"
	"context"
	"encoding/json"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// SysTenantUserList 列表数据
func SysTenantUserList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:系统用户:sys_user:SysTenantUserList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	id := cast.ToInt64(newCtx.Param("id"))

	clientTenant := tenant.NewSysTenantServiceClient(grpcClient)
	requestTenant := &tenant.SysTenantRequest{}
	requestTenant.Id = id
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
	scopeReq.TenantId = id
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
	request.TenantId = proto.Int64(id)      // 租户ID
	requestTotal.TenantId = proto.Int64(id) // 租户ID
	request.Deleted = proto.Int32(0)        // 删除状态
	requestTotal.Deleted = proto.Int32(0)   // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
			requestTotal.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val))      // 用户状态（0正常 1停用）
		requestTotal.Status = proto.Int32(cast.ToInt32(val)) // 用户状态（0正常 1停用）
	}
	if val, ok := newCtx.GetQuery("deptId"); ok {
		request.DeptId = proto.Int64(cast.ToInt64(val))      // 用户状态（0正常 1停用）
		requestTotal.DeptId = proto.Int64(cast.ToInt64(val)) // 用户状态（0正常 1停用）
	}
	if val, ok := newCtx.GetQuery("username"); ok {
		request.Username = proto.String(val)      // 用户名
		requestTotal.Username = proto.String(val) // 用户名
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
			"req": requestTotal,
			"err": err,
		}).Error("GrpcCall:系统用户:sys_user:SysTenantUserList")
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
	if resTotal.GetCode() == code.Success {
		total = resTotal.GetData()
	}
	// 执行服务
	res, err := client.SysUserList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:系统用户:sys_user:SysTenantUserList")
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
			userInfo := user.SysUserDao(item)
			userInfo.Password = nil
			list = append(list, userInfo)
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
