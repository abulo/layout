package user

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// sys_user 用户
// SysUserCreate 创建数据
func SysUserCreate(ctx context.Context, data dao.SysUser) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.SysUser{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysUserUpdate 更新数据
func SysUserUpdate(ctx context.Context, id int64, data dao.SysUser) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	result := db.WithContext(ctx).Model(&dao.SysUser{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysUserDelete 删除数据
func SysUserDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data := make(map[string]any)
	data["deleted"] = 1
	result := db.WithContext(ctx).Model(&dao.SysUser{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysUser 查询单条数据
func SysUser(ctx context.Context, id int64) (res dao.SysUser, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysUser{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysUserRecover 恢复数据
func SysUserRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data := make(map[string]any)
	data["deleted"] = 0
	result := db.WithContext(ctx).Model(&dao.SysUser{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysUserDrop 清理数据
func SysUserDrop(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	result := db.WithContext(ctx).Where("id = ?", id).Delete(&dao.SysUser{})
	return result.RowsAffected, result.Error
}

// SysUserLogin 查询单条数据
func SysUserLogin(ctx context.Context, condition map[string]any) (res dao.SysUser, err error) {
	if util.Empty(condition) {
		err = errors.New("condition is empty")
		return
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUser{})
	if val, ok := condition["username"]; ok {
		builder.Where("username = ?", val)
	}

	err = builder.Find(&res).Error
	return
}

// SysUserList 查询列表数据
func SysUserList(ctx context.Context, condition map[string]any) (res []dao.SysUser, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUser{})
	if val, ok := condition["username"]; ok {
		builder.Where("username = ?", val)
	}
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("deleted = ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["mobile"]; ok {
		builder.Where("mobile = ?", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("name = ?", val)
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

// SysUserListTotal 查询列表数据总量
func SysUserListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUser{})
	if val, ok := condition["username"]; ok {
		builder.Where("username = ?", val)
	}
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("deleted = ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["mobile"]; ok {
		builder.Where("mobile = ?", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("name = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
