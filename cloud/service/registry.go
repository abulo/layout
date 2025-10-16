package service

import (
	"cloud/service/sys"
	"cloud/service/tenant"
	"cloud/service/verify"

	"github.com/abulo/ratel/v3/server/xgrpc"
)

// Registry 注册服务
func Registry(server *xgrpc.Server) {
	// 验证码服务->verify
	verify.RegisterVerifyServiceServer(server.Server, &verify.SrvVerifyServiceServer{
		Server: server,
	})
	// 注册服务
	sys.Registry(server)
	// 租户->tenant
	tenant.RegisterTenantServiceServer(server.Server, &tenant.SrvTenantServiceServer{
		Server: server,
	})
	// 租户套餐->tenant_package
	tenant.RegisterTenantPackageServiceServer(server.Server, &tenant.SrvTenantPackageServiceServer{
		Server: server,
	})
}
