<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="开发日志列表"
      row-key="id"
      :columns="columns"
      :toolbar-right="['search', 'refresh', 'export', 'layout']"
      :request-api="getTableList"
      :request-auto="true"
      :pagination="ProTablePaginationEnum.BE"
      :search-col="12"
    >
      <!-- 菜单操作 -->
      <template #operation="scope">
        <el-button v-auth="'logger.SysLoggerDev'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-button
          v-auth="'logger.SysLoggerDevDelete'"
          type="primary"
          link
          :icon="Delete"
          @click="handleDelete(scope.row)"
        >
          删除
        </el-button>
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
import { ProTableInstance, ColumnProps } from '@/components/ProTable/interface'
import { Delete, View } from '@element-plus/icons-vue'
import { getSysLoggerDevListApi, deleteSysLoggerDevApi, getSysLoggerDevApi } from '@/api/modules/sysLoggerDev'
import { useHandleData } from '@/hooks/useHandleData'
import { HasAuth } from '@/utils/auth'
import { ProTablePaginationEnum } from '@/enums'
//禁用
const disabled = ref(true)
//弹出层标题
const title = ref('')
//列表数据
const proTable = ref<ProTableInstance>()
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

const columns: ColumnProps<ResSysLoggerDev>[] = [
  { prop: 'id', label: '编号' },
  { prop: 'host', label: '服务名', search: { el: 'input', span: 2 } },
  {
    prop: 'timestamp',
    label: '时间',
    search: {
      el: 'date-picker',
      span: 4,
      attrs: { type: 'datetimerange', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
  },
  { prop: 'file', label: '文件' },
  { prop: 'func', label: '方法名' },
  { prop: 'level', label: '等级', search: { el: 'input', span: 2 } },
  {
    prop: 'operation',
    label: '操作',
    width: 150,
    fixed: 'right',
    isShow: HasAuth('logger.SysLoggerDevDelete', 'logger.SysLoggerDev'),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
