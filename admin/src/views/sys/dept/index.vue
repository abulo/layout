<template>
  <div class="table-box">
    <ProTable
      v-if="refreshTable"
      ref="proTable"
      title="部门列表"
      row-key="id"
      :columns="columns"
      :toolbar-right="['search', 'refresh', 'export', 'layout']"
      :request-api="getTableList"
      :default-expand-all="isExpandAll"
      :data-callback="deptHandleTree"
      :request-auto="true"
      :pagination="ProTablePaginationEnum.NONE"
      :search-col="12"
      :indent="20"
    >
      <template #toolbarLeft>
        <el-button v-auth="'dept.SysDeptCreate'" type="primary" :icon="CirclePlus" @click="handleAdd()">新增</el-button>
        <el-button type="primary" :icon="Sort" @click="handleExpandAll">展开/折叠</el-button>
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
        <el-button v-auth="'dept.SysDept'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="[
              'dept.SysDeptCreate',
              'dept.SysDeptUpdate',
              'dept.SysDeptDelete',
              'dept.SysDeptRecover',
              'dept.SysDeptDrop',
            ]"
            type="primary"
            link
            :icon="DArrowRight"
          >
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <div v-auth="'dept.SysDeptCreate'">
                <el-dropdown-item :icon="CirclePlus" @click="handleAdd(scope.row)"> 新增 </el-dropdown-item>
              </div>
              <div v-auth="'dept.SysDeptUpdate'">
                <el-dropdown-item :icon="EditPen" @click="handleUpdate(scope.row)"> 编辑 </el-dropdown-item>
              </div>
              <div v-auth="'dept.SysDeptDelete'">
                <el-dropdown-item v-if="scope.row.deleted === 0" :icon="Delete" @click="handleDelete(scope.row)">
                  删除
                </el-dropdown-item>
              </div>
              <div v-auth="'dept.SysDeptRecover'">
                <el-dropdown-item v-if="scope.row.deleted === 1" :icon="Refresh" @click="handleRecover(scope.row)">
                  恢复
                </el-dropdown-item>
              </div>
              <div v-auth="'dept.SysDeptDrop'">
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
      @click="handleDialogClick"
    >
      <el-form ref="refSysDeptForm" :model="sysDeptForm" :rules="rulesSysDeptForm" label-width="100px">
        <el-form-item label="上级菜单" prop="parentId">
          <el-tree-select
            v-model="sysDeptForm.parentId"
            :data="deptOptions"
            :props="{ value: 'id', label: 'name' }"
            value-key="id"
            node-key="id"
            placeholder="请选择"
            check-strictly
            :disabled="disabled"
            :render-after-expand="false"
          />
        </el-form-item>
        <el-form-item label="名称" prop="name">
          <el-input v-model="sysDeptForm.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="sysDeptForm.sort" controls-position="right" :min="0" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="负责人" prop="userId">
          <el-popover placement="bottom-start" :width="800" :show-arrow="false" trigger="click" :visible="isUserOpen">
            <template #reference>
              <el-button class="mr-4" @click.stop="userOpen">{{ userItem }}</el-button>
            </template>
            <div class="table-box">
              <ProTable
                ref="userProTable"
                title="用户列表"
                row-key="id"
                :columns="userColumns"
                :request-api="getCustomSysUserListApi"
                :request-auto="false"
                :tool-button="false"
                :layout="'prev, pager, next'"
                :init-param="initUserParam"
                :pagination="ProTablePaginationEnum.BE"
              >
                <template #operation="scope">
                  <el-button type="primary" link :icon="CirclePlus" @click.stop="handleUser(scope.row)">确定</el-button>
                </template>
              </ProTable>
            </div>
          </el-popover>
        </el-form-item>
        <el-form-item label="联系电话" prop="phone">
          <el-input v-model="sysDeptForm.phone" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="邮件" prop="email">
          <el-input v-model="sysDeptForm.email" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="sysDeptForm.status">
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
          <el-button @click="resetForm(refSysDeptForm)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSysDeptForm)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx">
