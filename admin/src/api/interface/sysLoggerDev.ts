// sys_logger_dev 开发日志
export interface ReqSysLoggerDevParams extends RequestPage {
  beginTimestamp?: string // 时间
  finishTimestamp?: string // 时间
  host?: string // 服务名
  level?: string // 等级
}

export interface ResSysLoggerDev {
  id: number // 编号
  host: string | undefined // 服务名
  timestamp: string // 时间
  file: string | undefined //
  func: string | undefined // 方法名
  message: string | undefined // 消息
  level: string | undefined // 等级
  data: any | undefined // 数据
}
