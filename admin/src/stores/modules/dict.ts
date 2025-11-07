import { defineStore } from 'pinia'
import type { Dict } from '@/api/interface/dict'
import { getSysDictDataApi } from '@/api/modules/sysDict'

export const useDictStore = defineStore('admin-dict', () => {
  const state = reactive<Record<string, Dict[]>>({})

  const actions = {
    async setDict() {
      const data = await getSysDictDataApi()
      state = data
    },
  }
  const getters = {
    getDict: (code: string) => {
      return state[code]
    },
    getAllDict: () => {
      return state
    },
    clearDict: () => {
      state = {}
    },
  }
  return { ...actions, ...getters }
})
