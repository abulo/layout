<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="职位列表"
      row-key="id"
      :columns="columns"
      :toolbar-right="['search', 'refresh', 'export', 'layout']"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12"
    >
      <template #toolbarLeft>
        <el-button type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <template #operation="scope">
        <el-button type="primary" link :icon="View" @click="handleItem(scope.row)"> 查看 </el-button>
        <el-dropdown trigger="click">
          <el-button type="primary" link :icon="DArrowRight"> 更多 </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item :icon="EditPen" @click="handleUpdate(scope.row)"> 编辑 </el-dropdown-item>
              <el-dropdown-item v-if="scope.row.deleted === 0" :icon="Delete" @click="handleDelete(scope.row)">
                删除
              </el-dropdown-item>
              <el-dropdown-item v-if="scope.row.deleted === 1" :icon="Refresh" @click="handleRecover(scope.row)">
                恢复
              </el-dropdown-item>
              <el-dropdown-item v-if="scope.row.deleted === 1" :icon="DeleteFilled" @click="handleDrop(scope.row)">
                清理
              </el-dropdown-item>
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
      <el-form ref="refSysPostForm" :model="sysPostForm" :rules="rulesSysPostForm" label-width="100px">
        <el-form-item label="职位名称" prop="name">
          <el-input v-model="sysPostForm.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="显示顺序" prop="sort">
          <el-input-number v-model="sysPostForm.sort" controls-position="right" :min="0" :disabled="disabled" />
        </el-form-item>
      </el-form>
      <template v-if="!disabled" #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSysPostForm)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSysPostForm)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="tsx">
defineOptions({ name: 'SysPost' })
import type { ResSysPost } from '@/api/interface/sysPost'
import type { ProTableInstance, ColumnProps } from '@/components/ProTable/interface'
import { EditPen, CirclePlus, Delete, Refresh, DeleteFilled, View, DArrowRight } from '@element-plus/icons-vue'
import {
  getSysPostListApi,
  getSysPostApi,
  dropSysPostApi,
  deleteSysPostApi,
  recoverSysPostApi,
  updateSysPostApi,
  addSysPostApi,
} from '@/api/modules/sysPost'
import type { FormInstance, FormRules } from 'element-plus'
import { useHandleData, useHandleSet } from '@/hooks/useHandleData'
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
//数据接口
const sysPostForm = ref<ResSysPost>({
  id: 0, // 职位ID
  name: '', // 职位名称
  sort: 0, // 显示顺序
  status: 0, // 状态
  deleted: 0, // 是否删除
  tenantId: 0, // 租户ID
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
})
//表单
const refSysPostForm = ref<FormInstance>()
//校验
const rulesSysPostForm = reactive<FormRules>({
  name: [{ required: true, message: '岗位标题不能为空', trigger: 'blur' }],
  status: [{ required: true, message: '岗位状态不能为空', trigger: 'change' }],
})
/**
 * 获取表格数据列表
 * @param params 查询参数对象
 * @returns 返回列表数据的Promise对象
 */
const getTableList = (params: any) => {
  // 深拷贝参数对象，避免修改原始参数
  let newParams = JSON.parse(JSON.stringify(params))
  return getSysPostListApi(newParams)
}

/**
 * 重置数据
 * 该函数用于清空或初始化
 * @returns {void} 无返回值
 */
const resetSysPost = () => {
  // 设置sysPostForm为默认值
  Object.assign(sysPostForm.value, {
    id: 0, // 职位ID
    name: '', // 职位名称
    sort: 0, // 显示顺序
    status: 0, // 状态
    deleted: 0, // 是否删除
    tenantId: 0, // 租户ID
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
  resetSysPost()
  disabled.value = true
}

/**
 * 处理新增操作
 * 该函数用于初始化对话框的状态
 */
const handleAdd = () => {
  title.value = '新增职位'
  dialogVisible.value = true
  reset()
  disabled.value = false
}
/**
 * 处理更新操作
 * @param row - 要更新数据对象
 */
const handleUpdate = async (row: ResSysPost) => {
  title.value = '编辑职位'
  dialogVisible.value = true
  reset()
  const data = await getSysPostApi(Number(row.id))
  sysPostForm.value = data
  disabled.value = false
}
/**
 * 处理查看操作
 * @param row - 要查看数据对象
 */
const handleItem = async (row: ResSysPost) => {
  title.value = '查看职位'
  dialogVisible.value = true
  reset()
  const data = await getSysPostApi(Number(row.id))
  sysPostForm.value = data
  disabled.value = true
}
/**
 * 处理清理操作
 * @param row - 要清理数据对象
 */
const handleDrop = async (row: ResSysPost) => {
  await useHandleData(dropSysPostApi, Number(row.id), '清理职位')
  proTable.value?.getTableList()
}
/**
 * 处理删除操作
 * @param row - 要删除数据对象
 */
const handleDelete = async (row: ResSysPost) => {
  await useHandleData(deleteSysPostApi, Number(row.id), '删除职位')
  proTable.value?.getTableList()
}
/**
 * 处理恢复操作
 * @param row - 要恢复数据对象
 */
const handleRecover = async (row: ResSysPost) => {
  await useHandleData(recoverSysPostApi, Number(row.id), '恢复职位')
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
    const data = sysPostForm.value as unknown as ResSysPost
    if (data.id !== 0) {
      await useHandleSet(updateSysPostApi, data.id, data, '修改职位')
    } else {
      await useHandleData(addSysPostApi, data, '添加职位')
    }
    resetForm(formEl)
    loading.value = false
    proTable.value?.getTableList()
  })
}

/**
 * 定义表格列配置数组
 * 数组中的每个对象都遵循ColumnProps接口规范
 */
const columns: ColumnProps<ResSysPost>[] = [
  { prop: 'id', label: '编号', width: 100, fixed: 'left' },
  { prop: 'name', label: '岗位名称', search: { el: 'input', span: 2 } },
  { prop: 'sort', label: '岗位顺序' },
  { prop: 'creator', label: '创建人' },
  { prop: 'createTime', label: '创建时间' },
  { prop: 'updater', label: '更新人' },
  { prop: 'updateTime', label: '更新时间' },
  {
    prop: 'operation',
    label: '操作',
    width: 150,
    fixed: 'right',
  },
]
</script>

<style scoped lang="scss">
@use '@/styles/custom';
</style>
