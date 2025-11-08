import { useAuthStore } from '@/stores/modules/auth'

export const HasAuth = (...permissionList: string[]) => {
  if (permissionList.length === 1) {
    let userHasPermission = true
    const permissionItem = permissionList[0]
    if (!HasAuthItem(permissionItem)) {
      userHasPermission = false
    }
    return userHasPermission
  } else {
    let userHasPermission = false
    for (let i = 0; i < permissionList.length; i++) {
      const permissionItem = permissionList[i]
      if (HasAuthItem(permissionItem)) {
        userHasPermission = true
        break
      }
    }
    return userHasPermission
  }
}

export const HasAuthItem = (permission: string) => {
  const authStore = useAuthStore()
  let itemName = ''
  if (permission.indexOf('.') > -1) {
    const permissionArr = permission.split('.')
    itemName = permissionArr[0]
  } else {
    itemName = permission
  }
  // 修复：先检查 authStore.authButtonListGet 是否存在
  const authButtonList = authStore.authButtonListGet
  if (!authButtonList) {
    return false
  }
  const authButtons = authButtonList[itemName] || []
  // console.log('authButtons.includes(permission)', authButtons.includes(permission))
  return authButtons.includes(permission)
}
