import './assets/main.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import http, { setAuthToken } from '@/lib/http'

import ui from '@nuxt/ui/vue-plugin'
import App from './App.vue'
import router from './router'
import { useAuthStore } from './stores/auth'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(ui)

const authStore = useAuthStore()

// set auth token if exists
if (authStore.token) {
  setAuthToken(authStore.token)
}

// response error handler
http.interceptors.response.use(response => response, err => {
  if (err.response?.status === 401) {
    authStore.logout()
    router.push("/login");
  }

  return Promise.reject(err)
});

app.mount('#app')
