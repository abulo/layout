<template>
  <div class="table-box">
    <ProVxeTable
      ref="proTable"
      title="开发日志列表"
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
      <vxe-column field="id" title="编号" fixed="left" width="auto"></vxe-column>
      <vxe-column field="host" title="服务名"></vxe-column>
      <vxe-column field="timestamp" title="时间">
        <template #default="{ row }">
          {{ XEUtils.toDateString(row.timestamp, 'yyyy-MM-dd HH:mm:ss') }}
        </template>
      </vxe-column>
      <vxe-column field="file" title="文件" show-overflow></vxe-column>
      <vxe-column field="func" title="方法名" show-overflow></vxe-column>
      <vxe-column field="level" title="等级"></vxe-column>
      <vxe-column
        v-auth="['logger.SysLoggerDev', 'logger.SysLoggerDevDelete']"
        field="operation"
        fixed="right"
        title="操作"
        width="auto"
      >
        <template #default="{ row }">
          <el-button v-auth="'logger.SysLoggerDev'" type="primary" link :icon="View" @click="handleItem(row)">
            查看
          </el-button>
          <el-button v-auth="'logger.SysLoggerDevDelete'" type="primary" link :icon="Delete" @click="handleDelete(row)">
            删除
          </el-button>
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
        <el-descriptions-item label="编号"> {{ sysLoggerDevForm.id }}</el-descriptions-item>
        <el-descriptions-item label="服务名"> {{ sysLoggerDevForm.host }}</el-descriptions-item>
        <el-descriptions-item label="时间"> {{ sysLoggerDevForm.timestamp }}</el-descriptions-item>
        <el-descriptions-item label="文件"> {{ sysLoggerDevForm.file }}</el-descriptions-item>
        <el-descriptions-item label="方法名"> {{ sysLoggerDevForm.func }}</el-descriptions-item>
        <el-descriptions-item label="消息"> {{ sysLoggerDevForm.message }}</el-descriptions-item>
        <el-descriptions-item label="等级"> {{ sysLoggerDevForm.level }}</el-descriptions-item>
        <el-descriptions-item label="数据"> {{ sysLoggerDevForm.data }}</el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx">
defineOptions({ name: 'SysLoggerDev' })
import { ResSysLoggerDev } from '@/api/interface/sysLoggerDev'
import { ProVxeTableInstance, ProVxeColumnProps } from '@/components/ProVxeTable/interface'
import { Delete, View } from '@element-plus/icons-vue'
import { getSysLoggerDevListApi, deleteSysLoggerDevApi, getSysLoggerDevApi } from '@/api/modules/sysLoggerDev'
import { useHandleData } from '@/hooks/useHandleData'
import { ProTablePaginationEnum } from '@/enums'
import XEUtils from 'xe-utils'
// import { useLoadingStore } from '@/stores/modules/loading'
// import { storeToRefs } from 'pinia'
// 获取loading状态
// const { loading } = storeToRefs(useLoadingStore())
//禁用
const disabled = ref(true)
//弹出层标题
const title = ref('')
//列表数据
const proTable = ref<ProVxeTableInstance>()
//显示弹出层
const dialogVisible = ref(false)
//数据接口
const sysLoggerDevForm = ref<ResSysLoggerDev>({
  id: 0, // 编号
  host: undefined, // 服务名
  timestamp: '', // 时间
  file: undefined, //
  func: undefined, // 方法名
  message: undefined, // 消息
  level: undefined, // 等级
  data: undefined, // 数据
})
/**
 * 获取表格数据列表
 * @param params 查询参数对象
 * @returns 返回列表数据的Promise对象
 */
const getTableList = (params: any) => {
  // 深拷贝参数对象，避免修改原始参数
  let newParams = JSON.parse(JSON.stringify(params))
  newParams.timestamp && (newParams.beginTimestamp = newParams.timestamp[0])
  newParams.timestamp && (newParams.finishTimestamp = newParams.timestamp[1])
  delete newParams.timestamp
  return getSysLoggerDevListApi(newParams)
}

/**
 * 重置数据
 * 该函数用于清空或初始化
 * @returns {void} 无返回值
 */
const resetSysLoggerDev = () => {
  // 设置sysLoggerDevForm为默认值
  Object.assign(sysLoggerDevForm.value, {
    id: 0, // 编号
    host: undefined, // 服务名
    timestamp: '', // 时间
    file: undefined, //
    func: undefined, // 方法名
    message: undefined, // 消息
    level: undefined, // 等级
    data: undefined, // 数据
  })
}

/**
 * 重置函数
 * @returns {void}
 */
const reset = () => {
  // loading.value = false
  resetSysLoggerDev()
  disabled.value = true
}
/**
 * 处理查看操作
 * @param row - 要查看数据对象
 */
const handleItem = async (row: ResSysLoggerDev) => {
  title.value = '查看开发日志'
  dialogVisible.value = true
  reset()
  const { data } = await getSysLoggerDevApi(Number(row.id))
  sysLoggerDevForm.value = data
  disabled.value = true
}
/**
 * 处理删除操作
 * @param row - 要删除数据对象
 */
const handleDelete = async (row: ResSysLoggerDev) => {
  await useHandleData(deleteSysLoggerDevApi, Number(row.id), '删除开发日志')
  proTable.value?.getTableList()
}

const columns: ProVxeColumnProps[] = [
  { label: '服务名', prop: 'host', valueType: 'input' },
  {
    label: '时间',
    prop: 'timestamp',
    valueType: 'date-picker',
    fieldProps: {
      type: 'datetimerange',
      startPlaceholder: '请选择',
      endPlaceholder: '请选择',
      valueFormat: 'YYYY-MM-DD HH:mm:ss',
    },
  },
  { label: '等级', prop: 'level', valueType: 'input' },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
