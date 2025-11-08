// sys_dict 字典
export interface ReqSysDictParams extends RequestPage {
  dictId?: number // 字典类型
  status?: number // 状态:0正常/1停用
}

export interface ResSysDict {
  id: number // 编号
  sort: number // 排序
  label: string // 标签
  value: string // 键值
  dictId: number // 字典类型
  colorType: string | undefined // 颜色类型
  cssClass: string | undefined // CSS样式
  status: number // 状态:0正常/1停用
  remark: string // 备注
  creator: string | undefined // 创建人
  createTime: string | undefined // 创建时间
  updater: string | undefined // 更新人
  updateTime: string | undefined // 更新时间
}

// 字典列表
export interface Dict {
  [key: string]: ResSysDict[]
}
