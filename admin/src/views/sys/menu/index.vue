<template>
  <div class="table-box">
    <ProVxeTable
      ref="proTable"
      title="菜单列表"
      :columns="columns"
      :toolbar-right="['search', 'refresh']"
      :request-api="getTableList"
      :data-callback="menuHandleTree"
      :request-auto="true"
      :show-number="3"
      :border="true"
      :column-config="{ resizable: true, isCurrent: true, isHover: true }"
      :row-config="{ isCurrent: true, isHover: true }"
      :virtual-y-config="{ enabled: true, gt: 0 }"
      :tree-config="{}"
      :pagination="ProTablePaginationEnum.NONE"
      @data-rendered="dataRenderedEvent"
    >
      <template #toolbarLeft>
        <el-button v-auth="'menu.SysMenuCreate'" type="primary" :icon="CirclePlus" @click="handleAdd()">
          新增
        </el-button>
        <el-button type="primary" :icon="Sort" :loading="loadingStore.loading" @click="handleExpandAll">
          展开/折叠
        </el-button>
      </template>
      <vxe-column field="id" title="编号" fixed="left" width="auto"></vxe-column>
      <vxe-column field="name" title="名称" tree-node width="auto"></vxe-column>
      <vxe-column field="code" title="编码" width="auto"></vxe-column>
      <vxe-column field="type" title="类型" align="center">
        <template #default="{ row }">
          <DictTag type="menu.type" :value="row.type" />
        </template>
      </vxe-column>
      <vxe-column field="sort" title="排序" align="center"></vxe-column>
      <vxe-column field="path" title="地址" width="auto"></vxe-column>
      <vxe-column field="icon" title="图标" align="center">
        <template #default="{ row }">
          <Icon :icon="row.icon" :custom-key="row.id" :size="18" class="el-icon" />
        </template>
      </vxe-column>
      <vxe-column field="component" title="组件路径" width="auto"></vxe-column>
      <vxe-column field="componentName" title="组件名称"></vxe-column>
      <vxe-column field="status" title="状态" align="center">
        <template #default="{ row }">
          <DictTag type="status" :value="row.status" />
        </template>
      </vxe-column>
      <vxe-column
        v-auth="['menu.SysMenuUpdate', 'menu.SysMenuDelete', 'menu.SysMenuCreate', 'menu.SysMenu']"
        field="operation"
        fixed="right"
        title="操作"
        align="center"
        width="auto"
      >
        <template #default="{ row }">
          <el-button v-auth="'menu.SysMenu'" type="primary" link :icon="View" @click="handleItem(row)">
            查看
          </el-button>
          <el-dropdown trigger="click">
            <el-button
              v-auth="['menu.SysMenuUpdate', 'menu.SysMenuDelete', 'menu.SysMenuCreate']"
              type="primary"
              link
              :icon="DArrowRight"
            >
              更多
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <div v-auth="'menu.SysMenuCreate'">
                  <el-dropdown-item :icon="CirclePlus" @click="handleAdd(row)"> 新增 </el-dropdown-item>
                </div>
                <div v-auth="'menu.SysMenuUpdate'">
                  <el-dropdown-item :icon="EditPen" @click="handleUpdate(row)"> 编辑 </el-dropdown-item>
                </div>
                <div v-auth="'menu.SysMenuDelete'">
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
      <el-form ref="refSysMenuForm" :model="sysMenuForm" :rules="rulesSysMenuForm" label-width="100px">
        <el-row :gutter="1">
          <el-col :span="12">
            <el-form-item label="名称" prop="name">
              <el-input v-model="sysMenuForm.name" :disabled="disabled" />
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="上级菜单" prop="parentId">
              <el-tree-select
                v-model="sysMenuForm.parentId"
                :data="menuEnum"
                value-key="id"
                node-key="id"
                placeholder="请选择"
                check-strictly
                :disabled="disabled"
                :render-after-expand="false"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="1">
          <el-col :span="12">
            <el-form-item label="类型" prop="type">
              <el-radio-group v-model="sysMenuForm.type">
                <el-radio-button
                  v-for="dict in menuTypeEnum"
                  :key="dict.value"
                  :value="dict.value"
                  :disabled="disabled"
                >
                  {{ dict.label }}
                </el-radio-button>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="排序" prop="sort">
              <el-input-number v-model="sysMenuForm.sort" controls-position="right" :min="0" :disabled="disabled" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="1">
          <el-col :span="12">
            <el-form-item v-if="sysMenuForm.type !== 0" label="编码" prop="code">
              <el-input v-model="sysMenuForm.code" :disabled="disabled" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="1">
          <el-col :span="12">
            <el-form-item v-if="sysMenuForm.type !== 2" label="地址" prop="path">
              <el-input v-model="sysMenuForm.path" :disabled="disabled" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item v-if="sysMenuForm.type !== 2" label="图标" prop="icon">
              <SelectIcon v-model="sysMenuForm.icon" title="请选择图标" placeholder="搜索图标" :show-icon-name="true" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="1">
          <el-col :span="12">
            <el-form-item v-if="sysMenuForm.type === 1" label="组件路径" prop="component">
              <el-input v-model="sysMenuForm.component" :disabled="disabled" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item v-if="sysMenuForm.type === 1" label="组件名称" prop="componentName">
              <el-input v-model="sysMenuForm.componentName" :disabled="disabled" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="1">
          <el-col :span="12">
            <el-form-item v-if="sysMenuForm.type !== 2" label="重定向" prop="redirect">
              <el-input v-model="sysMenuForm.redirect" :disabled="disabled" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item v-if="sysMenuForm.type === 1" label="激活地址" prop="active">
              <el-input v-model="sysMenuForm.active" :disabled="disabled" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="1">
          <el-col :span="8">
            <el-form-item v-if="sysMenuForm.type !== 2" label="隐藏" prop="hide">
              <el-radio-group v-model="sysMenuForm.hide">
                <el-radio-button
                  v-for="dict in menuHideEnum"
                  :key="dict.value"
                  :value="dict.value"
                  :disabled="disabled"
                >
                  {{ dict.label }}
                </el-radio-button>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item v-if="sysMenuForm.type === 1" label="缓存" prop="cache">
              <el-radio-group v-model="sysMenuForm.cache">
                <el-radio-button
                  v-for="dict in menuCacheEnum"
                  :key="dict.value"
                  :value="dict.value"
                  :disabled="disabled"
                >
                  {{ dict.label }}
                </el-radio-button>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item v-if="sysMenuForm.type === 1" label="全屏" prop="full">
              <el-radio-group v-model="sysMenuForm.full">
                <el-radio-button
                  v-for="dict in menuFullEnum"
                  :key="dict.value"
                  :value="dict.value"
                  :disabled="disabled"
                >
                  {{ dict.label }}
                </el-radio-button>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="sysMenuForm.status">
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
        <el-form-item v-if="sysMenuForm.type !== 2" label="备注" prop="remark">
          <el-input v-model="sysMenuForm.remark" :disabled="disabled" />
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
import { ResSysMenu } from '@/api/interface/sysMenu'
import { ProVxeTableInstance, ProVxeColumnProps } from '@/components/ProVxeTable/interface'
import { EditPen, CirclePlus, Delete, View, DArrowRight, Sort } from '@element-plus/icons-vue'
import {
  getSysMenuListApi,
  deleteSysMenuApi,
  getSysMenuApi,
  addSysMenuApi,
  updateSysMenuApi,
  getSysMenuListSimpleApi,
} from '@/api/modules/sysMenu'
import { FormInstance, FormRules } from 'element-plus'
import { useHandleData, useHandleSet } from '@/hooks/useHandleData'
import { handleTree } from '@pureadmin/utils'
import { getIntDictOptions } from '@/utils/dict'
import { useLoadingStore } from '@/stores/modules/loading'
import { storeToRefs } from 'pinia'
import { ProTablePaginationEnum } from '@/enums'
import { useTimeoutFn } from '@vueuse/core'
import { HandleEnumList } from '@/utils'
import { VxeTableEvents } from 'vxe-table'

