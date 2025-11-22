import { defineStore } from 'pinia'
import { type DictState } from '@/stores/interface'
import { getSysDictDataApi } from '@/api/modules/sysDict'
import piniaPersistConfig from '@/stores/helper/persist'

export const useDictStore = defineStore('admin-dict', {
  state: (): DictState => ({
    dict: {},
  }),
  actions: {
    async getDictList() {
      const { data } = await getSysDictDataApi()
      this.dict = data
    },
    clearDict() {
      this.dict = {}
    },
  },
  getters: {
    getDict: state => {
      return (code: string) => state.dict[code]
    },
    getAllDict: state => {
      return state.dict
    },
  },
  persist: piniaPersistConfig('admin-dict'),
})
