<template>
  <div class="card content-box">
    <span class="text"> 按钮权限 🍓🍇🍈🍉</span>
    <el-alert class="mb-5" type="success" :closable="false" />
    <el-divider content-position="left"> 使用 Hooks 方式绑定权限 </el-divider>
    <el-divider content-position="left"> 使用 v-auth 指令绑定单个权限 </el-divider>
    <el-row class="mb-5">
      <el-button v-auth="'add'" type="primary" :icon="CirclePlus"> 新增 </el-button>
      <el-button v-auth="'edit'" type="warning" :icon="EditPen"> 编辑 </el-button>
      <el-button v-auth="'delete'" type="danger" plain :icon="Delete"> 删除 </el-button>
      <el-button v-auth="'import'" type="info" plain :icon="Upload"> 导入数据 </el-button>
      <el-button v-auth="'export'" type="info" plain :icon="Download"> 导出数据 </el-button>
    </el-row>
    <el-divider content-position="left"> 使用 v-auth 指令绑定多个权限 </el-divider>
    <el-row>
      <el-button v-auth="['add', 'edit', 'delete', 'import', 'export']" type="primary" :icon="CirclePlus">
        新增
      </el-button>
      <el-button v-auth="['add', 'edit', 'delete', 'import', 'export']" type="warning" :icon="EditPen">
        编辑
      </el-button>
      <el-button v-auth="['add', 'edit', 'delete', 'import', 'export']" type="danger" plain :icon="Delete">
        删除
      </el-button>
      <el-button v-auth="['add', 'edit', 'delete', 'import', 'export']" type="info" plain :icon="Upload">
        导入数据
      </el-button>
      <el-button v-auth="['add', 'edit', 'delete', 'import', 'export']" type="info" plain :icon="Download">
        导出数据
      </el-button>
    </el-row>
  </div>
</template>

<script setup lang="ts">
defineOptions({ name: 'AuthButton' })
import { CirclePlus, Delete, EditPen, Download, Upload } from '@element-plus/icons-vue'
</script>

<style scoped lang="scss">
@use './index';
</style>
