package dict

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
)

// sys_dict 字典
// SysDictCreate 创建数据
func SysDictCreate(ctx context.Context, data dao.SysDict) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.SysDict{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysDictUpdate 更新数据
func SysDictUpdate(ctx context.Context, id int64, data dao.SysDict) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.SysDict{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysDictDelete 删除数据
func SysDictDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysDict
	result := db.WithContext(ctx).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysDict 查询单条数据
func SysDict(ctx context.Context, id int64) (res dao.SysDict, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysDict{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysDictList 查询列表数据
func SysDictList(ctx context.Context, condition map[string]any) (res []dao.SysDict, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysDict{})
	if val, ok := condition["dictTypeId"]; ok {
		builder.Where("dict_type_id = ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
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

// SysDictListTotal 查询列表数据总量
func SysDictListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysDict{})
	if val, ok := condition["dictTypeId"]; ok {
		builder.Where("dict_type_id = ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
