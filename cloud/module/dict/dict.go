package dict

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
)

// dict 字典
// DictCreate 创建数据
func DictCreate(ctx context.Context, data dao.Dict) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.Dict{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// DictUpdate 更新数据
func DictUpdate(ctx context.Context, id int64, data dao.Dict) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.Dict{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// DictDelete 删除数据
func DictDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.Dict
	result := db.WithContext(ctx).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// Dict 查询单条数据
func Dict(ctx context.Context, id int64) (res dao.Dict, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.Dict{}).Where("id = ?", id).Find(&res).Error
	return
}

// DictList 查询列表数据
func DictList(ctx context.Context, condition map[string]any) (res []dao.Dict, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.Dict{})
	if val, ok := condition["dictType"]; ok {
		builder.Where("dict_type = ?", val)
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

// DictListTotal 查询列表数据总量
func DictListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.Dict{})
	if val, ok := condition["dictType"]; ok {
		builder.Where("dict_type = ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
