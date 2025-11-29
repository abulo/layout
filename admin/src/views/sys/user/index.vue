<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="用户列表"
      row-key="id"
      :columns="columns"
      :toolbar-right="['search', 'refresh', 'export', 'layout']"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="ProTablePaginationEnum.BE"
      :search-col="12"
    >
      <template #toolbarLeft>
        <el-button v-auth="'user.SysUserCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
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
        <el-button v-auth="'user.SysUser'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-button v-auth="'user.SysUserLogin'" type="primary" link :icon="Connection" @click="handleLogin(scope.row)">
          登录
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['user.SysUserUpdate', 'user.SysUserDelete', 'user.SysUserRecover', 'user.SysUserDrop']"
            type="primary"
            link
            :icon="DArrowRight"
          >
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <div v-auth="'user.SysUserUpdate'">
                <el-dropdown-item :icon="EditPen" @click="handleUpdate(scope.row)"> 编辑 </el-dropdown-item>
              </div>
              <div v-auth="'user.SysUserDelete'">
                <el-dropdown-item v-if="scope.row.deleted === 0" :icon="Delete" @click="handleDelete(scope.row)">
                  删除
                </el-dropdown-item>
              </div>
              <div v-auth="'user.SysUserRecover'">
                <el-dropdown-item v-if="scope.row.deleted === 1" :icon="Refresh" @click="handleRecover(scope.row)">
                  恢复
                </el-dropdown-item>
              </div>
              <div v-auth="'user.SysUserDrop'">
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
      width="60%"
      destroy-on-close
      align-center
      center
      append-to-body
      draggable
      :lock-scroll="false"
      class="dialog-settings"
    >
      <el-form ref="refSysUserForm" :model="sysUserForm" :rules="rulesSysUserForm" label-width="100px">
        <el-form-item label="姓名" prop="name">
          <el-input v-model="sysUserForm.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="手机号码" prop="mobile">
          <el-input v-model="sysUserForm.mobile" :disabled="disabled" />
        </el-form-item>
        <el-form-item v-if="sysUserForm.id === 0" label="用户名" prop="username">
          <el-input v-model="sysUserForm.username" :disabled="disabled" />
        </el-form-item>
        <el-form-item v-if="sysUserForm.id === 0" label="密码" prop="password">
          <el-input v-model="sysUserForm.password" show-password type="password" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="角色" prop="roleIds">
          <el-select v-model="sysUserForm.roleIds" multiple placeholder="请选择角色" :disabled="disabled">
            <el-option
              v-for="item in roleOptions"
              :key="item.id"
              :label="item.name"
              :value="item.id"
              :disabled="disabled"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="职位" prop="postIds">
          <el-select v-model="sysUserForm.postIds" multiple placeholder="请选择职位" :disabled="disabled">
            <el-option
              v-for="item in postOptions"
              :key="item.id"
              :label="item.name"
              :value="item.id"
              :disabled="disabled"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="部门" prop="deptIds">
          <el-card class="cardHeight">
            <template #header>
              全选/不选:
              <el-switch
                v-model="deptSelect"
                active-text="是"
                inactive-text="否"
                inline-prompt
                @change="handleDeptSelect"
              />
              展开/折叠:
              <el-switch
                v-model="deptExpand"
                active-text="展开"
                inactive-text="折叠"
                inline-prompt
                @change="handleDeptExpand"
              />
            </template>
            <el-tree
              ref="deptRef"
              :data="deptOptions"
              :props="{
                children: 'children',
                label: 'name',
              }"
              :list="sysUserForm.deptIds"
              empty-text="加载中，请稍候"
              node-key="id"
              show-checkbox
            />
          </el-card>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="sysUserForm.status">
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
          <el-button @click="resetForm(refSysUserForm)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSysUserForm)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx">
