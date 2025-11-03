import http from '@/utils/request'
import authMenuList from '@/assets/json/authMenuList.json'

interface UserInfo {
  id: string
  username: string
  email: string
  phone: string
  avatar: string
}
interface ResAuthButtons {
  [key: string]: string[]
}
// interface MenuOptions {
//   path: string
//   name: string
//   component: string
//   meta: {
//     icon: string
//     title: string
//     isLink: string
//     isHide: boolean
//     isFull: boolean
//     isAffix: boolean
//     isKeepAlive: boolean
//   }
//   children?: MenuOptions[]
// }
export const AccountAPI = {
  getUserInfo: () => http.get<UserInfo>(`/user/info?userCode=${localStorage.getItem('userCode')}`),
  // getUserMenu: () => http.get<MenuOptions[]>(`/user/menu?userCode=${localStorage.getItem('userCode')}`),
  getUserMenu: () => {
    // console.log('authMenuList', authMenuList)
    return authMenuList
    // return http.get<MenuOptions[]>(`/user/menu?userCode=${localStorage.getItem('userCode')}`)
  },
  getUserButtons: () => http.get<ResAuthButtons>(`/user/buttons?userCode=${localStorage.getItem('userCode')}`),
}
