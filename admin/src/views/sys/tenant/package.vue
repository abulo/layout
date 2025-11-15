<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="租户套餐列表"
      row-key="id"
      :columns="columns"
      :toolbar-right="['search', 'refresh', 'export', 'layout']"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12"
    >
      <template #toolbarLeft>
        <el-button v-auth="'tenant.SysTenantPackageCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">
          新增
        </el-button>
      </template>
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'tenant.SysTenantPackage'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['tenant.SysTenantPackageUpdate', 'tenant.SysTenantPackageDelete']"
            type="primary"
            link
            :icon="DArrowRight"
          >
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <div v-auth="'tenant.SysTenantPackageUpdate'">
                <el-dropdown-item :icon="EditPen" @click="handleUpdate(scope.row)"> 编辑 </el-dropdown-item>
              </div>
              <div v-auth="'tenant.SysTenantPackageDelete'">
                <el-dropdown-item v-if="scope.row.deleted === 0" :icon="Delete" @click="handleDelete(scope.row)">
                  删除
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
    >
      <el-form
        ref="refSysTenantPackageForm"
        :model="sysTenantPackageForm"
        :rules="rulesSysTenantPackageForm"
        label-width="100px"
      >
        <el-form-item label="名称" prop="name">
          <el-input v-model="sysTenantPackageForm.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="数据范围">
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
                value: 'id',
              }"
              :list="sysTenantPackageForm.scopeMenu"
              empty-text="加载中，请稍候"
              node-key="id"
              show-checkbox
            />
          </el-card>
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number
            v-model="sysTenantPackageForm.sort"
            controls-position="right"
            :min="0"
            :disabled="disabled"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="sysTenantPackageForm.status">
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
        <el-form-item label="备注" prop="remark">
          <el-input v-model="sysTenantPackageForm.remark" :disabled="disabled" />
        </el-form-item>
      </el-form>
      <template v-if="!disabled" #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSysTenantPackageForm)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSysTenantPackageForm)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx">
defineOptions({ name: 'SysTenantPackage' })
import type { ResSysTenantPackage } from '@/api/interface/sysTenantPackage'
import type { ProTableInstance, ColumnProps } from '@/components/ProTable/interface'
import { EditPen, CirclePlus, Delete, View, DArrowRight } from '@element-plus/icons-vue'
import {
  getSysTenantPackageListApi,
  deleteSysTenantPackageApi,
  getSysTenantPackageApi,
  addSysTenantPackageApi,
  updateSysTenantPackageApi,
} from '@/api/modules/sysTenantPackage'
import type { FormInstance, FormRules, ElTree } from 'element-plus'
import { useHandleData, useHandleSet } from '@/hooks/useHandleData'
import { HasAuth } from '@/utils/auth'
import { getIntDictOptions } from '@/utils/dict'
import type { ResSysMenu } from '@/api/interface/sysMenu'
import { getSysMenuListSimpleApi } from '@/api/modules/sysMenu'
import { handleTree } from '@pureadmin/utils'
import type Node from 'element-plus/es/components/tree/src/model/node'
import { useTimeoutFn } from '@vueuse/core'
//加载
const loading = ref(false)
//禁用
const disabled = ref(true)
//弹出层标题
const title = ref('')
//列表数据
const proTable = ref<ProTableInstance>()
//显示弹出层
const dialogVisible = ref(false)
// 获取选择
const menuSelect = ref(false)
// 菜单ref
const menuRef = ref<InstanceType<typeof ElTree>>()
// 菜单数据
const menuOptions = ref<ResSysMenu[]>([])
// 菜单展开/折叠
const menuExpand = ref(false)
// 状态枚举
const statusEnum = getIntDictOptions('status')
//数据接口
const sysTenantPackageForm = ref<ResSysTenantPackage>({
  id: 0, // 编号
  name: '', // 名称
  scopeMenu: undefined, // 数据范围
  sort: 0, // 排序
  status: 0, // 状态:0正常/1停用
  remark: '', // 备注
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
})
//表单
const refSysTenantPackageForm = ref<FormInstance>()
//校验
const rulesSysTenantPackageForm = reactive<FormRules>({
  id: [{ required: true, message: '编号不能为空', trigger: 'blur' }],
  name: [{ required: true, message: '名称不能为空', trigger: 'blur' }],
  sort: [{ required: true, message: '排序不能为空', trigger: 'blur' }],
  status: [{ required: true, message: '状态:0正常/1停用不能为空', trigger: 'blur' }],
  remark: [{ required: true, message: '备注不能为空', trigger: 'blur' }],
})

