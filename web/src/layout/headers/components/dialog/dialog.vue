<template>
  <el-dialog
    :model-value="dialogVisible"
    :title="dialogTitle"
    width="40%"
    @click="handleClose"
  >
    <el-form ref="formRef" :rules="rules" :model="form" label-width="120px">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="form.username" />
      </el-form-item>
      <el-form-item label="头像 Url" prop="avatar">
        <el-input v-model="form.avatar" />
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
import { ElMessage } from 'element-plus'
import { editMyself } from '@/api/users'
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
  uid: '',
  username: '',
  avatar: ''
})

const rules = ref({
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  avatar: [
    { required: true, message: '请输入头像 Url', trigger: 'blur' },
    { type: 'url', message: '请输入正确的头像 Url', trigger: 'blur' }
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

const emits = defineEmits(['update:modelValue', 'initUserInfo'])

const UpdateInfo = () => {
  localStorage.setItem('username', form.value.username)
  localStorage.setItem('avatar', form.value.avatar)
}
const handleClose = () => {
  emits('update:modelValue', false)
}

const handleConfirm = () => {
  formRef.value.validate(async (valid, fields) => {
    if (valid) {
      await editMyself(form.value)
      ElMessage({
        showClose: true,
        message: '提交成功',
        type: 'success'
      })
      UpdateInfo()
      handleClose()
      emits('initUserInfo')
    } else {
      console.log('error submit!', fields)
    }
  })
}
</script>

<style lang="scss" scoped></style>
