import http from '@/utils/request/sys'
import type { ResSysPost, ReqSysPostParams } from '@/api/interface/sysPost'

// 创建数据
export const addSysPostApi = (params: ResSysPost) => {
  return http.post(`/post`, params)
}
// 更新数据
export const updateSysPostApi = (id: number, params: ResSysPost) => {
  return http.put(`/post/${id}/update`, params)
}
// 删除数据
export const deleteSysPostApi = (id: number) => {
  return http.delete(`/post/${id}/delete`)
}
// 查询单条数据
export const getSysPostApi = (id: number) => {
  return http.get<ResSysPost>(`/post/${id}/item`)
}
// 恢复数据
export const recoverSysPostApi = (id: number) => {
  return http.put(`/post/${id}/recover`)
}
// 清理数据
export const dropSysPostApi = (id: number) => {
  return http.delete(`/post/${id}/drop`)
}
// 列表数据
export const getSysPostListApi = (params?: ReqSysPostParams) => {
  return http.get<ResultPage<ResSysPost>>(`/post`, params)
}
// 列表数据精简
export const getSysPostListSimpleApi = (params?: ReqSysPostParams) => {
  return http.get<ResultPage<ResSysPost>>(`/post/simple`, params)
}
