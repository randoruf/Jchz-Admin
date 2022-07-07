<template>
  <div class="avatar-container" id="avatar">
    <el-dropdown>
      <span class="el-dropdown-link">
        <el-avatar shape="square" :size="40" :src="avatar"></el-avatar>
      </span>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item>管理员：{{ username }}</el-dropdown-item>
          <el-dropdown-item @click="handleDialogValue"
            >编辑个人信息</el-dropdown-item
          >
          <el-dropdown-item @click="logout">退出</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </div>
  <Dialog
    v-model="dialogVisible"
    :dialogTitle="dialogTitle"
    v-if="dialogVisible"
    @initUserInfo="initUserInfo"
    :dialogTableValue="dialogTableValue"
  ></Dialog>
</template>

<script setup>
import { ref } from 'vue'
import { useStore } from 'vuex'
import Dialog from './dialog/dialog.vue'
const store = useStore()
const username = ref('')
const avatar = ref('')
const dialogVisible = ref(false)
const dialogTitle = ref('')
const dialogTableValue = ref({})
const logout = () => {
  store.dispatch('app/logout')
}

const getUsername = () => {
  return localStorage.getItem('username')
}

const getAvatar = () => {
  return localStorage.getItem('avatar')
}

const initUserInfo = () => {
  username.value = getUsername()
  avatar.value = getAvatar()
}
initUserInfo()
const handleDialogValue = () => {
  dialogTitle.value = '编辑个人信息'
  dialogTableValue.value = {
    id: localStorage.getItem('uid'),
    username: username.value,
    avatar: avatar.value
  }
  dialogVisible.value = true
}
</script>

<style lang="scss" scoped>
::v-deep .el-dropdown-menu__item {
  white-space: nowrap;
}
</style>
