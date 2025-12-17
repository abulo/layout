<template>
  <div class="table-box">
    <ProVxeTable
      ref="proTable"
      title="职位列表"
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
        <el-button v-auth="'post.SysPostCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <vxe-column field="id" title="编号" fixed="left" width="auto"> </vxe-column>
      <vxe-column field="name" title="名称"> </vxe-column>
      <vxe-column field="sort" title="排序"> </vxe-column>
      <vxe-column field="status" title="状态">
        <template #default="{ row }">
          <DictTag type="status" :value="row.status" />
        </template>
      </vxe-column>
      <vxe-column v-auth="'post.SysPostDelete'" field="deleted" title="删除">
        <template #default="{ row }">
          <DictTag type="deleted" :value="row.deleted" />
        </template>
      </vxe-column>
      <vxe-column field="creator" title="创建人"> </vxe-column>
      <vxe-column field="createTime" title="创建时间"> </vxe-column>
      <vxe-column field="updater" title="更新人"> </vxe-column>
      <vxe-column field="updateTime" title="更新时间"> </vxe-column>
      <vxe-column
        v-auth="['post.SysPostUpdate', 'post.SysPostDelete', 'post.SysPostDrop', 'post.SysPostRecover', 'post.SysPost']"
        field="operation"
        fixed="right"
        title="操作"
        width="auto"
      >
        <template #default="{ row }">
          <el-button v-auth="'post.SysPost'" type="primary" link :icon="View" @click="handleItem(row)">
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
                  <el-dropdown-item :icon="EditPen" @click="handleUpdate(row)"> 编辑 </el-dropdown-item>
                </div>
                <div v-auth="'post.SysPostDelete'">
                  <el-dropdown-item v-if="row.deleted === 0" :icon="Delete" @click="handleDelete(row)">
                    删除
                  </el-dropdown-item>
                </div>
                <div v-auth="'post.SysPostRecover'">
                  <el-dropdown-item v-if="row.deleted === 1" :icon="Refresh" @click="handleRecover(row)">
                    恢复
                  </el-dropdown-item>
                </div>
                <div v-auth="'post.SysPostDrop'">
                  <el-dropdown-item v-if="row.deleted === 1" :icon="DeleteFilled" @click="handleDrop(row)">
                    清理
                  </el-dropdown-item>
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
import { ResSysPost } from '@/api/interface/sysPost'
import { ProVxeTableInstance, ProVxeColumnProps } from '@/components/ProVxeTable/interface'
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
import { FormInstance, FormRules } from 'element-plus'
import { getIntDictOptions } from '@/utils/dict'
import { useHandleData, useHandleSet } from '@/hooks/useHandleData'
import { HasAuth } from '@/utils/auth'
import { useLoadingStore } from '@/stores/modules/loading'
import { storeToRefs } from 'pinia'
import { ProTablePaginationEnum } from '@/enums'
// 获取loading状态
const loadingStore = useLoadingStore()
const { loading } = storeToRefs(loadingStore)
//禁用
const disabled = ref(true)
//弹出层标题
const title = ref('')
// 状态枚举
const statusEnum = getIntDictOptions('status')
//列表数据
const proTable = ref<ProVxeTableInstance>()
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
  const { data } = await getSysPostApi(Number(row.id))
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
  const { data } = await getSysPostApi(Number(row.id))
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
    const data = sysPostForm.value as unknown as ResSysPost
    if (data.id !== 0) {
      await useHandleSet(updateSysPostApi, data.id, data, '修改职位')
    } else {
      await useHandleData(addSysPostApi, data, '添加职位')
    }
    resetForm(formEl)
    proTable.value?.getTableList()
  })
}

const columns: ProVxeColumnProps[] = [
  { prop: 'name', label: '名称', valueType: 'input' },
  { prop: 'status', label: '状态', valueType: 'select', options: statusEnum },
  {
    prop: 'deleted',
    label: '删除',
    valueType: 'switch',
    fieldProps: {
      activeValue: 1,
      inactiveValue: 0,
    },
    hideInSearch: !HasAuth('post.SysPostDelete'),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
