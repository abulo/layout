// sys_logger_dev 开发日志
import http from '@/utils/request'
import type { ResSysLoggerDev, ReqSysLoggerDevParams } from '@/api/interface/sysLoggerDev'
// 开发日志删除数据
export const deleteSysLoggerDevApi = (id: number) => {
  return http.delete(`/sys/logger/dev/${id}/delete`)
}
// 开发日志查询单条数据
export const getSysLoggerDevApi = (id: number) => {
  return http.get<ResSysLoggerDev>(`/sys/logger/dev/${id}/item`)
}
// 开发日志列表数据
export const getSysLoggerDevListApi = (params?: ReqSysLoggerDevParams) => {
  return http.get<ResultPage<ResSysLoggerDev>>(`/sys/logger/dev`, params)
}
