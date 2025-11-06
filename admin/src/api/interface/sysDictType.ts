// sys_dict_type 字典类型
export interface ReqSysDictTypeParams extends RequestPage {
  type?: string // 字典类型
  status?: number // 状态:0正常/1停用
  name?: string // 字典名称
}

export interface ResSysDictType {
  id: number // 编号
  name: string // 字典名称
  type: string // 字典类型
  status: number // 状态:0正常/1停用
  remark: string // 备注
  creator: string | undefined // 创建人
  createTime: string | undefined // 创建时间
  updater: string | undefined // 更新人
  updateTime: string | undefined // 更新时间
}
