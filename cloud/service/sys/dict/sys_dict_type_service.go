package dict

import (
	"cloud/code"
	"cloud/module/sys/dict"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// sys_dict_type 字典类型

// SrvSysDictTypeServiceServer 字典类型
type SrvSysDictTypeServiceServer struct {
	UnimplementedSysDictTypeServiceServer
	Server *xgrpc.Server
}

// SysDictTypeCreate 创建数据
func (srv SrvSysDictTypeServiceServer) SysDictTypeCreate(ctx context.Context, request *SysDictTypeCreateRequest) (*SysDictTypeCreateResponse, error) {
	req := SysDictTypeDao(request.GetData())
	data, err := dict.SysDictTypeCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:字典类型:sys_dict_type:SysDictTypeCreate")
		return &SysDictTypeCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDictTypeCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysDictTypeUpdate 更新数据
func (srv SrvSysDictTypeServiceServer) SysDictTypeUpdate(ctx context.Context, request *SysDictTypeUpdateRequest) (*SysDictTypeUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysDictTypeUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysDictTypeDao(request.GetData())
	_, err := dict.SysDictTypeUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:字典类型:sys_dict_type:SysDictTypeUpdate")
		return &SysDictTypeUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDictTypeUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysDictTypeDelete 删除数据
func (srv SrvSysDictTypeServiceServer) SysDictTypeDelete(ctx context.Context, request *SysDictTypeDeleteRequest) (*SysDictTypeDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysDictTypeDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := dict.SysDictTypeDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:字典类型:sys_dict_type:SysDictTypeDelete")
		return &SysDictTypeDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDictTypeDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysDictType 查询单条数据
func (srv SrvSysDictTypeServiceServer) SysDictType(ctx context.Context, request *SysDictTypeRequest) (*SysDictTypeResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysDictTypeResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := dict.SysDictType(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:字典类型:sys_dict_type:SysDictType")
		return &SysDictTypeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDictTypeResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysDictTypeProto(res),
	}, nil
}

// SysDictTypeType 查询单条数据
func (srv SrvSysDictTypeServiceServer) SysDictTypeType(ctx context.Context, request *SysDictTypeTypeRequest) (*SysDictTypeTypeResponse, error) {
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
		}).Error("Sql:字典类型:sys_dict_type:SysDictTypeType")
		return &SysDictTypeTypeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	res, err := dict.SysDictTypeType(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典类型:sys_dict_type:SysDictTypeType")
		return &SysDictTypeTypeResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDictTypeTypeResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysDictTypeProto(res),
	}, nil
}

// SysDictTypeList 列表数据
func (srv SrvSysDictTypeServiceServer) SysDictTypeList(ctx context.Context, request *SysDictTypeListRequest) (*SysDictTypeListResponse, error) {
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
	list, err := dict.SysDictTypeList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典类型:sys_dict_type:SysDictTypeList")
		return &SysDictTypeListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysDictTypeObject
	for _, item := range list {
		res = append(res, SysDictTypeProto(item))
	}
	return &SysDictTypeListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysDictTypeListTotal 获取总数
func (srv SrvSysDictTypeServiceServer) SysDictTypeListTotal(ctx context.Context, request *SysDictTypeListTotalRequest) (*SysDictTypeTotalResponse, error) {
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
	total, err := dict.SysDictTypeListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:字典类型:sys_dict_type:SysDictTypeListTotal")
		return &SysDictTypeTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysDictTypeTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
