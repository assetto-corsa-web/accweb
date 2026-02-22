<script setup lang="ts">
import { cloneInstance, deleteInstance } from '@/lib/accweb/client'
import { useAuthStore } from '@/stores/auth'
import InstanceStartStopBtn from '@/components/InstanceStartStopBtn.vue'
import type { ListServerItem } from '@/lib/accweb/types'

const state = useAuthStore()

const props = defineProps<{
  instance: ListServerItem
  refreshFunction: () => void
}>()

const optionItems = [
  [
    {
      label: 'Clone',
      icon: 'i-lucide-copy',
      disabled: !state.admin,
      onSelect: () => clone(props.instance.id),
    },
  ],
  [
    {
      label: 'Delete',
      color: 'error',
      icon: 'i-lucide-trash',
      disabled: props.instance.isRunning || !state.admin,
      onSelect: () => deleteI(props.instance.id),
    },
  ],
]

const deleteI = (id: string) => {
  deleteInstance(id)
    .then(() => {
      props.refreshFunction()
    })
    .catch((error) => {
      console.error('Error deleting instance:', error)
    })
}

const clone = (id: string) => {
  cloneInstance(id)
    .then(() => {
      props.refreshFunction()
    })
    .catch((error) => {
      console.error('Error cloning instance:', error)
    })
}
</script>

<template>
  <UCard class="min-w-32" variant="soft" :class="{ withDrivers: instance.nrClients > 0 }">
    <template #header>
      <RouterLink :to="`/instance/${instance.id}`">
        <UTooltip :text="instance.name">
          <div class="max-w-fit text-nowrap overflow-hidden text-ellipsis">
            {{ instance.name }}
          </div>
        </UTooltip>
      </RouterLink>
    </template>

    {{ instance.nrClients }} - {{ instance.serverState }} <br />
    <strong>Phase:</strong> {{ instance.sessionType }} ({{ instance.sessionPhase }}) [{{
      instance.sessionRemaining
    }}
    min]

    <template #footer>
      <div>
        <InstanceStartStopBtn
          :instanceId="instance.id"
          :isRunning="instance.isRunning"
          :refreshFunction="props.refreshFunction"
        />

        <UDropdownMenu :items="optionItems" class="float-right">
          <UButton icon="i-lucide-menu" color="neutral" variant="ghost" label="Options" />
        </UDropdownMenu>
      </div>
    </template>
  </UCard>
</template>

<style scoped>
.withDrivers {
  background-color: #0a1f2a !important;
}
</style>
