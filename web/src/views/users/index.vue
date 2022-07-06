<template>
  <el-card>
    <el-row :gutter="20" class="header">
      <el-col :span="7">
        <el-input
          :placeholder="$t('table.placeholder')"
          cleartable
          v-model="queryForm.query"
        ></el-input>
      </el-col>
      <el-button type="primary" :icon="Search" @click="initGetUsersList">{{
        $t('table.search')
      }}</el-button>
      <el-button type="primary" @click="handleDialogValue($data)">{{
        $t('table.adduser')
      }}</el-button>
    </el-row>
    <el-radio-group v-model="roleRadio" @change="handleRoleChange()">
      <el-radio label="user">{{ $t('radio.user') }}</el-radio>
      <el-radio label="companyUser">{{ $t('radio.companyUser') }}</el-radio>
      <el-radio label="admin">{{ $t('radio.admin') }}</el-radio>
    </el-radio-group>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column
        :width="item.width"
        :prop="item.prop"
        :label="$t(`table.${item.label}`)"
        v-for="(item, index) in options"
        :key="index"
      >
        <template #default="{ row }" v-if="item.prop === 'action'">
          <el-button
            type="primary"
            size="small"
            :icon="Edit"
            @click="handleDialogValue(row)"
          ></el-button>
          <el-button
            type="danger"
            size="small"
            :icon="Delete"
            @click="delUser(row)"
          ></el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      v-model:currentPage="queryForm.pagenum"
      v-model:page-size="queryForm.pagesize"
      :page-sizes="[5, 10, 15, 20]"
      :small="small"
      :disabled="disabled"
      :background="background"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />
  </el-card>
  <Dialog
    v-model="dialogVisible"
    :dialogTitle="dialogTitle"
    v-if="dialogVisible"
    @initUserList="initGetUsersList"
    @role="roleRadio"
    :dialogTableValue="dialogTableValue"
  ></Dialog>
</template>

<script setup>
import { Search, Edit, Delete } from '@element-plus/icons-vue'
import { ref } from 'vue'
import { getUsers, deleteUser, deleteComUser, deleteAdmin } from '@/api/users'
import { options } from './options'
import Dialog from './components/dialog.vue'
import { isNull } from '@/utils/filters'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useI18n } from 'vue-i18n'
import { store } from './userstore'
const i18n = useI18n()
const roleRadio = ref('user')

const queryForm = ref({
  query: '',
  pagenum: 1,
  pagesize: 5,
  role: 'user'
})

const total = ref(0)
const tableData = ref([])
const dialogVisible = ref(false)
const dialogTitle = ref('')
const dialogTableValue = ref({})

const initGetUsersList = async () => {
  const res = await getUsers(queryForm.value)
  console.log(res)
  total.value = res.total
  tableData.value = res.users
}
initGetUsersList()

const handleRoleChange = () => {
  store.setRole(roleRadio.value)
  queryForm.value.role = store.role
  initGetUsersList()
}
const handleSizeChange = (val) => {
  queryForm.value.pagesize = val
  initGetUsersList()
}

const handleCurrentChange = (val) => {
  queryForm.value.pagenum = val
  initGetUsersList()
}

const handleDialogValue = (row) => {
  if (isNull(row)) {
    dialogTitle.value = '添加用户'
    dialogTableValue.value = {}
  } else {
    dialogTitle.value = '编辑用户'
    dialogTableValue.value = JSON.parse(JSON.stringify(row))
  }
  dialogVisible.value = true
}

const delUser = (row) => {
  ElMessageBox.confirm(i18n.t('dialog.deleteTitle'), 'Warning', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    type: 'warning'
  })
    .then(async () => {
      if (roleRadio.value === 'user') {
        await deleteUser(row.id)
      } else if (roleRadio.value === 'companyUser') {
        await deleteComUser(row.id)
      } else if (roleRadio.value === 'admin') {
        await deleteAdmin(row.id)
      } else {
        ElMessage.error(i18n.t('dialog.deleteError'))
      }
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      initGetUsersList()
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '取消删除'
      })
    })
  console.log(row)
}
</script>

<style lang="scss" scoped>
.header {
  padding-bottom: 16px;
  box-sizing: border-box;
}

::v-deep .el-pagination {
  padding-top: 16px;
  box-sizing: border-box;
  text-align: right;
}
</style>
