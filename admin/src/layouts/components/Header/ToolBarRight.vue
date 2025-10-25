<template>
  <div class="tool-bar-ri">
    <div class="header-icon">
      <assembly-size id="assemblySize" />
      <language id="language" />
      <search-menu id="searchMenu" />
      <theme-setting id="themeSetting" />
      <message id="message" />
      <fullscreen id="fullscreen" />
    </div>
    <avatar />
  </div>
</template>

<script setup lang="ts">
import AssemblySize from './components/AssemblySize.vue'
import Language from './components/Language.vue'
import SearchMenu from './components/SearchMenu.vue'
import ThemeSetting from './components/ThemeSetting.vue'
import Message from './components/Message.vue'
import Fullscreen from './components/Fullscreen.vue'
import Avatar from './components/Avatar.vue'
</script>

<style scoped lang="scss">
.tool-bar-ri {
  display: flex;
  align-items: center;
  justify-content: center;
  padding-right: 20px;
  .header-icon {
    display: flex;
    gap: 0 20px;
    align-items: center;
    color: var(--el-header-text-color);
  }
}
</style>
