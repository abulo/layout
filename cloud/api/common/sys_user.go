package common

import (
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/token"
	"cloud/service/sys/menu"
	"cloud/service/sys/role"
	"cloud/service/sys/tenant"
	"cloud/service/sys/user"
	"context"
	"encoding/json"
	"errors"
	"time"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/util"
	"github.com/gogo/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// SysUserMenuDao 数据转换
func SysUserMenuDao(item *menu.SysMenuObject) *dao.SysMenuTree {
	daoItem := &dao.SysMenuTree{}
	if !util.Empty(item.Id) {
		daoItem.Id = item.GetId() // 菜单ID
	}
	if !util.Empty(item.ParentId) {
		daoItem.ParentId = item.GetParentId() // 父菜单ID
	}
	if !util.Empty(item.Code) {
		daoItem.Code = item.GetCode() // 菜单权限
	}
	if !util.Empty(item.Path) {
		daoItem.Path = item.GetPath() // 路由地址
	}
	if !util.Empty(item.Name) {
		daoItem.Name = item.GetComponentName() // 路由名称
	}
	if !util.Empty(item.Component) {
		daoItem.Component = item.GetComponent() // 组件路径
	}
	if !util.Empty(item.Redirect) {
		daoItem.Redirect = item.GetRedirect() // 路由重定向地址
	}
	if !util.Empty(item.Type) {
		daoItem.Type = item.GetType()
	}
	daoMetaItem := &dao.SysMenuTreeMeta{}
	if !util.Empty(item.Icon) {
		daoMetaItem.Icon = item.GetIcon() // 菜单图标
	}
	if !util.Empty(item.Name) {
		daoMetaItem.Title = item.GetName() // 菜单标题
	}
	if !util.Empty(item.Link) {
		daoMetaItem.IsLink = item.GetLink() // 是否外链
	}
	daoMetaItem.IsHide = cast.ToBool(cast.ToInt(item.GetHide()))       // 是否隐藏(0:否/1是)
	daoMetaItem.IsFull = cast.ToBool(cast.ToInt(item.GetFull()))       // 是否全屏
	daoMetaItem.IsAffix = false                                        // 是否固定
	daoMetaItem.IsKeepAlive = cast.ToBool(cast.ToInt(item.GetCache())) // 是否缓存
	if !util.Empty(item.Active) {
		daoMetaItem.ActiveMenu = item.GetActive() // 激活菜单
	}
	daoItem.Meta = daoMetaItem
	return daoItem
}

func SysUserTree(ctx context.Context, sysUserId, tenantId int64) (list []*dao.SysMenuTree, currentMenuIds []int64, err error) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:菜单权限表:sys_menu:SysUserMenu")
		return nil, nil, err
	}
	// 获取用户信息
	clientSysUser := user.NewSysUserServiceClient(grpcClient)
	requestSysUser := &user.SysUserRequest{}
	requestSysUser.Id = sysUserId
	// 执行服务
	res, err := clientSysUser.SysUser(ctx, requestSysUser)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": requestSysUser,
			"err": err,
		}).Error("GrpcCall:用户信息表:sys_user:SysUser")
		return nil, nil, err
	}

	userInfo := res.GetData()
	//链接服务
	tenantClient := tenant.NewSysTenantServiceClient(grpcClient)
	tenantRequest := &tenant.SysTenantRequest{}
	tenantRequest.Id = tenantId
	// 执行服务
	tenantRes, tenantErr := tenantClient.SysTenant(ctx, tenantRequest)
	if tenantErr != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": tenantRequest,
			"err": tenantErr,
		}).Error("GrpcCall:租户:sys_tenant:SysTenant")

		return nil, nil, tenantErr
	}
	// 判断这个有没有值
	if tenantRes.GetCode() != code.Success {
		return nil, nil, nil
	}
	var menuIds []int64
	// 这是租户的信息
	tenantItem := tenant.SysTenantDao(tenantRes.GetData())
	// 获取套餐服务
	tenantPackageId := tenantItem.PackageId
	//链接服务
	menuClient := menu.NewSysMenuServiceClient(grpcClient)
	// 构造查询条件
	menuRequest := &menu.SysMenuListRequest{}
	menuRequest.Status = proto.Int32(0)

	if cast.ToInt64(tenantPackageId) != 0 {
		//链接服务
		tenantPackageClient := tenant.NewSysTenantPackageServiceClient(grpcClient)
		systemTenantPackageId := cast.ToInt64(tenantPackageId)
		tenantPackageRequest := &tenant.SysTenantPackageRequest{}
		tenantPackageRequest.Id = systemTenantPackageId
		// 执行服务
		if res, err := tenantPackageClient.SysTenantPackage(ctx, tenantPackageRequest); err == nil {
			tenantPackageItem := tenant.SysTenantPackageDao(res.GetData())
			if tenantPackageItem.ScopeMenu.IsValid() {
				json.Unmarshal(tenantPackageItem.ScopeMenu.JSON, &menuIds)
			}
		}
	}

	// 执行服务
	menuRes, menuErr := menuClient.SysMenuList(ctx, menuRequest)
	if menuErr != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": menuRequest,
			"err": menuErr,
		}).Error("GrpcCall:系统菜单:sys_menu:SysMenuList")
		return nil, nil, err
	}
	if menuRes.GetCode() == code.Success {
		rpcList := menuRes.GetData()
		for _, item := range rpcList {
			currentMenuIds = append(currentMenuIds, *item.Id)
			if cast.ToInt64(tenantPackageId) != 0 {
				if !util.InArray(*item.Id, menuIds) {
					continue
				}
			}
			list = append(list, SysUserMenuDao(item))
		}
	}
	if sysUserId != cast.ToInt64(tenantItem.UserId.Ptr()) {
		userItem := user.SysUserDao(userInfo)
		var userRoleIds []int64
		err = json.Unmarshal(userItem.RoleIds.JSON, &userRoleIds)
		if err != nil {
			return nil, nil, err
		}
		roleClient := role.NewSysRoleServiceClient(grpcClient)
		var otherRole []int64
		for _, item := range userRoleIds {
			systemRoleId := item
			roleRequest := &role.SysRoleRequest{}
			roleRequest.Id = systemRoleId
			if roleItem, err := roleClient.SysRole(ctx, roleRequest); err == nil {
				if roleItem.GetCode() == code.Success {
					roleData := role.SysRoleDao(roleItem.GetData())
					var newCurrentMenuIds []int64
					if roleData.ScopeDept.IsValid() {
						json.Unmarshal(*roleData.ScopeDept.Ptr(), &newCurrentMenuIds)
					}
					otherRole = append(otherRole, currentMenuIds...)
				}
			}
		}
		currentMenuIds = otherRole
	}
	return list, currentMenuIds, nil
}

