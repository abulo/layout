import { defineStore } from 'pinia'
import type { Dict } from '@/api/interface/sysDict'
import { getSysDictDataApi } from '@/api/modules/sysDict'

export const useDictStore = defineStore('admin-dict', () => {
  const state = ref<Dict>({})

  const actions = {
    async setDict() {
      const data = await getSysDictDataApi()
      state.value = data
    },
  }
  const getters = {
    getDict: (code: string) => {
      return state.value[code]
    },
    getAllDict: () => {
      return state.value
    },
    clearDict: () => {
      state.value = {}
    },
  }
  return { ...actions, ...getters }
})