const proTable = ref<ProVxeTableInstance>()
// 获取loading状态
const loadingStore = useLoadingStore()
const { loading } = storeToRefs(loadingStore)
//禁用
const disabled = ref(true)
//弹出层标题
const title = ref('')
//列表数据
const dialogVisible = ref(false)
//是否展开，默认全部折叠
const isExpandAll = ref(false)
//菜单类别
const menuTypeEnum = getIntDictOptions('menu.type')
//菜单隐藏
const menuHideEnum = getIntDictOptions('menu.hide')
// 菜单缓存
const menuCacheEnum = getIntDictOptions('menu.cache')
//菜单全屏
const menuFullEnum = getIntDictOptions('menu.full')
// 状态枚举
const statusEnum = getIntDictOptions('status')
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
//下拉菜单选项
const menuOptions = ref<ResSysMenu[]>([])
//表单
const refSysMenuForm = ref<FormInstance>()
//校验
const rulesSysMenuForm = reactive<FormRules>({
  name: [{ required: true, message: '名称不能为空', trigger: 'blur' }],
  type: [{ required: true, message: '类型不能为空', trigger: 'blur' }],
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
  resetSysMenu()
  disabled.value = true
}
/**
 * 处理新增操作
 * 该函数用于初始化对话框的状态
 */
const handleAdd = (row?: ResSysMenu) => {
  title.value = '新增菜单'
  dialogVisible.value = true
  reset()
  disabled.value = false
  getMenuOptions()
  if (row != null && row.id) {
    sysMenuForm.value.parentId = row.id
  }
}
/**
 * 处理更新操作
 * @param row - 要更新数据对象
 */
const handleUpdate = async (row: ResSysMenu) => {
  title.value = '编辑菜单'
  dialogVisible.value = true
  reset()
  getMenuOptions()
  const { data } = await getSysMenuApi(Number(row.id))
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
  getMenuOptions()
  const { data } = await getSysMenuApi(Number(row.id))
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
    const data = sysMenuForm.value as unknown as ResSysMenu
    delete data.children
    if (data.id !== 0) {
      await useHandleSet(updateSysMenuApi, data.id, data, '修改菜单')
    } else {
      await useHandleData(addSysMenuApi, data, '添加菜单')
    }
    resetForm(formEl)
    proTable.value?.getTableList()
  })
}