defineOptions({ name: 'SysUser' })
import { ResSysUser } from '@/api/interface/sysUser'
import { ProTableInstance, ColumnProps, SearchProps } from '@/components/ProTable/interface'
import {
  EditPen,
  CirclePlus,
  Delete,
  Refresh,
  DeleteFilled,
  View,
  DArrowRight,
  Connection,
} from '@element-plus/icons-vue'
import {
  getSysUserListApi,
  deleteSysUserApi,
  dropSysUserApi,
  recoverSysUserApi,
  getSysUserApi,
  addSysUserApi,
  updateSysUserApi,
  postSysUserLogin,
} from '@/api/modules/sysUser'
import { FormInstance, FormRules, ElTree, ElNotification, ElMessage } from 'element-plus'
import { getIntDictOptions } from '@/utils/dict'
import { DictTag } from '@/components/DictTag'
import { useHandleData, useHandleSet } from '@/hooks/useHandleData'
import { HasAuth } from '@/utils/auth'
import { useLoadingStore } from '@/stores/modules/loading'
import { storeToRefs } from 'pinia'
import { ProTablePaginationEnum } from '@/enums'
import { handleTree } from '@pureadmin/utils'
import { getSysRoleListSimpleApi } from '@/api/modules/sysRole'
import { getSysDeptListSimpleApi } from '@/api/modules/sysDept'
import { getSysPostListSimpleApi } from '@/api/modules/sysPost'
import { ResSysDept } from '@/api/interface/sysDept'
import { ResSysPost } from '@/api/interface/sysPost'
import { ResSysRole } from '@/api/interface/sysRole'
import { useTimeoutFn } from '@vueuse/core'
import Node from 'element-plus/es/components/tree/src/model/node'
import { encryptPassword, getTimeState, parseRedirect } from '@/utils'
import { useUserStore } from '@/stores/modules/user'
import { useTabsStore } from '@/stores/modules/tabs'
import { useKeepAliveStore } from '@/stores/modules/keepAlive'
import { initDynamicRouter } from '@/routers/modules/dynamicRouter'
import { useRoute, useRouter } from 'vue-router'
const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const tabsStore = useTabsStore()
const keepAliveStore = useKeepAliveStore()
// 获取loading状态
const { loading } = storeToRefs(useLoadingStore())
//禁用
const disabled = ref(true)
//弹出层标题
const title = ref('')
//列表数据
const proTable = ref<ProTableInstance>()
//显示弹出层
const dialogVisible = ref(false)
// 状态枚举
const statusEnum = getIntDictOptions('status')
// 定义部门
const deptOptions = ref<ResSysDept[]>([])
// 定义岗位
const postOptions = ref<ResSysPost[]>([])
// 定义角色
const roleOptions = ref<ResSysRole[]>([])

// 获取选择
const deptSelect = ref(false)
// 菜单ref
const deptRef = ref<InstanceType<typeof ElTree>>()
// 菜单展开/折叠
const deptExpand = ref(false)

