<template>
  <div class="table-box">
    <ProVxeTable
      ref="proTable"
      title="登录日志列表"
      :columns="columns"
      :toolbar-right="['search', 'refresh']"
      :request-api="getTableList"
      :request-auto="true"
      :show-number="4"
      :border="true"
      :column-config="{ resizable: true, isCurrent: true, isHover: true }"
      :row-config="{ isCurrent: true, isHover: true }"
      :pagination="ProTablePaginationEnum.BE"
    >
      <template #toolbarLeft> </template>
      <vxe-column field="id" title="编号" fixed="left" width="auto"> </vxe-column>
      <vxe-column field="name" title="姓名"> </vxe-column>
      <vxe-column field="username" title="用户名"> </vxe-column>
      <vxe-column field="ua" title="UA" show-overflow> </vxe-column>
      <vxe-column field="loginTime" title="登录时间"> </vxe-column>
      <vxe-column field="channel" title="渠道"> </vxe-column>
      <vxe-column field="ip" title="IP"> </vxe-column>
      <vxe-column v-auth="'logger.SysLoggerLoginDelete'" field="deleted" title="删除">
        <template #default="{ row }">
          <DictTag type="deleted" :value="row.deleted" />
        </template>
      </vxe-column>
      <vxe-column
        v-auth="[
          'logger.SysLoggerLoginDelete',
          'logger.SysLoggerLoginDrop',
          'logger.SysLoggerLoginRecover',
          'logger.SysLoggerLogin',
        ]"
        fixed="right"
        title="操作"
        width="auto"
      >
        <template #default="{ row }">
          <!-- 菜单操作 -->
          <el-button v-auth="'logger.SysLoggerLogin'" type="primary" link :icon="View" @click="handleItem(row)">
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
                  <el-dropdown-item v-if="row.deleted === 0" :icon="Delete" @click="handleDelete(row)">
                    删除
                  </el-dropdown-item>
                </div>
                <div v-auth="'logger.SysLoggerLoginRecover'">
                  <el-dropdown-item v-if="row.deleted === 1" :icon="Refresh" @click="handleRecover(row)">
                    恢复
                  </el-dropdown-item>
                </div>
                <div v-auth="'logger.SysLoggerLoginDrop'">
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
import { ProVxeTableInstance, ProVxeColumnProps } from '@/components/ProVxeTable/interface'
import { Delete, Refresh, DeleteFilled, View, DArrowRight } from '@element-plus/icons-vue'
import {
  getSysLoggerLoginListApi,
  deleteSysLoggerLoginApi,
  dropSysLoggerLoginApi,
  recoverSysLoggerLoginApi,
  getSysLoggerLoginApi,
} from '@/api/modules/sysLoggerLogin'
// import { getIntDictOptions } from '@/utils/dict'
import { DictTag } from '@/components/DictTag'
import { useHandleData } from '@/hooks/useHandleData'
// import { useLoadingStore } from '@/stores/modules/loading'
// import { storeToRefs } from 'pinia'
import { ProTablePaginationEnum } from '@/enums'
import { HasAuth } from '@/utils/auth'
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
const sysLoggerLoginForm = ref<ResSysLoggerLogin>({
  id: 0, // 编号
  name: undefined, // 姓名
  username: '', // 用户名
  userId: 0, // 用户编号
  ua: undefined, // UA
  loginTime: '', // 登录时间
  channel: undefined, // 渠道
  ip: undefined, // IP
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
  newParams.loginTime && (newParams.beginLoginTime = newParams.loginTime[0])
  newParams.loginTime && (newParams.finishLoginTime = newParams.loginTime[1])
  delete newParams.loginTime
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
const columns: ProVxeColumnProps[] = [
  { prop: 'username', label: '用户名', valueType: 'input' },
  {
    prop: 'loginTime',
    label: '登录时间',
    valueType: 'date-picker',
    fieldProps: {
      type: 'datetimerange',
      startPlaceholder: '请选择',
      endPlaceholder: '请选择',
      valueFormat: 'YYYY-MM-DD HH:mm:ss',
    },
  },
  { prop: 'channel', label: '渠道', valueType: 'input' },
  {
    prop: 'deleted',
    label: '删除',
    valueType: 'switch',
    fieldProps: {
      activeValue: 1,
      inactiveValue: 0,
    },
    hideInSearch: !HasAuth('logger.SysLoggerLoginDelete'),
  },
]
</script>
<style scoped lang="scss">
@use '@/styles/custom';
</style>
