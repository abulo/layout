<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="租户列表"
      row-key="id"
      :columns="columns"
      :toolbar-right="['search', 'refresh', 'export', 'layout']"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12"
    >
      <template #toolbarLeft>
        <el-button v-auth="'tenant.SysTenantCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">
          新增
        </el-button>
      </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="deleted" :value="scope.row.deleted" />
      </template>
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'tenant.SysTenant'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="[
              'tenant.SysTenantUpdate',
              'tenant.SysTenantDelete',
              'tenant.SysTenantRecover',
              'tenant.SysTenantDrop',
            ]"
            type="primary"
            link
            :icon="DArrowRight"
          >
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <div v-auth="'tenant.SysTenantUpdate'">
                <el-dropdown-item :icon="EditPen" @click="handleUpdate(scope.row)"> 编辑 </el-dropdown-item>
              </div>
              <div v-auth="'tenant.SysTenantDelete'">
                <el-dropdown-item v-if="scope.row.deleted === 0" :icon="Delete" @click="handleDelete(scope.row)">
                  删除
                </el-dropdown-item>
              </div>
              <div v-auth="'tenant.SysTenantRecover'">
                <el-dropdown-item v-if="scope.row.deleted === 1" :icon="Refresh" @click="handleRecover(scope.row)">
                  恢复
                </el-dropdown-item>
              </div>
              <div v-auth="'tenant.SysTenantDrop'">
                <el-dropdown-item v-if="scope.row.deleted === 1" :icon="DeleteFilled" @click="handleDrop(scope.row)">
                  清理
                </el-dropdown-item>
              </div>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </template>
    </ProTable>
    <el-dialog
      v-model="dialogVisible"
      :title="title"
      width="40%"
      destroy-on-close
      align-center
      center
      append-to-body
      draggable
      :lock-scroll="false"
      class="dialog-settings"
      @click="handleDialogClick"
    >
      <el-form ref="refSysTenantForm" :model="sysTenantForm" :rules="rulesSysTenantForm" label-width="100px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="sysTenantForm.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item v-if="sysTenantForm.userId === undefined" label="用户名称" prop="username">
          <el-input v-model="sysTenantForm.username" placeholder="请输入用户名称" />
        </el-form-item>
        <el-form-item v-if="sysTenantForm.userId === undefined" label="用户密码" prop="password">
          <el-input v-model="sysTenantForm.password" placeholder="请输入用户密码" show-password type="password" />
        </el-form-item>
        <el-form-item v-if="sysTenantForm.userId !== 0" label="负责人" prop="userId">
          <el-popover placement="bottom-start" :width="600" :show-arrow="false" trigger="click" :visible="isUserOpen">
            <template #reference>
              <el-button class="mr-4" @click.stop="userOpen">{{ userItem }}</el-button>
            </template>
            <div class="table-box">
              <ProTable
                ref="userProTable"
                title="用户列表"
                row-key="id"
                :columns="userColumns"
                :request-api="getCustomSysTenantUserListApi"
                :request-auto="false"
                :tool-button="false"
                :pagination-layout="'prev, pager, next'"
                :init-param="initUserParam"
                :pagination="true"
              >
                <template #operation="scope">
                  <el-button type="primary" link :icon="CirclePlus" @click.stop="handleUser(scope.row)">确定</el-button>
                </template>
              </ProTable>
            </div>
          </el-popover>
        </el-form-item>
        <el-form-item label="联系人" prop="contactName">
          <el-input v-model="sysTenantForm.contactName" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="联系电话" prop="contactMobile">
          <el-input v-model="sysTenantForm.contactMobile" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="过期时间" prop="expireDate">
          <el-date-picker
            v-model="sysTenantForm.expireDate"
            clearable
            placeholder="请选择过期时间"
            type="date"
            :disabled="disabled"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        <el-form-item label="账号数量" prop="accountTotal">
          <el-input v-model="sysTenantForm.accountTotal" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="套餐编号" prop="packageId">
          <el-select v-model="sysTenantForm.packageId" clearable placeholder="请选择租户套餐">
            <el-option
              v-for="item in sysTenantPackageOptions"
              :key="item.id"
              :label="item.name"
              :value="item.id"
              :disabled="disabled"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="sysTenantForm.status">
            <el-radio-button
              v-for="dict in statusEnum"
              :key="Number(dict.value)"
              :value="dict.value"
              :disabled="disabled"
            >
              {{ dict.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template v-if="!disabled" #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSysTenantForm)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSysTenantForm)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx">
defineOptions({ name: 'SysTenant' })
import type { ResSysTenant } from '@/api/interface/sysTenant'
import type { ProTableInstance, ColumnProps, SearchProps } from '@/components/ProTable/interface'
import { EditPen, CirclePlus, Delete, Refresh, DeleteFilled, View, DArrowRight } from '@element-plus/icons-vue'
import {
  getSysTenantListApi,
  deleteSysTenantApi,
  dropSysTenantApi,
  recoverSysTenantApi,
  getSysTenantApi,
  addSysTenantApi,
  updateSysTenantApi,
  getSysTenantUserListApi,
} from '@/api/modules/sysTenant'
import type { FormInstance, FormRules } from 'element-plus'
import { ElDatePicker } from 'element-plus'
import { getIntDictOptions } from '@/utils/dict'
import { DictTag } from '@/components/DictTag'
import { useHandleData, useHandleSet } from '@/hooks/useHandleData'
import { HasAuth } from '@/utils/auth'
import type { ResSysTenantPackage } from '@/api/interface/sysTenantPackage'
import { getSysTenantPackageListSimpleApi } from '@/api/modules/sysTenantPackage'
import { useLoadingStore } from '@/stores/modules/loading'
import { storeToRefs } from 'pinia'
import type { ResSysUser } from '@/api/interface/sysUser'
import { getSysUserApi } from '@/api/modules/sysUser'
// 获取loading状态
const { loading } = storeToRefs(useLoadingStore())
//禁用
const disabled = ref(true)
//弹出层标题
const title = ref('')
//列表数据
const proTable = ref<ProTableInstance>()
const userProTable = ref<ProTableInstance>()
//显示弹出层
const dialogVisible = ref(false)
// 定义负责人
const userItem = ref('点击选择')
// 显示负责人
const isUserOpen = ref(false)
// 默认参数
const initUserParam = reactive({ pageSize: 5 })
// 状态枚举
const statusEnum = getIntDictOptions('status')
// 服务包
const sysTenantPackageOptions = ref<ResSysTenantPackage[]>([])
//数据接口
const sysTenantForm = ref<ResSysTenant>({
  id: 0, // 编号
  name: '', // 名称
  userId: undefined, // 用户
  contactName: undefined, // 联系人
  contactMobile: undefined, // 联系电话
  expireDate: '', // 过期时间
  accountTotal: 0, // 账号数量
  packageId: 0, // 套餐编号
  status: 0, // 状态:0正常/1停用
  deleted: 0, // 删除:0否/1是
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
  username: '',
  password: '',
})
//表单
const refSysTenantForm = ref<FormInstance>()
//校验
const rulesSysTenantForm = reactive<FormRules>({
  name: [{ required: true, message: '名称不能为空', trigger: 'blur' }],
  expireDate: [{ required: true, message: '过期时间不能为空', trigger: 'blur' }],
  accountTotal: [{ required: true, message: '账号数量不能为空', trigger: 'blur' }],
  packageId: [{ required: true, message: '套餐编号不能为空', trigger: 'blur' }],
  status: [{ required: true, message: '状态', trigger: 'blur' }],
})

/**
 * 获取表格数据列表
 * @param params 查询参数对象
 * @returns 返回列表数据的Promise对象
 */
const getTableList = (params: any) => {
  // 深拷贝参数对象，避免修改原始参数
  let newParams = JSON.parse(JSON.stringify(params))
  newParams.expireDate && (newParams.beginExpireDate = newParams.expireDate[0])
  newParams.expireDate && (newParams.finishExpireDate = newParams.expireDate[1])
  delete newParams.expireDate
  return getSysTenantListApi(newParams)
}

/**
 * 重置数据
 * 该函数用于清空或初始化
 * @returns {void} 无返回值
 */
const resetSysTenant = () => {
  // 设置sysTenantForm为默认值
  Object.assign(sysTenantForm.value, {
    id: 0, // 编号
    name: '', // 名称
    userId: undefined, // 用户
    contactName: undefined, // 联系人
    contactMobile: undefined, // 联系电话
    expireDate: '', // 过期时间
    accountTotal: 0, // 账号数量
    packageId: 0, // 套餐编号
    status: 0, // 状态:0正常/1停用
    deleted: 0, // 删除:0否/1是
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined, // 更新时间
    username: '',
    password: '',
  })
}

/**
 * 重置函数
 * @returns {void}
 */
const reset = () => {
  resetSysTenant()
  disabled.value = true
  userItem.value = '点击选择'
  isUserOpen.value = false
}
/**
 * 处理新增操作
 * 该函数用于初始化对话框的状态
 */
const handleAdd = async () => {
  title.value = '新增租户'
  dialogVisible.value = true
  reset()
  getSysTenantPackageOptions()
  disabled.value = false
}
/**
 * 处理更新操作
 * @param row - 要更新数据对象
 */
const handleUpdate = async (row: ResSysTenant) => {
  title.value = '编辑租户'
  dialogVisible.value = true
  reset()
  getSysTenantPackageOptions()
  const { data } = await getSysTenantApi(Number(row.id))
  sysTenantForm.value = data
  if (Number(data.userId) !== 0) {
    const { data: user } = await getSysUserApi(Number(data.userId))
    userItem.value = user.name as string
  }
  disabled.value = false
}
/**
 * 处理查看操作
 * @param row - 要查看数据对象
 */
const handleItem = async (row: ResSysTenant) => {
  title.value = '查看租户'
  dialogVisible.value = true
  reset()
  getSysTenantPackageOptions()
  const { data } = await getSysTenantApi(Number(row.id))
  sysTenantForm.value = data
  if (Number(data.userId) !== 0) {
    const { data: user } = await getSysUserApi(Number(data.userId))
    userItem.value = user.name as string
  }
  disabled.value = true
}
/**
 * 处理清理操作
 * @param row - 要清理数据对象
 */
const handleDrop = async (row: ResSysTenant) => {
  await useHandleData(dropSysTenantApi, Number(row.id), '清理租户')
  proTable.value?.getTableList()
}
/**
 * 处理删除操作
 * @param row - 要删除数据对象
 */
const handleDelete = async (row: ResSysTenant) => {
  await useHandleData(deleteSysTenantApi, Number(row.id), '删除租户')
  proTable.value?.getTableList()
}
/**
 * 处理恢复操作
 * @param row - 要恢复数据对象
 */
const handleRecover = async (row: ResSysTenant) => {
  await useHandleData(recoverSysTenantApi, Number(row.id), '恢复租户')
  proTable.value?.getTableList()
}

// 获取服务包
const getSysTenantPackageOptions = async () => {
  const { data } = await getSysTenantPackageListSimpleApi()
  sysTenantPackageOptions.value = data
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
    const data = sysTenantForm.value as unknown as ResSysTenant
    if (data.id !== 0) {
      await useHandleSet(updateSysTenantApi, data.id, data, '修改租户')
    } else {
      await useHandleData(addSysTenantApi, data, '添加租户')
    }
    resetForm(formEl)
    proTable.value?.getTableList()
  })
}

