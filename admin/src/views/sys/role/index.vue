<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="角色列表"
      row-key="id"
      :columns="columns"
      :toolbar-right="['search', 'refresh', 'export', 'layout']"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="ProTablePaginationEnum.BE"
      :search-col="12"
    >
      <template #toolbarLeft>
        <el-button v-auth="'role.SysRoleCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="deleted" :value="scope.row.deleted" />
      </template>
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <template #scope="scope">
        <DictTag type="role.scope" :value="scope.row.scope" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'role.SysRole'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['role.SysRoleUpdate', 'role.SysRoleDelete', 'role.SysRoleRecover', 'role.SysRoleDrop']"
            type="primary"
            link
            :icon="DArrowRight"
          >
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <div v-auth="'role.SysRoleUpdate'">
                <el-dropdown-item :icon="EditPen" @click="handleUpdate(scope.row)"> 编辑 </el-dropdown-item>
              </div>
              <div v-auth="'role.SysRoleDelete'">
                <el-dropdown-item v-if="scope.row.deleted === 0" :icon="Delete" @click="handleDelete(scope.row)">
                  删除
                </el-dropdown-item>
              </div>
              <div v-auth="'role.SysRoleRecover'">
                <el-dropdown-item v-if="scope.row.deleted === 1" :icon="Refresh" @click="handleRecover(scope.row)">
                  恢复
                </el-dropdown-item>
              </div>
              <div v-auth="'role.SysRoleDrop'">
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
      <el-form ref="refSysRoleForm" :model="sysRoleForm" :rules="rulesSysRoleForm" label-width="100px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="sysRoleForm.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="编码" prop="code">
          <el-input v-model="sysRoleForm.code" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="菜单权限">
          <el-card class="cardHeight">
            <template #header>
              全选/不选:
              <el-switch
                v-model="menuSelect"
                active-text="是"
                inactive-text="否"
                inline-prompt
                @change="handleMenuSelect"
              />
              展开/折叠:
              <el-switch
                v-model="menuExpand"
                active-text="展开"
                inactive-text="折叠"
                inline-prompt
                @change="handleMenuExpand"
              />
            </template>
            <el-tree
              ref="menuRef"
              :data="menuOptions"
              :props="{
                children: 'children',
                label: 'name',
              }"
              :list="sysRoleForm.menuIds"
              empty-text="加载中，请稍候"
              node-key="id"
              show-checkbox
            />
          </el-card>
        </el-form-item>
        <el-form-item label="数据范围" prop="scope">
          <el-select v-model="sysRoleForm.scope" :disabled="disabled">
            <el-option
              v-for="dict in roleScopeEnum"
              :key="Number(dict.value)"
              :label="dict.label"
              :value="Number(dict.value)"
            />
          </el-select>
        </el-form-item>
        <el-form-item v-if="sysRoleForm.scope === 2" label="权限范围" prop="scopeDept">
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
              :list="sysRoleForm.scopeDept"
              empty-text="加载中，请稍候"
              node-key="id"
              show-checkbox
            />
          </el-card>
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="sysRoleForm.sort" controls-position="right" :min="0" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="sysRoleForm.status">
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
          <el-button @click="resetForm(refSysRoleForm)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSysRoleForm)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx">
