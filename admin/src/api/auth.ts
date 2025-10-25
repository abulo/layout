// import authMenuList from '@/assets/json/authMenuList.json'
// import authButtonList from '@/assets/json/authButtonList.json'
import http from '@/utils/request'
import type { MenuOptions } from '@/api/system/menu'

export interface ReqLoginForm {
  username: string
  password: string
}
export interface ResLogin {
  access_token: string
}
export interface ResAuthButtons {
  [key: string]: string[]
}

/**
 * @name 登录模块
 */
// 用户登录
export const AuthApi = {
  login: (params: ReqLoginForm) => {
    return http.post<ResLogin>(`/login`, params, { loading: false }) // 正常 post json 请求  ==>  application/json
    // return http.post<ResLogin>(`/login`, params, { loading: false }); // 控制当前请求不显示 loading
    // return http.post<ResLogin>(`/login`, {}, { params }); // post 请求携带 query 参数  ==>  ?username=admin&password=123456
    // return http.post<ResLogin>(`/login`, qs.stringify(params)); // post 请求携带表单参数  ==>  application/x-www-form-urlencoded
    // return http.get<ResLogin>(`/login?${qs.stringify(params, { arrayFormat: "repeat" })}`); // get 请求可以携带数组等复杂参数
  },

  // 获取菜单列表
  getAuthMenuList: () => {
    return http.get<MenuOptions[]>(`/user/menu`, {}, { loading: false })
    // 如果想让菜单变为本地数据，注释上一行代码，并引入本地 authMenuList.json 数据
    // return authMenuList
  },

  // 获取按钮权限
  getAuthButtonList: () => {
    return http.get<ResAuthButtons>(`/user/buttons`, {}, { loading: false })
    // 如果想让按钮权限变为本地数据，注释上一行代码，并引入本地 authButtonList.json 数据
    // return authButtonList
  },

  // 用户退出登录
  logout: () => {
    return http.post(`/logout`)
  },
}
