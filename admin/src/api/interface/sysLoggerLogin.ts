// sys_logger_login 登录日志
export interface ReqSysLoggerLoginParams extends RequestPage {
  deleted?: number // 删除:0否/1是
  channel?: string // 渠道
  beginLoginTime?: string // 登录时间
  finishLoginTime?: string
  username?: string // 用户名
  deptId?: number // 部门
}

export interface ResSysLoggerLogin {
  id: number // 编号
  name: string | undefined // 姓名
  username: string // 用户名
  userId: number // 用户编号
  ua: string | undefined // UA
  loginTime: string // 登录时间
  channel: string | undefined // 渠道
  ip: string | undefined // IP
  deleted: number // 删除:0否/1是
  tenantId: number // 租户
  creator: string | undefined // 创建人
  createTime: string | undefined // 创建时间
  updater: string | undefined // 更新人
  updateTime: string | undefined // 更新时间
}
