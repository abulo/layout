// sys_dept 部门
import http from '@/utils/request'
import { ResSysDept, ReqSysDeptParams } from '@/api/interface/sysDept'
// 部门创建数据
export const addSysDeptApi = (params: ResSysDept) => {
  return http.post(`/v1/sys/dept`, params)
}
// 部门更新数据
export const updateSysDeptApi = (id: number, params: ResSysDept) => {
  return http.put(`/v1/sys/dept/${id}/update`, params)
}
// 部门删除数据
export const deleteSysDeptApi = (id: number) => {
  return http.delete(`/v1/sys/dept/${id}/delete`)
}
// 部门查询单条数据
export const getSysDeptApi = (id: number) => {
  return http.get<ResSysDept>(`/v1/sys/dept/${id}/item`)
}
// 部门恢复数据
export const recoverSysDeptApi = (id: number) => {
  return http.put(`/v1/sys/dept/${id}/recover`)
}
// 部门清理数据
export const dropSysDeptApi = (id: number) => {
  return http.delete(`/v1/sys/dept/${id}/drop`)
}
// 部门列表数据
export const getSysDeptListApi = (params?: ReqSysDeptParams) => {
  return http.get<ResSysDept[]>(`/v1/sys/dept`, params)
}

// 部门列表(精简)数据
export const getSysDeptListSimpleApi = (params?: ReqSysDeptParams) => {
  return http.get<ResSysDept[]>(`/v1/sys/dept/simple`, params)
}
