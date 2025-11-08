import { defineStore } from 'pinia'
import type { UserState } from '@/stores/interface'
import piniaPersistConfig from '@/stores/helper/persist'

export const useUserStore = defineStore('admin-user', {
  state: (): UserState => ({
    token: '',
    refreshToken: '',
    userName: '',
  }),
  getters: {},
  actions: {
    // Set Token
    setToken(token: string) {
      this.token = token
    },
    // Set setUserName
    setUserName(userName: string) {
      this.userName = userName
    },
    // Set RefreshToken
    setRefreshToken(refreshToken: string) {
      this.refreshToken = refreshToken
    },
  },
  persist: piniaPersistConfig('admin-user'),
})
