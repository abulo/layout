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
)

// dict_type 字典类型
// DictTypeCreate 创建数据
func DictTypeCreate(ctx context.Context, data dao.DictType) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.DictType{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// DictTypeUpdate 更新数据
func DictTypeUpdate(ctx context.Context, id int64, data dao.DictType) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.DictType{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// DictTypeDelete 删除数据
func DictTypeDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.DictType
	result := db.WithContext(ctx).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// DictType 查询单条数据
func DictType(ctx context.Context, id int64) (res dao.DictType, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.DictType{}).Where("id = ?", id).Find(&res).Error
	return
}

// DictTypeType 查询单条数据
func DictTypeType(ctx context.Context, condition map[string]any) (res dao.DictType, err error) {
	if util.Empty(condition) {
		err = errors.New("condition is empty")
		return
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.DictType{})
	if val, ok := condition["type"]; ok {
		builder.Where("type = ?", val)
	}

	err = builder.Find(&res).Error
	return
}

// DictTypeList 查询列表数据
func DictTypeList(ctx context.Context, condition map[string]any) (res []dao.DictType, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.DictType{})
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
	builder.Order("id")
	err = builder.Find(&res).Error
	return
}

// DictTypeListTotal 查询列表数据总量
func DictTypeListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.DictType{})
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
