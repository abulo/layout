<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="菜单列表"
      row-key="id"
      :columns="columns"
      :toolbar-right="['search', 'refresh', 'export', 'layout']"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="true"
      :search-col="12"
    >
      <template #toolbarLeft>
        <el-button v-auth="'menu.SysMenuCreate'" type="primary" :icon="CirclePlus" @click="handleAdd">新增</el-button>
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'menu.SysMenu'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button v-auth="['menu.SysMenuUpdate', 'menu.SysMenuDelete']" type="primary" link :icon="DArrowRight">
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-auth="'menu.SysMenuUpdate'" :icon="EditPen" @click="handleUpdate(scope.row)">
                编辑
              </el-dropdown-item>
              <el-dropdown-item
                v-if="scope.row.deleted === 0"
                v-auth="'menu.SysMenuDelete'"
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
      <el-form ref="refSysMenuForm" :model="sysMenuForm" :rules="rulesSysMenuForm" label-width="100px">
        <el-form-item label="编号" prop="id">
          <el-input v-model="sysMenuForm.id" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="名称" prop="name">
          <el-input v-model="sysMenuForm.name" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="编码" prop="code">
          <el-input v-model="sysMenuForm.code" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="类型:0 目录/1 菜单/2 按钮" prop="type">
          <el-input v-model="sysMenuForm.type" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input v-model="sysMenuForm.sort" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="上级" prop="parentId">
          <el-input v-model="sysMenuForm.parentId" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="地址" prop="path">
          <el-input v-model="sysMenuForm.path" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="图标" prop="icon">
          <el-input v-model="sysMenuForm.icon" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="组件路径" prop="component">
          <el-input v-model="sysMenuForm.component" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="组件名称" prop="componentName">
          <el-input v-model="sysMenuForm.componentName" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="隐藏:0 否/1 是" prop="hide">
          <el-input v-model="sysMenuForm.hide" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="外部地址" prop="link">
          <el-input v-model="sysMenuForm.link" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="缓存:0否/1 是" prop="cache">
          <el-input v-model="sysMenuForm.cache" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="sysMenuForm.remark" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="激活地址" prop="active">
          <el-input v-model="sysMenuForm.active" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="全屏:1 开/0 关" prop="full">
          <el-input v-model="sysMenuForm.full" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="重定向" prop="redirect">
          <el-input v-model="sysMenuForm.redirect" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="状态:0正常/1停用" prop="status">
          <el-input v-model="sysMenuForm.status" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="创建人" prop="creator">
          <el-input v-model="sysMenuForm.creator" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="创建时间" prop="createTime">
          <el-input v-model="sysMenuForm.createTime" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="更新人" prop="updater">
          <el-input v-model="sysMenuForm.updater" :disabled="disabled" />
        </el-form-item>
        <el-form-item label="更新时间" prop="updateTime">
          <el-input v-model="sysMenuForm.updateTime" :disabled="disabled" />
        </el-form-item>
      </el-form>
      <template v-if="!disabled" #footer>
        <span class="dialog-footer">
          <el-button @click="resetForm(refSysMenuForm)">取消</el-button>
          <el-button type="primary" :loading="loading" @click="submitForm(refSysMenuForm)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx">
defineOptions({ name: 'SysMenu' })
import type { ResSysMenu } from '@/api/interface/sysMenu'
import type { ProTableInstance, ColumnProps } from '@/components/ProTable/interface'
import { EditPen, CirclePlus, Delete, View, DArrowRight } from '@element-plus/icons-vue'
import {
  getSysMenuListApi,
  deleteSysMenuApi,
  getSysMenuApi,
  addSysMenuApi,
  updateSysMenuApi,
} from '@/api/modules/sysMenu'
import type { FormInstance, FormRules } from 'element-plus'
import { useHandleData, useHandleSet } from '@/hooks/useHandleData'
import { HasAuth } from '@/utils/auth'
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
const sysMenuForm = ref<ResSysMenu>({
  id: 0, // 编号
  name: '', // 名称
  code: undefined, // 编码
  type: 0, // 类型:0 目录/1 菜单/2 按钮
  sort: 0, // 排序
  parentId: 0, // 上级
  path: undefined, // 地址
  icon: undefined, // 图标
  component: undefined, // 组件路径
  componentName: undefined, // 组件名称
  hide: 0, // 隐藏:0 否/1 是
  link: undefined, // 外部地址
  cache: 0, // 缓存:0否/1 是
  remark: undefined, // 备注
  active: undefined, // 激活地址
  full: 0, // 全屏:1 开/0 关
  redirect: undefined, // 重定向
  status: 0, // 状态:0正常/1停用
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
})
//表单
const refSysMenuForm = ref<FormInstance>()
//校验
const rulesSysMenuForm = reactive<FormRules>({
  id: [{ required: true, message: '编号不能为空', trigger: 'blur' }],
  name: [{ required: true, message: '名称不能为空', trigger: 'blur' }],
  type: [{ required: true, message: '类型:0 目录/1 菜单/2 按钮不能为空', trigger: 'blur' }],
  sort: [{ required: true, message: '排序不能为空', trigger: 'blur' }],
  parentId: [{ required: true, message: '上级不能为空', trigger: 'blur' }],
  hide: [{ required: true, message: '隐藏:0 否/1 是不能为空', trigger: 'blur' }],
  cache: [{ required: true, message: '缓存:0否/1 是不能为空', trigger: 'blur' }],
  full: [{ required: true, message: '全屏:1 开/0 关不能为空', trigger: 'blur' }],
  status: [{ required: true, message: '状态:0正常/1停用不能为空', trigger: 'blur' }],
})

