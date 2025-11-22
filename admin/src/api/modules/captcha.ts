import http from '@/utils/request'
import type { ResCaptcha } from '@/api/interface/captcha'

export const getCaptchaApi = () => {
  return http.get<ResCaptcha>(`/v1/verify/generate`, {}, { loading: false })
}
