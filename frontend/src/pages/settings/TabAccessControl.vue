<script setup lang="ts">
import { addUser, deleteUser, listUsers, updateUser } from '@/lib/accweb/client'
import type { NewUserPayload, UpdateUserPayload, UserPayload } from '@/lib/accweb/types'
import type { FormError, TableColumn } from '@nuxt/ui'
import { onMounted, ref } from 'vue'

const data = ref<UserPayload[]>()
const newEntry = ref<NewUserPayload>({
  username: '',
  role: 'admin',
  password: '',
} as NewUserPayload)
const updateUserEntry = ref<NewUserPayload>()
const show = ref(false)
const roles = ref(['admin', 'moderator', 'read_only'])
const newUserOpen = ref(false)
const updateUserOpen = ref(false)

function loadData() {
  listUsers().then((res) => {
    data.value = res
  })
}

onMounted(() => {
  loadData()
})

const columns: TableColumn<UserPayload>[] = [
  {
    accessorKey: 'username',
    header: 'User Name',
  },
  {
    accessorKey: 'role',
    header: 'role',
  },
  {
    accessorKey: 'action',
    header: '',
  },
]

function onSubmit() {
  addUser(newEntry.value).then(() => {
    newUserOpen.value = false
    newEntry.value = {
      username: '',
      role: 'admin',
      password: '',
    } as NewUserPayload
    loadData()
  })
}

function validate(): FormError[] {
  return []
}

function onUpdateUser(username: string, role: string) {
  updateUserEntry.value = {
    username,
    role,
  } as NewUserPayload

  updateUserOpen.value = true
}

async function onDeleteUser(username: string) {
  return deleteUser(username).then(() => {
    loadData()
  })
}

function onUpdateSubmit() {
  if (!updateUserEntry.value) return

  const data = {
    role: updateUserEntry.value.role,
    password: updateUserEntry.value.password != '' ? updateUserEntry.value.password : null,
  } as UpdateUserPayload

  updateUser(updateUserEntry.value.username, data).then(() => {
    updateUserOpen.value = false
    loadData()
  })
}
</script>

<template>
  <div v-id="data">
    <UModal title="Add new user" v-model:open="newUserOpen">
      <UButton icon="i-lucide-plus" label="Add new user" color="primary" variant="subtle" />

      <template #body>
        <UForm :validate="validate" :state="newEntry" class="space-y-4" @submit="onSubmit">
          <div class="flex flex-col gap-2">
            <TFormField name="username" label="User Name">
              <UInput v-model="newEntry.username" class="w-50" />
            </TFormField>

            <TFormField name="password" label="Password">
              <UInput
                v-model="newEntry.password"
                placeholder="Password"
                :type="show ? 'text' : 'password'"
                class="w-50"
                :ui="{ trailing: 'pe-1' }"
              >
                <template #trailing>
                  <UButton
                    color="neutral"
                    variant="link"
                    size="sm"
                    :icon="show ? 'i-lucide-eye-off' : 'i-lucide-eye'"
                    :aria-label="show ? 'Hide password' : 'Show password'"
                    :aria-pressed="show"
                    aria-controls="password"
                    @click="show = !show"
                  />
                </template>
              </UInput>
            </TFormField>

            <TFormField name="role" label="Role">
              <USelect v-model="newEntry.role" :items="roles" class="w-50" />
            </TFormField>

            <UButton type="submit" icon="i-lucide-save" label="Submit" class="w-fit mt-3" />
          </div>
        </UForm>
      </template>
    </UModal>

    <UTable :data="data" :columns="columns" class="w-1/2">
      <template #action-cell="{ row }">
        <div class="flex gap-1">
          <UButton
            icon="i-lucide-pen"
            color="secondary"
            variant="ghost"
            @click="onUpdateUser(row.getValue('username'), row.getValue('role'))"
          />

          <UButton
            icon="i-lucide-trash"
            color="error"
            variant="ghost"
            loading-auto
            @click="onDeleteUser(row.getValue('username'))"
          />
        </div>
      </template>
    </UTable>

    <UModal :title="`Update user ${updateUserEntry?.username}`" v-model:open="updateUserOpen">
      <template #body>
        <UForm
          :validate="validate"
          :state="updateUserEntry"
          class="space-y-4"
          @submit="onUpdateSubmit"
        >
          <div class="flex flex-col gap-2" v-if="updateUserEntry">
            <TFormField name="password" label="Password">
              <UInput
                v-model="updateUserEntry.password"
                placeholder="Password"
                :type="show ? 'text' : 'password'"
                class="w-50"
                :ui="{ trailing: 'pe-1' }"
              >
                <template #trailing>
                  <UButton
                    color="neutral"
                    variant="link"
                    size="sm"
                    :icon="show ? 'i-lucide-eye-off' : 'i-lucide-eye'"
                    :aria-label="show ? 'Hide password' : 'Show password'"
                    :aria-pressed="show"
                    aria-controls="password"
                    @click="show = !show"
                  />
                </template>
              </UInput>
            </TFormField>

            <TFormField name="role" label="Role">
              <USelect v-model="updateUserEntry.role" :items="roles" class="w-50" />
            </TFormField>

            <UButton type="submit" icon="i-lucide-save" label="Submit" class="w-fit mt-3" />
          </div>
        </UForm>
      </template>
    </UModal>
  </div>
</template>
