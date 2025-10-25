export const SERVER_ENDPOINTS = {
  geeker: '/geeker',
  api: '/api',
}
// 字典缓存时间 30 天
export const DICT_CACHE_TIME = 1000 * 60 * 60 * 24 * 30

// 状态码提示信息
export const statusMessages: Record<number, string> = {
  400: '请求失败！请您稍后重试',
  401: '登录失效！请您重新登录',
  403: '当前账号无权限访问！',
  404: '你所访问的资源不存在！',
  405: '请求方式错误！请您稍后重试',
  408: '请求超时！请您稍后重试',
  500: '服务异常！',
  502: '网关错误！',
  503: '服务不可用！',
  504: '网关超时！',
}