//数据接口
const sysUserForm = ref<ResSysUser>({
  id: 0, // 编号
  name: undefined, // 姓名
  mobile: undefined, // 手机号码
  username: '', // 用户名
  password: '', // 密码
  status: 0, // 状态:0正常/1停用
  deleted: 0, // 删除:0否/1是
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
  roleIds: undefined, // 角色编号
  deptIds: undefined, // 部门编号
  postIds: undefined, // 岗位编号
})
//表单
const refSysUserForm = ref<FormInstance>()
//校验
const rulesSysUserForm = reactive<FormRules>({
  username: [{ required: true, message: '用户名不能为空', trigger: 'blur' }],
  password: [{ required: true, message: '密码不能为空', trigger: 'blur' }],
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
  return getSysUserListApi(newParams)
}

/**
 * 重置数据
 * 该函数用于清空或初始化
 * @returns {void} 无返回值
 */
const resetSysUser = () => {
  // 设置sysUserForm为默认值
  Object.assign(sysUserForm.value, {
    id: 0, // 编号
    name: undefined, // 姓名
    mobile: undefined, // 手机号码
    username: '', // 用户名
    password: '', // 密码
    status: 0, // 状态:0正常/1停用
    deleted: 0, // 删除:0否/1是
    tenantId: 0, // 租户
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined, // 更新时间
    roleIds: undefined, // 角色编号
    deptIds: undefined, // 部门编号
    postIds: undefined, // 岗位编号
  })
}

/**
 * 重置函数
 * @returns {void}
 */
const reset = () => {
  // loading.value = false
  resetSysUser()
  disabled.value = true
  deptSelect.value = false
  deptExpand.value = false
  deptRef.value?.setCheckedNodes([])
}
/**
 * 处理新增操作
 * 该函数用于初始化对话框的状态
 */
const handleAdd = () => {
  title.value = '新增用户'
  dialogVisible.value = true
  reset()
  getOptions()
  disabled.value = false
}
/**
 * 处理更新操作
 * @param row - 要更新数据对象
 */
const handleUpdate = async (row: ResSysUser) => {
  title.value = '编辑用户'
  dialogVisible.value = true
  reset()
  getOptions()
  const { data } = await getSysUserApi(Number(row.id))
  sysUserForm.value = data
  disabled.value = false
  useTimeoutFn(() => {
    data.deptIds?.forEach((deptId: number) => {
      deptRef.value?.setChecked(deptId, true, false)
    })
  }, 200)
}
/**
 * 处理查看操作
 * @param row - 要查看数据对象
 */
const handleItem = async (row: ResSysUser) => {
  title.value = '查看用户'
  dialogVisible.value = true
  reset()
  getOptions()
  const { data } = await getSysUserApi(Number(row.id))
  sysUserForm.value = data
  disabled.value = true
  useTimeoutFn(() => {
    data.deptIds?.forEach((deptId: number) => {
      deptRef.value?.setChecked(deptId, true, false)
    })
  }, 200)
}
/**
 * 处理清理操作
 * @param row - 要清理数据对象
 */
const handleDrop = async (row: ResSysUser) => {
  await useHandleData(dropSysUserApi, Number(row.id), '清理用户')
  proTable.value?.getTableList()
}
/**
 * 处理删除操作
 * @param row - 要删除数据对象
 */
const handleDelete = async (row: ResSysUser) => {
  await useHandleData(deleteSysUserApi, Number(row.id), '删除用户')
  proTable.value?.getTableList()
}
/**
 * 处理恢复操作
 * @param row - 要恢复数据对象
 */
const handleRecover = async (row: ResSysUser) => {
  await useHandleData(recoverSysUserApi, Number(row.id), '恢复用户')
  proTable.value?.getTableList()
}

// 获取菜单选项
const getOptions = async () => {
  const { data: deptItem } = await getSysDeptListSimpleApi()
  deptOptions.value = handleTree(deptItem)
  const { data: roleItem } = await getSysRoleListSimpleApi()
  roleOptions.value = roleItem
  const { data: postItem } = await getSysPostListSimpleApi()
  postOptions.value = handleTree(postItem)
}

/** 全选/全不选 */
const handleDeptSelect = () => {
  let data = deptSelect.value ? deptOptions.value : []
  deptRef.value!.setCheckedNodes(data as unknown as Node[])
}

/** 展开/折叠全部 */
const handleDeptExpand = () => {
  const nodes = deptRef.value?.store.nodesMap
  for (let node in nodes) {
    if (nodes[node].expanded === deptExpand.value) {
      continue
    }
    nodes[node].expanded = deptExpand.value
  }
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
    // loading.value = true
    const data = sysUserForm.value as unknown as ResSysUser
    data.deptIds = [
      ...(deptRef.value!.getCheckedKeys(false) as unknown as Array<number>), // 获得当前选中节点
      ...(deptRef.value!.getHalfCheckedKeys() as unknown as Array<number>), // 获得半选中的父节点
    ]
    if (data.id !== 0) {
      await useHandleSet(updateSysUserApi, data.id, data, '修改用户')
    } else {
      data.password = encryptPassword(data.password)
      await useHandleData(addSysUserApi, data, '添加用户')
    }
    resetForm(formEl)
    // loading.value = false
    proTable.value?.getTableList()
  })
}

// 登录
const handleLogin = async (row: ResSysUser) => {
  try {
    const { data } = await postSysUserLogin(Number(row.id))
    userStore.setUser(data)
    // 2.添加动态路由
    await initDynamicRouter()
    // 3.清空 tabs、keepAlive 数据
    tabsStore.setTabs([])
    keepAliveStore.setKeepAliveName([])
    // 4.跳转到首页
    const { path, queryParams } = parseRedirect(route.query)
    router.push({ path, query: queryParams })
    ElNotification({
      title: getTimeState(),
      message: '欢迎登录',
      type: 'success',
      duration: 3000,
    })
  } catch {
    ElMessage.error({ message: '登录失败' })
  }
}
//删除状态
const deletedEnum = getIntDictOptions('deleted')
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasAuth('user.SysUserDelete')
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

const columns: ColumnProps<ResSysUser>[] = [
  { prop: 'id', label: '编号' },
  { prop: 'name', label: '姓名', search: { el: 'input', span: 2 } },
  { prop: 'mobile', label: '手机号码', search: { el: 'input', span: 2 } },
  { prop: 'username', label: '用户名', search: { el: 'input', span: 2 } },
  { prop: 'status', label: '状态', tag: true, enum: statusEnum, search: { el: 'select', span: 2 } },
  { prop: 'deleted', label: '删除', tag: true, enum: deletedEnum, search: deleteSearch },
  { prop: 'creator', label: '创建人' },
  { prop: 'createTime', label: '创建时间' },
  { prop: 'updater', label: '更新人' },
  { prop: 'updateTime', label: '更新时间' },
  {
    prop: 'operation',
    label: '操作',
    width: 220,
    fixed: 'right',
    isShow: HasAuth(
      'user.SysUserUpdate',
      'user.SysUserDelete',
      'user.SysUserDrop',
      'user.SysUserRecover',
      'user.SysUser',
      'user.SysUserLogin'
    ),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
