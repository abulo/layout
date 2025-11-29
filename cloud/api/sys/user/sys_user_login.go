package user

import (
	"cloud/api/common"
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/sys/user"
	"cloud/service/verify"
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

func Login(ctx context.Context, newCtx *app.RequestContext) {
	//1.验证参数必传
	var req dao.SysUserLogin
	if err := newCtx.BindAndValidate(&req); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	//2.判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:用户信息表:sys_user:SysUserLogin")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接验证码服务
	clientCaptcha := verify.NewVerifyServiceClient(grpcClient)
	requestCaptcha := &verify.VerifyRequest{}
	requestCaptcha.VerifyCode = req.VerifyCode
	requestCaptcha.VerifyCodeId = req.VerifyCodeId
	// 执行服务
	_, err = clientCaptcha.Verify(ctx, requestCaptcha)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestCaptcha,
			"err": err,
		}).Error("GrpcCall:用户信息表:sys_user:SysUserLogin")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	// if resCaptcha.GetCode() != code.Success {
	// 	response.JSON(newCtx, consts.StatusOK, utils.H{
	// 		"code": resCaptcha.GetCode(),
	// 		"msg":  resCaptcha.GetMsg(),
	// 	})
	// 	return
	// }
	// if !resCaptcha.GetData().GetResult() {
	// 	response.JSON(newCtx, consts.StatusOK, utils.H{
	// 		"code": code.VerifyCodeError,
	// 		"msg":  code.StatusText(code.VerifyCodeError),
	// 	})
	// 	return
	// }

	userTokenItem, err := common.SysUserLogin(ctx, newCtx, req, true)
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

func SysUserLogin(ctx context.Context, newCtx *app.RequestContext) {
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
	userId := cast.ToInt64(newCtx.Param("id"))
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
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId"))
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
