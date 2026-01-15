<template>
  <!-- 分页组件 -->
  <el-pagination
    :background="true"
    :current-page="pageable.pageNum"
    :page-size="pageable.pageSize"
    :page-sizes="pageable.pageSizes"
    :page-count="Math.ceil(pageable.total / pageable.pageSize)"
    :total="pageable.total"
    :size="globalStore?.assemblySize ?? 'default'"
    :layout="layout"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
  />
</template>

<script setup lang="ts">
defineOptions({ name: 'ProVxePagination' })
import { useGlobalStore } from '@/stores/modules/global'
const globalStore = useGlobalStore()

interface Pageable {
  pageNum: number
  pageSize: number
  total: number
  pageSizes: number[]
}

interface PaginationProps {
  pageable: Pageable
  layout: string
  handleSizeChange: (_size: number) => void
  handleCurrentChange: (_currentPage: number) => void
}

defineProps<PaginationProps>()
</script>