/**
 * 获取表格数据列表
 * @param params 查询参数对象
 * @returns 返回列表数据的Promise对象
 */
const getTableList = (params: any) => {
  // 深拷贝参数对象，避免修改原始参数
  let newParams = JSON.parse(JSON.stringify(params))
  return getSysTenantPackageListApi(newParams)
}

/**
 * 重置数据
 * 该函数用于清空或初始化
 * @returns {void} 无返回值
 */
const resetSysTenantPackage = () => {
  // 设置sysTenantPackageForm为默认值
  Object.assign(sysTenantPackageForm.value, {
    id: 0, // 编号
    name: '', // 名称
    scopeMenu: undefined, // 数据范围
    sort: 0, // 排序
    status: 0, // 状态:0正常/1停用
    remark: '', // 备注
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
  loading.value = false
  resetSysTenantPackage()
  disabled.value = true
  menuSelect.value = false
  menuExpand.value = false
  menuRef.value?.setCheckedNodes([])
  sysTenantPackageForm.value?.resetFields()
}
/**
 * 处理新增操作
 * 该函数用于初始化对话框的状态
 */
const handleAdd = async () => {
  title.value = '新增租户套餐'
  dialogVisible.value = true
  reset()
  getMenuOptions()
  disabled.value = false
}
/**
 * 处理更新操作
 * @param row - 要更新数据对象
 */
const handleUpdate = async (row: ResSysTenantPackage) => {
  title.value = '编辑租户套餐'
  dialogVisible.value = true
  reset()
  getMenuOptions()
  const data = await getSysTenantPackageApi(Number(row.id))
  sysTenantPackageForm.value = data
  disabled.value = false
  useTimeoutFn(() => {
    data.scopeMenu?.forEach((menuId: number) => {
      menuRef.value?.setChecked(menuId, true, false)
    })
  }, 200)
}
/**
 * 处理查看操作
 * @param row - 要查看数据对象
 */
const handleItem = async (row: ResSysTenantPackage) => {
  title.value = '查看租户套餐'
  dialogVisible.value = true
  reset()
  getMenuOptions()
  const data = await getSysTenantPackageApi(Number(row.id))
  sysTenantPackageForm.value = data
  disabled.value = true
  useTimeoutFn(() => {
    data.scopeMenu?.forEach((menuId: number) => {
      menuRef.value?.setChecked(menuId, true, false)
    })
  }, 200)
}
/**
 * 处理删除操作
 * @param row - 要删除数据对象
 */
const handleDelete = async (row: ResSysTenantPackage) => {
  await useHandleData(deleteSysTenantPackageApi, Number(row.id), '删除租户套餐')
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
    loading.value = true
    const data = sysTenantPackageForm.value as unknown as ResSysTenantPackage
    data.scopeMenu = [
      ...(menuRef.value!.getCheckedKeys(false) as unknown as Array<number>), // 获得当前选中节点
      ...(menuRef.value!.getHalfCheckedKeys() as unknown as Array<number>), // 获得半选中的父节点
    ]
    if (data.id !== 0) {
      await useHandleSet(updateSysTenantPackageApi, data.id, data, '修改租户套餐')
    } else {
      await useHandleData(addSysTenantPackageApi, data, '添加租户套餐')
    }
    resetForm(formEl)
    loading.value = false
    proTable.value?.getTableList()
  })
}

// 获取菜单选项
const getMenuOptions = async () => {
  const data = await getSysMenuListSimpleApi()
  menuOptions.value = handleTree(data)
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

const columns: ColumnProps<ResSysTenantPackage>[] = [
  { prop: 'id', label: '编号' },
  { prop: 'name', label: '名称', search: { el: 'input', span: 2 } },
  { prop: 'scopeMenu', label: '数据范围' },
  { prop: 'sort', label: '排序' },
  { prop: 'status', label: '状态', tag: true, enum: statusEnum, search: { el: 'select', span: 2 } },
  { prop: 'remark', label: '备注' },
  { prop: 'creator', label: '创建人' },
  { prop: 'createTime', label: '创建时间' },
  { prop: 'updater', label: '更新人' },
  { prop: 'updateTime', label: '更新时间' },

  {
    prop: 'operation',
    label: '操作',
    width: 150,
    fixed: 'right',
    isShow: HasAuth('tenant.SysTenantPackageUpdate', 'tenant.SysTenantPackageDelete', 'tenant.SysTenantPackage'),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
