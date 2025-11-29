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
	if util.Empty(condition["userId"]) || util.Empty(condition["scope"]) || util.Empty(condition["scopeDept"]) || util.Empty(condition["tenantId"]) {
		return nil, errors.New("userId or scope or scopeDept or tenantId is empty")
	}
	scopeDept := condition["scopeDept"].([]int64)
	scope := cast.ToInt32(condition["scope"])
	userId := cast.ToInt64(condition["userId"])
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysLoggerOperate{}).Select("sys_logger_operate.*").Joins("LEFT JOIN sys_user_dept ON sys_user_dept.tenant_id = sys_logger_operate.tenant_id  AND sys_logger_operate.user_id = sys_user_dept.user_id")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("sys_logger_operate.tenant_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("sys_logger_operate.deleted = ?", val)
	}
	if val, ok := condition["result"]; ok {
		builder.Where("sys_logger_operate.result = ?", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("sys_logger_operate.channel = ?", val)
	}
	if val, ok := condition["method"]; ok {
		builder.Where("sys_logger_operate.method = ?", val)
	}
	if val, ok := condition["beginStartTime"]; ok {
		builder.Where("sys_logger_operate.start_time >= ?", val)
	}
	if val, ok := condition["finishStartTime"]; ok {
		builder.Where("sys_logger_operate.start_time <= ?", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("sys_logger_operate.username = ?", val)
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
	builder.Group("`sys_logger_operate`.`id`")
	builder.Order(clause.OrderBy{Columns: []clause.OrderByColumn{
		{Column: clause.Column{Name: "sys_logger_operate.id"}, Desc: true},
	}})
	err = builder.Find(&res).Error
	return
}

// SysLoggerOperateListTotal 查询列表数据总量
func SysLoggerOperateListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	if util.Empty(condition["userId"]) || util.Empty(condition["scope"]) || util.Empty(condition["scopeDept"]) || util.Empty(condition["tenantId"]) {
		return 0, errors.New("userId or scope or scopeDept or tenantId is empty")
	}
	scopeDept := condition["scopeDept"].([]int64)
	scope := cast.ToInt32(condition["scope"])
	userId := cast.ToInt64(condition["userId"])
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysLoggerOperate{}).Joins("LEFT JOIN sys_user_dept ON sys_user_dept.tenant_id = sys_logger_operate.tenant_id  AND sys_logger_operate.user_id = sys_user_dept.user_id")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("sys_logger_operate.tenant_id = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("sys_logger_operate.deleted = ?", val)
	}
	if val, ok := condition["result"]; ok {
		builder.Where("sys_logger_operate.result = ?", val)
	}
	if val, ok := condition["channel"]; ok {
		builder.Where("sys_logger_operate.channel = ?", val)
	}
	if val, ok := condition["method"]; ok {
		builder.Where("sys_logger_operate.method = ?", val)
	}
	if val, ok := condition["beginStartTime"]; ok {
		builder.Where("sys_logger_operate.start_time >= ?", val)
	}
	if val, ok := condition["finishStartTime"]; ok {
		builder.Where("sys_logger_operate.start_time <= ?", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("sys_logger_operate.username = ?", val)
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
	builder.Group("`sys_logger_operate`.`id`")

	err = db.Table("(?) as sys_logger_operate_table", builder).Count(&res).Error
	return
}
