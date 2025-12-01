<template>
  <el-dialog v-model="dialogVisible" title="修改密码" width="500px" draggable>
    <el-form
      ref="refSysUserPasswordForm"
      :model="sysUserPasswordForm"
      :rules="rulesSysUserPasswordForm"
      label-width="100px"
    >
      <el-form-item label="密码" prop="password">
        <el-input v-model="sysUserPasswordForm.password" show-password type="password" />
      </el-form-item>
      <el-form-item label="确认密码" prop="confirmPassword">
        <el-input v-model="sysUserPasswordForm.confirmPassword" show-password type="password" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="resetForm(refSysUserPasswordForm)">取消</el-button>
        <el-button type="primary" :loading="loading" @click="submitForm(refSysUserPasswordForm)">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
defineOptions({
  name: 'PasswordDialog',
})
import { ref } from 'vue'
import { ReqSysUserPassword } from '@/api/interface/sysUser'
import { FormInstance, FormRules } from 'element-plus'
import { useHandleData } from '@/hooks/useHandleData'
import { encryptPassword } from '@/utils'
import { resetSysUserPasswordApi } from '@/api/modules/sysUser'
import { useLoadingStore } from '@/stores/modules/loading'
import { storeToRefs } from 'pinia'
// 获取loading状态
const { loading } = storeToRefs(useLoadingStore())

const dialogVisible = ref(false)

const sysUserPasswordForm = ref<ReqSysUserPassword>({
  password: '', // 密码
  confirmPassword: '',
})

//表单
const refSysUserPasswordForm = ref<FormInstance>()
//校验
const rulesSysUserPasswordForm = reactive<FormRules>({
  password: [{ required: true, message: '密码不能为空', trigger: 'blur' }],
  confirmPassword: [
    { required: true, message: '确认密码不能为空', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== sysUserPasswordForm.value.password) {
          callback(new Error('两次输入的密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur',
    },
  ],
})
const openDialog = () => {
  dialogVisible.value = true
}

/**
 * 重置表单并关闭对话框
 * @param formEl - 表单实例对象，用于调用重置方法
 */
const resetForm = (formEl: FormInstance | undefined) => {
  dialogVisible.value = false
  if (!formEl) return
  formEl.resetFields()
}

// 提交数据
const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate(async valid => {
    if (!valid) return
    const data = sysUserPasswordForm.value as unknown as ReqSysUserPassword
    data.password = encryptPassword(data.password)
    data.confirmPassword = encryptPassword(data.confirmPassword)
    await useHandleData(resetSysUserPasswordApi, data, '添加用户')
    resetForm(formEl)
  })
}

defineExpose({ openDialog })
</script>
