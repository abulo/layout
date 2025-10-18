package dict

import (
	"github.com/abulo/ratel/v3/server/xgrpc"
)

// Registry 注册服务
func Registry(server *xgrpc.Server) {
	// 字典类型->dict_type
	RegisterDictTypeServiceServer(server.Server, &SrvDictTypeServiceServer{
		Server: server,
	})
	// 字典->dict
	RegisterDictServiceServer(server.Server, &SrvDictServiceServer{
		Server: server,
	})
}
