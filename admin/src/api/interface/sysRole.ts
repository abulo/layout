// sys_role 角色
export interface ReqSysRoleParams extends RequestPage {
  code?: string // 编码
  deleted?: number // 删除:0否/1是
  status?: number // 状态:0正常/1停用
  name?: string // 名称
}

export interface ResSysRole {
  id: number // 编号
  name: string // 名称
  code: string // 编码
  scope: number | undefined // 数据范围:1:全部数据权限/2:自定数据权限/3:本部门数据权限/4:本部门及以下数据权限
  scopeDept: any | undefined // 数据范围(指定部门数组)
  sort: number // 排序
  status: number // 状态:0正常/1停用
  deleted: number // 删除:0否/1是
  tenantId: number // 租户
  creator: string | undefined // 创建人
  createTime: string | undefined // 创建时间
  updater: string | undefined // 更新人
  updateTime: string | undefined // 更新时间
  menuIds: any | undefined // 菜单编号
}
