<template>
  <div class="table-box">
    <ProVxeTable
      ref="proTable"
      title="操作日志列表"
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
      <template #toolbarLeft> </template>
      <vxe-column field="id" title="编号" fixed="left" width="auto"> </vxe-column>
      <vxe-column field="username" title="用户名"> </vxe-column>
      <vxe-column field="module" title="模块名称" width="auto"> </vxe-column>
      <vxe-column field="method" title="请求方法">
        <template #default="{ row }">
          <DictTag type="operate.method" :value="row.method" />
        </template>
      </vxe-column>
      <vxe-column field="url" title="请求地址" width="auto" show-overflow> </vxe-column>
      <vxe-column field="ip" title="IP"> </vxe-column>
      <vxe-column field="startTime" title="开始时间"> </vxe-column>
      <vxe-column field="duration" title="执行时长"> </vxe-column>
      <vxe-column field="channel" title="渠道"> </vxe-column>
      <vxe-column field="result" title="结果">
        <template #default="{ row }">
          <DictTag type="operate.result" :value="row.result" />
        </template>
      </vxe-column>
      <vxe-column v-auth="'logger.SysLoggerOperateDelete'" field="deleted" title="删除">
        <template #default="{ row }">
          <DictTag type="deleted" :value="row.deleted" />
        </template>
      </vxe-column>
      <vxe-column
        v-auth="[
          'logger.SysLoggerOperateDelete',
          'logger.SysLoggerOperateDrop',
          'logger.SysLoggerOperateRecover',
          'logger.SysLoggerOperate',
        ]"
        fixed="right"
        width="auto"
        field="operation"
        title="操作"
      >
        <template #default="{ row }">
          <!-- 菜单操作 -->
          <el-button v-auth="'logger.SysLoggerOperate'" type="primary" link :icon="View" @click="handleItem(row)">
            查看
          </el-button>
          <el-dropdown trigger="click">
            <el-button
              v-auth="[
                'logger.SysLoggerOperateDelete',
                'logger.SysLoggerOperateRecover',
                'logger.SysLoggerOperateDrop',
              ]"
              type="primary"
              link
              :icon="DArrowRight"
            >
              更多
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <div v-auth="'logger.SysLoggerOperateDelete'">
                  <el-dropdown-item v-if="row.deleted === 0" :icon="Delete" @click="handleDelete(row)">
                    删除
                  </el-dropdown-item>
                </div>
                <div v-auth="'logger.SysLoggerOperateRecover'">
                  <el-dropdown-item v-if="row.deleted === 1" :icon="Refresh" @click="handleRecover(row)">
                    恢复
                  </el-dropdown-item>
                </div>
                <div v-auth="'logger.SysLoggerOperateDrop'">
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
      <el-descriptions :column="1" border>
        <el-descriptions-item label="编号"> {{ sysLoggerOperateForm.id }} </el-descriptions-item>
        <el-descriptions-item label="姓名"> {{ sysLoggerOperateForm.name }} </el-descriptions-item>
        <el-descriptions-item label="用户名"> {{ sysLoggerOperateForm.username }} </el-descriptions-item>
        <el-descriptions-item label="用户编号"> {{ sysLoggerOperateForm.userId }} </el-descriptions-item>
        <el-descriptions-item label="模块名称"> {{ sysLoggerOperateForm.module }} </el-descriptions-item>
        <el-descriptions-item label="请求方法"> {{ sysLoggerOperateForm.method }} </el-descriptions-item>
        <el-descriptions-item label="请求地址"> {{ sysLoggerOperateForm.url }} </el-descriptions-item>
        <el-descriptions-item label="IP"> {{ sysLoggerOperateForm.ip }} </el-descriptions-item>
        <el-descriptions-item label="UA"> {{ sysLoggerOperateForm.ua }} </el-descriptions-item>
        <el-descriptions-item label="方法名"> {{ sysLoggerOperateForm.goMethod }} </el-descriptions-item>
        <el-descriptions-item label="方法参数"> {{ sysLoggerOperateForm.goMethodArgs }} </el-descriptions-item>
        <el-descriptions-item label="开始时间"> {{ sysLoggerOperateForm.startTime }} </el-descriptions-item>
        <el-descriptions-item label="执行时长"> {{ sysLoggerOperateForm.duration }} </el-descriptions-item>
        <el-descriptions-item label="渠道"> {{ sysLoggerOperateForm.channel }} </el-descriptions-item>
        <el-descriptions-item label="结果"> {{ sysLoggerOperateForm.result }} </el-descriptions-item>
        <el-descriptions-item label="结果信息"> {{ sysLoggerOperateForm.resultMsg }} </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx">
