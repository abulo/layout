<template>
  <el-form ref="loginFormRef" :model="loginForm" :rules="loginRules" size="large">
    <el-form-item prop="username">
      <el-input v-model="loginForm.username" placeholder="用户名">
        <template #prefix>
          <el-icon class="el-input__icon">
            <user />
          </el-icon>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item prop="password">
      <el-input
        v-model="loginForm.password"
        type="password"
        placeholder="密码"
        show-password
        autocomplete="new-password"
      >
        <template #prefix>
          <el-icon class="el-input__icon">
            <lock />
          </el-icon>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item prop="verifyCode">
      <el-input v-model="loginForm.verifyCode" placeholder="验证码">
        <template #prefix>
          <el-icon class="el-input__icon"><Key /></el-icon>
        </template>
        <template #append>
          <el-image class="captchaImg" :src="resCaptcha.verifyImage" fit="fill" @click="createCaptcha" />
        </template>
      </el-input>
    </el-form-item>
  </el-form>
  <div class="login-btn">
    <el-button :icon="CircleClose" round size="large" @click="resetForm(loginFormRef)"> 重置 </el-button>
    <el-button :icon="UserFilled" round size="large" type="primary" :loading="loading" @click="login(loginFormRef)">
      登录
    </el-button>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import { encryptPassword, getTimeState } from '@/utils'
import { ElNotification } from 'element-plus'
import { useUserStore } from '@/stores/modules/user'
import { useAuthStore } from '@/stores/modules/auth'
import { useTabsStore } from '@/stores/modules/tabs'
import { useKeepAliveStore } from '@/stores/modules/keepAlive'
import { initDynamicRouter } from '@/routers/modules/dynamicRouter'
import { CircleClose, UserFilled } from '@element-plus/icons-vue'
import type { ElForm } from 'element-plus'
import { useLoadingStore } from '@/stores/modules/loading'
import { storeToRefs } from 'pinia'
import type { ResCaptcha } from '@/api/interface/captcha'
import { getCaptchaApi } from '@/api/modules/captcha'
import { type ReqSysUserLogin } from '@/api/interface/sysUser'
import { postSysUserLogin } from '@/api/modules/sysUser'

//验证码
const resCaptcha = reactive<ResCaptcha>({
  captchaId: '', // 验证码id
  captchaImage: '', // 验证码图片
})

// todo caps lock
// todo forget password
// todo remember me
const router = useRouter()
const userStore = useUserStore()
const tabsStore = useTabsStore()
const authStore = useAuthStore()
const keepAliveStore = useKeepAliveStore()

type FormInstance = InstanceType<typeof ElForm>
const loginFormRef = ref<FormInstance>()
const loginRules = reactive({
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  verifyCode: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
})

const { loading } = storeToRefs(useLoadingStore())
const loginForm = reactive<ReqSysUserLogin>({
  username: '',
  password: '',
  verifyCode: '',
  verifyCodeId: '',
})

// 验证码生成
const createCaptcha = async () => {
  if (loading.value) return
  const { data } = await getCaptchaApi()
  resCaptcha.verifyCodeId = data.verifyCodeId
  resCaptcha.verifyImage = data.verifyImage
  loginForm.verifyCodeId = resCaptcha.verifyCodeId
}

// login
const login = (formEl: FormInstance | undefined) => {
  if (!formEl) {
    return
  }
  formEl.validate(async valid => {
    if (!valid) {
      return
    }
    try {
      // 1.执行登录接口
      const hashedPassword = encryptPassword(loginForm.password)
      const { data } = await postSysUserLogin({ ...loginForm, password: hashedPassword })
      userStore.setUser(data)
      // 2.添加动态路由
      await initDynamicRouter()

      // 3.清空 tabs、keepAlive 数据
      tabsStore.setTabs([])
      keepAliveStore.setKeepAliveName([])
      const menuItem = authStore.showHomeMenu
      // 4.跳转到首页
      router.push(menuItem.path)
      ElNotification({
        title: getTimeState(),
        message: '欢迎登录',
        type: 'success',
        duration: 3000,
      })
    } finally {
      //登录失败，重置验证码
      createCaptcha()
    }
  })
}

// resetForm
const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) {
    return
  }
  formEl.resetFields()
}

onMounted(() => {
  // 1.获取验证码
  createCaptcha()
  // 监听 enter 事件（调用登录）
  document.onkeydown = (e: KeyboardEvent) => {
    if (e.code === 'Enter' || e.code === 'enter' || e.code === 'NumpadEnter') {
      if (loading.value) {
        return
      }
      login(loginFormRef.value)
    }
  }
})

onBeforeUnmount(() => {
  document.onkeydown = null
})
</script>

<style scoped lang="scss">
@use '../index';
:deep(.el-input-group__append) {
  padding: 0;
}
</style>
