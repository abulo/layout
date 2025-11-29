<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="操作日志列表"
      row-key="id"
      :columns="columns"
      :toolbar-right="['search', 'refresh', 'export', 'layout']"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="ProTablePaginationEnum.BE"
      :search-col="12"
    >
      <template #toolbarLeft> </template>
      <!-- 删除状态 -->
      <template #deleted="scope">
        <DictTag type="deleted" :value="scope.row.deleted" />
      </template>
      <template #status="scope">
        <DictTag type="status" :value="scope.row.status" />
      </template>
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'logger.SysLoggerOperate'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['logger.SysLoggerOperateDelete', 'logger.SysLoggerOperateRecover', 'logger.SysLoggerOperateDrop']"
            type="primary"
            link
            :icon="DArrowRight"
          >
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <div v-auth="'logger.SysLoggerOperateDelete'">
                <el-dropdown-item v-if="scope.row.deleted === 0" :icon="Delete" @click="handleDelete(scope.row)">
                  删除
                </el-dropdown-item>
              </div>
              <div v-auth="'logger.SysLoggerOperateRecover'">
                <el-dropdown-item v-if="scope.row.deleted === 1" :icon="Refresh" @click="handleRecover(scope.row)">
                  恢复
                </el-dropdown-item>
              </div>
              <div v-auth="'logger.SysLoggerOperateDrop'">
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
import type { ResSysLoggerOperate } from '@/api/interface/sysLoggerOperate'
import type { ProTableInstance, ColumnProps, SearchProps } from '@/components/ProTable/interface'
import { Delete, Refresh, DeleteFilled, View, DArrowRight } from '@element-plus/icons-vue'
import {
  getSysLoggerOperateListApi,
  deleteSysLoggerOperateApi,
  dropSysLoggerOperateApi,
  recoverSysLoggerOperateApi,
  getSysLoggerOperateApi,
} from '@/api/modules/sysLoggerOperate'
import { getIntDictOptions } from '@/utils/dict'
import { DictTag } from '@/components/DictTag'
import { useHandleData } from '@/hooks/useHandleData'
import { HasAuth } from '@/utils/auth'
import { ProTablePaginationEnum } from '@/enums'
// 获取loading状态
//禁用
const disabled = ref(true)
//弹出层标题
const title = ref('')
//列表数据
const proTable = ref<ProTableInstance>()
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
  result: undefined, // 结果:0 成功/1 失败
  resultMsg: undefined, // 结果信息
  deleted: 0, // 删除:0否/1是
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
  newParams.startTime && (newParams.beginStartTime = newParams.startTime[0])
  newParams.startTime && (newParams.finishStartTime = newParams.startTime[1])
  delete newParams.startTime
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
    result: undefined, // 结果:0 成功/1 失败
    resultMsg: undefined, // 结果信息
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

//删除状态
const deletedEnum = getIntDictOptions('deleted')
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasAuth('logger.SysLoggerOperateDelete')
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

const columns: ColumnProps<ResSysLoggerOperate>[] = [
  { prop: 'id', label: '编号' },
  { prop: 'username', label: '用户名', search: { el: 'input', span: 2 } },
  { prop: 'module', label: '模块名称' },
  { prop: 'method', label: '请求方法', search: { el: 'input', span: 2 } },
  { prop: 'url', label: '请求地址' },
  { prop: 'ip', label: 'IP' },
  {
    prop: 'startTime',
    label: '开始时间',
    search: {
      el: 'date-picker',
      span: 4,
      attrs: { type: 'datetimerange', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
  },
  { prop: 'duration', label: '执行时长' },
  { prop: 'channel', label: '渠道', search: { el: 'input', span: 2 } },
  { prop: 'result', label: '结果', search: { el: 'input', span: 2 } },
  { prop: 'deleted', label: '删除', tag: true, enum: deletedEnum, search: deleteSearch },

  {
    prop: 'operation',
    label: '操作',
    width: 150,
    fixed: 'right',
    isShow: HasAuth(
      'logger.SysLoggerOperateDelete',
      'logger.SysLoggerOperateDrop',
      'logger.SysLoggerOperateRecover',
      'logger.SysLoggerOperate'
    ),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
