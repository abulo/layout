package tenant

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
)

// tenant_package 租户套餐
// TenantPackageCreate 创建数据
func TenantPackageCreate(ctx context.Context, data dao.TenantPackage) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.TenantPackage{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// TenantPackageUpdate 更新数据
func TenantPackageUpdate(ctx context.Context, id int64, data dao.TenantPackage) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.TenantPackage{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// TenantPackageDelete 删除数据
func TenantPackageDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.TenantPackage
	result := db.WithContext(ctx).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// TenantPackage 查询单条数据
func TenantPackage(ctx context.Context, id int64) (res dao.TenantPackage, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.TenantPackage{}).Where("id = ?", id).Find(&res).Error
	return
}

// TenantPackageList 查询列表数据
func TenantPackageList(ctx context.Context, condition map[string]any) (res []dao.TenantPackage, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.TenantPackage{})
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
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

// TenantPackageListTotal 查询列表数据总量
func TenantPackageListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.TenantPackage{})
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("name = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
