/**
 * v-auth
 * 按钮权限指令
 */
import { HasAuthItem } from '@/utils/auth'

const auth: Directive = {
  mounted(el: HTMLElement, binding: DirectiveBinding) {
    const { value } = binding
    if (value instanceof Array && value.length) {
      let HasAuth = false
      for (let index = 0; index < value.length; index++) {
        HasAuth = HasAuthItem(value[index])
        if (HasAuth) break
      }
      if (!HasAuth) el.remove()
    } else {
      if (!HasAuthItem(value)) el.remove()
    }
  },
}

export default auth
