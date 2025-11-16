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
        <el-button v-auth="'post.SysPostCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
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
        <el-button v-auth="'post.SysPost'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['post.SysPostUpdate', 'post.SysPostDelete', 'post.SysPostRecover', 'post.SysPostDrop']"
            type="primary"
            link
            :icon="DArrowRight"
          >
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <div v-auth="'post.SysPostUpdate'">
                <el-dropdown-item :icon="EditPen" @click="handleUpdate(scope.row)"> 编辑 </el-dropdown-item>
              </div>
              <div v-auth="'post.SysPostDelete'">
                <el-dropdown-item v-if="scope.row.deleted === 0" :icon="Delete" @click="handleDelete(scope.row)">
                  删除
                </el-dropdown-item>
              </div>
              <div v-auth="'post.SysPostRecover'">
                <el-dropdown-item v-if="scope.row.deleted === 1" :icon="Refresh" @click="handleRecover(scope.row)">
                  恢复
                </el-dropdown-item>
              </div>
              <div v-auth="'post.SysPostDrop'">
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
    >
      <el-form ref="refSysPostForm" :model="sysPostForm" :rules="rulesSysPostForm" label-width="100px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="sysPostForm.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="sysPostForm.sort" controls-position="right" :min="0" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="sysPostForm.status">
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
import type { ProTableInstance, ColumnProps, SearchProps } from '@/components/ProTable/interface'
import { EditPen, CirclePlus, Delete, Refresh, DeleteFilled, View, DArrowRight } from '@element-plus/icons-vue'
import {
  getSysPostListApi,
  deleteSysPostApi,
  dropSysPostApi,
  recoverSysPostApi,
  getSysPostApi,
  addSysPostApi,
  updateSysPostApi,
} from '@/api/modules/sysPost'
import type { FormInstance, FormRules } from 'element-plus'
import { getIntDictOptions } from '@/utils/dict'
import { useHandleData, useHandleSet } from '@/hooks/useHandleData'
import { HasAuth } from '@/utils/auth'
//加载
const loading = ref(false)
//禁用
const disabled = ref(true)
//弹出层标题
const title = ref('')
// 状态枚举
const statusEnum = getIntDictOptions('status')
//列表数据
const proTable = ref<ProTableInstance>()
//显示弹出层
const dialogVisible = ref(false)
//数据接口
const sysPostForm = ref<ResSysPost>({
  id: 0, // 编号
  name: '', // 名称
  sort: 0, // 排序
  status: 0, // 状态:0正常/1停用
  deleted: 0, // 删除:0否/1是
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
})
//表单
const refSysPostForm = ref<FormInstance>()
//校验
const rulesSysPostForm = reactive<FormRules>({
  name: [{ required: true, message: '名称不能为空', trigger: 'blur' }],
  sort: [{ required: true, message: '排序不能为空', trigger: 'blur' }],
  status: [{ required: true, message: '状态不能为空', trigger: 'blur' }],
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
    id: 0, // 编号
    name: '', // 名称
    sort: 0, // 排序
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
//删除状态
const deletedEnum = getIntDictOptions('deleted')
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasAuth('post.SysPostDelete')
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

const columns: ColumnProps<ResSysPost>[] = [
  { prop: 'id', label: '编号', fixed: 'left' },
  { prop: 'name', label: '名称', search: { el: 'input', span: 2 } },
  { prop: 'sort', label: '排序' },
  { prop: 'status', label: '状态', tag: true, enum: statusEnum, search: { el: 'select', span: 2 } },
  { prop: 'deleted', label: '删除', tag: true, enum: deletedEnum, search: deleteSearch },
  { prop: 'tenantId', label: '租户' },
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
      'post.SysPostUpdate',
      'post.SysPostDelete',
      'post.SysPostDrop',
      'post.SysPostRecover',
      'post.SysPost'
    ),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
