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
      <el-form-item label="密码" prop="password">
        <el-input v-model="form.password" />
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="form.email" />
      </el-form-item>
      <el-form-item label="手机号" prop="phone">
        <el-input v-model="form.phone" />
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
import {
  addUser,
  editUser,
  addComUser,
  editComUser,
  addAdmin,
  editAdmin
} from '@/api/users'
import { ElMessage } from 'element-plus'
import { store } from '../userstore'
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
  username: '',
  password: '',
  email: '',
  phone: '',
  sex: '',
  hometown: '',
  avatar: '',
  birth: ''
})

const rules = ref({
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 16, message: '长度在 3 到 16 个字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 16, message: '长度在 6 到 16 个字符', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur' }
  ],
  phone: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    { min: 11, max: 11, message: '请输入 11 位手机号', trigger: 'blur' }
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

const emits = defineEmits(['update:modelValue', 'initUserList'])
const handleClose = () => {
  emits('update:modelValue', false)
}

const handleConfirm = () => {
  formRef.value.validate(async (valid, fields) => {
    if (valid) {
      if (store.role === 'user') {
        props.dialogTitle === '添加用户'
          ? await addUser(form.value)
          : await editUser(form.value)
      } else if (store.role === 'companyUser') {
        props.dialogTitle === '添加用户'
          ? await addComUser(form.value)
          : await editComUser(form.value)
      } else if (store.role === 'admin') {
        props.dialogTitle === '添加用户'
          ? await addAdmin(form.value)
          : await editAdmin(form.value)
      } else {
        ElMessage.error('角色错误')
      }
      ElMessage({
        showClose: true,
        message: '提交成功',
        type: 'success'
      })
      handleClose()
      emits('initUserList')
    } else {
      console.log('error submit!', fields)
    }
  })
}
</script>

<style lang="scss" scoped></style>
