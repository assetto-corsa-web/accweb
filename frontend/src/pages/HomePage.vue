<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue'
import { fetchServer } from '@/lib/accweb/client'
import type { ListServerItem } from '@/lib/accweb/types'
import InstanceCard from './home-page/InstanceCard.vue'
import _ from 'lodash'
import { useAuthStore } from '@/stores/auth'

const servers = ref<ListServerItem[]>([])
const state = useAuthStore()

let intervalId: NodeJS.Timeout

const loading = ref(false)

onMounted(async () => {
  refreshServers()

  // intervalId = setInterval(refreshServers, 10000)
})

onUnmounted(() => {
  clearInterval(intervalId)
})

const refreshServers = () => {
  if (loading.value) {
    return
  }

  loading.value = true
  fetchServer()
    .then((serversResp) => {
      servers.value = serversResp
      intervalId = setTimeout(refreshServers, 10000)
    })
    .catch((error) => {
      console.error('Error fetching servers:', error)
    })
    .finally(() => {
      loading.value = false
    })
}

function refreshNow() {
  loading.value = true
  fetchServer()
    .then((serversResp) => {
      servers.value = serversResp
    })
    .catch((error) => {
      console.error('Error fetching servers:', error)
    })
    .finally(() => {
      loading.value = false
    })
}

const sortedServerList = computed<ListServerItem[]>(() => {
  return _.sortBy(servers.value, ['name'])
})
</script>

<template>
  <div class="flex flex-col gap-3">
    <div class="flex-none flex flex-row mb-5">
      <div class="flex-auto">.</div>
      <div class="flex-none">
        <UButton to="/instance-new" icon="i-lucide-plus" label="Create New" v-if="state.admin" />
      </div>
    </div>

    <div class="flex-auto">
      <div v-if="servers.length">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3">
          <InstanceCard
            v-for="server in sortedServerList"
            :key="server.id + '-' + server.isRunning"
            :instance="server"
            :refreshFunction="refreshNow"
          />
        </div>
      </div>
      <div v-else>
        <p>No servers found.</p>
      </div>
    </div>
  </div>
</template>
