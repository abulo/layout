// sys_tenant 租户
export interface ReqSysTenantParams extends RequestPage {
  deleted?: number // 删除:0否/1是
  status?: number // 状态:0正常/1停用
  name?: string // 名称
  beginExpireDate?: string // 过期时间
  finishExpireDate?: string // 过期时间
  packageId?: number // 套餐编号
}

export interface ResSysTenant {
  id: number // 编号
  name: string // 名称
  userId: number | undefined // 用户
  contactName: string | undefined // 联系人
  contactMobile: string | undefined // 联系电话
  expireDate: string // 过期时间
  accountTotal: number // 账号数量
  packageId: number // 套餐编号
  status: number // 状态:0正常/1停用
  deleted: number // 删除:0否/1是
  creator: string | undefined // 创建人
  createTime: string | undefined // 创建时间
  updater: string | undefined // 更新人
  updateTime: string | undefined // 更新时间
  username?: string
  password?: string
}
