<template>
  <el-dropdown trigger="click">
    <div class="avatar">
      <span class="username">
        <AuthorImage :name="username" :size="23"></AuthorImage>
      </span>
    </div>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item @click="openDialog('passwordRef')">
          <el-icon><Edit /></el-icon>{{ $t('header.changePassword') }}
        </el-dropdown-item>
        <el-dropdown-item divided @click="logout">
          <el-icon><SwitchButton /></el-icon>{{ $t('header.logout') }}
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
  <!-- passwordDialog -->
  <password-dialog ref="passwordRef" />
</template>

<script setup lang="ts">
defineOptions({
  name: 'Avatar',
})
import { computed, ref } from 'vue'
import { LOGIN_URL } from '@/config'
import { useRouter } from 'vue-router'
// import { AuthApi } from '@/api/auth'
import { useUserStore } from '@/stores/modules/user'
import { ElMessageBox, ElMessage } from 'element-plus'
import PasswordDialog from './PasswordDialog.vue'

const router = useRouter()
const userStore = useUserStore()
const username = computed(() => userStore.userName)

// 退出登录
const logout = () => {
  ElMessageBox.confirm('您是否确认退出登录?', '温馨提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(async () => {
    // 2.清除 Token
    userStore.clearUserState()

    // 3.重定向到登陆页
    router.replace(LOGIN_URL)
    ElMessage.success('退出登录成功！')
  })
}

// 打开修改密码和个人信息弹窗
const passwordRef = ref<InstanceType<typeof PasswordDialog> | null>(null)
const openDialog = (ref: string) => {
  if (ref == 'passwordRef') {
    passwordRef.value?.openDialog()
  }
}
</script>

<style scoped lang="scss">
.avatar {
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  cursor: pointer;
  .username {
    margin-left: 20px;
    color: var(--el-header-text-color);
  }
}
</style>
