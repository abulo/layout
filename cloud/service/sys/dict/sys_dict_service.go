package dict

import (
	"cloud/code"
	"cloud/module/sys/dict"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// sys_dict 字典

// SrvSysDictServiceServer 字典
type SrvSysDictServiceServer struct {
	UnimplementedSysDictServiceServer
	Server *xgrpc.Server
}

// SysDictCreate 创建数据
func (srv SrvSysDictServiceServer) SysDictCreate(ctx context.Context, request *SysDictCreateRequest) (*SysDictCreateResponse, error) {
	req := SysDictDao(request.GetData())
	data, err := dict.SysDictCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:字典:sys_dict:SysDictCreate")
		return &SysDictCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDictCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysDictUpdate 更新数据
func (srv SrvSysDictServiceServer) SysDictUpdate(ctx context.Context, request *SysDictUpdateRequest) (*SysDictUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysDictUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysDictDao(request.GetData())
	_, err := dict.SysDictUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:字典:sys_dict:SysDictUpdate")
		return &SysDictUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDictUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysDictDelete 删除数据
func (srv SrvSysDictServiceServer) SysDictDelete(ctx context.Context, request *SysDictDeleteRequest) (*SysDictDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysDictDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dict.SysDictDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:字典:sys_dict:SysDictDelete")
		return &SysDictDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDictDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysDict 查询单条数据
func (srv SrvSysDictServiceServer) SysDict(ctx context.Context, request *SysDictRequest) (*SysDictResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysDictResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := dict.SysDict(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:字典:sys_dict:SysDict")
		return &SysDictResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDictResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysDictProto(res),
	}, nil
}

// SysDictList 列表数据
func (srv SrvSysDictServiceServer) SysDictList(ctx context.Context, request *SysDictListRequest) (*SysDictListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.DictId != nil {
		condition["dictId"] = request.GetDictId()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}

	paginationRequest := request.GetPagination()
	if paginationRequest != nil {
		// 当前页面
		pageNum := paginationRequest.GetPageNum()
		// 每页多少数据
		pageSize := paginationRequest.GetPageSize()
		if pageNum < 1 {
			pageNum = 1
		}
		if pageSize < 1 {
			pageSize = 10
		}
		// 分页数据
		offset := pageSize * (pageNum - 1)
		pagination := &sql.Pagination{
			Offset: &offset,
			Limit:  &pageSize,
		}
		condition["pagination"] = pagination
	}
	// 获取数据集合
	list, err := dict.SysDictList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典:sys_dict:SysDictList")
		return &SysDictListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysDictObject
	for _, item := range list {
		res = append(res, SysDictProto(item))
	}
	return &SysDictListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysDictListTotal 获取总数
func (srv SrvSysDictServiceServer) SysDictListTotal(ctx context.Context, request *SysDictListTotalRequest) (*SysDictTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.DictId != nil {
		condition["dictId"] = request.GetDictId()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}

	// 获取数据集合
	total, err := dict.SysDictListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典:sys_dict:SysDictListTotal")
		return &SysDictTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDictTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
