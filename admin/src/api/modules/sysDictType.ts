// sys_dict_type 字典类型
import http from '@/utils/request'
import type { ResSysDictType, ReqSysDictTypeParams } from '@/api/interface/sysDictType'
// 字典类型创建数据
export const addSysDictTypeApi = (params: ResSysDictType) => {
  return http.post(`/v1/sys/dict/type`, params)
}
// 字典类型更新数据
export const updateSysDictTypeApi = (id: number, params: ResSysDictType) => {
  return http.put(`/v1/sys/dict/type/${id}/update`, params)
}
// 字典类型删除数据
export const deleteSysDictTypeApi = (id: number) => {
  return http.delete(`/v1/sys/dict/type/${id}/delete`)
}
// 字典类型查询单条数据
export const getSysDictTypeApi = (id: number) => {
  return http.get<ResSysDictType>(`/v1/sys/dict/type/${id}/item`)
}
// 字典类型列表数据
export const getSysDictTypeListApi = (params?: ReqSysDictTypeParams) => {
  return http.get<ResultPage<ResSysDictType>>(`/v1/sys/dict/type`, params)
}
