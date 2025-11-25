package user

import (
	"cloud/api/common"
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/internal/token"
	"cloud/service/sys/user"
	"cloud/service/verify"
	"context"
	"encoding/json"
	"time"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
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

	//3.链接用户信息表服务
	clientSysUser := user.NewSysUserServiceClient(grpcClient)
	requestSysUser := &user.SysUserLoginRequest{}
	requestSysUser.Username = proto.String(req.Username)
	// 执行服务
	res, err := clientSysUser.SysUserLogin(ctx, requestSysUser)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestSysUser,
			"err": err,
		}).Error("GrpcCall:用户信息表:sys_user:SysUserLogin")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	userInfo := res.GetData()
	if userInfo.Id == nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.LoginFail,
			"msg":  code.StatusText(code.LoginFail),
		})
		return
	}
	// 比对密码
	if req.Password != userInfo.GetPassword() {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.LoginFail,
			"msg":  code.StatusText(code.LoginFail),
		})
		return
	}
	// 用户信息
	userItem := user.SysUserDao(userInfo)
	// 去判断这个用户有没有绑定租户
	userTenantClient := user.NewSysUserTenantServiceClient(grpcClient)
	requestUserTenant := &user.SysUserTenantBindRequest{}
	requestUserTenant.UserId = userItem.Id
	requestUserTenant.TenantId = userItem.TenantId
	resUserTenant, err := userTenantClient.SysUserTenantBind(ctx, requestUserTenant)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestUserTenant,
			"err": err,
		}).Error("GrpcCall:用户信息表:sys_user:SysUserLogin")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	userTenantInfo := resUserTenant.GetData()
	if userTenantInfo.Id == nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.LoginFail,
			"msg":  code.StatusText(code.LoginFail),
		})
		return
	}
	// 判断 userInfo.TenantId 是否和userTenantInfo.TenantId一致
	if cast.ToInt64(userInfo.TenantId) != cast.ToInt64(userTenantInfo.TenantId) {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.LoginFail,
			"msg":  code.StatusText(code.LoginFail),
		})
		return
	}

	userVerifyItem := dao.UserVerify{}
	userVerifyItem.UserId = cast.ToInt64(userItem.Id)          // 用户ID
	userVerifyItem.UserName = cast.ToString(userItem.Username) // 用户名
	userVerifyItem.TenantId = cast.ToInt64(userItem.TenantId)  // 租户ID
	userVerifyItem.Name = cast.ToString(userItem.Name.Ptr())   // 用户姓名
	userTokenItem := dao.UserToken{}
	userTokenItem.UserName = proto.String(cast.ToString(userItem.Username)) // 用户名
	dateTime := util.Now().Add(time.Duration(86000*30) * time.Second)
	accessToken, err := token.GenerateToken(userVerifyItem, "user", 86400*30)
	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.TokenError,
			"msg":  code.StatusText(code.TokenError),
		})
		return
	}
	userTokenItem.AccessToken = proto.String(accessToken)
	refreshToken, err := token.GenerateToken(userVerifyItem, "user", 86400*31)
	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.TokenError,
			"msg":  code.StatusText(code.TokenError),
		})
		return
	}
	userTokenItem.RefreshToken = proto.String(refreshToken)
	userTokenItem.Expires = proto.String(util.Date("Y-m-d H:i:s", dateTime))
	list, currentMenuIds, err := common.SysUserTree(ctx, cast.ToInt64(userItem.Id), cast.ToInt64(userItem.TenantId))
	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.TokenError,
			"msg":  code.StatusText(code.TokenError),
		})
		return
	}

	var permissionList []string
	for _, item := range list {
		if util.InArray(item.Id, currentMenuIds) {
			if !util.Empty(item.Code) {
				permissionList = append(permissionList, item.Code)
			}
		}
	}
	redisHandler := initial.Core.Store.LoadRedis("redis")
	keyMenu := util.NewReplacer(initial.Core.Config.String("Cache.SysUser.Code"), ":UserId", userInfo.GetId())
	permission, _ := json.Marshal(permissionList)
	redisHandler.Set(ctx, keyMenu, cast.ToString(permission), time.Duration(time.Duration(86000*30))*time.Second)
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": code.Success,
		"msg":  code.StatusText(code.Success),
		"data": userTokenItem,
	})
}
