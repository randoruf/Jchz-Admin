<template>
  <el-dialog
    :model-value="dialogVisible"
    :title="dialogTitle"
    width="40%"
    @click="handleClose"
  >
    <el-form ref="formRef" :rules="rules" :model="form" label-width="120px">
      <el-form-item label="标签名称" prop="tagname">
        <el-input v-model="form.tagname" />
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
import { addTag, editTag } from '@/api/tags'
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
  tagname: ''
})

const rules = ref({
  tagname: [{ required: true, message: '请输入标签名', trigger: 'blur' }]
})

watch(
  () => props.dialogTableValue,
  () => {
    console.log(props.dialogTableValue)
    form.value = props.dialogTableValue
  },
  { deep: true, immediate: true }
)

const emits = defineEmits(['update:modelValue', 'initTagsList'])
const handleClose = () => {
  emits('update:modelValue', false)
}

const handleConfirm = () => {
  formRef.value.validate(async (valid, fields) => {
    if (valid) {
      props.dialogTitle === '添加标签'
        ? await addTag(form.value)
        : await editTag(form.value)
      ElMessage({
        showClose: true,
        message: '提交成功',
        type: 'success'
      })
      handleClose()
      emits('initTagsList')
    } else {
      console.log('error submit!', fields)
    }
  })
}
</script>

<style lang="scss" scoped></style>
