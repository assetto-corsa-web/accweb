<script setup lang="ts">
import { onMounted } from 'vue'
import { login, setAuthToken } from '@/lib/accweb/client'
import * as z from 'zod'
import type { FormSubmitEvent, AuthFormField } from '@nuxt/ui'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const state = useAuthStore()

onMounted(async () => {})

const toast = useToast()
const router = useRouter()

const fields: AuthFormField[] = [
  {
    name: 'username',
    label: 'Username',
    type: 'text',
    placeholder: 'Enter your username',
    required: true,
  },
  {
    name: 'password',
    label: 'Password',
    type: 'password',
    placeholder: 'Enter your password',
    required: true,
  },
]

const schema = z.object({
  username: z.string('Username is required'),
  password: z.string('Password is required'),
})

type Schema = z.output<typeof schema>

async function onSubmit(payload: FormSubmitEvent<Schema>) {
  login(payload.data.username, payload.data.password)
    .then((resp) => {
      setAuthToken(resp.token)
      state.login(resp)
      router.push({ path: '/' })
    })
    .catch((error) => {
      console.error('Login error:', error)
      toast.add({
        title: 'Uh oh! Something went wrong.',
        description: error.response?.data?.message || 'Please try again later.',
        color: 'error',
      })
    })
}
</script>

<template>
  <div class="flex flex-col items-center justify-center gap-4 p-4">
    <UPageCard class="w-full max-w-md">
      <UAuthForm
        :schema="schema"
        :fields="fields"
        title="Welcome back!"
        icon="i-lucide-lock"
        @submit="onSubmit"
      />
    </UPageCard>
  </div>
</template>
