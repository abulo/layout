// sys_menu 菜单
export interface ReqSysMenuParams extends RequestPage {
  status?: number // 状态:0正常/1停用
  type?: number // 类型:0 目录/1 菜单/2 按钮
  status?: number // 状态:0正常/1停用
  type?: number // 类型:0 目录/1 菜单/2 按钮
}

export interface ResSysMenu {
  id: number // 编号
  name: string // 名称
  code: string | undefined // 编码
  type: number // 类型:0 目录/1 菜单/2 按钮
  sort: number // 排序
  parentId: number // 上级
  path: string | undefined // 地址
  icon: string | undefined // 图标
  component: string | undefined // 组件路径
  componentName: string | undefined // 组件名称
  hide: number // 隐藏:0 否/1 是
  link: string | undefined // 外部地址
  cache: number // 缓存:0否/1 是
  remark: string | undefined // 备注
  active: string | undefined // 激活地址
  full: number // 全屏:1 开/0 关
  redirect: string | undefined // 重定向
  status: number // 状态:0正常/1停用
  creator: string | undefined // 创建人
  createTime: string | undefined // 创建时间
  updater: string | undefined // 更新人
  updateTime: string | undefined // 更新时间
}
