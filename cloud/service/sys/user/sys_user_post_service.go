package user

import (
	"cloud/code"
	"cloud/module/sys/user"
	"context"

	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// sys_user_post 用户职位

// SrvSysUserPostServiceServer 用户职位
type SrvSysUserPostServiceServer struct {
	UnimplementedSysUserPostServiceServer
	Server *xgrpc.Server
}

// SysUserPostCreate 创建数据
func (srv SrvSysUserPostServiceServer) SysUserPostCreate(ctx context.Context, request *SysUserPostCreateRequest) (*SysUserPostCreateResponse, error) {
	req := SysUserPostDao(request.GetData())
	data, err := user.SysUserPostCreate(ctx, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:用户职位:sys_user_post:SysUserPostCreate")
		return &SysUserPostCreateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserPostCreateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: data,
	}, nil
}

// SysUserPostUpdate 更新数据
func (srv SrvSysUserPostServiceServer) SysUserPostUpdate(ctx context.Context, request *SysUserPostUpdateRequest) (*SysUserPostUpdateResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserPostUpdateResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	req := SysUserPostDao(request.GetData())
	_, err := user.SysUserPostUpdate(ctx, id, *req)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": req,
			"err": err,
		}).Error("Sql:用户职位:sys_user_post:SysUserPostUpdate")
		return &SysUserPostUpdateResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserPostUpdateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysUserPostDelete 删除数据
func (srv SrvSysUserPostServiceServer) SysUserPostDelete(ctx context.Context, request *SysUserPostDeleteRequest) (*SysUserPostDeleteResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserPostDeleteResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	_, err := user.SysUserPostDelete(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:用户职位:sys_user_post:SysUserPostDelete")
		return &SysUserPostDeleteResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserPostDeleteResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
	}, nil
}

// SysUserPost 查询单条数据
func (srv SrvSysUserPostServiceServer) SysUserPost(ctx context.Context, request *SysUserPostRequest) (*SysUserPostResponse, error) {
	id := request.GetId()
	if id < 1 {
		return &SysUserPostResponse{}, status.Error(code.ConvertToGrpc(code.ParamInvalid), code.StatusText(code.ParamInvalid))
	}
	res, err := user.SysUserPost(ctx, id)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": id,
			"err": err,
		}).Error("Sql:用户职位:sys_user_post:SysUserPost")
		return &SysUserPostResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserPostResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysUserPostProto(res),
	}, nil
}

// SysUserPostItem 查询单条数据
func (srv SrvSysUserPostServiceServer) SysUserPostItem(ctx context.Context, request *SysUserPostItemRequest) (*SysUserPostItemResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.PostId != nil {
		condition["postId"] = request.GetPostId()
	}

	if util.Empty(condition) {
		err := errors.New("condition is empty")
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户职位:sys_user_post:SysUserPostItem")
		return &SysUserPostItemResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	res, err := user.SysUserPostItem(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户职位:sys_user_post:SysUserPostItem")
		return &SysUserPostItemResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserPostItemResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: SysUserPostProto(res),
	}, nil
}

// SysUserPostList 列表数据
func (srv SrvSysUserPostServiceServer) SysUserPostList(ctx context.Context, request *SysUserPostListRequest) (*SysUserPostListResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.PostId != nil {
		condition["postId"] = request.GetPostId()
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
	list, err := user.SysUserPostList(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户职位:sys_user_post:SysUserPostList")
		return &SysUserPostListResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	var res []*SysUserPostObject
	for _, item := range list {
		res = append(res, SysUserPostProto(item))
	}
	return &SysUserPostListResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: res,
	}, nil
}

// SysUserPostListTotal 获取总数
func (srv SrvSysUserPostServiceServer) SysUserPostListTotal(ctx context.Context, request *SysUserPostListTotalRequest) (*SysUserPostTotalResponse, error) {
	// 数据库查询条件
	condition := make(map[string]any)
	// 构造查询条件
	if request.TenantId != nil {
		condition["tenantId"] = request.GetTenantId()
	}
	if request.UserId != nil {
		condition["userId"] = request.GetUserId()
	}
	if request.PostId != nil {
		condition["postId"] = request.GetPostId()
	}

	// 获取数据集合
	total, err := user.SysUserPostListTotal(ctx, condition)
	if sql.Acceptable(err) != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": condition,
			"err": err,
		}).Error("Sql:用户职位:sys_user_post:SysUserPostListTotal")
		return &SysUserPostTotalResponse{}, status.Error(code.ConvertToGrpc(code.SqlError), err.Error())
	}
	return &SysUserPostTotalResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: total,
	}, nil
}