func SysUserLogin(ctx context.Context, req dao.SysUserLogin, verifyPassword bool) (userTokenItem dao.UserToken, err error) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:用户信息表:sys_user:SysUserLogin")

		return userTokenItem, err
	}
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
		return userTokenItem, err
	}

	userInfo := res.GetData()
	if userInfo.Id == nil {
		return userTokenItem, errors.New("用户不存在")
	}
	if verifyPassword {
		// 比对密码
		if req.Password != userInfo.GetPassword() {
			return userTokenItem, errors.New("密码错误")
		}
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
		return userTokenItem, err
	}
	userTenantInfo := resUserTenant.GetData()
	if userTenantInfo.Id == nil {
		return userTokenItem, errors.New("租户不存在")
	}
	if cast.ToInt64(userInfo.TenantId) != cast.ToInt64(userTenantInfo.TenantId) {
		return userTokenItem, errors.New("用户租户获取失败")
	}
	userVerifyItem := dao.UserVerify{}
	userVerifyItem.UserId = cast.ToInt64(userItem.Id)                       // 用户ID
	userVerifyItem.UserName = cast.ToString(userItem.Username)              // 用户名
	userVerifyItem.TenantId = cast.ToInt64(userItem.TenantId)               // 租户ID
	userVerifyItem.Name = cast.ToString(userItem.Name.Ptr())                // 用户姓名
	userTokenItem.UserName = proto.String(cast.ToString(userItem.Username)) // 用户名
	dateTime := util.Now().Add(time.Duration(86000*30) * time.Second)
	accessToken, err := token.GenerateToken(userVerifyItem, "user", 86400*30)
	if err != nil {
		return userTokenItem, err
	}
	userTokenItem.AccessToken = proto.String(accessToken)
	refreshToken, err := token.GenerateToken(userVerifyItem, "user", 86400*31)
	if err != nil {
		return userTokenItem, err
	}
	userTokenItem.RefreshToken = proto.String(refreshToken)
	userTokenItem.Expires = proto.String(util.Date("Y-m-d H:i:s", dateTime))
	list, currentMenuIds, err := SysUserTree(ctx, cast.ToInt64(userItem.Id), cast.ToInt64(userItem.TenantId))
	if err != nil {
		return userTokenItem, err
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
	return userTokenItem, nil
}
