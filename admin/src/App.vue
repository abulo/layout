<template>
  <ElConfigProvider :locale="locale" :size="assemblySize" :button="buttonConfig">
    <RouterView />
  </ElConfigProvider>
</template>

<script setup lang="ts">
import { onMounted, reactive, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { getBrowserLang } from '@/utils'
import { useTheme } from '@/hooks/useTheme'
import { ElConfigProvider } from 'element-plus'
import { LanguageType } from './stores/interface'
import { useGlobalStore } from '@/stores/modules/global'
import en from 'element-plus/es/locale/lang/en'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import { VxeUI } from 'vxe-table'
import zhCN from 'vxe-table/lib/locale/lang/zh-CN' // 简体中文
import enUS from 'vxe-table/lib/locale/lang/en-US' // 英语(美国)
import plusZhCn from 'plus-pro-components/es/locale/lang/zh-cn'
import plusEn from 'plus-pro-components/es/locale/lang/en'
// 注册语言
VxeUI.setI18n('zh-CN', zhCN)
VxeUI.setI18n('en-US', enUS)

const zhCnLocales = {
  ...zhCn,
  ...plusZhCn,
}
const enLocales = {
  ...en,
  ...plusEn,
}

const globalStore = useGlobalStore()

// init theme
const { initTheme } = useTheme()
initTheme()

// init language
const i18n = useI18n()
onMounted(() => {
  const language = globalStore.language ?? (getBrowserLang() as LanguageType)
  i18n.locale.value = language || ''
  globalStore.language = language
  if (language == 'zh') {
    VxeUI.setLanguage('zh-CN')
  } else {
    VxeUI.setLanguage('en-US')
  }
})

// element language
const locale = computed(() => {
  if (globalStore.language == 'zh') {
    return zhCnLocales
  }
  if (globalStore.language == 'en') {
    return enLocales
  }
  return getBrowserLang() == 'zh' ? zhCnLocales : enLocales
})

// element assemblySize
const assemblySize = computed(() => globalStore.assemblySize)

// element button config
const buttonConfig = reactive({ autoInsertSpace: false })
</script>
