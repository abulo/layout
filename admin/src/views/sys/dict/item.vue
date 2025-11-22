<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="字典列表"
      row-key="id"
      :columns="columns"
      :toolbar-right="['search', 'refresh', 'export', 'layout']"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12"
    >
      <template #toolbarLeft>
        <el-button v-auth="'dict.SysDictCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
        <el-button type="primary" :icon="Remove" @click="closeTab">关闭</el-button>
      </template>
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'dict.SysDict'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button v-auth="['dict.SysDictUpdate', 'dict.SysDictDelete']" type="primary" link :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <div v-auth="'dict.SysDictUpdate'">
                <el-dropdown-item :icon="EditPen" @click="handleUpdate(scope.row)"> 编辑 </el-dropdown-item>
              </div>
              <div v-auth="'dict.SysDictDelete'">
                <el-dropdown-item :icon="Delete" @click="handleDelete(scope.row)"> 删除 </el-dropdown-item>
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
      <el-form ref="refSysDictForm" :model="sysDictForm" :rules="rulesSysDictForm" label-width="100px">
        <el-form-item label="标签" prop="label">
          <el-input v-model="sysDictForm.label" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="键值" prop="value">
          <el-input v-model="sysDictForm.value" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="sysDictForm.sort" controls-position="right" :min="0" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="颜色类型" prop="colorType">
          <el-select v-model="sysDictForm.colorType" clearable :disabled="disabled">
            <el-option
              v-for="item in colorTypeOptions"
              :key="item.value"
              :label="item.label + '(' + item.value + ')'"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="CSS样式" prop="cssClass">
          <el-input v-model="sysDictForm.cssClass" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="sysDictForm.status">
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
          <el-input v-model="sysDictForm.remark" :disabled="disabled" />
        </el-form-item>
      </el-form>
      <template v-if="!disabled" #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSysDictForm)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSysDictForm)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx">
