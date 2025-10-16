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

// sys_user_tenant 租户用户
// SysUserTenantCreate 创建数据
func SysUserTenantCreate(ctx context.Context, data dao.SysUserTenant) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.SysUserTenant{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysUserTenantUpdate 更新数据
func SysUserTenantUpdate(ctx context.Context, id int64, data dao.SysUserTenant) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	result := db.WithContext(ctx).Model(&dao.SysUserTenant{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysUserTenantDelete 删除数据
func SysUserTenantDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	result := db.WithContext(ctx).Where("id = ?", id).Delete(&dao.SysUserTenant{})
	return result.RowsAffected, result.Error
}

// SysUserTenant 查询单条数据
func SysUserTenant(ctx context.Context, id int64) (res dao.SysUserTenant, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysUserTenant{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysUserTenantBind 查询单条数据
func SysUserTenantBind(ctx context.Context, condition map[string]any) (res dao.SysUserTenant, err error) {
	if util.Empty(condition) {
		err = errors.New("condition is empty")
		return
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUserTenant{})
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}

	err = builder.Find(&res).Error
	return
}

// SysUserTenantList 查询列表数据
func SysUserTenantList(ctx context.Context, condition map[string]any) (res []dao.SysUserTenant, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUserTenant{})
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
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

// SysUserTenantListTotal 查询列表数据总量
func SysUserTenantListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUserTenant{})
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
