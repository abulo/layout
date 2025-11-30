<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      title="登录日志列表"
      row-key="id"
      :columns="columns"
      :toolbar-right="['search', 'refresh']"
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
        <el-button v-auth="'logger.SysLoggerLogin'" type="primary" link :icon="View" @click="handleItem(scope.row)">
          查看
        </el-button>
        <el-dropdown trigger="click">
          <el-button
            v-auth="['logger.SysLoggerLoginDelete', 'logger.SysLoggerLoginRecover', 'logger.SysLoggerLoginDrop']"
            type="primary"
            link
            :icon="DArrowRight"
          >
            更多
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <div v-auth="'logger.SysLoggerLoginDelete'">
                <el-dropdown-item v-if="scope.row.deleted === 0" :icon="Delete" @click="handleDelete(scope.row)">
                  删除
                </el-dropdown-item>
              </div>
              <div v-auth="'logger.SysLoggerLoginRecover'">
                <el-dropdown-item v-if="scope.row.deleted === 1" :icon="Refresh" @click="handleRecover(scope.row)">
                  恢复
                </el-dropdown-item>
              </div>
              <div v-auth="'logger.SysLoggerLoginDrop'">
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
        <el-descriptions-item label="编号"> {{ sysLoggerLoginForm.id }} </el-descriptions-item>
        <el-descriptions-item label="姓名"> {{ sysLoggerLoginForm.name }} </el-descriptions-item>
        <el-descriptions-item label="用户名"> {{ sysLoggerLoginForm.username }} </el-descriptions-item>
        <el-descriptions-item label="用户编号"> {{ sysLoggerLoginForm.userId }} </el-descriptions-item>
        <el-descriptions-item label="UA"> {{ sysLoggerLoginForm.ua }} </el-descriptions-item>
        <el-descriptions-item label="登录时间"> {{ sysLoggerLoginForm.loginTime }} </el-descriptions-item>
        <el-descriptions-item label="渠道"> {{ sysLoggerLoginForm.channel }} </el-descriptions-item>
        <el-descriptions-item label="IP"> {{ sysLoggerLoginForm.ip }} </el-descriptions-item>
        <el-descriptions-item label="删除"> {{ sysLoggerLoginForm.deleted }} </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>
<script setup lang="tsx">
defineOptions({ name: 'SysLoggerLogin' })
import { ResSysLoggerLogin } from '@/api/interface/sysLoggerLogin'
import { ProTableInstance, ColumnProps, SearchProps } from '@/components/ProTable/interface'
import { Delete, Refresh, DeleteFilled, View, DArrowRight } from '@element-plus/icons-vue'
import {
  getSysLoggerLoginListApi,
  deleteSysLoggerLoginApi,
  dropSysLoggerLoginApi,
  recoverSysLoggerLoginApi,
  getSysLoggerLoginApi,
} from '@/api/modules/sysLoggerLogin'
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
const sysLoggerLoginForm = ref<ResSysLoggerLogin>({
  id: 0, // 编号
  name: undefined, // 姓名
  username: '', // 用户名
  userId: 0, // 用户编号
  ua: undefined, // UA
  loginTime: '', // 登录时间
  channel: undefined, // 渠道
  ip: undefined, // IP
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
  return getSysLoggerLoginListApi(newParams)
}

/**
 * 重置数据
 * 该函数用于清空或初始化
 * @returns {void} 无返回值
 */
const resetSysLoggerLogin = () => {
  // 设置sysLoggerLoginForm为默认值
  Object.assign(sysLoggerLoginForm.value, {
    id: 0, // 编号
    name: undefined, // 姓名
    username: '', // 用户名
    userId: 0, // 用户编号
    ua: undefined, // UA
    loginTime: '', // 登录时间
    channel: undefined, // 渠道
    ip: undefined, // IP
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
  resetSysLoggerLogin()
  disabled.value = true
}
/**
 * 处理查看操作
 * @param row - 要查看数据对象
 */
const handleItem = async (row: ResSysLoggerLogin) => {
  title.value = '查看登录日志'
  dialogVisible.value = true
  reset()
  const { data } = await getSysLoggerLoginApi(Number(row.id))
  sysLoggerLoginForm.value = data
  disabled.value = true
}
/**
 * 处理清理操作
 * @param row - 要清理数据对象
 */
const handleDrop = async (row: ResSysLoggerLogin) => {
  await useHandleData(dropSysLoggerLoginApi, Number(row.id), '清理登录日志')
  proTable.value?.getTableList()
}
/**
 * 处理删除操作
 * @param row - 要删除数据对象
 */
const handleDelete = async (row: ResSysLoggerLogin) => {
  await useHandleData(deleteSysLoggerLoginApi, Number(row.id), '删除登录日志')
  proTable.value?.getTableList()
}
/**
 * 处理恢复操作
 * @param row - 要恢复数据对象
 */
const handleRecover = async (row: ResSysLoggerLogin) => {
  await useHandleData(recoverSysLoggerLoginApi, Number(row.id), '恢复登录日志')
  proTable.value?.getTableList()
}

//删除状态
const deletedEnum = getIntDictOptions('deleted')
// 表格配置项
const deleteSearch = reactive<SearchProps>(
  HasAuth('logger.SysLoggerLoginDelete')
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

const columns: ColumnProps<ResSysLoggerLogin>[] = [
  { prop: 'id', label: '编号' },
  { prop: 'name', label: '姓名' },
  { prop: 'username', label: '用户名', search: { el: 'input', span: 2 } },
  { prop: 'ua', label: 'UA' },
  {
    prop: 'loginTime',
    label: '登录时间',
    search: {
      el: 'date-picker',
      span: 4,
      attrs: { type: 'datetimerange', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
  },
  { prop: 'channel', label: '渠道', search: { el: 'input', span: 2 } },
  { prop: 'ip', label: 'IP' },
  { prop: 'deleted', label: '删除', tag: true, enum: deletedEnum, search: deleteSearch },
  {
    prop: 'operation',
    label: '操作',
    width: 150,
    fixed: 'right',
    isShow: HasAuth(
      'logger.SysLoggerLoginDelete',
      'logger.SysLoggerLoginDrop',
      'logger.SysLoggerLoginRecover',
      'logger.SysLoggerLogin'
    ),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
