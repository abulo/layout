// sys_role 角色
import http from '@/utils/request'
import type { ResSysRole, ReqSysRoleParams } from '@/api/interface/sysRole'
// 角色创建数据
export const addSysRoleApi = (params: ResSysRole) => {
  return http.post(`/v1/sys/role`, params)
}
// 角色更新数据
export const updateSysRoleApi = (id: number, params: ResSysRole) => {
  return http.put(`/v1/sys/role/${id}/update`, params)
}
// 角色删除数据
export const deleteSysRoleApi = (id: number) => {
  return http.delete(`/v1/sys/role/${id}/delete`)
}
// 角色查询单条数据
export const getSysRoleApi = (id: number) => {
  return http.get<ResSysRole>(`/v1/sys/role/${id}/item`)
}
// 角色恢复数据
export const recoverSysRoleApi = (id: number) => {
  return http.put(`/v1/sys/role/${id}/recover`)
}
// 角色清理数据
export const dropSysRoleApi = (id: number) => {
  return http.delete(`/v1/sys/role/${id}/drop`)
}
// 角色列表数据
export const getSysRoleListApi = (params?: ReqSysRoleParams) => {
  return http.get<ResultPage<ResSysRole>>(`/v1/sys/role`, params)
}

// 角色列表(精简)数据
export const getSysRoleListSimpleApi = (params?: ReqSysRoleParams) => {
  return http.get<ResSysRole[]>(`/v1/sys/role/simple`, params)
}
