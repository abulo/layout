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
	tenant.Registry(server)
}
