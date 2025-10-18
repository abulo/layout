package dict

import (
	"cloud/code"
	"cloud/module/dict"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// dict 字典

// SrvDictServiceServer 字典
type SrvDictServiceServer struct {
	UnimplementedDictServiceServer
	Server *xgrpc.Server
}

// DictCreate 创建数据
func (srv SrvDictServiceServer) DictCreate(ctx context.Context, request *DictCreateRequest) (*DictCreateResponse, error) {
	req := DictDao(request.GetData())
	data, err := dict.DictCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:字典:dict:DictCreate")
		return &DictCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &DictCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// DictUpdate 更新数据
func (srv SrvDictServiceServer) DictUpdate(ctx context.Context, request *DictUpdateRequest) (*DictUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &DictUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := DictDao(request.GetData())
	_, err := dict.DictUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:字典:dict:DictUpdate")
		return &DictUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &DictUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// DictDelete 删除数据
func (srv SrvDictServiceServer) DictDelete(ctx context.Context, request *DictDeleteRequest) (*DictDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &DictDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dict.DictDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:字典:dict:DictDelete")
		return &DictDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &DictDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// Dict 查询单条数据
func (srv SrvDictServiceServer) Dict(ctx context.Context, request *DictRequest) (*DictResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &DictResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := dict.Dict(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:字典:dict:Dict")
		return &DictResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &DictResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: DictProto(res),
	}, nil
}

// DictList 列表数据
func (srv SrvDictServiceServer) DictList(ctx context.Context, request *DictListRequest) (*DictListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.DictType != nil {
		condition["dictType"] = request.GetDictType()
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
	list, err := dict.DictList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典:dict:DictList")
		return &DictListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*DictObject
	for _, item := range list {
		res = append(res, DictProto(item))
	}
	return &DictListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// DictListTotal 获取总数
func (srv SrvDictServiceServer) DictListTotal(ctx context.Context, request *DictListTotalRequest) (*DictTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.DictType != nil {
		condition["dictType"] = request.GetDictType()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}

	// 获取数据集合
	total, err := dict.DictListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典:dict:DictListTotal")
		return &DictTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &DictTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
