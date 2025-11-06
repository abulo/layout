// sys_post 职位
import http from '@/utils/request'
import type { ResSysPost, ReqSysPostParams } from '@/api/interface/sysPost'
// 职位创建数据
export const addSysPostApi = (params: ResSysPost) => {
  return http.post(`/v1/sys/post`, params)
}
// 职位更新数据
export const updateSysPostApi = (id: number, params: ResSysPost) => {
  return http.put(`/v1/sys/post/${id}/update`, params)
}
// 职位删除数据
export const deleteSysPostApi = (id: number) => {
  return http.delete(`/v1/sys/post/${id}/delete`)
}
// 职位查询单条数据
export const getSysPostApi = (id: number) => {
  return http.get<ResSysPost>(`/v1/sys/post/${id}/item`)
}
// 职位恢复数据
export const recoverSysPostApi = (id: number) => {
  return http.put(`/v1/sys/post/${id}/recover`)
}
// 职位清理数据
export const dropSysPostApi = (id: number) => {
  return http.delete(`/v1/sys/post/${id}/drop`)
}
// 职位列表数据
export const getSysPostListApi = (params?: ReqSysPostParams) => {
  return http.get<ResultPage<ResSysPost>>(`/v1/sys/post`, params)
}

// 职位列表(精简)数据
export const getSysPostListSimpleApi = (params?: ReqSysPostParams) => {
  return http.get<ResultPage<ResSysPost>>(`/v1/sys/post/simple`, params)
}
