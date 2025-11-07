<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="字典类型列表"
      row-key="id"
      :columns="columns"
      :toolbar-right="['search', 'refresh', 'export', 'layout']"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12"
    >
      <template #toolbarLeft>
        <el-button v-auth="'dict.SysDictTypeCreate'" type="primary" :icon="CirclePlus" @click="handleAdd"
          >新增</el-button
        >
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'dict.SysDictType'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['dict.SysDictTypeUpdate', 'dict.SysDictTypeDelete']"
            type="primary"
            link
            :icon="DArrowRight"
          >
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'dict.SysDictTypeUpdate'" :icon="EditPen" @click="handleUpdate(scope.row)">
                编辑
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 0"
                v-auth="'dict.SysDictTypeDelete'"
                :icon="Delete"
                @click="handleDelete(scope.row)"
              >
                删除
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
      <el-form ref="refSysDictTypeForm" :model="sysDictTypeForm" :rules="rulesSysDictTypeForm" label-width="100px">
        <el-form-item label="编号" prop="id">
          <el-input v-model="sysDictTypeForm.id" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="字典名称" prop="name">
          <el-input v-model="sysDictTypeForm.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="字典类型" prop="type">
          <el-input v-model="sysDictTypeForm.type" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="状态:0正常/1停用" prop="status">
          <el-input v-model="sysDictTypeForm.status" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="sysDictTypeForm.remark" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="创建人" prop="creator">
          <el-input v-model="sysDictTypeForm.creator" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="创建时间" prop="createTime">
          <el-input v-model="sysDictTypeForm.createTime" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="更新人" prop="updater">
          <el-input v-model="sysDictTypeForm.updater" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="更新时间" prop="updateTime">
          <el-input v-model="sysDictTypeForm.updateTime" :disabled="disabled" />
        </el-form-item>
      </el-form>
      <template v-if="!disabled" #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSysDictTypeForm)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSysDictTypeForm)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx">
defineOptions({ name: 'SysDictType' })
import type { ResSysDictType } from '@/api/interface/sysDictType'
import type { ProTableInstance, ColumnProps, SearchProps } from '@/components/ProTable/interface'
import { EditPen, CirclePlus, Delete, View, DArrowRight } from '@element-plus/icons-vue'
import {
  getSysDictTypeListApi,
  deleteSysDictTypeApi,
  getSysDictTypeApi,
  addSysDictTypeApi,
  updateSysDictTypeApi,
} from '@/api/modules/sysDictType'
import type { FormInstance, FormRules } from 'element-plus'
import { useHandleData, useHandleSet } from '@/hooks/useHandleData'
import { HasPermission } from '@/utils/permission'
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
const sysDictTypeForm = ref<ResSysDictType>({
  id: 0, // 编号
  name: '', // 字典名称
  type: '', // 字典类型
  status: 0, // 状态:0正常/1停用
  remark: '', // 备注
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
})
//表单
const refSysDictTypeForm = ref<FormInstance>()
//校验
const rulesSysDictTypeForm = reactive<FormRules>({
  id: [{ required: true, message: '编号不能为空', trigger: 'blur' }],
  name: [{ required: true, message: '字典名称不能为空', trigger: 'blur' }],
  type: [{ required: true, message: '字典类型不能为空', trigger: 'blur' }],
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
  return getSysDictTypeListApi(newParams)
}

/**
 * 重置数据
 * 该函数用于清空或初始化
 * @returns {void} 无返回值
 */
const resetSysDictType = () => {
  // 设置sysDictTypeForm为默认值
  Object.assign(sysDictTypeForm.value, {
    id: 0, // 编号
    name: '', // 字典名称
    type: '', // 字典类型
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
  resetSysDictType()
  disabled.value = true
}
/**
 * 处理新增操作
 * 该函数用于初始化对话框的状态
 */
const handleAdd = () => {
  title.value = '新增字典类型'
  dialogVisible.value = true
  reset()
  disabled.value = false
}
/**
 * 处理更新操作
 * @param row - 要更新数据对象
 */
const handleUpdate = async (row: ResSysDictType) => {
  title.value = '编辑字典类型'
  dialogVisible.value = true
  reset()
  const data = await getSysDictTypeApi(Number(row.id))
  sysDictTypeForm.value = data
  disabled.value = false
}
/**
 * 处理查看操作
 * @param row - 要查看数据对象
 */
const handleItem = async (row: ResSysDictType) => {
  title.value = '查看字典类型'
  dialogVisible.value = true
  reset()
  const data = await getSysDictTypeApi(Number(row.id))
  sysDictTypeForm.value = data
  disabled.value = true
}
/**
 * 处理删除操作
 * @param row - 要删除数据对象
 */
const handleDelete = async (row: ResSysDictType) => {
  await useHandleData(deleteSysDictTypeApi, Number(row.id), '删除字典类型')
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
    const data = sysDictTypeForm.value as unknown as ResSysDictType
    if (data.id !== 0) {
      await useHandleSet(updateSysDictTypeApi, data.id, data, '修改字典类型')
    } else {
      await useHandleData(addSysDictTypeApi, data, '添加字典类型')
    }
    resetForm(formEl)
    loading.value = false
    proTable.value?.getTableList()
  })
}

const columns: ColumnProps<ResSysDictType>[] = [
  { prop: 'id', label: '编号' },
  { prop: 'name', label: '字典名称', search: { valueType: 'input' } },
  { prop: 'type', label: '字典类型', search: { valueType: 'input' } },
  { prop: 'status', label: '状态:0正常/1停用', search: { valueType: 'input' } },
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
    isShow: HasPermission('dict.SysDictTypeUpdate', 'dict.SysDictTypeDelete', 'dict.SysDictType'),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
