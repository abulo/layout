// sys_logger_login 登录日志
import http from '@/utils/request'
import { ResSysLoggerLogin, ReqSysLoggerLoginParams } from '@/api/interface/sysLoggerLogin'
// 登录日志删除数据
export const deleteSysLoggerLoginApi = (id: number) => {
  return http.delete(`/v1/sys/logger/login/${id}/delete`)
}
// 登录日志查询单条数据
export const getSysLoggerLoginApi = (id: number) => {
  return http.get<ResSysLoggerLogin>(`/v1/sys/logger/login/${id}/item`)
}
// 登录日志恢复数据
export const recoverSysLoggerLoginApi = (id: number) => {
  return http.put(`/v1/sys/logger/login/${id}/recover`)
}
// 登录日志清理数据
export const dropSysLoggerLoginApi = (id: number) => {
  return http.delete(`/v1/sys/logger/login/${id}/drop`)
}
// 登录日志列表数据
export const getSysLoggerLoginListApi = (params?: ReqSysLoggerLoginParams) => {
  return http.get<ResultPage<ResSysLoggerLogin>>(`/v1/sys/logger/login`, params)
}
