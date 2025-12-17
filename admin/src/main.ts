import { createApp } from 'vue'
import App from './App.vue'
// CSS common style sheet
import '@/styles/common.scss'
import '@/styles/tailwind.css'
// element css
import 'element-plus/dist/index.css'
// element dark css
import 'element-plus/theme-chalk/dark/css-vars.css'
// custom element dark css
import '@/styles/element-dark.scss'
// custom element css
import '@/styles/element.scss'
// element plus
import ElementPlus from 'element-plus'
// element icons
import * as Icons from '@element-plus/icons-vue'
// custom directives
import directives from '@/directives/index'
// vue Router
import router from '@/routers'
// vue i18n
import I18n from '@/locales/index'
// pinia store
import pinia from '@/stores'
// errorHandler
import errorHandler from '@/utils/errorHandler'

import VxeUIBase from 'vxe-pc-ui'
import 'vxe-pc-ui/es/style.css'

import VxeUITable from 'vxe-table'
import 'vxe-table/es/style.css'
import PlusProComponents from 'plus-pro-components'
import 'plus-pro-components/index.css'

const app = createApp(App)

app.config.errorHandler = errorHandler
app.config.warnHandler = () => null

// register the element Icons component
Object.keys(Icons).forEach(key => {
  app.component(key, Icons[key as keyof typeof Icons])
})

app
  .use(ElementPlus)
  .use(VxeUIBase)
  .use(VxeUITable)
  .use(PlusProComponents)
  .use(directives)
  .use(router)
  .use(I18n)
  .use(pinia)
  .mount('#app')
