import { defineStore } from 'pinia'
import type { AuthState } from '@/stores/interface'
import { getFlatMenuList, getShowMenuList, getAllBreadcrumbList } from '@/utils'
import { useRoute } from 'vue-router'
import { computed, reactive, toRefs } from 'vue'
import { getSysUserMenuApi, getSysUserBtnApi } from '@/api/modules/sysUser'

export const useAuthStore = defineStore('admin-auth', () => {
  const route = useRoute()
  const state = reactive<AuthState>({
    // 全部的按钮权限列表
    allAuthButtonList: {},
    // 菜单权限列表
    authMenuList: [],
    // 当前页面的 router name，用来做按钮权限筛选
    routeName: '',
  })

  const authButtonList = computed(() => {
    return state.allAuthButtonList[route.name as string] || []
  })

  const showMenuListGet = computed(() => getShowMenuList(state.authMenuList))
  const flatMenuListGet = computed(() => getFlatMenuList(state.authMenuList))
  const breadcrumbListGet = computed(() => getAllBreadcrumbList(state.authMenuList))
  const authMenuListGet = computed(() => state.authMenuList)
  const authButtonListGet = computed(() => state.allAuthButtonList)
  const showHomeMenu = computed(() => {
    const menuList = getFlatMenuList(state.authMenuList)
    return menuList.find(item => item.meta.isAffix && !item.meta.isHide && !item.meta.isFull)
  })

  const actions = {
    async getAuthButtonList() {
      const { data } = await getSysUserBtnApi()
      state.allAuthButtonList = data
    },
    async getAuthMenuList() {
      const { data } = await getSysUserMenuApi()
      this.authMenuList = data
    },
    async setRouteName(name: string) {
      state.routeName = name
    },
  }

  return {
    ...toRefs(state),
    authButtonList,
    authButtonListGet,
    showMenuListGet,
    flatMenuListGet,
    breadcrumbListGet,
    authMenuListGet,
    showHomeMenu,
    ...actions,
  }
})
