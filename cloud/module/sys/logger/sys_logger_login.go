package logger

import (
	"cloud/dao"
	"cloud/initial"
	"cloud/module/sys/user"
	"context"
	"errors"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"
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
	if util.Empty(condition["userId"]) || util.Empty(condition["scope"]) || util.Empty(condition["scopeDept"]) || util.Empty(condition["tenantId"]) {
		return nil, errors.New("userId or scope or scopeDept or tenantId is empty")
	}
	scopeDept := condition["scopeDept"].([]int64)
	scope := cast.ToInt32(condition["scope"])
	userId := cast.ToInt64(condition["userId"])
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysLoggerLogin{}).Select("sys_logger_login.*").Joins("LEFT JOIN sys_user_dept ON sys_user_dept.tenant_id = sys_logger_login.tenant_id  AND sys_logger_login.user_id = sys_user_dept.user_id")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("sys_logger_login.tenant_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("sys_logger_login.deleted = ?", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("sys_logger_login.channel = ?", val)
	}
	if val, ok := condition["beginLoginTime"]; ok {
		builder.Where("sys_logger_login.login_time >= ?", val)
	}
	if val, ok := condition["finishLoginTime"]; ok {
		builder.Where("sys_logger_login.login_time <= ?", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("sys_logger_login.username = ?", val)
	}
	builder.Where("sys_user_dept.dept_id IN ? ", scopeDept)
	switch scope {
	case 1: // 全部数据
	case 2: // 指定部门数据
	case 4: // 本部门及以下数据
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				if deptList, err := user.SysDeptRecursive(ctx, deptId); err == nil {
					builder.Where("sys_user_dept.dept_id IN ? ", deptList)
				}
			}
		}
	case 3: // 本部门数据
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				builder.Where("sys_user_dept.dept_id= ?", deptId)
			}
		}
	case 5: // 仅本人数据
		builder.Where("sys_user_dept.user_id= ?", userId)
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
	builder.Group("`sys_logger_login`.`id`")
	builder.Order(clause.OrderBy{Columns: []clause.OrderByColumn{
		{Column: clause.Column{Name: "sys_logger_login.id"}, Desc: true},
	}})
	err = builder.Find(&res).Error
	return
}

// SysLoggerLoginListTotal 查询列表数据总量
func SysLoggerLoginListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	if util.Empty(condition["userId"]) || util.Empty(condition["scope"]) || util.Empty(condition["scopeDept"]) || util.Empty(condition["tenantId"]) {
		return 0, errors.New("userId or scope or scopeDept or tenantId is empty")
	}
	scopeDept := condition["scopeDept"].([]int64)
	scope := cast.ToInt32(condition["scope"])
	userId := cast.ToInt64(condition["userId"])
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysLoggerLogin{}).Joins("LEFT JOIN sys_user_dept ON sys_user_dept.tenant_id = sys_logger_login.tenant_id  AND sys_logger_login.user_id = sys_user_dept.user_id")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("sys_logger_login.tenant_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("sys_logger_login.deleted = ?", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("sys_logger_login.channel = ?", val)
	}
	if val, ok := condition["beginLoginTime"]; ok {
		builder.Where("sys_logger_login.login_time >= ?", val)
	}
	if val, ok := condition["finishLoginTime"]; ok {
		builder.Where("sys_logger_login.login_time <= ?", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("sys_logger_login.username = ?", val)
	}
	builder.Where("sys_user_dept.dept_id IN ? ", scopeDept)
	switch scope {
	case 1: // 全部数据
	case 2: // 指定部门数据
	case 4: // 本部门及以下数据
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				if deptList, err := user.SysDeptRecursive(ctx, deptId); err == nil {
					builder.Where("sys_user_dept.dept_id IN ? ", deptList)
				}
			}
		}
	case 3: // 本部门数据
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				builder.Where("sys_user_dept.dept_id= ?", deptId)
			}
		}
	case 5: // 仅本人数据
		builder.Where("sys_user_dept.user_id= ?", userId)
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
	builder.Group("`sys_logger_login`.`id`")
	err = db.Table("(?) as sys_logger_login_table", builder).Count(&res).Error
	return
}
