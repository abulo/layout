// sys_menu 菜单
import http from '@/utils/request'
import type { ResSysMenu, ReqSysMenuParams } from '@/api/interface/sysMenu'
// 菜单创建数据
export const addSysMenuApi = (params: ResSysMenu) => {
  return http.post(`/v1/sys/menu`, params)
}
// 菜单更新数据
export const updateSysMenuApi = (id: number, params: ResSysMenu) => {
  return http.put(`/v1/sys/menu/${id}/update`, params)
}
// 菜单删除数据
export const deleteSysMenuApi = (id: number) => {
  return http.delete(`/v1/sys/menu/${id}/delete`)
}
// 菜单查询单条数据
export const getSysMenuApi = (id: number) => {
  return http.get<ResSysMenu>(`/v1/sys/menu/${id}/item`)
}
// 菜单列表数据
export const getSysMenuListApi = (params?: ReqSysMenuParams) => {
  return http.get<ResSysMenu[]>(`/v1/sys/menu`, params)
}
// 菜单列表(精简)数据
export const getSysMenuListSimpleApi = (params?: ReqSysMenuParams) => {
  return http.get<ResSysMenu[]>(`/v1/sys/menu/simple`, params)
}
