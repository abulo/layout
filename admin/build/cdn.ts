import { Plugin as importToCDN } from 'vite-plugin-cdn-import'

/**
 * @description 打包时采用`cdn`模式，仅限外网使用（默认不采用，如果需要采用cdn模式，请在 .env.production 文件，将 VITE_CDN 设置成true）
 * 平台采用国内cdn：https://www.bootcdn.cn，当然你也可以选择 https://unpkg.com 或者 https://www.jsdelivr.com
 * 注意：上面提到的仅限外网使用也不是完全肯定的，如果你们公司内网部署的有相关js、css文件，也可以将下面配置对应改一下，整一套内网版cdn
 */
export const cdn = importToCDN({
  //（prodUrl解释： name: 对应下面modules的name，version: 自动读取本地package.json中dependencies依赖中对应包的版本号，path: 对应下面modules的path，当然也可写完整路径，会替换prodUrl）
  prodUrl: 'https://cdn.bootcdn.net/ajax/libs/{name}/{version}/{path}',
  modules: [
    {
      name: 'vue',
      var: 'Vue',
      path: 'vue.global.prod.min.js',
    },
    {
      name: 'vue-router',
      var: 'VueRouter',
      path: 'vue-router.global.min.js',
    },
    {
      name: 'vue-i18n',
      var: 'VueI18n',
      path: 'vue-i18n.runtime.global.prod.min.js',
    },
    {
      name: 'pinia',
      var: 'Pinia',
      path: 'pinia.iife.min.js',
    },
    {
      name: 'element-plus',
      var: 'ElementPlus',
      path: 'index.full.min.js',
      css: 'index.min.css',
    },
    {
      name: 'axios',
      var: 'axios',
      path: 'axios.min.js',
    },
    {
      name: 'echarts',
      var: 'echarts',
      path: 'echarts.min.js',
    },
  ],
})

// "@element-plus/icons-vue": "^2.3.2",
// "@icon-space/vue-next": "^1.0.7",
// "@tailwindcss/vite": "^4.1.13",
// "@vueuse/core": "^13.9.0",
// "@wangeditor/editor": "^5.1.23",
// "@wangeditor/editor-for-vue": "^5.1.12",
// "axios": "^1.12.2",
// "driver.js": "^1.3.6",
// "echarts": "^6.0.0",
// "echarts-liquidfill": "^3.1.0",
// "element-plus": "^2.11.3",
// "lodash-es": "^4.17.21",
// "nprogress": "^0.2.0",
// "pinia": "^3.0.3",
// "pinia-plugin-persistedstate": "^4.5.0",
// "qs": "^6.14.0",
// "screenfull": "^6.0.2",
// "sortablejs": "^1.15.6",
// "tailwindcss": "^4.1.13",
// "vue": "^3.5.22",
// "vue-i18n": "^11.1.12",
// "vue-router": "^4.5.1",
// "vuedraggable": "^4.1.0"
