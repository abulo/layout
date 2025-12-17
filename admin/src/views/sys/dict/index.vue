<template>
  <div class="table-box">
    <ProVxeTable
      ref="proTable"
      title="字典类型列表"
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
        <el-button v-auth="'dict.SysDictTypeCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">
          新增
        </el-button>
      </template>
      <vxe-column field="id" title="编号" fixed="left" width="auto"> </vxe-column>
      <vxe-column field="name" title="字典名称"> </vxe-column>
      <vxe-column field="type" title="字典类型"> </vxe-column>
      <vxe-column field="status" title="状态">
        <template #default="{ row }">
          <DictTag type="status" :value="row.status" />
        </template>
      </vxe-column>
      <vxe-column field="remark" title="备注"> </vxe-column>
      <vxe-column field="creator" title="创建人"> </vxe-column>
      <vxe-column field="createTime" title="创建时间"> </vxe-column>
      <vxe-column field="updater" title="更新人"> </vxe-column>
      <vxe-column field="updateTime" title="更新时间"> </vxe-column>
      <vxe-column
        v-auth="['dict.SysDictTypeUpdate', 'dict.SysDictTypeDelete', 'dict.SysDictType', 'dict.SysDictList']"
        field="operation"
        fixed="right"
        title="操作"
        width="auto"
      >
        <template #default="{ row }">
          <el-button v-auth="'dict.SysDictType'" type="primary" link :icon="View" @click="handleItem(row)">
            查看
          </el-button>
          <el-dropdown trigger="click">
            <el-button
              v-auth="['dict.SysDictTypeUpdate', 'dict.SysDictTypeDelete', 'dict.SysDictList']"
              type="primary"
              link
              :icon="DArrowRight"
            >
              更多
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <div v-auth="'dict.SysDictList'">
                  <el-dropdown-item :icon="DataBoard" @click="toLinkDict(row)"> 字典列表 </el-dropdown-item>
                </div>
                <div v-auth="'dict.SysDictTypeUpdate'">
                  <el-dropdown-item :icon="EditPen" @click="handleUpdate(row)"> 编辑 </el-dropdown-item>
                </div>
                <div v-auth="'dict.SysDictTypeDelete'">
                  <el-dropdown-item :icon="Delete" @click="handleDelete(row)"> 删除 </el-dropdown-item>
                </div>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
      </vxe-column>
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
      <el-form ref="refSysDictTypeForm" :model="sysDictTypeForm" :rules="rulesSysDictTypeForm" label-width="100px">
        <el-form-item label="字典名称" prop="name">
          <el-input v-model="sysDictTypeForm.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="字典类型" prop="type">
          <el-input v-model="sysDictTypeForm.type" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="sysDictTypeForm.status">
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
          <el-input v-model="sysDictTypeForm.remark" :disabled="disabled" />
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
import { ResSysDictType } from '@/api/interface/sysDictType'
import { ProVxeTableInstance, ProVxeColumnProps } from '@/components/ProVxeTable/interface'
import { EditPen, CirclePlus, Delete, View, DArrowRight, DataBoard } from '@element-plus/icons-vue'
import {
  getSysDictTypeListApi,
  deleteSysDictTypeApi,
  getSysDictTypeApi,
  addSysDictTypeApi,
  updateSysDictTypeApi,
} from '@/api/modules/sysDictType'
import { FormInstance, FormRules } from 'element-plus'
import { useHandleData, useHandleSet } from '@/hooks/useHandleData'
import { useRouter } from 'vue-router'
import { getIntDictOptions } from '@/utils/dict'
import { useLoadingStore } from '@/stores/modules/loading'
import { storeToRefs } from 'pinia'
import { ProTablePaginationEnum } from '@/enums'
// 获取loading状态
const { loading } = storeToRefs(useLoadingStore())
// 路由
const router = useRouter()
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
//数据接口
const sysDictTypeForm = ref<ResSysDictType>({
  id: 0, // 编号
  name: '', // 字典名称
  type: '', // 字典类型
  status: 0, // 状态:0正常/1停用
  remark: undefined, // 备注
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
})
//表单
const refSysDictTypeForm = ref<FormInstance>()
//校验
const rulesSysDictTypeForm = reactive<FormRules>({
  name: [{ required: true, message: '字典名称不能为空', trigger: 'blur' }],
  type: [{ required: true, message: '字典类型不能为空', trigger: 'blur' }],
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
    remark: undefined, // 备注
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
  const { data } = await getSysDictTypeApi(Number(row.id))
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
  const { data } = await getSysDictTypeApi(Number(row.id))
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
    const data = sysDictTypeForm.value as unknown as ResSysDictType
    if (data.id !== 0) {
      await useHandleSet(updateSysDictTypeApi, data.id, data, '修改字典类型')
    } else {
      await useHandleData(addSysDictTypeApi, data, '添加字典类型')
    }
    resetForm(formEl)
    proTable.value?.getTableList()
  })
}

// 跳转链接
const toLinkDict = (row: ResSysDictType) => {
  router.push({
    name: 'SysDict',
    params: {
      dictTypeId: row.id,
    },
  })
}

const columns: ProVxeColumnProps[] = [
  { prop: 'name', label: '字典名称', valueType: 'input' },
  { prop: 'type', label: '字典类型', valueType: 'input' },
  { prop: 'status', label: '状态', valueType: 'select', options: statusEnum },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
