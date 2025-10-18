package logger

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
)

// sys_logger_login 登录日志
// SysLoggerLoginCreate 创建数据
func SysLoggerLoginCreate(ctx context.Context, data dao.SysLoggerLogin) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.SysLoggerLogin{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysLoggerLoginUpdate 更新数据
func SysLoggerLoginUpdate(ctx context.Context, id int64, data dao.SysLoggerLogin) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.SysLoggerLogin{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysLoggerLoginDelete 删除数据
func SysLoggerLoginDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysLoggerLogin
	data.Id = proto.Int64(id)
	data.Deleted = proto.Int32(1)
	result := db.WithContext(ctx).Model(&dao.SysLoggerLogin{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysLoggerLogin 查询单条数据
func SysLoggerLogin(ctx context.Context, id int64) (res dao.SysLoggerLogin, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysLoggerLogin{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysLoggerLoginRecover 恢复数据
func SysLoggerLoginRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysLoggerLogin
	data.Id = proto.Int64(id)
	data.Deleted = proto.Int32(0)
	result := db.WithContext(ctx).Model(&dao.SysLoggerLogin{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysLoggerLoginDrop 清理数据
func SysLoggerLoginDrop(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysLoggerLogin
	result := db.WithContext(ctx).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysLoggerLoginList 查询列表数据
func SysLoggerLoginList(ctx context.Context, condition map[string]any) (res []dao.SysLoggerLogin, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysLoggerLogin{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("deleted = ?", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("channel = ?", val)
	}
	if val, ok := condition["loginTime"]; ok {
		builder.Where("login_time = ?", val)
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
	builder.Order("id")
	err = builder.Find(&res).Error
	return
}

// SysLoggerLoginListTotal 查询列表数据总量
func SysLoggerLoginListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysLoggerLogin{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("deleted = ?", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("channel = ?", val)
	}
	if val, ok := condition["loginTime"]; ok {
		builder.Where("login_time = ?", val)
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
