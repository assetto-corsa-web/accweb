<script setup lang="ts">
import { startInstance, stopInstance, stopInstanceAfterRaceWeekend } from '@/lib/accweb/client'
import { useAuthStore } from '@/stores/auth'

const state = useAuthStore()

const props = defineProps<{
  instanceId: string
  isRunning: boolean
  refreshFunction: () => void
}>()

const stopItems = [
  {
    label: 'Stop now',
    onSelect: () => stop(props.instanceId),
  },
  {
    label: 'Stop after race weekend',
    onSelect: () => stopAfterWeekend(props.instanceId),
  },
]

const start = (id: string) => {
  startInstance(id)
    .then(() => {
      props.refreshFunction()
    })
    .catch((error) => {
      console.error('Error starting instance:', error)
    })
}

const stop = (id: string) => {
  stopInstance(id)
    .then(() => {
      props.refreshFunction()
    })
    .catch((error) => {
      console.error('Error stopping instance:', error)
    })
}

const stopAfterWeekend = (id: string) => {
  stopInstanceAfterRaceWeekend(id)
    .then(() => {
      props.refreshFunction()
    })
    .catch((error) => {
      console.error('Error stopping instance:', error)
    })
}
</script>

<template>
  <UButton
    icon="i-lucide-play-circle"
    color="primary"
    label="Start"
    @click="start(instanceId)"
    v-if="!isRunning"
    :disabled="!state.mod"
  />
  <UDropdownMenu :items="stopItems" :disabled="!state.mod" v-else>
    <UButton icon="i-lucide-stop-circle" color="error" label="Stop" />
  </UDropdownMenu>
</template>
