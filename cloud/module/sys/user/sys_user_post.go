package user

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
)

// sys_user_post 用户职位
// SysUserPostCreate 创建数据
func SysUserPostCreate(ctx context.Context, data dao.SysUserPost) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.SysUserPost{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysUserPostUpdate 更新数据
func SysUserPostUpdate(ctx context.Context, id int64, data dao.SysUserPost) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.SysUserPost{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysUserPostDelete 删除数据
func SysUserPostDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysUserPost
	result := db.WithContext(ctx).Model(&dao.SysUserPost{}).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysUserPost 查询单条数据
func SysUserPost(ctx context.Context, id int64) (res dao.SysUserPost, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysUserPost{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysUserPostItem 查询单条数据
func SysUserPostItem(ctx context.Context, condition map[string]any) (res dao.SysUserPost, err error) {
	if util.Empty(condition) {
		err = errors.New("condition is empty")
		return
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUserPost{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["postId"]; ok {
		builder.Where("post_id = ?", val)
	}

	err = builder.Find(&res).Error
	return
}

// SysUserPostList 查询列表数据
func SysUserPostList(ctx context.Context, condition map[string]any) (res []dao.SysUserPost, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUserPost{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["postId"]; ok {
		builder.Where("post_id = ?", val)
	}

	if val, ok := condition["pagination"]; ok {
		pagination := val.(*sql.Pagination)
		if val, err := pagination.GetOffset(); err == nil {
			builder.Offset(cast.ToInt(val))
		}
		if val, err := pagination.GetLimit(); err == nil {
			builder.Limit(cast.ToInt(val))
		}
	}
	builder.Order("id")
	err = builder.Find(&res).Error
	return
}

// SysUserPostListTotal 查询列表数据总量
func SysUserPostListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUserPost{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["postId"]; ok {
		builder.Where("post_id = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
