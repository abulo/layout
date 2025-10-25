<template>
  <el-icon v-if="iconType == 'ep' && iconName" :key="customKey" :class="className" :size="size">
    <component :is="iconName" />
  </el-icon>
  <icon-space
    v-if="iconType == 'ip' && iconName"
    :key="customKey"
    :class="className"
    :size="size"
    :type="iconName"
    tag="i"
  />
</template>
<script lang="ts" setup>
defineOptions({ name: 'Icon' })
import { IconSpace } from '@icon-space/vue-next/es/all'
import { toPascalCase } from '@/utils'
import { computed } from 'vue'
interface IconProps {
  size?: number // 图标的大小 ==> 非必传
  icon: string // 图标的名称 ==> 必传
  iconType?: string // 图标类别 ==> 非必传（默认为 "ep"）
  iconName?: string // 图标的名称 ==> 非必传
  customKey?: string // 图标的 key ==> 非必传
  class?: string // 图标样式 ==> 非必传
}
// 接受父组件参数
const props = withDefaults(defineProps<IconProps>(), {
  size: 18,
})

const size = computed(() => props.size) // 图标大小
// const icon = computed(() => props.icon) // 图标名称

// 需要判断一下这里面有没有 ":" ,没有则默认 ep
const iconType = computed(() => {
  if (props.icon.includes(':')) {
    return props.icon.substring(0, props.icon.indexOf(':'))
  }
  return 'ep'
}) // 图标类别
// const iconType = computed(() => props.icon.substring(0, props.icon.indexOf(':'))) // 图标类别
const iconName = computed(() => toPascalCase(props.icon.substring(props.icon.indexOf(':') + 1))) // 图标名称
const customKey = computed(() => (props.customKey ? props.customKey : props.icon)) // 图标 key
const className = computed(() => props.class) // 图标样式
</script>

<style scoped lang="scss">
.i-icon {
  vertical-align: unset;
}
</style>
