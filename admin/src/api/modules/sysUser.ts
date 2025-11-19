// sys_user 用户
import http from '@/utils/request'
import type { ResSysUser, ReqSysUserParams } from '@/api/interface/sysUser'
// 用户创建数据
export const addSysUserApi = (params: ResSysUser) => {
  return http.post(`/v1/sys/user`, params)
}
// 用户更新数据
export const updateSysUserApi = (id: number, params: ResSysUser) => {
  return http.put(`/v1/sys/user/${id}/update`, params)
}
// 用户删除数据
export const deleteSysUserApi = (id: number) => {
  return http.delete(`/v1/sys/user/${id}/delete`)
}
// 用户查询单条数据
export const getSysUserApi = (id: number) => {
  return http.get<ResSysUser>(`/v1/sys/user/${id}/item`)
}
// 用户恢复数据
export const recoverSysUserApi = (id: number) => {
  return http.put(`/v1/sys/user/${id}/recover`)
}
// 用户清理数据
export const dropSysUserApi = (id: number) => {
  return http.delete(`/v1/sys/user/${id}/drop`)
}
// 用户列表数据
export const getSysUserListApi = (params?: ReqSysUserParams) => {
  return http.get<ResultPage<ResSysUser>>(`/v1/sys/user`, params)
}

// 用户列表(精简)数据
export const getSysUserListSimpleApi = (params?: ReqSysUserParams) => {
  return http.get<ResultPage<ResSysUser>>(`/v1/sys/user/simple`, params)
}
