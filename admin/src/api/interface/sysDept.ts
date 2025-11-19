// sys_dept 部门
export interface ReqSysDeptParams {
  deleted?: number // 删除:0否/1是
  status?: number // 状态:0正常/1停用
  name?: string // 名称
}

export interface ResSysDept {
  id: number // 编号
  name: string // 名称
  parentId: number // 上级
  sort: number // 排序
  userId: number | undefined // 负责人
  phone: string | undefined // 联系电话
  email: string | undefined // 邮件
  status: number // 状态:0正常/1停用
  deleted: number // 删除:0否/1是
  tenantId: number // 租户
  creator: string | undefined // 创建人
  createTime: string | undefined // 创建时间
  updater: string | undefined // 更新人
  updateTime: string | undefined // 更新时间
  children?: ResSysDept[]
}
