<template>
  <ProVxeSearchForm
    v-show="isShowSearch"
    ref="searchFormRef"
    :search="_search"
    :reset="_reset"
    :columns="searchColumns"
    :show-number="showNumber"
    :default-values="defaultValues"
    :search-param="searchParam"
  />
  <!-- 表格主体 -->
  <div class="card table-main">
    <!-- 表格头部 操作按钮 -->
    <div
      v-if="toolbarLeftArr.length > 0 || toolbarMiddle || toolbarRightArr.length > 0"
      class="table-header mb-2 flex justify-between"
    >
      <div class="header-button-left">
        <template v-for="item in toolbarLeftArr" :key="item.auth">
          <el-button
            :icon="item.icon"
            :type="item.type"
            :auth="item.auth"
            :loading="loadingStore.loading"
            v-bind="item.attrs"
            @click="handleToolbarClick(item.name)"
          >
            {{ item.text }}
          </el-button>
        </template>
        <template v-if="toolbarLeftArr.length === 0">
          <slot name="toolbarLeft" />
        </template>
      </div>
      <div v-if="toolbarMiddle" class="header-button-middle">
        <component :is="toolbarMiddle" />
      </div>
      <div class="header-button-right">
        <template v-for="item in toolbarRightArr" :key="item.auth">
          <el-button
            :icon="item.icon"
            :auth="item.auth"
            :title="item.text"
            :loading="loadingStore.loading"
            circle
            @click="handleToolbarClick(item.name)"
          />
        </template>
      </div>
    </div>
    <!-- 表格主体 -->
    <div class="h-full">
      <vxe-table
        :id="pageId"
        ref="tableRef"
        v-bind="$attrs"
        :border="props.border"
        height="100%"
        :cell-config="{ height: size }"
        :header-cell-config="{ height: size }"
        :data="tableData"
      >
        <slot />
      </vxe-table>
    </div>
    <!-- 分页组件 -->
    <slot name="pagination">
      <ProVxePagination
        v-if="pagination !== ProTablePaginationEnum.NONE"
        :pageable="pageable"
        :layout="layout"
        :handle-size-change="handleSizeChange"
        :handle-current-change="handleCurrentChange"
      />
    </slot>
  </div>
</template>
<script setup lang="ts">
defineOptions({ name: 'ProVxeTable' })
import { VxeTableInstance } from 'vxe-table'
import { toolbarButtonsConfig } from '@/utils/proTable'
import { ProTablePaginationEnum } from '@/enums'
import { useLoadingStore } from '@/stores/modules/loading'
import { ProVxeTableProps, ProVxeColumnProps } from './interface'
// import { ProVxeTableProps, FieldValues, ProVxeColumnProps } from './interface'
// import { PlusColumn } from 'plus-pro-components'
import { useGlobalStore } from '@/stores/modules/global'
import ProVxeSearchForm from './components/ProVxeSearchForm.vue'
import ProVxePagination from './components/ProVxePagination.vue'
import { ElMessage } from 'element-plus'
import { useTable } from '@/hooks/useTable'
import { useI18n } from 'vue-i18n'
// import { cloneDeep } from 'lodash-es'

// 接受父组件参数，配置默认值
const props = withDefaults(defineProps<ProVxeTableProps>(), {
  columns: () => [],
  requestAuto: true,
  pagination: ProTablePaginationEnum.BE,
  initParam: () => ({}),
  border: true,
  layout: 'total, sizes, prev, pager, next, jumper',
  showNumber: 3,
})

const { t } = useI18n()

const emit = defineEmits<{
  search: []
  reset: []
  toolbarClick: [{ name: string; payload?: any }]
}>()

const _search = () => {
  search()
  emit('search')
}

const defaultValues = computed(() => {
  const res = {}
  searchColumns.value.map(column => {
    if (column?.defaultValue !== undefined && column?.defaultValue !== null) {
      // 添加到对象中
      res[column.prop] = column.defaultValue
    }
  })
  return res
})

const _reset = () => {
  for (const key in defaultValues.value) {
    if (Object.prototype.hasOwnProperty.call(defaultValues.value, key)) {
      const defaultValue = defaultValues.value[key]
      // 判断是否为 Promise 或返回 Promise 的函数
      if (defaultValue !== undefined && defaultValue !== null) {
        // 直接赋值非空的普通值
        searchParam.value[key] = defaultValue
        searchInitParam.value[key] = defaultValue
      }
    }
  }
  reset()
  emit('reset')
}

