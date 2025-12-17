<template>
  <div class="table-box">
    <ProVxeTable
      ref="proTable"
      title="角色列表"
      :columns="columns"
      :toolbar-right="['search', 'refresh']"
      :request-api="getTableList"
      :request-auto="true"
      :show-number="3"
      :border="true"
      :column-config="{ resizable: true, isCurrent: true, isHover: true }"
      :row-config="{ isCurrent: true, isHover: true }"
      :pagination="ProTablePaginationEnum.BE"
    >
      <template #toolbarLeft>
        <el-button v-auth="'role.SysRoleCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <vxe-column field="id" title="编号" fixed="left" width="auto"> </vxe-column>
      <vxe-column field="name" title="名称"> </vxe-column>
      <vxe-column field="code" title="编码"> </vxe-column>
      <vxe-column field="scope" title="数据范围">
        <template #default="{ row }">
          <DictTag type="role.scope" :value="row.scope" />
        </template>
      </vxe-column>
      <vxe-column field="sort" title="排序"> </vxe-column>
      <vxe-column field="status" title="状态">
        <template #default="{ row }">
          <DictTag type="status" :value="row.status" />
        </template>
      </vxe-column>
      <vxe-column v-auth="'role.SysRoleDelete'" field="deleted" title="删除">
        <template #default="{ row }">
          <DictTag type="deleted" :value="row.deleted" />
        </template>
      </vxe-column>
      <vxe-column field="creator" title="创建人"> </vxe-column>
      <vxe-column field="createTime" title="创建时间"></vxe-column>
      <vxe-column field="updater" title="更新人"> </vxe-column>
      <vxe-column field="updateTime" title="更新时间"></vxe-column>
      <vxe-column
        v-auth="['role.SysRoleUpdate', 'role.SysRoleDelete', 'role.SysRoleDrop', 'role.SysRoleRecover', 'role.SysRole']"
        field="operation"
        title="操作"
        width="auto"
        fixed="right"
      >
        <template #default="{ row }">
          <el-button v-auth="'role.SysRole'" type="primary" link :icon="View" @click="handleItem(row)">
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
                  <el-dropdown-item :icon="EditPen" @click="handleUpdate(row)"> 编辑 </el-dropdown-item>
                </div>
                <div v-auth="'role.SysRoleDelete'">
                  <el-dropdown-item v-if="row.deleted === 0" :icon="Delete" @click="handleDelete(row)">
                    删除
                  </el-dropdown-item>
                </div>
                <div v-auth="'role.SysRoleRecover'">
                  <el-dropdown-item v-if="row.deleted === 1" :icon="Refresh" @click="handleRecover(row)">
                    恢复
                  </el-dropdown-item>
                </div>
                <div v-auth="'role.SysRoleDrop'">
                  <el-dropdown-item v-if="row.deleted === 1" :icon="DeleteFilled" @click="handleDrop(row)">
                    清理
                  </el-dropdown-item>
                </div>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
      </vxe-column>
      <!-- 菜单操作 -->
    </ProVxeTable>
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
import { ResSysRole } from '@/api/interface/sysRole'
import { ProVxeTableInstance, ProVxeColumnProps } from '@/components/ProVxeTable/interface'
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
import { FormInstance, FormRules, ElTree } from 'element-plus'
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
import Node from 'element-plus/es/components/tree/src/model/node'
// 获取loading状态
const { loading } = storeToRefs(useLoadingStore())
//禁用
const disabled = ref(true)
//弹出层标题
const title = ref('')
//列表数据
const proTable = ref<ProVxeTableInstance>()
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

const columns: ProVxeColumnProps[] = [
  { prop: 'name', label: '名称', valueType: 'input' },
  { prop: 'code', label: '编码', valueType: 'input' },
  { prop: 'status', label: '状态', valueType: 'select', options: statusEnum },
  {
    prop: 'deleted',
    label: '删除',
    valueType: 'switch',
    fieldProps: {
      activeValue: 1,
      inactiveValue: 0,
    },
    hideInSearch: !HasAuth('role.SysRoleDelete'),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
