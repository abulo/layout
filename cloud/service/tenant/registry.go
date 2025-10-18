package tenant

import (
	"github.com/abulo/ratel/v3/server/xgrpc"
)

// Registry 注册服务
func Registry(server *xgrpc.Server) {
	// 租户->tenant
	RegisterTenantServiceServer(server.Server, &SrvTenantServiceServer{
		Server: server,
	})
	// 租户套餐->tenant_package
	RegisterTenantPackageServiceServer(server.Server, &SrvTenantPackageServiceServer{
		Server: server,
	})
}
