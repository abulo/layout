<template>
  <div class="fullscreen">
    <component
      :is="isFullscreen ? MaterialSymbolsFullscreenExit : MaterialSymbolsFullscreen"
      class="cursor-pointer"
      @click="handleFullScreen"
    />
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import screenfull from 'screenfull'
import MaterialSymbolsFullscreenExit from '~icons/material-symbols/fullscreen-exit?width=20px&height=20px'
import MaterialSymbolsFullscreen from '~icons/material-symbols/fullscreen?width=20px&height=20px'

const isFullscreen = ref(screenfull.isFullscreen)

onMounted(() => {
  screenfull.on('change', () => {
    if (screenfull.isFullscreen) {
      isFullscreen.value = true
    } else {
      isFullscreen.value = false
    }
  })
})

const handleFullScreen = () => {
  if (!screenfull.isEnabled) {
    ElMessage.warning('当前您的浏览器不支持全屏 ❌')
  }
  screenfull.toggle()
}
</script>
