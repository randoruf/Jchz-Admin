import { reactive } from 'vue'

export const store = reactive({
  role: 'user',
  setRole(role) {
    this.role = role
  }
})
