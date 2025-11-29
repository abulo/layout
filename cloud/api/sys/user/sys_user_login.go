package user

import (
	"cloud/api/common"
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/verify"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

func SysUserLogin(ctx context.Context, newCtx *app.RequestContext) {
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