defineOptions({ name: 'SysDept' })
import { ResSysDept } from '@/api/interface/sysDept'
import { ProTableInstance, ColumnProps, SearchProps } from '@/components/ProTable/interface'
import { EditPen, CirclePlus, Delete, Refresh, DeleteFilled, View, DArrowRight, Sort } from '@element-plus/icons-vue'
import {
  getSysDeptListApi,
  deleteSysDeptApi,
  dropSysDeptApi,
  recoverSysDeptApi,
  getSysDeptApi,
  addSysDeptApi,
  updateSysDeptApi,
  getSysDeptListSimpleApi,
} from '@/api/modules/sysDept'
import { FormInstance, FormRules } from 'element-plus'
import { getIntDictOptions } from '@/utils/dict'
import { DictTag } from '@/components/DictTag'
import { useHandleData, useHandleSet } from '@/hooks/useHandleData'
import { HasAuth } from '@/utils/auth'
import { useLoadingStore } from '@/stores/modules/loading'
import { storeToRefs } from 'pinia'
import { handleTree } from '@pureadmin/utils'
import { ResSysUser } from '@/api/interface/sysUser'
import { getSysUserApi, getSysUserListSimpleApi } from '@/api/modules/sysUser'
import { ProTablePaginationEnum } from '@/enums'
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
//是否展开，默认全部折叠
const isExpandAll = ref(false)
//重新渲染表格状态
const refreshTable = ref(true)
// 定义负责人
const userItem = ref('点击选择')
// 显示负责人
const isUserOpen = ref(false)
// 默认参数
const initUserParam = reactive({ pageSize: 5 })
// 状态枚举
const statusEnum = getIntDictOptions('status')
//数据接口
const sysDeptForm = ref<ResSysDept>({
  id: 0, // 编号
  name: '', // 名称
  parentId: 0, // 上级
  sort: 0, // 排序
  userId: undefined, // 负责人
  phone: undefined, // 联系电话
  email: undefined, // 邮件
  status: 0, // 状态:0正常/1停用
  deleted: 0, // 删除:0否/1是
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
})
//下拉菜单选项
const deptOptions = ref<ResSysDept[]>([])
//表单
const refSysDeptForm = ref<FormInstance>()
//校验
const rulesSysDeptForm = reactive<FormRules>({
  name: [{ required: true, message: '名称不能为空', trigger: 'blur' }],
  parentId: [{ required: true, message: '上级不能为空', trigger: 'blur' }],
  sort: [{ required: true, message: '排序不能为空', trigger: 'blur' }],
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
  return getSysDeptListApi(newParams)
}

/**
 * 重置数据
 * 该函数用于清空或初始化
 * @returns {void} 无返回值
 */
const resetSysDept = () => {
  // 设置sysDeptForm为默认值
  Object.assign(sysDeptForm.value, {
    id: 0, // 编号
    name: '', // 名称
    parentId: 0, // 上级
    sort: 0, // 排序
    userId: undefined, // 负责人
    phone: undefined, // 联系电话
    email: undefined, // 邮件
    status: 0, // 状态:0正常/1停用
    deleted: 0, // 删除:0否/1是
    tenantId: 0, // 租户
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined, // 更新时间
  })
}

/**
 * 重置函数
 * @returns {void}
 */
const reset = () => {
  // loading.value = false
  resetSysDept()
  disabled.value = true
}
/**
 * 处理新增操作
 * 该函数用于初始化对话框的状态
 */
const handleAdd = (row?: ResSysDept) => {
  title.value = '新增部门'
  dialogVisible.value = true
  reset()
  getDeptOptions()
  disabled.value = false
  if (row != null && row.id) {
    sysDeptForm.value.parentId = row.id
  }
}
/**
 * 处理更新操作
 * @param row - 要更新数据对象
 */
const handleUpdate = async (row: ResSysDept) => {
  title.value = '编辑部门'
  dialogVisible.value = true
  reset()
  getDeptOptions()
  const { data } = await getSysDeptApi(Number(row.id))
  sysDeptForm.value = data
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
const handleItem = async (row: ResSysDept) => {
  title.value = '查看部门'
  dialogVisible.value = true
  reset()
  getDeptOptions()
  const { data } = await getSysDeptApi(Number(row.id))
  sysDeptForm.value = data
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
const handleDrop = async (row: ResSysDept) => {
  await useHandleData(dropSysDeptApi, Number(row.id), '清理部门')
  proTable.value?.getTableList()
}
/**
 * 处理删除操作
 * @param row - 要删除数据对象
 */
const handleDelete = async (row: ResSysDept) => {
  await useHandleData(deleteSysDeptApi, Number(row.id), '删除部门')
  proTable.value?.getTableList()
}
/**
 * 处理恢复操作
 * @param row - 要恢复数据对象
 */
const handleRecover = async (row: ResSysDept) => {
  await useHandleData(recoverSysDeptApi, Number(row.id), '恢复部门')
  proTable.value?.getTableList()
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
    const data = sysDeptForm.value as unknown as ResSysDept
    if (data.id !== 0) {
      await useHandleSet(updateSysDeptApi, data.id, data, '修改部门')
    } else {
      await useHandleData(addSysDeptApi, data, '添加部门')
    }
    resetForm(formEl)
    // loading.value = false
    proTable.value?.getTableList()
  })
}

// 设置展开合并
const handleExpandAll = () => {
  refreshTable.value = false
  isExpandAll.value = !isExpandAll.value
  nextTick(() => {
    refreshTable.value = true
    deptOptions.value = []
  })
}

// 获取部门列表
const getDeptOptions = async () => {
  const { data } = await getSysDeptListSimpleApi()
  deptOptions.value = [
    {
      id: 0, // 编号
      name: '主类目', // 名称
      parentId: 0, // 上级
      sort: 0, // 排序
      userId: undefined, // 负责人
      phone: undefined, // 联系电话
      email: undefined, // 邮件
      status: 0, // 状态:0正常/1停用
      deleted: 0, // 删除:0否/1是
      tenantId: 0, // 租户
      creator: undefined, // 创建人
      createTime: undefined, // 创建时间
      updater: undefined, // 更新人
      updateTime: undefined, // 更新时间
      children: handleTree(data),
    },
  ] as unknown as ResSysDept[]
}

// 在 el-dialog 上添加点击事件监听器
const handleDialogClick = () => {
  isUserOpen.value = false
}

const userOpen = () => {
  if (disabled.value) return
  isUserOpen.value = true
  userProTable.value?.getTableList()
}

const getCustomSysUserListApi = (params: any) => {
  const newParams = Object.assign(params, JSON.parse(JSON.stringify(initUserParam)))
  return getSysUserListSimpleApi(newParams)
}

const handleUser = (row: ResSysUser) => {
  userItem.value = row.name as string
  sysDeptForm.value.userId = Number(row.id)
  isUserOpen.value = false
}

const deptHandleTree = (data: ResSysDept[]) => {
  return handleTree(data)
}

const userColumns: ColumnProps<ResSysUser>[] = [
  { prop: 'id', label: '编码', fixed: 'left' },
  { prop: 'name', label: '姓名', search: { el: 'input', span: 1 } },
  { prop: 'mobile', label: '手机号码', search: { el: 'input', span: 1 } },
  { prop: 'username', label: '用户名', search: { el: 'input', span: 1 } },
  {
    prop: 'operation',
    label: '操作',
    fixed: 'right',
  },
]

//删除状态
const deletedEnum = getIntDictOptions('deleted')
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasAuth('dept.SysDeptDelete')
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

const columns: ColumnProps<ResSysDept>[] = [
  { prop: 'name', label: '名称', fixed: 'left', align: 'left' },
  { prop: 'sort', label: '排序' },
  { prop: 'phone', label: '联系电话' },
  { prop: 'email', label: '邮件' },
  { prop: 'status', label: '状态', tag: true, enum: statusEnum, search: { el: 'select', span: 2 } },
  { prop: 'deleted', label: '删除', tag: true, enum: deletedEnum, search: deleteSearch },
  {
    prop: 'operation',
    label: '操作',
    width: 150,
    fixed: 'right',
    isShow: HasAuth(
      'dept.SysDeptCreate',
      'dept.SysDeptUpdate',
      'dept.SysDeptDelete',
      'dept.SysDeptDrop',
      'dept.SysDeptRecover',
      'dept.SysDept'
    ),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
@use '@/styles/popover';
</style>
