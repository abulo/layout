package role

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

// sys_role_menu 角色菜单
// SysRoleMenuCreate 创建数据
func SysRoleMenuCreate(ctx context.Context, data dao.SysRoleMenu) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.SysRoleMenu{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysRoleMenuUpdate 更新数据
func SysRoleMenuUpdate(ctx context.Context, id int64, data dao.SysRoleMenu) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.SysRoleMenu{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysRoleMenuDelete 删除数据
func SysRoleMenuDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysRoleMenu
	result := db.WithContext(ctx).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysRoleMenu 查询单条数据
func SysRoleMenu(ctx context.Context, id int64) (res dao.SysRoleMenu, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysRoleMenu{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysRoleMenuItem 查询单条数据
func SysRoleMenuItem(ctx context.Context, condition map[string]any) (res dao.SysRoleMenu, err error) {
	if util.Empty(condition) {
		err = errors.New("condition is empty")
		return
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysRoleMenu{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["roleId"]; ok {
		builder.Where("role_id = ?", val)
	}
	if val, ok := condition["menuId"]; ok {
		builder.Where("menu_id = ?", val)
	}

	err = builder.Find(&res).Error
	return
}

// SysRoleMenuList 查询列表数据
func SysRoleMenuList(ctx context.Context, condition map[string]any) (res []dao.SysRoleMenu, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysRoleMenu{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["roleId"]; ok {
		builder.Where("role_id = ?", val)
	}
	if val, ok := condition["menuId"]; ok {
		builder.Where("menu_id = ?", val)
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

// SysRoleMenuListTotal 查询列表数据总量
func SysRoleMenuListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysRoleMenu{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["roleId"]; ok {
		builder.Where("role_id = ?", val)
	}
	if val, ok := condition["menuId"]; ok {
		builder.Where("menu_id = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
