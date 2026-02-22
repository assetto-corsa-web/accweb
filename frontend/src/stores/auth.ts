import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { AuthToken } from '@/lib/accweb/types'

export const useAuthStore = defineStore('auth', () => {
  const username = ref<string | null>(window.localStorage.getItem('user_name'))
  const role = ref<string | null>(window.localStorage.getItem('role'))
  const token = ref<string | null>(window.localStorage.getItem('auth_token'))
  const admin = ref(window.localStorage.getItem('admin') === 'true')
  const mod = ref(window.localStorage.getItem('mod') === 'true')
  const read_only = ref(window.localStorage.getItem('read_only') === 'true')

  function login(data: AuthToken) {
    username.value = data.user_name
    role.value = data.role
    admin.value = data.admin
    mod.value = data.mod
    read_only.value = data.read_only

    window.localStorage.setItem('user_name', data.user_name)
    window.localStorage.setItem('role', data.role)
    window.localStorage.setItem('auth_token', data.token)
    window.localStorage.setItem('admin', String(admin.value))
    window.localStorage.setItem('mod', String(mod.value))
    window.localStorage.setItem('read_only', String(read_only.value))
  }

  function logout() {
    username.value = null
    role.value = null
    token.value = null
    admin.value = false
    mod.value = false
    read_only.value = false


    window.localStorage.removeItem('user_name')
    window.localStorage.removeItem('role')
    window.localStorage.removeItem('auth_token')
    window.localStorage.removeItem('admin')
    window.localStorage.removeItem('mod')
    window.localStorage.removeItem('read_only')
  }

  return { username, role, token, admin, mod, read_only, login, logout }
})
