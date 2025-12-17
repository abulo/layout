<template>
  <div v-if="columns.length" class="card table-search pb-4!">
    <PlusSearch
      ref="formRef"
      :model-value="searchParam"
      :columns="columns"
      :default-values="defaultValues"
      :show-number="showNumber"
      @update:model-value="handleUpdateModelValue"
      @search="search"
      @reset="reset"
    />
  </div>
</template>
<script setup lang="ts">
defineOptions({ name: 'ProVxeSearchForm' })
import { ProSearchProps } from '../interface'

// 定义事件发射器
const emit = defineEmits<{
  (_event: 'update:searchParam', _value: Record<string, any>): void
}>()

// 默认值
withDefaults(defineProps<ProSearchProps>(), {
  columns: () => [],
  searchParam: () => ({}),
  showNumber: 3,
  defaultValues: () => ({}),
})

// 处理 model 值更新
const handleUpdateModelValue = (value: Record<string, any>) => {
  emit('update:searchParam', value)
}
</script>

<style scoped lang="scss">
:deep(.plus-form__row) {
  margin-right: unset !important;
  margin-left: unset !important;
  .el-col {
    padding-right: unset !important;
  }
}
:deep(.plus-search__button__wrapper) {
  margin-right: unset !important;
}
</style>
