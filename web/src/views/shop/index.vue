<template>
  <el-card>
    <el-row :gutter="20" class="header">
      <el-col :span="7">
        <el-input
          :placeholder="$t('table.shopplaceholder')"
          cleartable
          v-model="queryForm.query"
        ></el-input>
      </el-col>
      <el-button type="primary" :icon="Search" @click="initGetShopList">{{
        $t('table.search')
      }}</el-button>
      <el-button type="primary" @click="handleDialogValue($data)">{{
        $t('table.addShop')
      }}</el-button>
    </el-row>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column
        :width="item.width"
        :prop="item.prop"
        :label="$t(`shop.${item.label}`)"
        v-for="(item, index) in options"
        :key="index"
      >
        <template #default="{ row }" v-if="item.prop === 'name'">
          {{ TrimString(row.name, 16) }}
        </template>
        <template #default="{ row }" v-else-if="item.prop === 'content'">
          {{ TrimString(row.content, 32) }}
        </template>
        <template #default="{ row }" v-else-if="item.prop === 'address'">
          {{ TrimString(row.address, 32) }}
        </template>
        <template #default="{ row }" v-else-if="item.prop === 'action'">
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
            @click="delShop(row)"
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
    @initShopsList="initGetShopList"
    :dialogTableValue="dialogTableValue"
  ></Dialog>
</template>

<script setup>
import { Search, Edit, Delete } from '@element-plus/icons-vue'
import { ref } from 'vue'
import { getShops, deleteShop } from '@/api/shop'
import { options } from './options'
import Dialog from './components/dialog.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useI18n } from 'vue-i18n'
import { isNull, TrimString } from '../../utils/filters'
const i18n = useI18n()

const queryForm = ref({
  query: '',
  pagenum: 1,
  pagesize: 5
})

const total = ref(0)
const tableData = ref([])
const dialogVisible = ref(false)
const dialogTitle = ref('')
const dialogTableValue = ref({})

const initGetShopList = async () => {
  const res = await getShops(queryForm.value)
  console.log(res)
  total.value = res.total
  tableData.value = res.shops
}
initGetShopList()

const handleSizeChange = (val) => {
  queryForm.value.pagesize = val
  initGetShopList()
}

const handleCurrentChange = (val) => {
  queryForm.value.pagenum = val
  initGetShopList()
}

const handleDialogValue = (row) => {
  if (isNull(row)) {
    dialogTitle.value = '添加店铺'
    dialogTableValue.value = {}
  } else {
    dialogTitle.value = '编辑店铺'
    dialogTableValue.value = JSON.parse(JSON.stringify(row))
  }
  dialogVisible.value = true
}

const delShop = (row) => {
  ElMessageBox.confirm(i18n.t('dialog.deleteTitle'), 'Warning', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    type: 'warning'
  })
    .then(async () => {
      await deleteShop(row.id)
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      initGetShopList()
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
