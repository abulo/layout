package dict

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"
)

// sys_dict_type 字典类型
// SysDictTypeCreate 创建数据
func SysDictTypeCreate(ctx context.Context, data dao.SysDictType) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.SysDictType{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysDictTypeUpdate 更新数据
func SysDictTypeUpdate(ctx context.Context, id int64, data dao.SysDictType) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.SysDictType{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysDictTypeDelete 删除数据
func SysDictTypeDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysDictType
	result := db.WithContext(ctx).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysDictType 查询单条数据
func SysDictType(ctx context.Context, id int64) (res dao.SysDictType, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysDictType{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysDictTypeType 查询单条数据
func SysDictTypeType(ctx context.Context, condition map[string]any) (res dao.SysDictType, err error) {
	if util.Empty(condition) {
		err = errors.New("condition is empty")
		return
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysDictType{})
	if val, ok := condition["type"]; ok {
		builder.Where("type = ?", val)
	}

	err = builder.Find(&res).Error
	return
}

// SysDictTypeList 查询列表数据
func SysDictTypeList(ctx context.Context, condition map[string]any) (res []dao.SysDictType, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysDictType{})
	if val, ok := condition["type"]; ok {
		builder.Where("type = ?", val)
	}
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
		{Column: clause.Column{Name: "id"}, Desc: true},
	}})
	err = builder.Find(&res).Error
	return
}

// SysDictTypeListTotal 查询列表数据总量
func SysDictTypeListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysDictType{})
	if val, ok := condition["type"]; ok {
		builder.Where("type = ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("name = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
