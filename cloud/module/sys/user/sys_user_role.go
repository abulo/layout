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

// sys_user_role 用户角色
// SysUserRoleCreate 创建数据
func SysUserRoleCreate(ctx context.Context, data dao.SysUserRole) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.SysUserRole{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysUserRoleUpdate 更新数据
func SysUserRoleUpdate(ctx context.Context, id int64, data dao.SysUserRole) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.SysUserRole{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysUserRoleDelete 删除数据
func SysUserRoleDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysUserRole
	result := db.WithContext(ctx).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysUserRole 查询单条数据
func SysUserRole(ctx context.Context, id int64) (res dao.SysUserRole, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysUserRole{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysUserRoleItem 查询单条数据
func SysUserRoleItem(ctx context.Context, condition map[string]any) (res dao.SysUserRole, err error) {
	if util.Empty(condition) {
		err = errors.New("condition is empty")
		return
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUserRole{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["roleId"]; ok {
		builder.Where("role_id = ?", val)
	}

	err = builder.Find(&res).Error
	return
}

// SysUserRoleList 查询列表数据
func SysUserRoleList(ctx context.Context, condition map[string]any) (res []dao.SysUserRole, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUserRole{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["roleId"]; ok {
		builder.Where("role_id = ?", val)
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

// SysUserRoleListTotal 查询列表数据总量
func SysUserRoleListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUserRole{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["userId"]; ok {
		builder.Where("user_id = ?", val)
	}
	if val, ok := condition["roleId"]; ok {
		builder.Where("role_id = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
