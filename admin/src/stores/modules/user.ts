import { defineStore } from 'pinia'
import { UserState } from '@/stores/interface'
import piniaPersistConfig from '@/stores/helper/persist'
import { type ResSysUserToken } from '@/api/interface/sysUser'

export const useUserStore = defineStore('admin-user', {
  state: (): UserState => ({
    accessToken: '',
    refreshToken: '',
    userName: '',
    expires: '',
  }),
  getters: {},
  actions: {
    setUser(data: ResSysUserToken) {
      this.setAccessToken(data.accessToken)
      this.setUserName(data.userName)
      this.setRefreshToken(data.refreshToken)
      this.setExpires(data.expires)
    },
    // Set Token
    setAccessToken(accessToken: string) {
      this.accessToken = accessToken
    },
    // Set setUserName
    setUserName(userName: string) {
      this.userName = userName
    },
    // Set RefreshToken
    setRefreshToken(refreshToken: string) {
      this.refreshToken = refreshToken
    },
    setExpires(expires: string) {
      this.expires = expires
    },
    clearUserState() {
      this.accessToken = ''
      this.refreshToken = ''
      this.userName = ''
      this.expires = ''
    },
  },
  persist: piniaPersistConfig('admin-user'),
})