defineOptions({ name: 'SysDict' })
import type { ResSysDict } from '@/api/interface/sysDict'
import type { ProTableInstance, ColumnProps } from '@/components/ProTable/interface'
import { EditPen, CirclePlus, Delete, View, DArrowRight, Remove } from '@element-plus/icons-vue'
import {
  getSysDictListApi,
  deleteSysDictApi,
  getSysDictApi,
  addSysDictApi,
  updateSysDictApi,
} from '@/api/modules/sysDict'
import type { FormInstance, FormRules } from 'element-plus'
import { useHandleData, useHandleSet } from '@/hooks/useHandleData'
import { HasAuth } from '@/utils/auth'
import { useRoute } from 'vue-router'
import { useTabsStore } from '@/stores/modules/tabs'
import { useKeepAliveStore } from '@/stores/modules/keepAlive'
import { getSysDictTypeListSimpleApi } from '@/api/modules/sysDictType'
import { getIntDictOptions } from '@/utils/dict'
import type { ResSysDictType } from '@/api/interface/sysDictType'
import { useLoadingStore } from '@/stores/modules/loading'
import { storeToRefs } from 'pinia'
// 获取loading状态
const { loading } = storeToRefs(useLoadingStore())
// 获取当前路由信息
const route = useRoute()
// 获取标签页信息
const tabStore = useTabsStore()
// 缓存信息
const keepAliveStore = useKeepAliveStore()
// 字典类型id
const dictTypeId = ref(Number(route.params.dictTypeId))
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
// 字典类型所有数据
const sysDictTypeList = ref<ResSysDictType[]>([])
//数据接口
const sysDictForm = ref<ResSysDict>({
  id: 0, // 编号
  sort: 0, // 排序
  label: '', // 标签
  value: '', // 键值
  dictTypeId: dictTypeId.value, // 字典类型
  colorType: undefined, // 颜色类型
  cssClass: undefined, // CSS样式
  status: 0, // 状态:0正常/1停用
  remark: undefined, // 备注
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
})
//表单
const refSysDictForm = ref<FormInstance>()
//校验
const rulesSysDictForm = reactive<FormRules>({
  label: [{ required: true, message: '标签不能为空', trigger: 'blur' }],
  value: [{ required: true, message: '键值不能为空', trigger: 'blur' }],
  dictTypeId: [{ required: true, message: '字典类型不能为空', trigger: 'blur' }],
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
  if (newParams.dictTypeId) {
    dictTypeId.value = Number(newParams.dictTypeId)
    sysDictForm.value.dictTypeId = dictTypeId.value
  }
  return getSysDictListApi(newParams)
}

/**
 * 重置数据
 * 该函数用于清空或初始化
 * @returns {void} 无返回值
 */
const resetSysDict = () => {
  // 设置sysDictForm为默认值
  Object.assign(sysDictForm.value, {
    id: 0, // 编号
    sort: 0, // 排序
    label: '', // 标签
    value: '', // 键值
    dictTypeId: dictTypeId.value, // 字典类型
    colorType: undefined, // 颜色类型
    cssClass: undefined, // CSS样式
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
  resetSysDict()
  disabled.value = true
}
/**
 * 处理新增操作
 * 该函数用于初始化对话框的状态
 */
const handleAdd = () => {
  title.value = '新增字典'
  dialogVisible.value = true
  reset()
  disabled.value = false
}
/**
 * 处理更新操作
 * @param row - 要更新数据对象
 */
const handleUpdate = async (row: ResSysDict) => {
  title.value = '编辑字典'
  dialogVisible.value = true
  reset()
  const { data } = await getSysDictApi(Number(row.id))
  sysDictForm.value = data
  disabled.value = false
}
/**
 * 处理查看操作
 * @param row - 要查看数据对象
 */
const handleItem = async (row: ResSysDict) => {
  title.value = '查看字典'
  dialogVisible.value = true
  reset()
  const { data } = await getSysDictApi(Number(row.id))
  sysDictForm.value = data
  disabled.value = true
}
/**
 * 处理删除操作
 * @param row - 要删除数据对象
 */
const handleDelete = async (row: ResSysDict) => {
  await useHandleData(deleteSysDictApi, Number(row.id), '删除字典')
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
    const data = sysDictForm.value as unknown as ResSysDict
    if (data.id !== 0) {
      await useHandleSet(updateSysDictApi, data.id, data, '修改字典')
    } else {
      await useHandleData(addSysDictApi, data, '添加字典')
    }
    resetForm(formEl)
    proTable.value?.getTableList()
  })
}

// 关闭当前页
const closeTab = () => {
  if (route.meta.isAffix) return
  tabStore.removeTabs(route.fullPath)
  keepAliveStore.removeKeepAliveName(route.name as string)
}

// 数据标签回显样式
const colorTypeOptions = ref([
  {
    value: 'default',
    label: '默认',
  },
  {
    value: 'primary',
    label: '主要',
  },
  {
    value: 'success',
    label: '成功',
  },
  {
    value: 'info',
    label: '信息',
  },
  {
    value: 'warning',
    label: '警告',
  },
  {
    value: 'danger',
    label: '危险',
  },
])

// 获取字典类型列表
onMounted(async () => {
  const { data } = await getSysDictTypeListSimpleApi()
  sysDictTypeList.value = data
})

const columns: ColumnProps<ResSysDict>[] = [
  { prop: 'id', label: '编号', fixed: 'left' },
  { prop: 'label', label: '标签' },
  { prop: 'value', label: '键值' },
  { prop: 'sort', label: '排序' },
  {
    prop: 'dictTypeId',
    label: '字典类型',
    enum: sysDictTypeList,
    tag: true,
    fieldNames: { label: 'name', value: 'id' },
    search: { el: 'select', span: 2, defaultValue: dictTypeId.value },
  },
  { prop: 'colorType', label: '颜色类型' },
  { prop: 'cssClass', label: 'CSS样式' },
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
    isShow: HasAuth('dict.SysDictUpdate', 'dict.SysDictDelete', 'dict.SysDict'),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
