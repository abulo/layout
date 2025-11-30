package tenant

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/null"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// sys_tenant 租户
// SysTenantCreate 创建数据
func SysTenantCreate(ctx context.Context, data dao.SysTenant) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: false}).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = tx.WithContext(ctx).Model(&dao.SysTenant{}).Create(&data).Error
	if err != nil {
		tx.Rollback()
		return
	}
	res = cast.ToInt64(data.Id)
	// 插入租户管理员信息
	var user dao.SysUser
	user.Name = data.ContactName
	user.Mobile = data.ContactMobile
	user.Username = data.Username
	user.Password = data.Password
	user.Creator = data.Creator
	user.CreateTime = data.CreateTime
	user.Updater = data.Updater
	user.UpdateTime = data.UpdateTime
	user.TenantId = data.Id
	user.Status = proto.Int32(0)
	user.Deleted = proto.Int32(0)
	err = tx.WithContext(ctx).Model(&dao.SysUser{}).Create(&user).Error
	if err != nil {
		tx.Rollback()
		return
	}
	// 绑定用户和租户的关系
	var userTenant dao.SysUserTenant
	userTenant.UserId = user.Id
	userTenant.TenantId = data.Id
	err = tx.WithContext(ctx).Model(&dao.SysUserTenant{}).Create(&userTenant).Error
	if err != nil {
		tx.Rollback()
		return
	}
	//创建部门
	var dept dao.SysDept
	dept.Name = data.Name
	dept.ParentId = proto.Int64(0)
	dept.Sort = proto.Int32(1)
	dept.UserId = null.Int64From(cast.ToInt64(user.Id))
	dept.Status = proto.Int32(0)
	dept.TenantId = data.Id
	dept.Deleted = proto.Int32(0)
	dept.Creator = data.Creator
	dept.CreateTime = data.CreateTime
	dept.Updater = data.Updater
	dept.UpdateTime = data.UpdateTime
	err = tx.WithContext(ctx).Model(&dao.SysDept{}).Create(&dept).Error
	if err != nil {
		tx.Rollback()
		return
	}
	// 创建部门关系
	var userDept dao.SysUserDept
	userDept.DeptId = dept.Id
	userDept.UserId = user.Id
	userDept.TenantId = data.Id
	err = tx.WithContext(ctx).Model(&dao.SysUserDept{}).Create(&userDept).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}

// SysTenantUpdate 更新数据
func SysTenantUpdate(ctx context.Context, id int64, data dao.SysTenant) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data.Id = proto.Int64(id)
	result := db.WithContext(ctx).Model(&dao.SysTenant{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysTenantDelete 删除数据
func SysTenantDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysTenant
	data.Id = proto.Int64(id)
	data.Deleted = proto.Int32(1)
	result := db.WithContext(ctx).Model(&dao.SysTenant{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysTenant 查询单条数据
func SysTenant(ctx context.Context, id int64) (res dao.SysTenant, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysTenant{}).Select("`sys_tenant`.*", "`sys_user`.username", "`sys_user`.`password`").Joins("LEFT JOIN `sys_user` ON `sys_tenant`.user_id = `sys_user`.id").Where("`sys_tenant`.id = ?", id).Find(&res).Error
	return
}

// SysTenantRecover 恢复数据
func SysTenantRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysTenant
	data.Id = proto.Int64(id)
	data.Deleted = proto.Int32(0)
	result := db.WithContext(ctx).Model(&dao.SysTenant{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysTenantDrop 清理数据
func SysTenantDrop(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysTenant
	result := db.WithContext(ctx).Model(&dao.SysTenant{}).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysTenantList 查询列表数据
func SysTenantList(ctx context.Context, condition map[string]any) (res []dao.SysTenant, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysTenant{}).Select("`sys_tenant`.*", "`sys_user`.username", "`sys_user`.password").Joins("LEFT JOIN `sys_user` ON `sys_tenant`.user_id = `sys_user`.id")

	if val, ok := condition["deleted"]; ok {
		builder.Where("`sys_tenant`.deleted = ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`sys_tenant`.status = ?", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("`sys_tenant`.`name` LIKE ?", "%"+cast.ToString(val)+"%")
	}
	if val, ok := condition["beginExpireDate"]; ok {
		builder.Where("`sys_tenant`.`expire_date` >= ?", val)
	}
	if val, ok := condition["finishExpireDate"]; ok {
		builder.Where("`sys_tenant`.`expire_date` <= ?", val)
	}
	if val, ok := condition["packageId"]; ok {
		builder.Where("`sys_tenant`.package_id = ?", val)
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
		{Column: clause.Column{Name: "sys_tenant.id"}, Desc: true},
	}})
	err = builder.Find(&res).Error
	return
}

// SysTenantListTotal 查询列表数据总量
func SysTenantListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysTenant{})
	if val, ok := condition["deleted"]; ok {
		builder.Where("deleted = ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("status = ?", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Where("`name` LIKE ?", "%"+cast.ToString(val)+"%")
	}
	if val, ok := condition["beginExpireDate"]; ok {
		builder.Where("`expire_date` >= ?", val)
	}
	if val, ok := condition["finishExpireDate"]; ok {
		builder.Where("`expire_date` <= ?", val)
	}
	if val, ok := condition["packageId"]; ok {
		builder.Where("package_id = ?", val)
	}

	err = builder.Count(&res).Error
	return
}
