<template>
  <div class="maximize" @click="exitMaximize">
    <tdesign-fullscreen-exit class="color-white absolute top-1/2 left-1/5 cursor-pointer" />
  </div>
</template>

<script setup lang="ts">
import { useGlobalStore } from '@/stores/modules/global'
import TdesignFullscreenExit from '~icons/tdesign/fullscreen-exit?width=20px&height=20px'

const globalStore = useGlobalStore()
const exitMaximize = () => {
  globalStore.maximize = false
}
</script>

<style scoped lang="scss">
.maximize {
  position: fixed;
  top: -25px;
  right: -25px;
  z-index: 999;
  width: 55px;
  height: 55px;
  cursor: pointer;
  background-color: var(--el-color-info);
  border-radius: 50%;
  opacity: 0.9;
  &:hover {
    background-color: var(--el-color-info-dark-2);
  }
}
</style>
