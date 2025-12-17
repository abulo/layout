import { ProTablePaginationEnum } from '@/enums'
import { PlusColumn } from 'plus-pro-components'
import { JSX } from 'vue/jsx-runtime'
import { ButtonProps } from 'element-plus'
import { ComponentPublicInstance, VNodeChild } from 'vue'
import ProVxeTable from './index.vue'

type PropertyKey = string | number | symbol
type DefaultRow = Record<PropertyKey, any>

export type ProVxeColumnProps = PlusColumn

export interface ProSearchProps {
  columns?: PlusColumn[] // 搜索配置列
  searchParam?: IObject // 搜索参数
  defaultValues?: IObject // 搜索项默认值
  showNumber?: number // 搜索项每列占比配置 ==> 非必传
  search: (_params: IObject) => void // 搜索方法
  reset: (_params: IObject) => void // 重置方法
}

// export interface ProVxeColumnProps<T extends DefaultRow = any> extends PlusColumn {}

export type ProVxeTableInstance = Omit<
  InstanceType<typeof ProVxeTable>,
  keyof ComponentPublicInstance | keyof ProVxeTableProps
>
export interface ProVxeTableProps<Query = any, Item extends DefaultRow = any, ExtraQuery = IObject> {
  loading?: boolean // 表格是否加载中 ==> 非必传（默认为false）
  toolbarLeft?: Array<
    | 'add'
    | 'delete'
    | 'import'
    | {
        auth: string
        icon?: string
        name: string
        text: string
        type: ButtonProps['type']
        attrs?: Partial<Omit<ButtonProps, 'icon' | 'text' | 'type'>>
      }
  >
  // 表格工具栏右侧图标
  toolbarRight?: Array<
    | 'refresh'
    | 'layout'
    | 'export'
    | 'download'
    | 'search'
    | {
        name: string
        icon: string
        text?: string
        auth?: string
        attrs?: Partial<Omit<ButtonProps, 'icon' | 'text' | 'circle'>>
      }
  >
  toolbarMiddle?: () => VNodeChild | Component | JSX.Element | JSX.Element[] // 表格工具栏中间内容
  columns: ProVxeColumnProps[]
  fePaginationFilterMethod?: (_query: IObject) => IObject[] // 前端分页过滤方法 ==> 非必传
  pagination?: ProTablePaginationEnum // 是否需要分页组件 ==> 非必传（默认为1）
  requestApi: (_params: Query & ExtraQuery) => Promise<ResultData<Item>> | Promise<Item[]> // 请求表格数据的api ==> 非必传
  requestAuto?: boolean // 是否自动执行请求 api ==> 非必传（默认为true）
  dataCallback?: (_data: Item[]) => IObject[] // 返回数据的回调函数，可以对数据进行处理 ==> 非必传
  title?: string // 表格标题 ==> 非必传
  initParam?: Partial<Query> // 初始化请求参数 ==> 非必传（默认为{}）
  layout?: string
  showNumber?: number // 搜索项每列占比配置 ==> 非必传
  border?: boolean
  defaultValues?: IObject // 搜索项初始化参数 ==> 非必传（默认为{}）
}
