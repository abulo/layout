// sys_user 用户
export interface ReqSysUserParams extends RequestPage {
  username?: string // 用户名
  deleted?: number // 删除:0否/1是
  status?: number // 状态:0正常/1停用
  mobile?: string // 手机号码
  name?: string // 姓名
  deptId?: number // 部门
}

export interface ReqSysUserLogin {
  username: string // 用户名
  password: string // 密码
  verifyCode: string // 验证码
  verifyCodeId: string // 验证码id
}

export interface ResSysUserToken {
  accessToken: string // 令牌
  refreshToken: string // 刷新令牌
  userName: string // 用户名
  expires: string // 过期时间
}

export interface ResSysUser {
  id: number // 编号
  name: string | undefined // 姓名
  mobile: string | undefined // 手机号码
  username: string // 用户名
  password: string // 密码
  status: number // 状态:0正常/1停用
  deleted: number // 删除:0否/1是
  tenantId: number // 租户
  creator: string | undefined // 创建人
  createTime: string | undefined // 创建时间
  updater: string | undefined // 更新人
  updateTime: string | undefined // 更新时间
}
