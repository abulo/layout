package tenant

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
)

// tenant 租户
// TenantCreate 创建数据
func TenantCreate(ctx context.Context, data dao.Tenant) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.Tenant{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// TenantUpdate 更新数据
func TenantUpdate(ctx context.Context, id int64, data dao.Tenant) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	result := db.WithContext(ctx).Model(&dao.Tenant{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// TenantDelete 删除数据
func TenantDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data := make(map[string]any)
	data["deleted"] = 1
	result := db.WithContext(ctx).Model(&dao.Tenant{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// Tenant 查询单条数据
func Tenant(ctx context.Context, id int64) (res dao.Tenant, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.Tenant{}).Where("id = ?", id).Find(&res).Error
	return
}

// TenantRecover 恢复数据
func TenantRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data := make(map[string]any)
	data["deleted"] = 0
	result := db.WithContext(ctx).Model(&dao.Tenant{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// TenantDrop 清理数据
func TenantDrop(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	result := db.WithContext(ctx).Where("id = ?", id).Delete(&dao.Tenant{})
	return result.RowsAffected, result.Error
}

// TenantList 查询列表数据
func TenantList(ctx context.Context, condition map[string]any) (res []dao.Tenant, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.Tenant{})
	if val, ok := condition["packageId"]; ok {
		builder.Where("package_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("deleted = ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("name = ?", val)
	}
	if val, ok := condition["expireDate"]; ok {
		builder.Where("expire_date = ?", val)
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

// TenantListTotal 查询列表数据总量
func TenantListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.Tenant{})
	if val, ok := condition["packageId"]; ok {
		builder.Where("package_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("deleted = ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("name = ?", val)
	}
	if val, ok := condition["expireDate"]; ok {
		builder.Where("expire_date = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
