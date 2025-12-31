package role

import (
	"cloud/dao"
	"cloud/initial"
	"context"
	"encoding/json"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// sys_role 角色
// SysRoleCreate 创建数据
func SysRoleCreate(ctx context.Context, data dao.SysRole) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Write()
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: false}).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = tx.WithContext(ctx).Model(&dao.SysRole{}).Create(&data).Error
	res = cast.ToInt64(data.Id)

	var menuIds []int64
	var menuList []dao.SysRoleMenu
	err = json.Unmarshal(data.MenuIds.JSON, &menuIds)
	if err != nil {
		tx.Rollback()
		return
	}
	for _, menuId := range menuIds {
		if menuId == 0 {
			continue
		}
		menuList = append(menuList, dao.SysRoleMenu{
			RoleId:   data.Id,
			MenuId:   proto.Int64(menuId),
			TenantId: data.TenantId,
		})
	}
	if len(menuList) > 0 {
		err = tx.WithContext(ctx).Model(&dao.SysRoleMenu{}).Create(menuList).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	err = tx.Commit().Error
	return
}

// SysRoleUpdate 更新数据
func SysRoleUpdate(ctx context.Context, id int64, data dao.SysRole) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Write()
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: false}).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	data.Id = proto.Int64(id)
	result := tx.WithContext(ctx).Model(&dao.SysRole{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		tx.Rollback()
		err = result.Error
		return
	}
	res = cast.ToInt64(result.RowsAffected)

	var menuIds []int64
	var menuList []dao.SysRoleMenu
	err = json.Unmarshal(data.MenuIds.JSON, &menuIds)
	if err != nil {
		tx.Rollback()
		return
	}
	for _, menuId := range menuIds {
		if menuId == 0 {
			continue
		}
		menuList = append(menuList, dao.SysRoleMenu{
			RoleId:   data.Id,
			MenuId:   proto.Int64(menuId),
			TenantId: data.TenantId,
		})
	}
	err = tx.WithContext(ctx).Model(&dao.SysRoleMenu{}).Where("tenant_id = ?", data.TenantId).Where("role_id = ?", id).Delete(&dao.SysUserDept{}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	if len(menuList) > 0 {
		err = tx.WithContext(ctx).Model(&dao.SysRoleMenu{}).Create(&menuList).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	err = tx.Commit().Error
	return
}

// SysRoleDelete 删除数据
func SysRoleDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Write()
	var data dao.SysRole
	data.Id = proto.Int64(id)
	data.Deleted = proto.Int32(1)
	result := db.WithContext(ctx).Model(&dao.SysRole{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysRole 查询单条数据
func SysRole(ctx context.Context, id int64) (res dao.SysRole, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Read()
	err = db.WithContext(ctx).Model(&dao.SysRole{}).Select("sys_role.*", "JSON_ARRAYAGG(sys_role_menu.menu_id) AS menu_ids").Joins("LEFT JOIN sys_role_menu ON sys_role.id = sys_role_menu.role_id  AND sys_role.tenant_id = sys_role_menu.tenant_id").Where("sys_role.id = ?", id).Group("sys_role.id").Find(&res).Error
	return
}

// SysRoleRecover 恢复数据
func SysRoleRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Write()
	var data dao.SysRole
	data.Id = proto.Int64(id)
	data.Deleted = proto.Int32(0)
	result := db.WithContext(ctx).Model(&dao.SysRole{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysRoleDrop 清理数据
func SysRoleDrop(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Write()
	var data dao.SysRole
	result := db.WithContext(ctx).Model(&dao.SysRole{}).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysRoleCode 查询单条数据
func SysRoleCode(ctx context.Context, condition map[string]any) (res dao.SysRole, err error) {
	if util.Empty(condition) {
		err = errors.New("condition is empty")
		return
	}
	db := initial.Core.Store.LoadSQL("postgres").Read()
	builder := db.WithContext(ctx).Model(&dao.SysRole{}).Select("sys_role.*", "JSON_ARRAYAGG(sys_role_menu.menu_id) AS menu_ids").Joins("LEFT JOIN sys_role_menu ON sys_role.id = sys_role_menu.role_id  AND sys_role.tenant_id = sys_role_menu.tenant_id")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("sys_role.tenant_id = ?", val)
	}
	if val, ok := condition["code"]; ok {
		builder.Where("sys_role.code = ?", val)
	}
	builder.Group("sys_role.id")

	err = builder.Find(&res).Error
	return
}

// SysRoleList 查询列表数据
func SysRoleList(ctx context.Context, condition map[string]any) (res []dao.SysRole, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Read()
	builder := db.WithContext(ctx).Model(&dao.SysRole{}).Select("sys_role.*", "JSON_ARRAYAGG(sys_role_menu.menu_id) AS menu_ids").Joins("LEFT JOIN sys_role_menu ON sys_role.id = sys_role_menu.role_id  AND sys_role.tenant_id = sys_role_menu.tenant_id")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("sys_role.tenant_id = ?", val)
	}
	if val, ok := condition["code"]; ok {
		builder.Where("sys_role.code = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("sys_role.deleted = ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("sys_role.status = ?", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("sys_role.name = ?", val)
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
	builder.Group("sys_role.id")
	builder.Order(clause.OrderBy{Columns: []clause.OrderByColumn{
		{Column: clause.Column{Name: "sys_role.sort"}, Desc: false},
		{Column: clause.Column{Name: "sys_role.id"}, Desc: true},
	}})

	err = builder.Find(&res).Error
	return
}

// SysRoleListTotal 查询列表数据总量
func SysRoleListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("postgres").Read()
	builder := db.WithContext(ctx).Model(&dao.SysRole{})
	if val, ok := condition["tenantId"]; ok {
		builder.Where("tenant_id = ?", val)
	}
	if val, ok := condition["code"]; ok {
		builder.Where("code = ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("deleted = ?", val)
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
