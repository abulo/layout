// sys_tenant 租户
import http from '@/utils/request'
import type { ResSysTenant, ReqSysTenantParams } from '@/api/interface/sysTenant'
// 租户创建数据
export const addSysTenantApi = (params: ResSysTenant) => {
  return http.post(`/v1/sys/tenant`, params)
}
// 租户更新数据
export const updateSysTenantApi = (id: number, params: ResSysTenant) => {
  return http.put(`/v1/sys/tenant/${id}/update`, params)
}
// 租户删除数据
export const deleteSysTenantApi = (id: number) => {
  return http.delete(`/v1/sys/tenant/${id}/delete`)
}
// 租户查询单条数据
export const getSysTenantApi = (id: number) => {
  return http.get<ResSysTenant>(`/v1/sys/tenant/${id}/item`)
}
// 租户恢复数据
export const recoverSysTenantApi = (id: number) => {
  return http.put(`/v1/sys/tenant/${id}/recover`)
}
// 租户清理数据
export const dropSysTenantApi = (id: number) => {
  return http.delete(`/v1/sys/tenant/${id}/drop`)
}
// 租户列表数据
export const getSysTenantListApi = (params?: ReqSysTenantParams) => {
  return http.get<ResultPage<ResSysTenant>>(`/v1/sys/tenant`, params)
}

// 租户列表(精简)数据
export const getSysTenantListSimpleApi = (params?: ReqSysTenantParams) => {
  return http.get<ResultPage<ResSysTenant>>(`/v1/sys/tenant`, params)
}
