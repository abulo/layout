package tenant

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"
)

// sys_tenant_package 租户套餐
// SysTenantPackageCreate 创建数据
func SysTenantPackageCreate(ctx context.Context, data dao.SysTenantPackage) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Write()
	err = db.WithContext(ctx).Model(&dao.SysTenantPackage{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysTenantPackageUpdate 更新数据
func SysTenantPackageUpdate(ctx context.Context, id int64, data dao.SysTenantPackage) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.SysTenantPackage{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysTenantPackageDelete 删除数据
func SysTenantPackageDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Write()
	var data dao.SysTenantPackage
	result := db.WithContext(ctx).Model(&dao.SysTenantPackage{}).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysTenantPackage 查询单条数据
func SysTenantPackage(ctx context.Context, id int64) (res dao.SysTenantPackage, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Read()
	err = db.WithContext(ctx).Model(&dao.SysTenantPackage{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysTenantPackageList 查询列表数据
func SysTenantPackageList(ctx context.Context, condition map[string]any) (res []dao.SysTenantPackage, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Read()
	builder := db.WithContext(ctx).Model(&dao.SysTenantPackage{})
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
	builder.Order(clause.OrderBy{Columns: []clause.OrderByColumn{
		{Column: clause.Column{Name: "sort"}, Desc: false},
		{Column: clause.Column{Name: "id"}, Desc: true},
	}})
	err = builder.Find(&res).Error
	return
}

// SysTenantPackageListTotal 查询列表数据总量
func SysTenantPackageListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Read()
	builder := db.WithContext(ctx).Model(&dao.SysTenantPackage{})
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("name = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
