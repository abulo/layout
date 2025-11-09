/**
 * v-auth
 * 按钮权限指令
 */

import type { Directive, DirectiveBinding } from 'vue'
import { HasAuthItem } from '@/utils/auth'

const auth: Directive = {
  mounted(el: HTMLElement, binding: DirectiveBinding) {
    const { value } = binding
    let hasAuth = false
    if (value && value instanceof Array && value.length > 0) {
      for (let index = 0; index < value.length; index++) {
        hasAuth = HasAuthItem(value[index])
        if (hasAuth) break
      }
      if (!hasAuth) {
        el.remove()
      }
    } else {
      if (!HasAuthItem(value)) {
        el.remove()
      }
    }
  },
}

export default auth
