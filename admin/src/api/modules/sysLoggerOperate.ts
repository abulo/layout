// sys_logger_operate 操作日志
import http from '@/utils/request'
import type { ResSysLoggerOperate, ReqSysLoggerOperateParams } from '@/api/interface/sysLoggerOperate'
// 操作日志删除数据
export const deleteSysLoggerOperateApi = (id: number) => {
  return http.delete(`/v1/sys/logger/operate/${id}/delete`)
}
// 操作日志查询单条数据
export const getSysLoggerOperateApi = (id: number) => {
  return http.get<ResSysLoggerOperate>(`/v1/sys/logger/operate/${id}/item`)
}
// 操作日志恢复数据
export const recoverSysLoggerOperateApi = (id: number) => {
  return http.put(`/v1/sys/logger/operate/${id}/recover`)
}
// 操作日志清理数据
export const dropSysLoggerOperateApi = (id: number) => {
  return http.delete(`/v1/sys/logger/operate/${id}/drop`)
}
// 操作日志列表数据
export const getSysLoggerOperateListApi = (params?: ReqSysLoggerOperateParams) => {
  return http.get<ResultPage<ResSysLoggerOperate>>(`/v1/sys/logger/operate`, params)
}