defineOptions({ name: 'SysRole' })
import type { ResSysRole } from '@/api/interface/sysRole'
import type { ProTableInstance, ColumnProps, SearchProps } from '@/components/ProTable/interface'
import { EditPen, CirclePlus, Delete, Refresh, DeleteFilled, View, DArrowRight } from '@element-plus/icons-vue'
import {
  getSysRoleListApi,
  deleteSysRoleApi,
  dropSysRoleApi,
  recoverSysRoleApi,
  getSysRoleApi,
  addSysRoleApi,
  updateSysRoleApi,
} from '@/api/modules/sysRole'
import type { FormInstance, FormRules, ElTree } from 'element-plus'
import { getIntDictOptions } from '@/utils/dict'
import { DictTag } from '@/components/DictTag'
import { useHandleData, useHandleSet } from '@/hooks/useHandleData'
import { HasAuth } from '@/utils/auth'
import { useLoadingStore } from '@/stores/modules/loading'
import { storeToRefs } from 'pinia'
import { ProTablePaginationEnum } from '@/enums'
import { ResSysMenu } from '@/api/interface/sysMenu'
import { getSysMenuListSimpleApi } from '@/api/modules/sysMenu'
import { ResSysDept } from '@/api/interface/sysDept'
import { getSysDeptListSimpleApi } from '@/api/modules/sysDept'
import { handleTree } from '@pureadmin/utils'
import { useTimeoutFn } from '@vueuse/core'
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
// 数据范围枚举
const roleScopeEnum = getIntDictOptions('role.scope')
// 获取选择
const menuSelect = ref(false)
// 菜单ref
const menuRef = ref<InstanceType<typeof ElTree>>()
// 菜单数据
const menuOptions = ref<ResSysMenu[]>([])
// 菜单展开/折叠
const menuExpand = ref(false)

// 获取选择
const deptSelect = ref(false)
// 菜单ref
const deptRef = ref<InstanceType<typeof ElTree>>()
// 菜单数据
const deptOptions = ref<ResSysDept[]>([])
// 菜单展开/折叠
const deptExpand = ref(false)

