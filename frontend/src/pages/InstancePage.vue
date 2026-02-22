<script setup lang="ts">
import { fetchServer, getInstance, liveInstance } from '@/lib/accweb/client'
import type { InstancePayload } from '@/lib/accweb/types'
import type { TabsItem } from '@nuxt/ui'
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import type { SelectMenuItem } from '@nuxt/ui'
import _ from 'lodash'
import TabOverview from './instance-page/TabOverview.vue'
import TabDefinitions from './instance-page/TabDefinitions.vue'
import TabResults from './instance-page/TabResults.vue'

const route = useRoute()
const router = useRouter()

const instanceItems = ref<SelectMenuItem[]>([])
const instanceId = ref<string>(route.params.id as string)
const instance = ref<InstancePayload | null>(null)

function loadInstance(id: string) {
  getInstance(id)
    .then((resp) => {
      instance.value = resp
    })
    .catch((err) => console.error('Error fetching instance:', err))
}

function loadInstanceList() {
  fetchServer()
    .then((resp) => {
      const items = resp.map((server) => ({
        id: server.id,
        label: server.name,
      }))

      instanceItems.value = _.sortBy(items, (item) => item.label)
    })
    .catch((error) => {
      console.error('Error fetching servers:', error)
    })
}

onMounted(() => {
  loadInstance(route.params.id as string)
  loadInstanceList()

  liveInstance(route.params.id as string).catch((error) => {
    console.error('Error starting live instance:', error)
  })

  document.title = 'Account Settings - ACC Web'
})

watch(
  () => route.params.id,
  (id) => {
    if (id) {
      instanceId.value = id as string
      loadInstance(id as string)
    }
  },
)

const tabItems = ref<TabsItem[]>([
  {
    label: 'Overview',
    icon: 'i-lucide-monitor',
    content: 'This is the account content.',
    slot: 'overview',
  },
  {
    label: 'Definitions',
    icon: 'i-lucide-settings',
    content: 'This is the password content.',
    slot: 'definitions',
  },
  {
    label: 'Results',
    icon: 'i-lucide-notebook-pen',
    content: 'This is the results content.',
    slot: 'results',
  },
  {
    label: 'Statistics',
    icon: 'i-lucide-bar-chart-4',
    content: 'This is the statistics content.',
    slot: 'statistics',
  },
])
</script>

<template>
  <div class="flex gap-4">
    <div class="flex-auto">
      <USelectMenu
        v-model="instanceId"
        size="xl"
        value-key="id"
        :items="instanceItems"
        class="w-full mb-5 hover:border"
        variant="none"
        @change="router.push({ path: `/instance/${instanceId}` })"
      />
    </div>
    <div class="flex-none">
      <InstanceStartStopBtn
        v-if="instance"
        :instanceId="instanceId"
        :isRunning="instance.is_running"
        :refreshFunction="() => loadInstance(instanceId)"
      />
    </div>
  </div>

  <UTabs :items="tabItems" color="secondary" :unmount-on-hide="false" class="w-full gap-4">
    <template #overview v-if="instance">
      <TabOverview :instance="instance" @instance-updated="loadInstance(instanceId)" />
    </template>

    <template #definitions v-if="instance">
      <TabDefinitions :instance="instance" @instance-updated="loadInstanceList()" />
    </template>

    <template #results v-if="instance">
      <TabResults :instance="instance" />
    </template>
  </UTabs>
</template>
