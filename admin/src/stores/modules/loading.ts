import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useLoadingStore = defineStore('geeker-loading', () => {
  const loading = ref(false)

  const setLoading = (_state: boolean) => {
    loading.value = _state
  }

  return {
    loading,
    setLoading,
  }
})