defineOptions({ name: 'SysLoggerOperate' })
import { ResSysLoggerOperate } from '@/api/interface/sysLoggerOperate'
import { ProVxeTableInstance, ProVxeColumnProps } from '@/components/ProVxeTable/interface'
import { Delete, Refresh, DeleteFilled, View, DArrowRight } from '@element-plus/icons-vue'
import {
  getSysLoggerOperateListApi,
  deleteSysLoggerOperateApi,
  dropSysLoggerOperateApi,
  recoverSysLoggerOperateApi,
  getSysLoggerOperateApi,
} from '@/api/modules/sysLoggerOperate'
import { getIntDictOptions, getStrDictOptions } from '@/utils/dict'
import { DictTag } from '@/components/DictTag'
import { useHandleData } from '@/hooks/useHandleData'
import { ProTablePaginationEnum } from '@/enums'
import { HasAuth } from '@/utils/auth'
// 获取loading状态
//禁用
const disabled = ref(true)
//弹出层标题
const title = ref('')
//列表数据
const proTable = ref<ProVxeTableInstance>()
//显示弹出层
const dialogVisible = ref(false)
//数据接口
const sysLoggerOperateForm = ref<ResSysLoggerOperate>({
  id: 0, // 编号
  name: undefined, // 姓名
  username: '', // 用户名
  userId: 0, // 用户编号
  module: undefined, // 模块名称
  method: undefined, // 请求方法
  url: undefined, // 请求地址
  ip: undefined, // IP
  ua: undefined, // UA
  goMethod: undefined, // 方法名
  goMethodArgs: undefined, // 方法参数
  startTime: undefined, // 开始时间
  duration: undefined, // 执行时长
  channel: undefined, // 渠道
  result: undefined, // 结果
  resultMsg: undefined, // 结果信息
  deleted: 0, // 删除
  tenantId: 0, // 租户
  creator: undefined, // 创建人
  createTime: undefined, // 创建时间
  updater: undefined, // 更新人
  updateTime: undefined, // 更新时间
})
/**
 * 获取表格数据列表
 * @param params 查询参数对象
 * @returns 返回列表数据的Promise对象
 */
const getTableList = (params: any) => {
  // 深拷贝参数对象，避免修改原始参数
  let newParams = JSON.parse(JSON.stringify(params))
  return getSysLoggerOperateListApi(newParams)
}

/**
 * 重置数据
 * 该函数用于清空或初始化
 * @returns {void} 无返回值
 */
const resetSysLoggerOperate = () => {
  // 设置sysLoggerOperateForm为默认值
  Object.assign(sysLoggerOperateForm.value, {
    id: 0, // 编号
    name: undefined, // 姓名
    username: '', // 用户名
    userId: 0, // 用户编号
    module: undefined, // 模块名称
    method: undefined, // 请求方法
    url: undefined, // 请求地址
    ip: undefined, // IP
    ua: undefined, // UA
    goMethod: undefined, // 方法名
    goMethodArgs: undefined, // 方法参数
    startTime: undefined, // 开始时间
    duration: undefined, // 执行时长
    channel: undefined, // 渠道
    result: undefined, // 结果
    resultMsg: undefined, // 结果信息
    deleted: 0, // 删除
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
  // loading.value = false
  resetSysLoggerOperate()
  disabled.value = true
}
/**
 * 处理查看操作
 * @param row - 要查看数据对象
 */
const handleItem = async (row: ResSysLoggerOperate) => {
  title.value = '查看操作日志'
  dialogVisible.value = true
  reset()
  const { data } = await getSysLoggerOperateApi(Number(row.id))
  sysLoggerOperateForm.value = data
  disabled.value = true
}
/**
 * 处理清理操作
 * @param row - 要清理数据对象
 */
const handleDrop = async (row: ResSysLoggerOperate) => {
  await useHandleData(dropSysLoggerOperateApi, Number(row.id), '清理操作日志')
  proTable.value?.getTableList()
}
/**
 * 处理删除操作
 * @param row - 要删除数据对象
 */
const handleDelete = async (row: ResSysLoggerOperate) => {
  await useHandleData(deleteSysLoggerOperateApi, Number(row.id), '删除操作日志')
  proTable.value?.getTableList()
}
/**
 * 处理恢复操作
 * @param row - 要恢复数据对象
 */
const handleRecover = async (row: ResSysLoggerOperate) => {
  await useHandleData(recoverSysLoggerOperateApi, Number(row.id), '恢复操作日志')
  proTable.value?.getTableList()
}

// 结果
const resultEnum = getIntDictOptions('operate.result')
// 请求方法
const methodEnum = getStrDictOptions('operate.method')

const columns: ProVxeColumnProps[] = [
  { prop: 'username', label: '用户名', valueType: 'input' },
  { prop: 'method', label: '请求方法', valueType: 'select', options: methodEnum },
  {
    prop: 'startTime',
    label: '开始时间',
    valueType: 'date-picker',
    fieldProps: {
      type: 'datetimerange',
      startPlaceholder: '请选择',
      endPlaceholder: '请选择',
      valueFormat: 'YYYY-MM-DD HH:mm:ss',
    },
  },
  { prop: 'channel', label: '渠道', valueType: 'input' },
  { prop: 'result', label: '结果', valueType: 'select', options: resultEnum },
  {
    prop: 'deleted',
    label: '删除',
    valueType: 'switch',
    fieldProps: {
      activeValue: 1,
      inactiveValue: 0,
    },
    hideInSearch: !HasAuth('logger.SysLoggerOperateDelete'),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
