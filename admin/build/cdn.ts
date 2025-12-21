import { Plugin as importToCDN } from 'vite-plugin-cdn-import'

/**
 * @description 打包时采用`cdn`模式，仅限外网使用（默认不采用，如果需要采用cdn模式，请在 .env.production 文件，将 VITE_CDN 设置成true）
 * 平台采用国内cdn：https://www.bootcdn.cn，当然你也可以选择 https://unpkg.com 或者 https://www.jsdelivr.com
 * 注意：上面提到的仅限外网使用也不是完全肯定的，如果你们公司内网部署的有相关js、css文件，也可以将下面配置对应改一下，整一套内网版cdn
 */
export const cdn = importToCDN({
  //（prodUrl解释： name: 对应下面modules的name，version: 自动读取本地package.json中dependencies依赖中对应包的版本号，path: 对应下面modules的path，当然也可写完整路径，会替换prodUrl）
  // prodUrl: 'https://cdn.bootcdn.net/ajax/libs/{name}/{version}/{path}',
  prodUrl: 'https://cdn.jsdelivr.net/npm/{name}@{version}/{path}',
  // <script src="https://cdn.jsdelivr.net/npm/vue@3.5.16/dist/vue.global.prod.js"></script>
  modules: [
    {
      name: '@element-plus/icons-vue',
      var: 'ElementPlusIcons',
      path: 'dist/index.iife.min.js',
    },
    // {
    //   name: '@pureadmin/utils',
    //   var: '@pureadmin/utils',
    //   path: 'dist/iife.global.min.js',
    // },
    // {
    //   name: '@vueuse/core',
    //   var: '@vueuse/core',
    //   path: 'dist/index.iife.min.js',
    // },
    {
      name: '@wangeditor/editor',
      var: '@wangeditor/editor',
      path: 'dist/index.min.js',
      css: 'dist/css/style.min.css',
    },
    {
      name: '@wangeditor/editor-for-vue',
      var: '@wangeditor/editor-for-vue',
      path: 'dist/index.min.js',
    },
    {
      name: 'axios',
      var: 'axios',
      path: 'dist/axios.min.js',
    },
    {
      name: 'driver.js',
      var: 'driver.js',
      path: 'dist/driver.js.iife.min.js',
      css: 'dist/driver.min.css',
    },
    {
      name: 'echarts',
      var: 'echarts',
      path: 'dist/echarts.min.js',
    },
    {
      name: 'echarts-liquidfill',
      var: 'echarts-liquidfill',
      path: 'dist/echarts-liquidfill.min.js',
    },
    {
      name: 'element-plus',
      var: 'ElementPlus',
      path: 'dist/index.full.min.js',
      css: 'dist/index.min.css',
    },
    {
      name: 'lodash-es',
      var: 'lodash-es',
      path: 'lodash.min.js',
    },
    {
      name: 'md5',
      var: 'md5',
      path: 'dist/md5.min.js',
    },
    {
      name: 'nprogress',
      var: 'nprogress',
      path: 'nprogress.min.js',
      css: 'nprogress.min.css',
    },
    {
      name: 'pinia',
      var: 'Pinia',
      path: 'dist/pinia.iife.min.js',
    },
    {
      name: 'pinia-plugin-persistedstate',
      var: 'pinia-plugin-persistedstate',
      path: 'dist/index.iife.min.js',
    },
    {
      name: 'plus-pro-components',
      var: 'plus-pro-components',
      path: 'index.min.js',
      css: 'index.min.css',
    },
    {
      name: 'qs',
      var: 'qs',
      path: 'dist/qs.min.js',
    },
    {
      name: 'screenfull',
      var: 'screenfull',
      path: 'index.min.js',
    },
    {
      name: 'sortablejs',
      var: 'sortablejs',
      path: 'Sortable.min.js',
    },
    {
      name: 'tailwindcss',
      var: 'tailwindcss',
      path: 'dist/lib.min.js',
      css: 'index.min.css',
    },
    {
      name: 'vue',
      var: 'Vue',
      path: 'dist/vue.global.prod.min.js',
    },
    {
      name: 'vue-i18n',
      var: 'VueI18n',
      path: 'dist/vue-i18n.global.min.js',
    },
    {
      name: 'vue-router',
      var: 'VueRouter',
      path: 'dist/vue-router.global.min.js',
    },
    {
      name: 'vuedraggable',
      var: 'vuedraggable',
      path: 'dist/vuedraggable.umd.min.js',
    },
    {
      name: 'vxe-pc-ui',
      var: 'vxe-pc-ui',
      path: 'lib/index.umd.min.js',
      css: 'lib/style.min.css',
    },
    {
      name: 'vxe-table',
      var: 'vxe-table',
      path: 'lib/index.umd.min.js',
      css: 'lib/style.min.css',
    },
    {
      name: 'xe-utils',
      var: 'xe-utils',
      path: 'dist/xe-utils.umd.min.js',
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
