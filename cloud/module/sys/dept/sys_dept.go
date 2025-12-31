package dept

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"
)

// sys_dept 部门
// SysDeptCreate 创建数据
func SysDeptCreate(ctx context.Context, data dao.SysDept) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Write()
	err = db.WithContext(ctx).Model(&dao.SysDept{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysDeptUpdate 更新数据
func SysDeptUpdate(ctx context.Context, id int64, data dao.SysDept) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.SysDept{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysDeptDelete 删除数据
func SysDeptDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Write()
	var data dao.SysDept
	data.Id = proto.Int64(id)
	data.Deleted = proto.Int32(1)
	result := db.WithContext(ctx).Model(&dao.SysDept{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysDept 查询单条数据
func SysDept(ctx context.Context, id int64) (res dao.SysDept, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Read()
	err = db.WithContext(ctx).Model(&dao.SysDept{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysDeptRecover 恢复数据
func SysDeptRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Write()
	var data dao.SysDept
	data.Id = proto.Int64(id)
	data.Deleted = proto.Int32(0)
	result := db.WithContext(ctx).Model(&dao.SysDept{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysDeptDrop 清理数据
func SysDeptDrop(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Write()
	var data dao.SysDept
	result := db.WithContext(ctx).Model(&dao.SysDept{}).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysDeptList 查询列表数据
func SysDeptList(ctx context.Context, condition map[string]any) (res []dao.SysDept, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Read()
	builder := db.WithContext(ctx).Model(&dao.SysDept{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("deleted = ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["parentId"]; ok {
		builder.Where("parent_id = ?", val)
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
		{Column: clause.Column{Name: "parent_id"}, Desc: false},
		{Column: clause.Column{Name: "id"}, Desc: true},
	}})
	err = builder.Find(&res).Error
	return
}

// SysDeptListTotal 查询列表数据总量
func SysDeptListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Read()
	builder := db.WithContext(ctx).Model(&dao.SysDept{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("deleted = ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["parentId"]; ok {
		builder.Where("parent_id = ?", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("name = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
