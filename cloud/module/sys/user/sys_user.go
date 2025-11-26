package user

import (
	"cloud/dao"
	"cloud/initial"
	"cloud/module/sys/dept"
	"cloud/module/sys/tenant"
	"context"
	"encoding/json"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/abulo/ratel/v3/util"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

// sys_user 用户
// SysUserCreate 创建数据
func SysUserCreate(ctx context.Context, data dao.SysUser) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: false}).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = tx.WithContext(ctx).Model(&dao.SysUser{}).Create(&data).Error
	if err != nil {
		tx.Rollback()
		return
	}
	res = cast.ToInt64(data.Id)
	var userDeptIds []int64
	var userDeptList []dao.SysUserDept
	err = json.Unmarshal(data.DeptIds.JSON, &userDeptIds)
	if err != nil {
		tx.Rollback()
		return
	}
	for _, deptId := range userDeptIds {
		if deptId == 0 {
			continue
		}
		userDeptList = append(userDeptList, dao.SysUserDept{
			UserId:   data.Id,
			DeptId:   proto.Int64(deptId),
			TenantId: data.TenantId,
		})
	}
	if len(userDeptList) > 0 {
		err = tx.WithContext(ctx).Model(&dao.SysUserDept{}).Create(&userDeptList).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	var userPostList []dao.SysUserPost
	var userPostIds []int64
	err = json.Unmarshal(data.PostIds.JSON, &userPostIds)
	if err != nil {
		tx.Rollback()
		return
	}
	for _, postId := range userPostIds {
		if postId == 0 {
			continue
		}
		userPostList = append(userPostList, dao.SysUserPost{
			UserId:   data.Id,
			PostId:   proto.Int64(postId),
			TenantId: data.TenantId,
		})
	}
	if len(userPostList) > 0 {
		err = tx.WithContext(ctx).Model(&dao.SysUserPost{}).Create(&userPostList).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	var userRoleList []dao.SysUserRole
	var userRoleIds []int64
	err = json.Unmarshal(data.RoleIds.JSON, &userRoleIds)
	if err != nil {
		tx.Rollback()
		return
	}
	for _, roleId := range userRoleIds {
		if roleId == 0 {
			continue
		}
		userRoleList = append(userRoleList, dao.SysUserRole{
			UserId:   data.Id,
			RoleId:   proto.Int64(roleId),
			TenantId: data.TenantId,
		})
	}
	if len(userRoleList) > 0 {
		err = tx.WithContext(ctx).Model(&dao.SysUserRole{}).Create(&userRoleList).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	err = tx.Commit().Error
	return
}

// SysUserUpdate 更新数据
func SysUserUpdate(ctx context.Context, id int64, data dao.SysUser) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: false}).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	data.Id = proto.Int64(id)
	result := tx.WithContext(ctx).Model(&dao.SysUser{}).Where("id = ?", id).Updates(data)

	if result.Error != nil {
		tx.Rollback()
		err = result.Error
		return
	}
	res = cast.ToInt64(result.RowsAffected)

	var userDeptIds []int64
	var userDeptList []dao.SysUserDept
	err = json.Unmarshal(data.DeptIds.JSON, &userDeptIds)
	if err != nil {
		tx.Rollback()
		return
	}
	for _, deptId := range userDeptIds {
		if deptId == 0 {
			continue
		}
		userDeptList = append(userDeptList, dao.SysUserDept{
			UserId:   data.Id,
			DeptId:   proto.Int64(deptId),
			TenantId: data.TenantId,
		})
	}

	err = tx.WithContext(ctx).Model(&dao.SysUserDept{}).Where("tenant_id = ?", data.TenantId).Where("user_id = ?", id).Delete(&dao.SysUserDept{}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	if len(userDeptList) > 0 {
		err = tx.WithContext(ctx).Model(&dao.SysUserDept{}).Create(&userDeptList).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	var userPostList []dao.SysUserPost
	var userPostIds []int64
	err = json.Unmarshal(data.PostIds.JSON, &userPostIds)
	if err != nil {
		tx.Rollback()
		return
	}
	for _, postId := range userPostIds {
		if postId == 0 {
			continue
		}
		userPostList = append(userPostList, dao.SysUserPost{
			UserId:   data.Id,
			PostId:   proto.Int64(postId),
			TenantId: data.TenantId,
		})
	}
	err = tx.WithContext(ctx).Model(&dao.SysUserPost{}).Where("tenant_id = ?", data.TenantId).Where("user_id = ?", id).Delete(&dao.SysUserPost{}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	if len(userPostList) > 0 {
		err = tx.WithContext(ctx).Model(&dao.SysUserPost{}).Create(&userPostList).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	var userRoleList []dao.SysUserRole
	var userRoleIds []int64
	err = json.Unmarshal(data.RoleIds.JSON, &userRoleIds)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.WithContext(ctx).Model(&dao.SysUserRole{}).Where("tenant_id = ?", data.TenantId).Where("user_id = ?", id).Delete(&dao.SysUserRole{}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	for _, roleId := range userRoleIds {
		if roleId == 0 {
			continue
		}
		userRoleList = append(userRoleList, dao.SysUserRole{
			UserId:   data.Id,
			RoleId:   proto.Int64(roleId),
			TenantId: data.TenantId,
		})
	}
	if len(userRoleList) > 0 {
		err = tx.WithContext(ctx).Model(&dao.SysUserRole{}).Create(&userRoleList).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	err = tx.Commit().Error
	return
}

// SysUserDelete 删除数据
func SysUserDelete(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysUser
	data.Id = proto.Int64(id)
	data.Deleted = proto.Int32(1)
	result := db.WithContext(ctx).Model(&dao.SysUser{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysUser 查询单条数据
func SysUser(ctx context.Context, id int64) (res dao.SysUser, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	err = db.WithContext(ctx).Model(&dao.SysUser{}).Select("`sys_user`.*", "JSON_ARRAYAGG(sys_user_dept.dept_id) AS dept_ids",
		"JSON_ARRAYAGG(sys_user_post.post_id) AS post_ids",
		"JSON_ARRAYAGG(sys_user_role.role_id) AS role_ids").Joins("LEFT JOIN `sys_user_dept` ON `sys_user`.`id` = `sys_user_dept`.`user_id`  AND sys_user.tenant_id = sys_user_dept.tenant_id  LEFT JOIN `sys_dept` ON `sys_dept`.tenant_id = sys_user_dept.tenant_id  AND sys_dept.deleted = 0 LEFT JOIN  `sys_user_post` ON `sys_user`.`id` = `sys_user_post`.`user_id`  AND sys_user.tenant_id = sys_user_post.tenant_id LEFT JOIN  `sys_user_role` ON `sys_user`.`id` = `sys_user_role`.`user_id`  AND sys_user.tenant_id = sys_user_role.tenant_id").Where("sys_user.id = ?", id).Group("`sys_user`.id").Find(&res).Error
	return
}

// SysUserRecover 恢复数据
func SysUserRecover(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysUser
	data.Id = proto.Int64(id)
	data.Deleted = proto.Int32(0)
	result := db.WithContext(ctx).Model(&dao.SysUser{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}

// SysUserDrop 清理数据
func SysUserDrop(ctx context.Context, id int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	var data dao.SysUser
	result := db.WithContext(ctx).Where("id = ?", id).First(&data).Delete(&data)
	return result.RowsAffected, result.Error
}

// SysUserLogin 查询单条数据
func SysUserLogin(ctx context.Context, condition map[string]any) (res dao.SysUser, err error) {
	if util.Empty(condition) {
		err = errors.New("condition is empty")
		return
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUser{}).Select("`sys_user`.*", "JSON_ARRAYAGG(sys_user_dept.dept_id) AS dept_ids",
		"JSON_ARRAYAGG(sys_user_post.post_id) AS post_ids",
		"JSON_ARRAYAGG(sys_user_role.role_id) AS role_ids").Joins("LEFT JOIN `sys_user_dept` ON `sys_user`.`id` = `sys_user_dept`.`user_id`  AND sys_user.tenant_id = sys_user_dept.tenant_id  LEFT JOIN `sys_dept` ON `sys_dept`.tenant_id = sys_user_dept.tenant_id  AND sys_dept.deleted = 0 LEFT JOIN  `sys_user_post` ON `sys_user`.`id` = `sys_user_post`.`user_id`  AND sys_user.tenant_id = sys_user_post.tenant_id LEFT JOIN  `sys_user_role` ON `sys_user`.`id` = `sys_user_role`.`user_id`  AND sys_user.tenant_id = sys_user_role.tenant_id")
	if val, ok := condition["username"]; ok {
		builder.Where("sys_user.username = ?", val)
	}
	err = builder.Find(&res).Error
	return
}

// SysUserList 查询列表数据
func SysUserList(ctx context.Context, condition map[string]any) (res []dao.SysUser, err error) {
	if util.Empty(condition["userId"]) || util.Empty(condition["scope"]) || util.Empty(condition["scopeDept"]) || util.Empty(condition["tenantId"]) {
		return nil, errors.New("userId or scope or scopeDept or tenantId is empty")
	}
	scopeDept := condition["scopeDept"].([]int64)
	scope := cast.ToInt32(condition["scope"])
	userId := cast.ToInt64(condition["userId"])
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUser{}).Select("`sys_user`.*", "JSON_ARRAYAGG(sys_user_dept.dept_id) AS dept_ids",
		"JSON_ARRAYAGG(sys_user_post.post_id) AS post_ids",
		"JSON_ARRAYAGG(sys_user_role.role_id) AS role_ids").Joins("LEFT JOIN `sys_user_dept` ON `sys_user`.`id` = `sys_user_dept`.`user_id`  AND sys_user.tenant_id = sys_user_dept.tenant_id  LEFT JOIN `sys_dept` ON `sys_dept`.tenant_id = sys_user_dept.tenant_id  AND sys_dept.deleted = 0 LEFT JOIN  `sys_user_post` ON `sys_user`.`id` = `sys_user_post`.`user_id`  AND sys_user.tenant_id = sys_user_post.tenant_id LEFT JOIN  `sys_user_role` ON `sys_user`.`id` = `sys_user_role`.`user_id`  AND sys_user.tenant_id = sys_user_role.tenant_id")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`sys_user`.`tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`sys_user`.`deleted`= ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`sys_user`.`status`= ?", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("`sys_user`.username LIKE ?", "%"+cast.ToString(val)+"%")
	}
	if val, ok := condition["mobile"]; ok {
		builder.Where("`sys_user`.mobile LIKE ?", "%"+cast.ToString(val)+"%")
	}
	if val, ok := condition["name"]; ok {
		builder.Where("`sys_user`.name LIKE ?", "%"+cast.ToString(val)+"%")
	}
	builder.Where("sys_user_dept.dept_id IN ? ", scopeDept)
	builder.Group("`sys_user`.id")
	builder.Order("`sys_user`.`id`")

	newBuilder := db.Table("(?) as newTable", builder).Select("`newTable`.*").Joins("LEFT JOIN sys_user_dept ON newTable.id = sys_user_dept.user_id AND newTable.tenant_id = sys_user_dept.tenant_id")
	switch scope {
	case 1: // 全部数据
	case 2: // 指定部门数据
	case 4: // 本部门及以下数据
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				if deptList, err := SysDeptRecursive(ctx, deptId); err == nil {
					newBuilder.Where("sys_user_dept.dept_id IN ? ", deptList)
				}
			}
		}
	case 3: // 本部门数据
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				newBuilder.Where("sys_user_dept.dept_id= ?", deptId)
			}
		}
	case 5: // 仅本人数据
		newBuilder.Where("sys_user_dept.user_id= ?", userId)
	}
	newBuilder.Group("newTable.id")
	if val, ok := condition["pagination"]; ok {
		pagination := val.(*sql.Pagination)
		if val, err := pagination.GetOffset(); err == nil {
			newBuilder.Offset(cast.ToInt(val))
		}
		if val, err := pagination.GetLimit(); err == nil {
			newBuilder.Limit(cast.ToInt(val))
		}
	}
	newBuilder.Scan(&res)
	return
}

// SysUserListTotal 查询列表数据总量
func SysUserListTotal(ctx context.Context, condition map[string]any) (res int64, err error) {
	if util.Empty(condition["userId"]) || util.Empty(condition["scope"]) || util.Empty(condition["scopeDept"]) || util.Empty(condition["tenantId"]) {
		return 0, errors.New("userId or scope or scopeDept or tenantId is empty")
	}
	scopeDept := condition["scopeDept"].([]int64)
	scope := cast.ToInt32(condition["scope"])
	userId := cast.ToInt64(condition["userId"])
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := db.WithContext(ctx).Model(&dao.SysUser{}).Select("`sys_user`.*", "JSON_ARRAYAGG(sys_user_dept.dept_id) AS dept_ids",
		"JSON_ARRAYAGG(sys_user_post.post_id) AS post_ids",
		"JSON_ARRAYAGG(sys_user_role.role_id) AS role_ids").Joins("LEFT JOIN `sys_user_dept` ON `sys_user`.`id` = `sys_user_dept`.`user_id`  AND sys_user.tenant_id = sys_user_dept.tenant_id  LEFT JOIN `sys_dept` ON `sys_dept`.tenant_id = sys_user_dept.tenant_id  AND sys_dept.deleted = 0 LEFT JOIN  `sys_user_post` ON `sys_user`.`id` = `sys_user_post`.`user_id`  AND sys_user.tenant_id = sys_user_post.tenant_id LEFT JOIN  `sys_user_role` ON `sys_user`.`id` = `sys_user_role`.`user_id`  AND sys_user.tenant_id = sys_user_role.tenant_id")
	if val, ok := condition["tenantId"]; ok {
		builder.Where("`sys_user`.`tenant_id`= ?", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`sys_user`.`deleted`= ?", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`sys_user`.`status`= ?", val)
	}
	if val, ok := condition["username"]; ok {
		builder.Where("`sys_user`.username LIKE ?", "%"+cast.ToString(val)+"%")
	}
	if val, ok := condition["mobile"]; ok {
		builder.Where("`sys_user`.mobile LIKE ?", "%"+cast.ToString(val)+"%")
	}
	if val, ok := condition["name"]; ok {
		builder.Where("`sys_user`.name LIKE ?", "%"+cast.ToString(val)+"%")
	}
	builder.Where("sys_user_dept.dept_id IN ? ", scopeDept)
	builder.Group("`sys_user`.id")
	builder.Order("`sys_user`.`id`")

	newBuilder := db.Table("(?) as newTable", builder).Joins("LEFT JOIN sys_user_dept ON newTable.id = sys_user_dept.user_id AND newTable.tenant_id = sys_user_dept.tenant_id")
	switch scope {
	case 1: // 全部数据
	case 2: // 指定部门数据
	case 4: // 本部门及以下数据
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				if deptList, err := SysDeptRecursive(ctx, deptId); err == nil {
					newBuilder.Where("sys_user_dept.dept_id IN ? ", deptList)
				}
			}
		}
	case 3: // 本部门数据
		if val, ok := condition["deptId"]; ok {
			deptId := cast.ToInt64(val)
			if deptId > 0 {
				newBuilder.Where("sys_user_dept.dept_id= ?", deptId)
			}
		}
	case 5: // 仅本人数据
		newBuilder.Where("sys_user_dept.user_id= ?", userId)
	}
	newBuilder.Group("newTable.id")
	newBuilder.Count(&res)
	return
}

// SysDeptRecursive 递归查询
func SysDeptRecursive(ctx context.Context, id int64) (res []int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	query := "WITH RECURSIVE filter_sys_dept AS (SELECT id,name,parent_id,sort,user_id,phone,email,status,deleted,tenant_id,creator,create_time,updater,update_time FROM `sys_dept` WHERE `id`=? UNION ALL SELECT d.* FROM sys_dept AS d INNER JOIN filter_sys_dept f ON d.parent_id=f.id) SELECT filter_sys_dept.* FROM filter_sys_dept"
	var deptList []dao.SysDept
	db.WithContext(ctx).Raw(query).Scan(&deptList)
	for _, v := range deptList {
		if cast.ToInt32(v.Deleted) == 0 {
			res = append(res, cast.ToInt64(v.Id))
		}
	}
	return
}

// SysUserScope 用户数据权限
func SysUserScope(ctx context.Context, tenantId, userId int64) (res dao.SysUserScope, err error) {
	tenantItem, err := tenant.SysTenant(ctx, tenantId)
	if err != nil {
		return res, err
	}
	deptList, err := dept.SysDeptList(ctx, map[string]any{"tenantId": tenantId, "deleted": 0, "status": 0})
	if err != nil {
		return res, err
	}
	superDept := make([]int64, 0)
	for _, v := range deptList {
		superDept = append(superDept, *v.Id)
	}
	if len(superDept) < 1 {
		superDept = append(superDept, 0)
	}
	if tenantItem.UserId.IsValid() && cast.ToInt64(tenantItem.UserId.Ptr()) == userId {
		// 如果是租户管理员, 则数据权限为全部数据权限
		res.Scope = proto.Int32(1)
		res.ScopeDept = superDept
		return res, nil
	}
	db := initial.Core.Store.LoadSQL("mysql").Read()
	roleBuilder := db.WithContext(ctx).Model(&dao.SysRole{}).Select("sys_role.*").Joins("LEFT JOIN sys_user_role ON `sys_role`.`id` = `sys_user_role`.`role_id` AND `sys_role`.`tenant_id` = `sys_user_role`.`tenant_id`")
	roleBuilder.Where("`sys_user_role`.`tenant_id` =? ", tenantId)
	roleBuilder.Where("`sys_user_role`.`user_id` =? ", userId)
	roleBuilder.Where("`sys_role`.`tenant_id` =? ", tenantId)
	roleBuilder.Where("`sys_role`.`deleted` =? ", 0)
	roleBuilder.Where("`sys_role`.`status` =? ", 0)
	roleBuilder.Group("`sys_role`.`id`")
	roleBuilder.Order("`sys_role`.`scope`")
	var sysRoleList []dao.SysRole
	err = roleBuilder.Find(&sysRoleList).Error
	if err != nil {
		return res, err
	}
	// sysRoleList 是一个数组, 这里需要遍历数据, scope就取数据的第一个对象里面的 scope 值, scopeDept 需要获取整个数组的值
	var scope int32
	var scopeDept []int64
	if len(sysRoleList) > 0 {
		if sysRoleList[0].Scope.IsValid() {
			scope = cast.ToInt32(sysRoleList[0].Scope.Ptr())
		}
		for _, v := range sysRoleList {
			if v.ScopeDept.IsValid() {
				tmp := make([]int64, 0)
				if err := json.Unmarshal(*v.ScopeDept.Ptr(), &tmp); err == nil {
					scopeDept = append(scopeDept, tmp...)
				}
			}
		}
	}
	deptBuilder := db.WithContext(ctx).Model(&dao.SysDept{}).Select("sys_dept.*").Joins("LEFT JOIN `sys_user_dept` ON `sys_dept`.`id` = `sys_user_dept`.`dept_id` AND `sys_dept`.`tenant_id` = `sys_user_dept`.`tenant_id`")
	deptBuilder.Where("`sys_user_dept`.`tenant_id` =? ", tenantId)
	deptBuilder.Where("`sys_user_dept`.`user_id` =? ", userId)
	deptBuilder.Where("`sys_dept`.`tenant_id` =? ", tenantId)
	deptBuilder.Where("`sys_dept`.`deleted` =? ", 0)
	deptBuilder.Where("`sys_dept`.`status` =? ", 0)
	var sysDeptList []dao.SysDept
	err = deptBuilder.Find(&sysDeptList).Error
	if err != nil {
		return res, err
	}
	var deptAll []int64
	var deptTree []int64
	for _, v := range sysDeptList {
		deptAll = append(deptAll, *v.Id)
		deptTree = append(deptTree, *v.Id)
		if deptList, err := SysDeptRecursive(ctx, *v.Id); err == nil {
			for _, v := range deptList {
				if !util.InArray(v, deptTree) {
					deptTree = append(deptTree, v)
				}
			}
		}
	}
	// 查询本部门数据权限
	switch scope {
	case 1: // 全部数据权限
		res.ScopeDept = superDept
	case 2: // 指定部门数据权限
		res.ScopeDept = scopeDept
	case 3: // 本部门数据权限
		res.ScopeDept = deptAll
	case 4: // 本部门及以下数据权限
		res.ScopeDept = deptTree
	default: // 仅仅自己
		res.ScopeDept = deptAll
	}
	if len(res.ScopeDept) < 1 {
		res.ScopeDept = append(res.ScopeDept, 0)
	}
	res.Scope = proto.Int32(scope)
	return res, nil
}

func SysUserPassword(ctx context.Context, id int64, password string) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	data, err := SysUser(ctx, id)
	if err != nil {
		return 0, err
	}
	data.Password = proto.String(password)
	result := db.WithContext(ctx).Model(&dao.SysUser{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}
