// sys_tenant_package 租户套餐
import http from '@/utils/request'
import type { ResSysTenantPackage, ReqSysTenantPackageParams } from '@/api/interface/sysTenantPackage'
// 租户套餐创建数据
export const addSysTenantPackageApi = (params: ResSysTenantPackage) => {
  return http.post(`/v1/sys/tenant/package`, params)
}
// 租户套餐更新数据
export const updateSysTenantPackageApi = (id: number, params: ResSysTenantPackage) => {
  return http.put(`/v1/sys/tenant/package/${id}/update`, params)
}
// 租户套餐删除数据
export const deleteSysTenantPackageApi = (id: number) => {
  return http.delete(`/v1/sys/tenant/package/${id}/delete`)
}
// 租户套餐查询单条数据
export const getSysTenantPackageApi = (id: number) => {
  return http.get<ResSysTenantPackage>(`/v1/sys/tenant/package/${id}/item`)
}
// 租户套餐列表数据
export const getSysTenantPackageListApi = (params?: ReqSysTenantPackageParams) => {
  return http.get<ResultPage<ResSysTenantPackage>>(`/v1/sys/tenant/package`, params)
}

// 租户套餐列表(精简)数据
export const getSysTenantPackageListSimpleApi = (params?: ReqSysTenantPackageParams) => {
  return http.get<ResSysTenantPackage[]>(`/v1/sys/tenant/package/simple`, params)
}
