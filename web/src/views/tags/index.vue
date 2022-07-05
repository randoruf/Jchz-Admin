<template>
  <el-card>
    <el-row :gutter="20" class="header">
      <el-col :span="7">
        <el-input
          :placeholder="$t('table.tagplaceholder')"
          cleartable
          v-model="queryForm.query"
        ></el-input>
      </el-col>
      <el-button type="primary" :icon="Search" @click="initGetTagsList">{{
        $t('table.search')
      }}</el-button>
      <el-button type="primary" @click="handleDialogValue($data)">{{
        $t('table.addtag')
      }}</el-button>
    </el-row>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column
        :width="item.width"
        :prop="item.prop"
        :label="$t(`tag.${item.label}`)"
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
            @click="delTag(row)"
          ></el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      v-model:currentPage="queryForm.pagenum"
      v-model:page-size="queryForm.pagesize"
      :page-sizes="[2, 5, 10, 15]"
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
    @initTagsList="initGetTagsList"
    :dialogTableValue="dialogTableValue"
  ></Dialog>
</template>

<script setup>
import { Search, Edit, Delete } from '@element-plus/icons-vue'
import { ref } from 'vue'
import { getTags, deleteTag } from '@/api/tags'
import { options } from './options'
import Dialog from './components/dialog.vue'
import { isNull } from '@/utils/filters'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useI18n } from 'vue-i18n'
const i18n = useI18n()

const queryForm = ref({
  query: '',
  pagenum: 1,
  pagesize: 2
})

const total = ref(0)
const tableData = ref([])
const dialogVisible = ref(false)
const dialogTitle = ref('')
const dialogTableValue = ref({})

const initGetTagsList = async () => {
  const res = await getTags(queryForm.value)
  console.log(res)
  total.value = res.total
  tableData.value = res.tags
}
initGetTagsList()

const handleSizeChange = (val) => {
  queryForm.value.pagesize = val
  initGetTagsList()
}

const handleCurrentChange = (val) => {
  queryForm.value.pagenum = val
  initGetTagsList()
}

const handleDialogValue = (row) => {
  if (isNull(row)) {
    dialogTitle.value = '添加标签'
    dialogTableValue.value = {}
  } else {
    dialogTitle.value = '编辑标签'
    dialogTableValue.value = JSON.parse(JSON.stringify(row))
  }
  dialogVisible.value = true
}

const delTag = (row) => {
  ElMessageBox.confirm(i18n.t('dialog.deleteTitle'), 'Warning', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    type: 'warning'
  })
    .then(async () => {
      deleteTag(row.id)
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      initGetTagsList()
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
