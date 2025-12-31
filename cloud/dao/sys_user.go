package dao

import (
	"github.com/abulo/ratel/v3/stores/null"
	"github.com/golang-jwt/jwt/v5"
)

// SysUser 用户 sys_user
type SysUser struct {
	Id         *int64        `gorm:"primaryKey;autoIncrement;column:id" json:"id"` //bigint 编号,PRI
	Name       null.String   `gorm:"column:name" json:"name"`                      //varchar 姓名
	Mobile     null.String   `gorm:"column:mobile" json:"mobile"`                  //varchar 手机号码
	Username   *string       `gorm:"column:username" json:"username"`              //varchar 用户名
	Password   *string       `gorm:"column:password" json:"password"`              //varchar 密码
	Status     *int32        `gorm:"column:status" json:"status"`                  //tinyint 状态:0正常/1停用
	Deleted    *int32        `gorm:"column:deleted" json:"deleted"`                //tinyint 删除:0否/1是
	TenantId   *int64        `gorm:"column:tenant_id" json:"tenantId"`             //bigint 租户
	Creator    null.String   `gorm:"column:creator" json:"creator"`                //varchar 创建人
	CreateTime null.TimeStamp `gorm:"column:create_time" json:"createTime"`         //datetime 创建时间
	Updater    null.String   `gorm:"column:updater" json:"updater"`                //varchar 更新人
	UpdateTime null.TimeStamp `gorm:"column:update_time" json:"updateTime"`         //datetime 更新时间
	DeptIds    null.JSON     `gorm:"column:dept_ids;<-:false" json:"deptIds"`      // 部门ID
	RoleIds    null.JSON     `gorm:"column:role_ids;<-:false" json:"roleIds"`      // 角色ID
	PostIds    null.JSON     `gorm:"column:post_ids;<-:false" json:"postIds"`      // 岗位ID
}

type SysUserScope struct {
	Scope     *int32  `gorm:"column:scope" json:"scope"`          //tinyint 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	ScopeDept []int64 `gorm:"column:scope_dept" json:"scopeDept"` //json 数据范围(指定部门数组)
}

type SysUserLogin struct {
	Username     string `json:"username,required"`     // 用户名
	Password     string `json:"password,required"`     // 密码
	VerifyCode   string `json:"verifyCode,required"`   // 验证码
	VerifyCodeId string `json:"verifyCodeId,required"` // 验证码ID
}

// UserVerify 用户令牌
type UserVerify struct {
	UserId   int64  `json:"userId"`   // 用户ID
	Name     string `json:"name"`     // 姓名
	UserName string `json:"userName"` // 用户名
	TenantId int64  `json:"tenantId"` // 租户ID
	jwt.RegisteredClaims
}

type UserToken struct {
	UserName     *string `json:"userName"`     // 用户名
	AccessToken  *string `json:"accessToken"`  // `token`
	RefreshToken *string `json:"refreshToken"` // `refreshToken`
	Expires      *string `json:"expires"`      // 过期时间
}

type RefreshUserToken struct {
	AccessToken  *string `json:"accessToken"`  // `token`
	RefreshToken *string `json:"refreshToken"` // `refreshToken`
	Expires      *string `json:"expires"`      // 过期时间
	UserName     *string `json:"userName"`     // 用户名
}

// 用户密码
type SysUserPassword struct {
	Password *string `json:"password,required"` // 密码
}

func (SysUser) TableName() string {
	return "sys_user"
}
