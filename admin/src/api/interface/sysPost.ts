// sys_post 职位
export interface ReqSysPostParams extends RequestPage {
  deleted?: number // 删除:0否/1是
  status?: number // 状态:0正常/1停用
  name?: string // 名称
}

export interface ResSysPost {
  id: number // 编号
  name: string // 名称
  sort: number // 排序
  status: number // 状态:0正常/1停用
  deleted: number // 删除:0否/1是
  tenantId: number // 租户
  creator: string | undefined // 创建人
  createTime: string | undefined // 创建时间
  updater: string | undefined // 更新人
  updateTime: string | undefined // 更新时间
}
