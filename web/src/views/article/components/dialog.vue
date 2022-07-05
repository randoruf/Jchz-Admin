<template>
  <el-dialog
    :model-value="dialogVisible"
    :title="dialogTitle"
    width="40%"
    @click="handleClose"
  >
    <el-form ref="formRef" :rules="rules" :model="form" label-width="120px">
      <el-form-item label="文章标题" prop="title">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="文章内容" prop="content" class="artContent">
        <el-input
          v-model="form.content"
          :autosize="{ minRows: 2, maxRows: 4 }"
          type="textarea"
          placeholder="请输入..."
        />
      </el-form-item>
      <el-form-item label="封面图片" prop="cover">
        <el-input v-model="form.cover" />
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
import { editArticle } from '@/api/article'
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
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
  content: [{ required: true, message: '请输入内容', trigger: 'blur' }],
  cover: [{ required: true, message: '请输入封面图片URL', trigger: 'blur' }]
})

watch(
  () => props.dialogTableValue,
  () => {
    console.log(props.dialogTableValue)
    form.value = props.dialogTableValue
  },
  { deep: true, immediate: true }
)

const emits = defineEmits(['update:modelValue', 'initArticlesList'])
const handleClose = () => {
  emits('update:modelValue', false)
}

const handleConfirm = () => {
  formRef.value.validate(async (valid, fields) => {
    if (valid) {
      await editArticle(form.value)
      ElMessage({
        showClose: true,
        message: '提交成功',
        type: 'success'
      })
      handleClose()
      emits('initArticlesList')
    } else {
      console.log('error submit!', fields)
    }
  })
}
</script>

<style lang="scss" scoped></style>
