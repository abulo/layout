package menu

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"
)

// sys_menu 菜单
// SysMenuCreate 创建数据
func SysMenuCreate(ctx context.Context, data dao.SysMenu) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	err = db.WithContext(ctx).Model(&dao.SysMenu{}).Create(&data).Error
	res = cast.ToInt64(data.Id)
	return
}

// SysMenuUpdate 更新数据
func SysMenuUpdate(ctx context.Context, id int64, data dao.SysMenu) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.SysMenu{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysMenuDelete 删除数据
func SysMenuDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysMenu
	result := db.WithContext(ctx).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysMenu 查询单条数据
func SysMenu(ctx context.Context, id int64) (res dao.SysMenu, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysMenu{}).Where("id = ?", id).Find(&res).Error
	return
}

// SysMenuParent 查询列表数据
func SysMenuParent(ctx context.Context, condition map[string]any) (res []dao.SysMenu, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysMenu{})
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["type"]; ok {
		builder.Where("type = ?", val)
	}
	if val, ok := condition["parentId"]; ok {
		builder.Where("parent_id = ?", val)
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

// SysMenuParentTotal 查询列表数据总量
func SysMenuParentTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysMenu{})
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["type"]; ok {
		builder.Where("type = ?", val)
	}
	if val, ok := condition["parentId"]; ok {
		builder.Where("parent_id = ?", val)
	}

	err = builder.Count(&res).Error
	return
}

// SysMenuList 查询列表数据
func SysMenuList(ctx context.Context, condition map[string]any) (res []dao.SysMenu, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysMenu{})
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["type"]; ok {
		builder.Where("type = ?", val)
	}
	if val, ok := condition["parentId"]; ok {
		builder.Where("parent_id = ?", val)
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
		{Column: clause.Column{Name: "sort"}, Desc: false},
		{Column: clause.Column{Name: "parent_id"}, Desc: false},
		{Column: clause.Column{Name: "id"}, Desc: true},
	}})
	err = builder.Find(&res).Error
	return
}

// SysMenuListTotal 查询列表数据总量
func SysMenuListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysMenu{})
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["type"]; ok {
		builder.Where("type = ?", val)
	}
	if val, ok := condition["parentId"]; ok {
		builder.Where("parent_id = ?", val)
	}

	err = builder.Count(&res).Error
	return
}

// SystemMenuListRecursive 递归查询向上查询, 用于在计划任务中使用
func SysMenuListRecursive(ctx context.Context, id int64) (res []dao.SysMenu, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	db.WithContext(ctx).Raw("WITH RECURSIVE filter_sys_menu (id,name,code,type,sort,parent_id,path,icon,component,component_name,hide,link,cache,remark,active,full,redirect,status,creator,create_time,updater,update_time) AS ( SELECT * FROM sys_menu WHERE id = ? UNION ALL SELECT t.* FROM sys_menu t INNER JOIN filter_sys_menu ON filter_sys_menu.parent_id=t.id) SELECT DISTINCT * FROM filter_sys_menu ORDER BY filter_sys_menu.sort ASC,filter_sys_menu.parent_id ASC,filter_sys_menu.id DESC", id).Scan(&res)
	err = nil
	return
}
