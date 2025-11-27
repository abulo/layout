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

// sys_logger_operate 操作日志
// SysLoggerOperateCreate 创建数据
func SysLoggerOperateCreate(ctx context.Context, data dao.SysLoggerOperate) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.SysLoggerOperate{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysLoggerOperateUpdate 更新数据
func SysLoggerOperateUpdate(ctx context.Context, id int64, data dao.SysLoggerOperate) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.SysLoggerOperate{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysLoggerOperateDelete 删除数据
func SysLoggerOperateDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysLoggerOperate
	data.Id = proto.Int64(id)
	data.Deleted = proto.Int32(1)
	result := db.WithContext(ctx).Model(&dao.SysLoggerOperate{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysLoggerOperate 查询单条数据
func SysLoggerOperate(ctx context.Context, id int64) (res dao.SysLoggerOperate, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysLoggerOperate{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysLoggerOperateRecover 恢复数据
func SysLoggerOperateRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysLoggerOperate
	data.Id = proto.Int64(id)
	data.Deleted = proto.Int32(0)
	result := db.WithContext(ctx).Model(&dao.SysLoggerOperate{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysLoggerOperateDrop 清理数据
func SysLoggerOperateDrop(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysLoggerOperate
	result := db.WithContext(ctx).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysLoggerOperateList 查询列表数据
func SysLoggerOperateList(ctx context.Context, condition map[string]any) (res []dao.SysLoggerOperate, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysLoggerOperate{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("deleted = ?", val)
	}
	if val, ok := condition["result"]; ok {
		builder.Where("result = ?", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("channel = ?", val)
	}
	if val, ok := condition["method"]; ok {
		builder.Where("method = ?", val)
	}
	if val, ok := condition["startTime"]; ok {
		builder.Where("start_time = ?", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("username = ?", val)
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

// SysLoggerOperateListTotal 查询列表数据总量
func SysLoggerOperateListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysLoggerOperate{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("deleted = ?", val)
	}
	if val, ok := condition["result"]; ok {
		builder.Where("result = ?", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("channel = ?", val)
	}
	if val, ok := condition["method"]; ok {
		builder.Where("method = ?", val)
	}
	if val, ok := condition["startTime"]; ok {
		builder.Where("start_time = ?", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("username = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