//数据接口
const sysRoleForm = ref<ResSysRole>({
  id: 0, // 编号
  name: '', // 名称
  code: '', // 编码
  scope: undefined, // 数据范围
  scopeDept: undefined, // 数据范围
  sort: 0, // 排序
  status: 0, // 状态:0正常/1停用
  deleted: 0, // 删除:0否/1是
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
  menuIds: undefined,
})
//表单
const refSysRoleForm = ref<FormInstance>()
//校验
const rulesSysRoleForm = reactive<FormRules>({
  name: [{ required: true, message: '名称不能为空', trigger: 'blur' }],
  code: [{ required: true, message: '编码不能为空', trigger: 'blur' }],
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
  return getSysRoleListApi(newParams)
}

/**
 * 重置数据
 * 该函数用于清空或初始化
 * @returns {void} 无返回值
 */
const resetSysRole = () => {
  // 设置sysRoleForm为默认值
  Object.assign(sysRoleForm.value, {
    id: 0, // 编号
    name: '', // 名称
    code: '', // 编码
    scope: undefined, // 数据范围:1:全部数据权限/2:自定数据权限/3:本部门数据权限/4:本部门及以下数据权限
    scopeDept: undefined, // 数据范围
    sort: 0, // 排序
    status: 0, // 状态:0正常/1停用
    deleted: 0, // 删除:0否/1是
    tenantId: 0, // 租户
    creator: undefined, // 创建人
    createTime: undefined, // 创建时间
    updater: undefined, // 更新人
    updateTime: undefined, // 更新时间
    menuIds: undefined,
  })
}

/**
 * 重置函数
 * @returns {void}
 */
const reset = () => {
  // loading.value = false
  resetSysRole()
  disabled.value = true
  menuSelect.value = false
  menuExpand.value = false
  menuRef.value?.setCheckedNodes([])

  deptSelect.value = false
  deptExpand.value = false
  deptRef.value?.setCheckedNodes([])
}
/**
 * 处理新增操作
 * 该函数用于初始化对话框的状态
 */
const handleAdd = () => {
  title.value = '新增角色'
  dialogVisible.value = true
  reset()
  getOptions()
  disabled.value = false
}
/**
 * 处理更新操作
 * @param row - 要更新数据对象
 */
const handleUpdate = async (row: ResSysRole) => {
  title.value = '编辑角色'
  dialogVisible.value = true
  reset()
  getOptions()
  const { data } = await getSysRoleApi(Number(row.id))
  sysRoleForm.value = data
  disabled.value = false
  useTimeoutFn(() => {
    data.menuIds?.forEach((menuId: number) => {
      menuRef.value?.setChecked(menuId, true, false)
    })
    data.scopeDept?.forEach((deptId: number) => {
      deptRef.value?.setChecked(deptId, true, false)
    })
  }, 200)
}
/**
 * 处理查看操作
 * @param row - 要查看数据对象
 */
const handleItem = async (row: ResSysRole) => {
  title.value = '查看角色'
  dialogVisible.value = true
  reset()
  getOptions()
  const { data } = await getSysRoleApi(Number(row.id))
  sysRoleForm.value = data
  disabled.value = true
  useTimeoutFn(() => {
    data.menuIds?.forEach((menuId: number) => {
      menuRef.value?.setChecked(menuId, true, false)
    })
    data.scopeDept?.forEach((deptId: number) => {
      deptRef.value?.setChecked(deptId, true, false)
    })
  }, 200)
}
/**
 * 处理清理操作
 * @param row - 要清理数据对象
 */
const handleDrop = async (row: ResSysRole) => {
  await useHandleData(dropSysRoleApi, Number(row.id), '清理角色')
  proTable.value?.getTableList()
}
/**
 * 处理删除操作
 * @param row - 要删除数据对象
 */
const handleDelete = async (row: ResSysRole) => {
  await useHandleData(deleteSysRoleApi, Number(row.id), '删除角色')
  proTable.value?.getTableList()
}
/**
 * 处理恢复操作
 * @param row - 要恢复数据对象
 */
const handleRecover = async (row: ResSysRole) => {
  await useHandleData(recoverSysRoleApi, Number(row.id), '恢复角色')
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
    const data = sysRoleForm.value as unknown as ResSysRole
    data.menuIds = [
      ...(menuRef.value!.getCheckedKeys(false) as unknown as Array<number>), // 获得当前选中节点
      ...(menuRef.value!.getHalfCheckedKeys() as unknown as Array<number>), // 获得半选中的父节点
    ]
    data.scopeDept = [
      ...(deptRef.value!.getCheckedKeys(false) as unknown as Array<number>), // 获得当前选中节点
      ...(deptRef.value!.getHalfCheckedKeys() as unknown as Array<number>), // 获得半选中的父节点
    ]
    if (data.scope !== 2) {
      data.scopeDept = []
    }
    if (data.id !== 0) {
      await useHandleSet(updateSysRoleApi, data.id, data, '修改角色')
    } else {
      await useHandleData(addSysRoleApi, data, '添加角色')
    }
    resetForm(formEl)
    // loading.value = false
    proTable.value?.getTableList()
  })
}
// 获取菜单选项
const getOptions = async () => {
  const { data } = await getSysMenuListSimpleApi()
  menuOptions.value = handleTree(data)
  const { data: deptList } = await getSysDeptListSimpleApi()
  deptOptions.value = handleTree(deptList)
}

/** 全选/全不选 */
const handleMenuSelect = () => {
  let data = menuSelect.value ? menuOptions.value : []
  menuRef.value!.setCheckedNodes(data as unknown as Node[])
}

/** 展开/折叠全部 */
const handleMenuExpand = () => {
  const nodes = menuRef.value?.store.nodesMap
  for (let node in nodes) {
    if (nodes[node].expanded === menuExpand.value) {
      continue
    }
    nodes[node].expanded = menuExpand.value
  }
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
//删除状态
const deletedEnum = getIntDictOptions('deleted')
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasAuth('role.SysRoleDelete')
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

const columns: ColumnProps<ResSysRole>[] = [
  { prop: 'id', label: '编号' },
  { prop: 'name', label: '名称', search: { el: 'input', span: 2 } },
  { prop: 'code', label: '编码', search: { el: 'input', span: 2 } },
  { prop: 'scope', label: '数据范围' },
  { prop: 'sort', label: '排序' },
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
      'role.SysRoleUpdate',
      'role.SysRoleDelete',
      'role.SysRoleDrop',
      'role.SysRoleRecover',
      'role.SysRole'
    ),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
