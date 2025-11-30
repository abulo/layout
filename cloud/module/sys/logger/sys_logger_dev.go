package logger

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"
)

// sys_logger_dev 开发日志
// SysLoggerDevCreate 创建数据
func SysLoggerDevCreate(ctx context.Context, data dao.SysLoggerDev) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.SysLoggerDev{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysLoggerDevUpdate 更新数据
func SysLoggerDevUpdate(ctx context.Context, id int64, data dao.SysLoggerDev) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.SysLoggerDev{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysLoggerDevDelete 删除数据
func SysLoggerDevDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysLoggerDev
	result := db.WithContext(ctx).Model(&dao.SysLoggerDev{}).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysLoggerDev 查询单条数据
func SysLoggerDev(ctx context.Context, id int64) (res dao.SysLoggerDev, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysLoggerDev{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysLoggerDevList 查询列表数据
func SysLoggerDevList(ctx context.Context, condition map[string]any) (res []dao.SysLoggerDev, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysLoggerDev{})
	if val, ok := condition["beginTimestamp"]; ok {
		builder.Where("timestamp >= ?", val)
	}
	if val, ok := condition["finishTimestamp"]; ok {
		builder.Where("timestamp <= ?", val)
	}
	if val, ok := condition["host"]; ok {
		builder.Where("host = ?", val)
	}
	if val, ok := condition["level"]; ok {
		builder.Where("level = ?", val)
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

// SysLoggerDevListTotal 查询列表数据总量
func SysLoggerDevListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysLoggerDev{})
	if val, ok := condition["beginTimestamp"]; ok {
		builder.Where("timestamp >= ?", val)
	}
	if val, ok := condition["finishTimestamp"]; ok {
		builder.Where("timestamp <= ?", val)
	}
	if val, ok := condition["host"]; ok {
		builder.Where("host = ?", val)
	}
	if val, ok := condition["level"]; ok {
		builder.Where("level = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