const globalStore = useGlobalStore()
// 设置表头和单元格的高度
const size = computed(() => {
  switch (globalStore.assemblySize) {
    case 'large':
      return 50
    case 'small':
      return 40
    default: // 包括 'default' 和其他未匹配的情况
      return 45
  }
})

const loadingStore = useLoadingStore()
const searchParamDefaultValuePromises: { key: string; promise: Promise<any> }[] = []

// 过滤需要搜索的配置项 && 排序
const searchColumns = computed(() => {
  return props.columns as ProVxeColumnProps[]
})

// 如果是前端分页，且有筛选参数，但是没有 fePaginationFilterMethod，则抛出错误
if (
  props.pagination === ProTablePaginationEnum.FE &&
  searchColumns.value.length !== 0 &&
  !props.fePaginationFilterMethod
) {
  ElMessage.error(t('error.fePaginationFilterMethodIsRequired'))
}

// 是否显示搜索模块
const isShowSearch = ref(true)

const pageId = computed(() => `id-${crypto.randomUUID()}`)
// table 实例
const tableRef = ref<VxeTableInstance>()

const toolbarLeftArr = computed(() => {
  if (!props.toolbarLeft) {
    return []
  }
  return props.toolbarLeft.map(item => {
    if (typeof item === 'string') {
      return toolbarButtonsConfig[item]
    } else {
      return item
    }
  })
})

const toolbarRightArr = computed(() => {
  if (!props.toolbarRight) {
    return []
  }
  return props.toolbarRight.map(item => {
    if (typeof item === 'string') {
      return toolbarButtonsConfig[item]
    } else {
      return item
    }
  })
})

// 表格操作 Hooks
const {
  tableData,
  pageable,
  searchParam,
  searchInitParam,
  getTableList,
  search,
  reset,
  handleSizeChange,
  handleCurrentChange,
} = useTable(props.requestApi, props.initParam, props.pagination, props.fePaginationFilterMethod, props.dataCallback)

// 处理 toolbar 点击事件
const handleToolbarClick = (name: string) => {
  const payload: { name: string; params?: any } = { name }
  switch (name) {
    case 'refresh':
      search()
      break
    case 'search':
      isShowSearch.value = !isShowSearch.value
      break
    default:
      payload.name = name
      break
  }
  emit('toolbarClick', payload)
}

// 遍历一下 initParam 的值 并且设置到 searchParam 中
for (const key in props.initParam) {
  if (Object.prototype.hasOwnProperty.call(props.initParam, key)) {
    const defaultValue = props.initParam[key]
    // 判断是否为 Promise 或返回 Promise 的函数
    if (
      defaultValue instanceof Promise ||
      (typeof defaultValue === 'function' && defaultValue.constructor.name === 'AsyncFunction')
    ) {
      searchParamDefaultValuePromises.push({ key, promise: defaultValue })
    } else if (defaultValue !== undefined && defaultValue !== null) {
      // 直接赋值非空的普通值
      searchParam.value[key] = defaultValue
      searchInitParam.value[key] = defaultValue
    }
  }
}

// 遍历一下 initParam 的值 并且设置到 searchParam 中
for (const key in defaultValues.value) {
  if (Object.prototype.hasOwnProperty.call(defaultValues.value, key)) {
    const defaultValue = defaultValues.value[key]
    // 判断是否为 Promise 或返回 Promise 的函数
    if (defaultValue !== undefined && defaultValue !== null) {
      // 直接赋值非空的普通值
      searchParam.value[key] = defaultValue
    }
  }
}

const setSearchParamForm = (key: string, value: any) => {
  searchParam.value[key] = value
}

// 初始化表格数据
onMounted(() => {
  Promise.all(searchParamDefaultValuePromises.map(item => item.promise))
    .then(res => {
      res.forEach((value, index) => {
        const { key } = searchParamDefaultValuePromises[index]
        searchParam.value[key] = value
        searchInitParam.value[key] = value
      })
    })
    .then(() => {
      props.requestAuto && getTableList()
    })
})

// 监听页面 initParam 改化，重新获取表格数据
watch(
  () => props.initParam,
  () => getTableList(false),
  { deep: true }
)

// 暴露给父组件的参数和方法
defineExpose({
  element: tableRef,
  tableData,
  pageable,
  searchParam,
  searchInitParam,
  setSearchParamForm,
  getTableList,
  search,
  reset,
  handleSizeChange,
  handleCurrentChange,
})
</script>