/**
 * 获取表格数据列表
 * @param params 查询参数对象
 * @returns 返回列表数据的Promise对象
 */
const getTableList = (params: any) => {
  // 深拷贝参数对象，避免修改原始参数
  let newParams = JSON.parse(JSON.stringify(params))
  return getSysMenuListApi(newParams)
}

/**
 * 重置数据
 * 该函数用于清空或初始化
 * @returns {void} 无返回值
 */
const resetSysMenu = () => {
  // 设置sysMenuForm为默认值
  Object.assign(sysMenuForm.value, {
    id: 0, // 编号
    name: '', // 名称
    code: undefined, // 编码
    type: 0, // 类型:0 目录/1 菜单/2 按钮
    sort: 0, // 排序
    parentId: 0, // 上级
    path: undefined, // 地址
    icon: undefined, // 图标
    component: undefined, // 组件路径
    componentName: undefined, // 组件名称
    hide: 0, // 隐藏:0 否/1 是
    link: undefined, // 外部地址
    cache: 0, // 缓存:0否/1 是
    remark: undefined, // 备注
    active: undefined, // 激活地址
    full: 0, // 全屏:1 开/0 关
    redirect: undefined, // 重定向
    status: 0, // 状态:0正常/1停用
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
  resetSysMenu()
  disabled.value = true
}
/**
 * 处理新增操作
 * 该函数用于初始化对话框的状态
 */
const handleAdd = () => {
  title.value = '新增菜单'
  dialogVisible.value = true
  reset()
  disabled.value = false
}
/**
 * 处理更新操作
 * @param row - 要更新数据对象
 */
const handleUpdate = async (row: ResSysMenu) => {
  title.value = '编辑菜单'
  dialogVisible.value = true
  reset()
  const data = await getSysMenuApi(Number(row.id))
  sysMenuForm.value = data
  disabled.value = false
}
/**
 * 处理查看操作
 * @param row - 要查看数据对象
 */
const handleItem = async (row: ResSysMenu) => {
  title.value = '查看菜单'
  dialogVisible.value = true
  reset()
  const data = await getSysMenuApi(Number(row.id))
  sysMenuForm.value = data
  disabled.value = true
}
/**
 * 处理删除操作
 * @param row - 要删除数据对象
 */
const handleDelete = async (row: ResSysMenu) => {
  await useHandleData(deleteSysMenuApi, Number(row.id), '删除菜单')
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
    const data = sysMenuForm.value as unknown as ResSysMenu
    if (data.id !== 0) {
      await useHandleSet(updateSysMenuApi, data.id, data, '修改菜单')
    } else {
      await useHandleData(addSysMenuApi, data, '添加菜单')
    }
    resetForm(formEl)
    loading.value = false
    proTable.value?.getTableList()
  })
}

const columns: ColumnProps<ResSysMenu>[] = [
  { prop: 'id', label: '编号' },
  { prop: 'name', label: '名称' },
  { prop: 'code', label: '编码' },
  { prop: 'type', label: '类型:0 目录/1 菜单/2 按钮', search: { el: 'input' } },
  { prop: 'sort', label: '排序' },
  { prop: 'parentId', label: '上级' },
  { prop: 'path', label: '地址' },
  { prop: 'icon', label: '图标' },
  { prop: 'component', label: '组件路径' },
  { prop: 'componentName', label: '组件名称' },
  { prop: 'hide', label: '隐藏:0 否/1 是' },
  { prop: 'link', label: '外部地址' },
  { prop: 'cache', label: '缓存:0否/1 是' },
  { prop: 'remark', label: '备注' },
  { prop: 'active', label: '激活地址' },
  { prop: 'full', label: '全屏:1 开/0 关' },
  { prop: 'redirect', label: '重定向' },
  { prop: 'status', label: '状态:0正常/1停用', search: { el: 'input' } },
  { prop: 'creator', label: '创建人' },
  { prop: 'createTime', label: '创建时间' },
  { prop: 'updater', label: '更新人' },
  { prop: 'updateTime', label: '更新时间' },

  {
    prop: 'operation',
    label: '操作',
    width: 150,
    fixed: 'right',
    isShow: HasAuth('menu.SysMenuUpdate', 'menu.SysMenuDelete', 'menu.SysMenu'),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
