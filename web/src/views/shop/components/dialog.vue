<template>
  <el-dialog
    :model-value="dialogVisible"
    :title="dialogTitle"
    width="40%"
    @click="handleClose"
  >
    <el-form ref="formRef" :rules="rules" :model="form" label-width="120px">
      <el-form-item label="店铺名称" prop="name">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="店铺描述" prop="content">
        <el-input
          v-model="form.content"
          :autosize="{ minRows: 2, maxRows: 4 }"
          type="textarea"
          placeholder="请输入..."
        />
      </el-form-item>
      <el-form-item label="店铺地址" prop="address">
        <el-input v-model="form.address" />
      </el-form-item>
      <el-form-item label="店铺电话" prop="phone">
        <el-input v-model="form.phone" />
      </el-form-item>
      <el-form-item label="店铺用户 id" prop="com_id">
        <el-input v-model="form.com_id" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" @click="handleConfirm">确认</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import { defineEmits, ref, defineProps, watch } from 'vue'
import { addShop, editShop, CheckComID } from '@/api/shop'
import { ElMessage } from 'element-plus'
const props = defineProps({
  dialogTitle: {
    type: String,
    default: '',
    required: true
  },
  dialogTableValue: {
    type: Object,
    default: () => {}
  }
})

const formRef = ref(null)

const form = ref({
  name: '',
  content: '',
  address: '',
  phone: '',
  com_id: ''
})

const validateID = async (rule, value, callback) => {
  const reg = /^[1-9]\d*$/
  if (!reg.test(value)) {
    callback(new Error('商家用户 id 必须是正整数'))
  } else {
    const res = await CheckComID(value)
    if (res != null && res.result === 'true') {
      return
    }
    callback(new Error('商家用户不存在'))
  }
}

const rules = ref({
  name: [{ required: true, message: '请输入店铺名称', trigger: 'blur' }],
  content: [{ required: true, message: '请输入店铺描述', trigger: 'blur' }],
  address: [{ required: true, message: '请输入店铺地址', trigger: 'blur' }],
  phone: [{ required: true, message: '请输入店铺电话', trigger: 'blur' }],
  com_id: [
    {
      required: true,
      message: '请输入商家用户 id',
      trigger: 'blur'
    },
    {
      validator: validateID,
      trigger: 'blur'
    }
  ]
})

watch(
  () => props.dialogTableValue,
  () => {
    console.log(props.dialogTableValue)
    form.value = props.dialogTableValue
  },
  { deep: true, immediate: true }
)

const emits = defineEmits(['update:modelValue', 'initShopsList'])
const handleClose = () => {
  emits('update:modelValue', false)
}

const handleConfirm = () => {
  formRef.value.validate(async (valid, fields) => {
    if (valid) {
      props.dialogTitle === '添加店铺'
        ? await addShop(form.value)
        : await editShop(form.value)
      ElMessage({
        showClose: true,
        message: '提交成功',
        type: 'success'
      })
      handleClose()
      emits('initShopsList')
    } else {
      console.log('error submit!', fields)
    }
  })
}
</script>

<style lang="scss" scoped></style>
