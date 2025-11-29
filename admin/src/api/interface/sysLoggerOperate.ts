// sys_logger_operate 操作日志
export interface ReqSysLoggerOperateParams extends RequestPage {
  tenantId?: number // 租户
  deleted?: number // 删除:0否/1是
  result?: number // 结果:0 成功/1 失败
  channel?: string // 渠道
  method?: string // 请求方法
  username?: string // 用户名
  deptId?: number // 部门
  beginStartTime?: string // 时间
  finishStartTime?: string // 时间
}

export interface ResSysLoggerOperate {
  id: number // 编号
  name: string | undefined // 姓名
  username: string // 用户名
  userId: number // 用户编号
  module: string | undefined // 模块名称
  method: string | undefined // 请求方法
  url: string | undefined // 请求地址
  ip: string | undefined // IP
  ua: string | undefined // UA
  goMethod: string | undefined // 方法名
  goMethodArgs: any | undefined // 方法参数
  startTime: string | undefined // 开始时间
  duration: number | undefined // 执行时长
  channel: string | undefined // 渠道
  result: number | undefined // 结果:0 成功/1 失败
  resultMsg: string | undefined // 结果信息
  deleted: number // 删除:0否/1是
  tenantId: number // 租户
  creator: string | undefined // 创建人
  createTime: string | undefined // 创建时间
  updater: string | undefined // 更新人
  updateTime: string | undefined // 更新时间
}
