// sys_tenant_package 租户套餐
export interface ReqSysTenantPackageParams extends RequestPage {
  status?: number // 状态:0正常/1停用
  name?: string // 名称
}

export interface ResSysTenantPackage {
  id: number // 编号
  name: string // 名称
  scopeMenu: any | undefined // 数据范围
  sort: number // 排序
  status: number // 状态:0正常/1停用
  remark: string // 备注
  creator: string | undefined // 创建人
  createTime: string | undefined // 创建时间
  updater: string | undefined // 更新人
  updateTime: string | undefined // 更新时间
}