const userOpen = () => {
  if (disabled.value) return
  isUserOpen.value = true
  userProTable.value?.getTableList()
}

// 在 el-dialog 上添加点击事件监听器
const handleDialogClick = () => {
  isUserOpen.value = false
}

// 当用户被选择
const handleUser = (row: ResSysUser) => {
  userItem.value = row.name as string
  sysTenantForm.value.userId = Number(row.id)
  isUserOpen.value = false
}

// 用户列表数据接口
const getCustomSysTenantUserListApi = (params: any) => {
  const newParams = Object.assign(params, JSON.parse(JSON.stringify(initUserParam)))
  return getSysTenantUserListApi(Number(sysTenantForm.value.id), newParams)
}

const userColumns: ColumnProps<ResSysUser>[] = [
  { prop: 'id', label: '编码', fixed: 'left' },
  { prop: 'name', label: '姓名', search: { el: 'input', span: 2 } },
  { prop: 'mobile', label: '手机号码', search: { el: 'input', span: 2 } },
  { prop: 'username', label: '用户名', search: { el: 'input', span: 2 } },
  {
    prop: 'operation',
    label: '操作',
    fixed: 'right',
  },
]

onMounted(() => {
  getSysTenantPackageOptions()
})

//删除状态
const deletedEnum = getIntDictOptions('deleted')
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasAuth('tenant.SysTenantDelete')
    ? {
        el: 'switch',
        span: 2,
        attrs: {
          activeValue: 1,
          inactiveValue: 0,
        },
      }
    : {}
)

