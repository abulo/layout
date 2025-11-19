package dept

import (
	"context"

	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/sys/dept"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// sys_dept 部门
// SysDeptItem 查询单条数据
func SysDeptItem(ctx context.Context, newCtx *app.RequestContext, id int64) (*dept.SysDeptResponse, error) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:sys_dept:SysDeptItem")
		return nil, status.Error(code.ConvertToGrpc(code.RPCError), code.StatusText(code.RPCError))
	}
	//链接服务
	client := dept.NewSysDeptServiceClient(grpcClient)
	request := &dept.SysDeptRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysDept(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:sys_dept:SysDeptItem")
		return nil, err
	}
	return res, nil
}

// SysDeptCreate 创建数据
func SysDeptCreate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:sys_dept:SysDeptCreate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dept.NewSysDeptServiceClient(grpcClient)
	request := &dept.SysDeptCreateRequest{}
	// 数据绑定
	var reqInfo dao.SysDept
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Id = nil
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId"))
	reqInfo.Creator = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.CreateTime = null.DateTimeFrom(util.Now())
	request.Data = dept.SysDeptProto(reqInfo)
	// 执行服务
	res, err := client.SysDeptCreate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:sys_dept:SysDeptCreate")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SysDeptUpdate 更新数据
func SysDeptUpdate(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysDeptItem(ctx, newCtx, id); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:sys_dept:SysDeptUpdate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dept.NewSysDeptServiceClient(grpcClient)
	request := &dept.SysDeptUpdateRequest{}
	request.Id = id
	// 数据绑定
	var reqInfo dao.SysDept
	if err := newCtx.BindAndValidate(&reqInfo); err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ParamInvalid,
			"msg":  code.StatusText(code.ParamInvalid),
		})
		return
	}
	reqInfo.Id = nil
	reqInfo.TenantId = proto.Int64(newCtx.GetInt64("tenantId"))
	reqInfo.Updater = null.StringFrom(newCtx.GetString("userName"))
	reqInfo.UpdateTime = null.DateTimeFrom(util.Now())
	reqInfo.Creator = null.StringFromPtr(nil)
	reqInfo.CreateTime = null.DateTimeFromPtr(nil)
	request.Data = dept.SysDeptProto(reqInfo)
	// 执行服务
	res, err := client.SysDeptUpdate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:sys_dept:SysDeptUpdate")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SysDeptDelete 删除数据
func SysDeptDelete(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysDeptItem(ctx, newCtx, id); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:sys_dept:SysDeptDelete")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dept.NewSysDeptServiceClient(grpcClient)
	request := &dept.SysDeptDeleteRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysDeptDelete(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:sys_dept:SysDeptDelete")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SysDept 查询单条数据
func SysDept(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	// 执行服务
	res, err := SysDeptItem(ctx, newCtx, id)
	if err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": dept.SysDeptDao(res.GetData()),
	})
}

// SysDeptRecover 恢复数据
func SysDeptRecover(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysDeptItem(ctx, newCtx, id); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:sys_dept:SysDeptRecover")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dept.NewSysDeptServiceClient(grpcClient)
	request := &dept.SysDeptRecoverRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysDeptRecover(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:sys_dept:SysDeptRecover")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SysDeptDrop 清理数据
func SysDeptDrop(ctx context.Context, newCtx *app.RequestContext) {
	id := cast.ToInt64(newCtx.Param("id"))
	if _, err := SysDeptItem(ctx, newCtx, id); err != nil {
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:sys_dept:SysDeptDrop")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dept.NewSysDeptServiceClient(grpcClient)
	request := &dept.SysDeptDropRequest{}
	request.Id = id
	// 执行服务
	res, err := client.SysDeptDrop(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:sys_dept:SysDeptDrop")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
	})
}

// SysDeptList 列表数据
func SysDeptList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:sys_dept:SysDeptList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dept.NewSysDeptServiceClient(grpcClient)
	// 构造查询条件
	request := &dept.SysDeptListRequest{}
	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                            // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val)) // 状态:0正常/1停用
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val) // 名称
	}

	// 执行服务
	res, err := client.SysDeptList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:sys_dept:SysDeptList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SysDept
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, dept.SysDeptDao(item))
		}
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": list,
	})
}

// SysDeptListSimple 列表精简数据
func SysDeptListSimple(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:部门:sys_dept:SysDeptListSimple")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dept.NewSysDeptServiceClient(grpcClient)
	// 构造查询条件
	request := &dept.SysDeptListRequest{}

	request.TenantId = proto.Int64(newCtx.GetInt64("tenantId")) // 租户ID
	request.Deleted = proto.Int32(0)                            // 删除状态
	if val, ok := newCtx.GetQuery("deleted"); ok {
		if cast.ToBool(val) {
			request.Deleted = nil
		}
	}
	if val, ok := newCtx.GetQuery("status"); ok {
		request.Status = proto.Int32(cast.ToInt32(val)) // 状态:0正常/1停用
	}
	if val, ok := newCtx.GetQuery("name"); ok {
		request.Name = proto.String(val) // 名称
	}

	// 执行服务
	res, err := client.SysDeptList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:部门:sys_dept:SysDeptListSimple")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	var list []*dao.SysDept
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		for _, item := range rpcList {
			list = append(list, dept.SysDeptDao(item))
		}
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": list,
	})
}
