import type { ButtonProps } from 'element-plus'

/**
 * @description 表格工具栏按钮配置
 */
export const toolbarButtonsConfig: Record<
  string,
  {
    auth: string
    icon: string
    name: string
    text: string
    type: ButtonProps['type']
    attrs?: Omit<ButtonProps, 'icon' | 'text' | 'type'>
  }
> = {
  // 默认左边按钮
  add: {
    auth: 'add',
    icon: 'Plus',
    name: 'add',
    text: '新增',
    type: 'primary',
  },
  delete: {
    auth: 'delete',
    icon: 'Delete',
    name: 'delete',
    text: '删除',
    type: 'danger',
  },
  import: {
    auth: 'import',
    icon: 'Upload',
    name: 'import',
    text: '导入',
    type: 'primary',
  },
  // 默认右边按钮
  refresh: {
    auth: 'refresh',
    icon: 'Refresh',
    name: 'refresh',
    type: 'primary',
    text: '刷新',
  },
  layout: {
    auth: 'layout',
    icon: 'Setting',
    name: 'layout',
    type: 'primary',
    text: '列布局',
  },
  upload: {
    auth: 'upload',
    icon: 'Upload',
    name: 'upload',
    type: 'primary',
    text: '上传',
  },
  export: {
    auth: 'export',
    icon: 'Download',
    name: 'export',
    type: 'primary',
    text: '下载',
  },
  search: {
    auth: 'search',
    icon: 'Search',
    name: 'search',
    text: '搜索',
    type: 'primary',
  },
}
