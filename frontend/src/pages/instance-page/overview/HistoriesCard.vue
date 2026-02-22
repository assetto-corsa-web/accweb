<script setup lang="ts">
import { computed } from 'vue'
import moment from 'moment'
import _ from 'lodash'
import type {
  HistoryChat,
  HistoryDamage,
  HistoryNewConnection,
  HistoryRemoveConnection,
  HistorySession,
  LiveServerInstancePayload,
} from '@/lib/accweb/types'

const { live } = defineProps<{ live: LiveServerInstancePayload }>()

const histories = computed(() => {
  return _.orderBy(live?.live?.history || [], ['ts'], ['desc']).slice(0, 50)
})
</script>

<template>
  <UCard variant="subtle" :ui="{ header: 'p-2 sm:px-3', body: 'p2 sm:p-3' }">
    <template #header>
      <div class="font-bold">History</div>

      <div class="text-sm mt-4 flex flex-col gap-2">
        <div v-for="item in histories" :key="item.id" class="flex flex-col">
          <div v-if="item.type == 'session'" class="flex-none text-yellow-200">
            <div class="inline text-xs">{{ moment(item.ts).format('HH:mm:ss') }}</div>
            <div class="inline font-bold ml-2">
              {{ (item.data as HistorySession).sessionType }}:
            </div>
            <div class="inline">
              {{ (item.data as HistorySession).sessionPhase }} ({{
                (item.data as HistorySession).sessionRemaining
              }}m)
            </div>
          </div>

          <div v-else-if="item.type == 'chat'" class="flex-none">
            <div class="inline text-xs">{{ moment(item.ts).format('HH:mm:ss') }}</div>
            <div class="inline font-bold mx-2">{{ (item.data as HistoryChat).name }}:</div>
            <div class="inline">{{ (item.data as HistoryChat).message }}</div>
          </div>

          <div v-else-if="item.type == 'new_connection'" class="flex-none text-green-400">
            <div class="inline text-xs">{{ moment(item.ts).format('HH:mm:ss') }}</div>
            <div class="inline font-bold mx-2">Player Joined:</div>
            <div class="inline">{{ (item.data as HistoryNewConnection).name }}</div>
          </div>

          <div v-else-if="item.type == 'remove_connection'" class="flex-none text-red-300">
            <div class="inline text-xs">{{ moment(item.ts).format('HH:mm:ss') }}</div>
            <div class="inline font-bold mx-2">Player Left:</div>
            <div class="inline">{{ (item.data as HistoryRemoveConnection).name }}</div>
          </div>

          <div v-else-if="item.type == 'damage'" class="flex-none text-stone-400">
            <div class="inline text-xs">{{ moment(item.ts).format('HH:mm:ss') }}</div>
            <div class="inline font-bold ml-2 mr-1">{{ (item.data as HistoryDamage).name }}:</div>
            <div class="inline">Damage Report #{{ (item.data as HistoryDamage).raceNumber }}</div>
          </div>
        </div>
      </div>
    </template>
  </UCard>
</template>