// 设置展开合并
const handleExpandAll = () => {
  loadingStore.setLoading(true)
  isExpandAll.value = !isExpandAll.value
  if (isExpandAll.value) {
    proTable.value?.element?.setAllTreeExpand(true)
  } else {
    proTable.value?.element?.clearTreeExpand()
  }
  useTimeoutFn(() => {
    loadingStore.setLoading(false)
  }, 500)
}

const dataRenderedEvent: VxeTableEvents.DataRendered<any> = () => {
  if (isExpandAll.value) {
    proTable.value?.element?.setAllTreeExpand(true)
  } else {
    proTable.value?.element?.clearTreeExpand()
  }
}

// 获取菜单选项
const getMenuOptions = async () => {
  const { data } = await getSysMenuListSimpleApi()
  menuOptions.value = [
    {
      id: 0,
      name: '主类目',
      code: undefined,
      type: 0,
      sort: 1,
      parentId: 0,
      path: undefined,
      icon: undefined,
      component: undefined,
      componentName: undefined,
      hide: 0,
      link: undefined,
      cache: 0,
      remark: undefined,
      active: undefined,
      full: 0,
      redirect: undefined,
      status: 0,
      creator: undefined,
      createTime: undefined,
      updater: undefined,
      updateTime: undefined,
      // children: handleTree(data),
    },
    ...data,
  ] as unknown as ResSysMenu[]
}

const menuEnum = computed(() => {
  if (Array.isArray(menuOptions.value)) {
    return HandleEnumList(handleTree(menuOptions.value), {
      label: 'name',
      value: 'id',
      children: 'children',
    })
  }
  return []
})

const menuHandleTree = (data: ResSysMenu[]) => {
  return handleTree(data)
}

// 定义列配置项
const columns: ProVxeColumnProps[] = []
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
