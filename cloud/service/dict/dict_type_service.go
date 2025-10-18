package dict

import (
	"cloud/code"
	"cloud/module/dict"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// dict_type 字典类型

// SrvDictTypeServiceServer 字典类型
type SrvDictTypeServiceServer struct {
	UnimplementedDictTypeServiceServer
	Server *xgrpc.Server
}

// DictTypeCreate 创建数据
func (srv SrvDictTypeServiceServer) DictTypeCreate(ctx context.Context, request *DictTypeCreateRequest) (*DictTypeCreateResponse, error) {
	req := DictTypeDao(request.GetData())
	data, err := dict.DictTypeCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:字典类型:dict_type:DictTypeCreate")
		return &DictTypeCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &DictTypeCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// DictTypeUpdate 更新数据
func (srv SrvDictTypeServiceServer) DictTypeUpdate(ctx context.Context, request *DictTypeUpdateRequest) (*DictTypeUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &DictTypeUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := DictTypeDao(request.GetData())
	_, err := dict.DictTypeUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:字典类型:dict_type:DictTypeUpdate")
		return &DictTypeUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &DictTypeUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// DictTypeDelete 删除数据
func (srv SrvDictTypeServiceServer) DictTypeDelete(ctx context.Context, request *DictTypeDeleteRequest) (*DictTypeDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &DictTypeDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dict.DictTypeDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:字典类型:dict_type:DictTypeDelete")
		return &DictTypeDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &DictTypeDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// DictType 查询单条数据
func (srv SrvDictTypeServiceServer) DictType(ctx context.Context, request *DictTypeRequest) (*DictTypeResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &DictTypeResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := dict.DictType(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:字典类型:dict_type:DictType")
		return &DictTypeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &DictTypeResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: DictTypeProto(res),
	}, nil
}

// DictTypeType 查询单条数据
func (srv SrvDictTypeServiceServer) DictTypeType(ctx context.Context, request *DictTypeTypeRequest) (*DictTypeTypeResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Type != nil {
		condition["type"] = request.GetType()
	}

	if util.Empty(condition) {
		err := errors.New("condition is empty")
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典类型:dict_type:DictTypeType")
		return &DictTypeTypeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	res, err := dict.DictTypeType(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典类型:dict_type:DictTypeType")
		return &DictTypeTypeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &DictTypeTypeResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: DictTypeProto(res),
	}, nil
}

// DictTypeList 列表数据
func (srv SrvDictTypeServiceServer) DictTypeList(ctx context.Context, request *DictTypeListRequest) (*DictTypeListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
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
	list, err := dict.DictTypeList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典类型:dict_type:DictTypeList")
		return &DictTypeListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*DictTypeObject
	for _, item := range list {
		res = append(res, DictTypeProto(item))
	}
	return &DictTypeListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// DictTypeListTotal 获取总数
func (srv SrvDictTypeServiceServer) DictTypeListTotal(ctx context.Context, request *DictTypeListTotalRequest) (*DictTypeTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.Type != nil {
		condition["type"] = request.GetType()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}

	// 获取数据集合
	total, err := dict.DictTypeListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典类型:dict_type:DictTypeListTotal")
		return &DictTypeTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &DictTypeTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
