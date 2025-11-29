package tenant

import (
	"cloud/api/common"
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/sys/tenant"
	"cloud/service/sys/user"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func SysTenantUserLogin(ctx context.Context, newCtx *app.RequestContext) {
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
	userId := cast.ToInt64(tenantItem.UserId)
	userRes, err := common.SysUserItem(ctx, newCtx, userId)

	if err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	reqInfo := user.SysUserDao(userRes.GetData())
	reqInfo.Password = nil
	//链接服务
	client := user.NewSysUserServiceClient(grpcClient)
	request := &user.SysUserUpdateRequest{}
	request.Id = userId
	// 数据绑定
	reqInfo.TenantId = proto.Int64(id)
	request.Data = user.SysUserProto(*reqInfo)
	// 执行服务
	_, err = client.SysUserUpdate(ctx, request)
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
	reqLogin := dao.SysUserLogin{
		Username: *reqInfo.Username,
	}
	userTokenItem, err := common.SysUserLogin(ctx, newCtx, reqLogin, false)

	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.LoginFail,
			"msg":  err.Error(),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": code.Success,
		"msg":  code.StatusText(code.Success),
		"data": userTokenItem,
	})
}