const columns: ColumnProps<ResSysTenant>[] = [
  { prop: 'id', label: '编号' },
  { prop: 'name', label: '名称', search: { el: 'input', span: 2 } },
  { prop: 'contactName', label: '联系人' },
  { prop: 'contactMobile', label: '联系电话' },
  {
    prop: 'expireDate',
    label: '过期时间',
    search: {
      el: 'date-picker',
      span: 4,
      attrs: { type: 'daterange', valueFormat: 'YYYY-MM-DD' },
    },
  },
  { prop: 'accountTotal', label: '账号数量' },
  {
    prop: 'packageId',
    label: '套餐编号',
    enum: sysTenantPackageOptions,
    fieldNames: { label: 'name', value: 'id' },
    search: { el: 'select', span: 2 },
  },
  { prop: 'status', label: '状态', tag: true, enum: statusEnum, search: { el: 'select', span: 2 } },
  { prop: 'deleted', label: '删除', tag: true, enum: deletedEnum, search: deleteSearch },
  { prop: 'creator', label: '创建人' },
  { prop: 'createTime', label: '创建时间' },
  { prop: 'updater', label: '更新人' },
  { prop: 'updateTime', label: '更新时间' },

  {
    prop: 'operation',
    label: '操作',
    width: 150,
    fixed: 'right',
    isShow: HasAuth(
      'tenant.SysTenantUpdate',
      'tenant.SysTenantDelete',
      'tenant.SysTenantDrop',
      'tenant.SysTenantRecover',
      'tenant.SysTenant'
    ),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
@use '@/styles/popover';
</style>
