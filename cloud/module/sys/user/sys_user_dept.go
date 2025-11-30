package user

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

// sys_user_dept 用户部门
// SysUserDeptCreate 创建数据
func SysUserDeptCreate(ctx context.Context, data dao.SysUserDept) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.SysUserDept{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysUserDeptUpdate 更新数据
func SysUserDeptUpdate(ctx context.Context, id int64, data dao.SysUserDept) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.SysUserDept{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysUserDeptDelete 删除数据
func SysUserDeptDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysUserDept
	result := db.WithContext(ctx).Model(&dao.SysUserDept{}).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysUserDept 查询单条数据
func SysUserDept(ctx context.Context, id int64) (res dao.SysUserDept, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysUserDept{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysUserDeptItem 查询单条数据
func SysUserDeptItem(ctx context.Context, condition map[string]any) (res dao.SysUserDept, err error) {
	if util.Empty(condition) {
		err = errors.New("condition is empty")
		return
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUserDept{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["deptId"]; ok {
		builder.Where("dept_id = ?", val)
	}

	err = builder.Find(&res).Error
	return
}

// SysUserDeptList 查询列表数据
func SysUserDeptList(ctx context.Context, condition map[string]any) (res []dao.SysUserDept, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUserDept{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["deptId"]; ok {
		builder.Where("dept_id = ?", val)
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

// SysUserDeptListTotal 查询列表数据总量
func SysUserDeptListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUserDept{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["deptId"]; ok {
		builder.Where("dept_id = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
