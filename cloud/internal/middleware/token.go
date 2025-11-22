package middleware

import (
	"context"
	"encoding/json"
	"time"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/internal/token"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cast"
)

type Response struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func TokenMiddleware() app.HandlerFunc {
	return func(ctx context.Context, newCtx *app.RequestContext) {
		//加入开始时间
		startTime := util.Now()
		authHeader := newCtx.Request.Header.Get("X-Access-Token")
		if util.Empty(authHeader) {
			response.JSON(newCtx, consts.StatusUnauthorized, utils.H{
				"code": code.TokenEmptyError,
				"msg":  code.StatusText(code.TokenEmptyError),
			})
			newCtx.Abort()
			return
		}
		// 按空格分割
		parts := util.Explode(".", authHeader)
		if len(parts) != 3 {
			response.JSON(newCtx, consts.StatusUnauthorized, utils.H{
				"code": code.TokenInvalidError,
				"msg":  code.StatusText(code.TokenInvalidError),
			})
			newCtx.Abort()
			return
		}
		rsp, err := token.ParseToken(authHeader)
		if err != nil {
			response.JSON(newCtx, consts.StatusUnauthorized, utils.H{
				"code": code.TokenInvalidError,
				"msg":  code.StatusText(code.TokenInvalidError),
			})
			newCtx.Abort()
			return
		}
		channelList, _ := rsp.RegisteredClaims.Audience.MarshalJSON()
		var channel []string
		json.Unmarshal(channelList, &channel)
		if len(channel) < 1 {
			channel = append(channel, "unknown")
		}
		newCtx.Set("userId", rsp.UserId)                             // 用户ID
		newCtx.Set("userName", rsp.UserName)                         // 用户名
		newCtx.Set("tenantId", rsp.TenantId)                         // 租户
		newCtx.Set("name", rsp.Name)                                 // 用户姓名
		newCtx.Set("startTime", util.Date("Y-m-d H:i:s", startTime)) // 开始时间
		newCtx.Set("auth", true)                                     // 需要验证
		newCtx.Set("channel", channel[0])                            // 渠道
		newCtx.Next(ctx)
	}
}

func MissMiddleware() app.HandlerFunc {
	return func(ctx context.Context, newCtx *app.RequestContext) {
		newCtx.Set("auth", false) // 需要验证
		newCtx.Next(ctx)
	}
}

// AuthMiddleware 验证token
func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, newCtx *app.RequestContext) {
		//加入开始时间
		startTime := cast.ToTimeInDefaultLocation(newCtx.GetString("startTime"), time.Local) // 开始时间
		userId := newCtx.GetInt64("userId")                                                  // 用户ID
		userName := newCtx.GetString("userName")                                             // 用户名
		name := newCtx.GetString("name")                                                     // 用户姓名
		channel := newCtx.GetString("channel")                                               // 渠道
		tenantId := newCtx.GetInt64("tenantId")                                              // 租户编码
		// 判断是否需要验证

		handlerName := newCtx.HandlerName()
		method := util.Explode("/", handlerName)
		methodName := method[len(method)-1]
		redisHandler := initial.Core.Store.LoadRedis("redis")
		if newCtx.GetBool("auth") {
			// 检查用户有没有该权限
			menuKey := util.NewReplacer(initial.Core.Config.String("Cache.SysMenu.Code"))
			if exist, err := redisHandler.SIsMember(ctx, menuKey, methodName); err == nil {
				if !exist {
					newCtx.Next(ctx)
				} else {
					// 判断一下这个用户的的权限
					key := util.NewReplacer(initial.Core.Config.String("Cache.SysUser.Code"), ":UserId", userId)
					// 获取用户的权限
					if permission, err := redisHandler.Get(ctx, key); err == nil {
						var permissionList []string
						json.Unmarshal([]byte(permission), &permissionList)
						if util.InArray(methodName, permissionList) {
							newCtx.Next(ctx)
						} else {
							response.JSON(newCtx, consts.StatusForbidden, utils.H{
								"code": code.TokenInvalidError,
								"msg":  code.StatusText(code.TokenInvalidError),
							})
							newCtx.Abort()
							return
						}
					}
				}
			}
		} else {
			newCtx.Next(ctx)
		}
		//添加日志收集流程
		var response Response
		json.Unmarshal(newCtx.Response.Body(), &response)

		var sysLoggerOperate dao.SysLoggerOperate
		sysLoggerOperate.Name = null.StringFrom(name)                                                  // 姓名
		sysLoggerOperate.Username = proto.String(userName)                                             // 用户名
		sysLoggerOperate.UserId = proto.Int64(userId)                                                  // 用户编码
		sysLoggerOperate.Module = null.StringFromPtr(nil)                                              // 模块名称
		sysLoggerOperate.Method = null.StringFrom(cast.ToString(newCtx.Request.Method()))              // 请求方法
		sysLoggerOperate.Url = null.StringFrom(cast.ToString(newCtx.Request.URI().RequestURI()))       // 请求地址
		sysLoggerOperate.Ip = null.StringFrom(newCtx.ClientIP())                                       // ip
		sysLoggerOperate.Ua = null.StringFrom(cast.ToString(newCtx.Request.Header.UserAgent()))        // ua
		sysLoggerOperate.GoMethod = null.StringFrom(newCtx.HandlerName())                              // 方法名
		sysLoggerOperate.GoMethodArgs = null.JSONFrom(newCtx.Request.Body())                           // 方法参数
		sysLoggerOperate.StartTime = null.DateTimeFrom(startTime)                                      // 开始时间
		sysLoggerOperate.Duration = null.Int32From(cast.ToInt32(time.Since(startTime).Milliseconds())) // 执行时长
		sysLoggerOperate.Channel = null.StringFrom(channel)                                            // 渠道
		sysLoggerOperate.Result = null.Int32From(0)
		if response.Code != 200 {
			sysLoggerOperate.Result = null.Int32From(1)
		}
		sysLoggerOperate.ResultMsg = null.StringFrom(response.Msg)
		sysLoggerOperate.Deleted = proto.Int32(0)
		sysLoggerOperate.TenantId = proto.Int64(tenantId)
		sysLoggerOperate.Creator = null.StringFrom(userName)       //创建者
		sysLoggerOperate.CreateTime = null.DateTimeFrom(startTime) //创建时间
		sysLoggerOperate.Updater = null.StringFrom(userName)       //更新者
		sysLoggerOperate.UpdateTime = null.DateTimeFrom(startTime) //更新时间
		// 将这些数据需要全部存储在消息列队中,然后后台去执行消息列队
		key := util.NewReplacer(initial.Core.Config.String("Cache.SysLoggerOperate.Queue"))
		bytes, _ := json.Marshal(sysLoggerOperate)
		redisHandler.LPush(ctx, key, cast.ToString(bytes))
	}
}
