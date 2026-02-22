<script setup lang="ts">
import {
  addToGlobalBanList,
  deleteFromGlobalBanList,
  getGlobalBanList,
  globalBanEnableToggle,
} from '@/lib/accweb/client'
import type { AccwebGlobalBanEntryJson, AccwebGlobalBanlistJson } from '@/lib/accweb/types'
import type { FormError, TableColumn } from '@nuxt/ui'
import { computed, onMounted, ref } from 'vue'

const data = ref<AccwebGlobalBanlistJson>()
const toggleButton = computed(() => {
  if (data.value && data.value.enabled) {
    return { color: 'error', label: 'Disable' }
  } else {
    return { color: 'success', label: 'Enable' }
  }
})

const newEntry = ref<AccwebGlobalBanEntryJson>({
  playerId: '',
  playerName: '',
} as AccwebGlobalBanEntryJson)

const toast = useToast()

const columns: TableColumn<AccwebGlobalBanEntryJson>[] = [
  {
    accessorKey: 'playerId',
    header: 'Player Id',
  },
  {
    accessorKey: 'playerName',
    header: 'Player Name',
  },
  {
    accessorKey: 'action',
    header: '',
  },
]

function loadData() {
  getGlobalBanList().then((res) => {
    data.value = res
  })
}

onMounted(() => {
  loadData()
})

function validate(): FormError[] {
  return []
}

async function onSubmit() {
  if (!data.value) {
    return
  }

  addToGlobalBanList(newEntry.value)
    .then(() => {
      toast.add({ title: 'Success', description: 'The form has been submitted.', color: 'success' })
      loadData()
    })
    .catch((err) => {
      toast.add({
        title: 'Failed to save',
        description: err.response.data.error,
        color: 'error',
      })
    })
}

async function toggleEnabled() {
  return globalBanEnableToggle().then(() => {
    toast.add({
      title: 'Success',
      description: 'The Global Ban list has been toggled.',
      color: 'success',
    })
    loadData()
  })
}

async function deleteEntry(id: number) {
  return deleteFromGlobalBanList(id)
    .then(() => {
      toast.add({ title: 'Success', description: 'The entry has been deleted.', color: 'success' })
      loadData()
    })
    .catch((err) => {
      toast.add({
        title: 'Failed to delete entry',
        description: err.response.data.error,
        color: 'error',
      })
    })
}
</script>

<template>
  <div class="flex flex-col gap-3" v-if="data">
    <p>Share a list of banned players across all your servers that have this feature enabled.</p>

    <UButton
      :label="toggleButton.label"
      :color="toggleButton.color"
      class="flex-none w-fit"
      loading-auto
      @click="toggleEnabled()"
    />

    <USeparator />

    <div class="flex flex-row gap-3">
      <div class="flex-1">
        <UForm :validate="validate" :state="newEntry" class="space-y-4" @submit="onSubmit">
          <UCard>
            <template #header>Add New Entry</template>

            <template #default>
              <div class="flex flex-row gap-2 items-end">
                <TFormField name="playerName" label="Player Name" help="Player Name">
                  <UInput v-model="newEntry.playerName" class="w-full" />
                </TFormField>

                <TFormField name="playerId" label="Player ID" help="Player ID">
                  <UInput v-model="newEntry.playerId" class="w-full" />
                </TFormField>

                <UButton type="submit" icon="i-lucide-save" variant="subtle">Add</UButton>
              </div>
            </template>
          </UCard>
        </UForm>
      </div>

      <div class="flex-1">
        <UTable :data="data.entries" :columns="columns" class="">
          <template #action-cell="{ row }">
            <UButton
              icon="i-lucide-trash"
              color="error"
              variant="ghost"
              loading-auto
              @click="deleteEntry(row.index)"
            />
          </template>
        </UTable>
      </div>
    </div>
  </div>
</template>
