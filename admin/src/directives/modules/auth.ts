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
        if (el.classList.contains('el-dropdown-menu__item')) {
          // 对于下拉菜单项，需要隐藏整个 li 元素
          el.parentNode && el.parentNode.removeChild(el)
        } else {
          el.remove()
        }
      }
    } else {
      if (!HasAuthItem(value)) {
        // 处理 el-dropdown-item 特殊情况
        if (el.classList.contains('el-dropdown-menu__item')) {
          // 对于下拉菜单项，需要隐藏整个 li 元素
          el.parentNode && el.parentNode.removeChild(el)
        } else {
          el.remove()
        }
      }
    }
  },
}

export default auth
