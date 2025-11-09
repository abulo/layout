package dict

import (
	"cloud/code"
	"cloud/dao"
	"cloud/initial"
	"cloud/internal/response"
	"cloud/service/sys/dict"
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

func DictList(ctx context.Context, newCtx *app.RequestContext) {
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:字典类型:sys_dict_type:DictList")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := dict.NewSysDictTypeServiceClient(grpcClient)
	// 构造查询条件
	request := &dict.SysDictTypeListRequest{}
	request.Status = proto.Int32(0) // 状态:0正常/1停用
	// 执行服务
	res, err := client.SysDictTypeList(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:字典类型:sys_dict_type:DictList")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	dictMap := make(map[string][]*dao.SysDict)
	if res.GetCode() == code.Success {
		rpcList := res.GetData()
		//链接服务
		clientDict := dict.NewSysDictServiceClient(grpcClient)
		// 构造查询条件
		requestDict := &dict.SysDictListRequest{}
		requestDict.Status = proto.Int32(0) // 状态（0正常 1停用）
		for _, itemType := range rpcList {
			requestDict.DictTypeId = itemType.Id
			requestDict.Status = itemType.Status
			// 执行服务
			resDict, errDict := clientDict.SysDictList(ctx, requestDict)
			if errDict != nil {
				continue
			}
			var dictList []*dao.SysDict
			if resDict.GetCode() == code.Success {
				rpcDictList := resDict.GetData()
				for _, item := range rpcDictList {
					dictList = append(dictList, dict.SysDictDao(item))
				}
			}
			if len(dictList) > 0 {
				dictMap[cast.ToString(itemType.Type)] = dictList
			}
		}
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": res.GetCode(),
		"msg":  res.GetMsg(),
		"data": dictMap,
	})
}
