package post

import (
	"cloud/code"
	"cloud/module/sys/post"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// sys_post 职位

// SrvSysPostServiceServer 职位
type SrvSysPostServiceServer struct {
	UnimplementedSysPostServiceServer
	Server *xgrpc.Server
}

// SysPostCreate 创建数据
func (srv SrvSysPostServiceServer) SysPostCreate(ctx context.Context, request *SysPostCreateRequest) (*SysPostCreateResponse, error) {
	req := SysPostDao(request.GetData())
	data, err := post.SysPostCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:职位:sys_post:SysPostCreate")
		return &SysPostCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysPostCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysPostUpdate 更新数据
func (srv SrvSysPostServiceServer) SysPostUpdate(ctx context.Context, request *SysPostUpdateRequest) (*SysPostUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysPostUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysPostDao(request.GetData())
	_, err := post.SysPostUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:职位:sys_post:SysPostUpdate")
		return &SysPostUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysPostUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysPostDelete 删除数据
func (srv SrvSysPostServiceServer) SysPostDelete(ctx context.Context, request *SysPostDeleteRequest) (*SysPostDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysPostDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := post.SysPostDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:职位:sys_post:SysPostDelete")
		return &SysPostDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysPostDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysPost 查询单条数据
func (srv SrvSysPostServiceServer) SysPost(ctx context.Context, request *SysPostRequest) (*SysPostResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysPostResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := post.SysPost(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:职位:sys_post:SysPost")
		return &SysPostResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysPostResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysPostProto(res),
	}, nil
}

// SysPostRecover 恢复数据
func (srv SrvSysPostServiceServer) SysPostRecover(ctx context.Context, request *SysPostRecoverRequest) (*SysPostRecoverResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysPostRecoverResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := post.SysPostRecover(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:职位:sys_post:SysPostRecover")
		return &SysPostRecoverResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysPostRecoverResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysPostDrop 清理数据
func (srv SrvSysPostServiceServer) SysPostDrop(ctx context.Context, request *SysPostDropRequest) (*SysPostDropResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysPostDropResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := post.SysPostDrop(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:职位:sys_post:SysPostDrop")
		return &SysPostDropResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysPostDropResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysPostList 列表数据
func (srv SrvSysPostServiceServer) SysPostList(ctx context.Context, request *SysPostListRequest) (*SysPostListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
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
	list, err := post.SysPostList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:职位:sys_post:SysPostList")
		return &SysPostListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysPostObject
	for _, item := range list {
		res = append(res, SysPostProto(item))
	}
	return &SysPostListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysPostListTotal 获取总数
func (srv SrvSysPostServiceServer) SysPostListTotal(ctx context.Context, request *SysPostListTotalRequest) (*SysPostTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.Deleted != nil {
		condition["deleted"] = request.GetDeleted()
	}
	if request.Status != nil {
		condition["status"] = request.GetStatus()
	}
	if request.Name != nil {
		condition["name"] = request.GetName()
	}

	// 获取数据集合
	total, err := post.SysPostListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:职位:sys_post:SysPostListTotal")
		return &SysPostTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysPostTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
