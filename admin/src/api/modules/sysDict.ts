// sys_dict 字典
import http from '@/utils/request'
import type { ResSysDict, ReqSysDictParams, Dict } from '@/api/interface/sysDict'
// 字典创建数据
export const addSysDictApi = (params: ResSysDict) => {
  return http.post(`/v1/sys/dict`, params)
}
// 字典更新数据
export const updateSysDictApi = (id: number, params: ResSysDict) => {
  return http.put(`/v1/sys/dict/${id}/update`, params)
}
// 字典删除数据
export const deleteSysDictApi = (id: number) => {
  return http.delete(`/v1/sys/dict/${id}/delete`)
}
// 字典查询单条数据
export const getSysDictApi = (id: number) => {
  return http.get<ResSysDict>(`/v1/sys/dict/${id}/item`)
}
// 字典列表数据
export const getSysDictListApi = (params?: ReqSysDictParams) => {
  return http.get<ResultPage<ResSysDict>>(`/v1/sys/dict`, params)
}
// 字典列表(精简)数据
export const getSysDictListSimpleApi = (params?: ReqSysDictParams) => {
  return http.get<ResultPage<ResSysDict>>(`/v1/sys/dict/simple`, params)
}
// 字典列表
export const getSysDictDataApi = () => {
  return http.get<Dict>(`/v1/dict`)
}
