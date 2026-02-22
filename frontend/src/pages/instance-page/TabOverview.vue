<script lang="ts" setup>
import type { LiveServerInstancePayload, InstancePayload } from '@/lib/accweb/types'
import { liveInstance } from '@/lib/accweb/client'
import { computed, onMounted, onUnmounted, ref, watch } from 'vue'
import moment from 'moment'
import HistoriesCard from './overview/HistoriesCard.vue'
import GridTable from './overview/GridTable.vue'

const props = defineProps<{ instance: InstancePayload }>()

const emit = defineEmits(['instanceUpdated'])

let intervalId: NodeJS.Timeout

const live = ref<LiveServerInstancePayload | null>(null)

function refreshLiveData(id: string) {
  return liveInstance(id)
    .then((resp) => {
      live.value = resp
    })
    .catch((err) => console.error('Error fetching live instance:', err))
}

function loadLiveInstance(id: string) {
  refreshLiveData(id).then(() => {
    intervalId = setTimeout(() => loadLiveInstance(id), 2000)
  })
}

const liveTs = computed(() => {
  return moment(live.value?.live?.updatedAt).format('MM/DD/YYYY, HH:mm:ss A')
})

const onInit = () => {
  clearInterval(intervalId)
  if (props.instance.is_running) {
    loadLiveInstance(props.instance.id)
  } else {
    clearInterval(intervalId)
    refreshLiveData(props.instance.id)
  }
}

watch(
  () => props.instance,
  () => onInit(),
)

watch(
  () => live.value?.serverState,
  (newValue, oldValue) => {
    if (newValue === 'offline' && oldValue === 'stopping') {
      emit('instanceUpdated')
    }
  },
)

onMounted(() => {
  onInit()
})

onUnmounted(() => {
  clearInterval(intervalId)
})
</script>

<template>
  <div class="flex flex-col gap-4">
    <div class="flex flex-auto flex-row gap-4 border rounded-lg p-3 text-sm">
      <div class="flex-1"><strong>Status:</strong> {{ live?.serverState }}</div>

      <div class="flex-1"><strong>Track:</strong> {{ live?.track }}</div>

      <div class="flex-2">
        <strong>Phase:</strong> {{ live?.sessionType }} ({{ live?.sessionPhase }}) [{{
          live?.sessionRemaining
        }}
        min]
      </div>

      <div class="flex-1"><strong>Drivers:</strong> {{ live?.nrClients }}</div>

      <div class="flex-2"><strong>Last Update:</strong> {{ liveTs }}</div>
    </div>

    <div class="flex-none">
      <div class="flex flex-row gap2">
        <div class="flex-3">
          <div v-if="live" class="flex flex-col gap-1">
            <GridTable :cars="live.live!.cars" />
          </div>
          <div v-else>Not running...</div>
        </div>

        <div class="flex-1">
          <HistoriesCard :live="live" v-if="live" />
        </div>
      </div>
    </div>
  </div>
</template>
