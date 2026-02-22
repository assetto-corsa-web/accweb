<script setup lang="ts">
import type { CarState, DriverState } from '@/lib/accweb/types'
import type { TableColumn } from '@nuxt/ui'
import { computed } from 'vue'
import _ from 'lodash'
import moment from 'moment'
import type { Row, TableMeta } from '@tanstack/vue-table'

const { cars } = defineProps<{ cars: Record<string, CarState> }>()

const data = computed(() => {
  const rs = _.orderBy(_.values(cars), ['position'])

  // console.log('GridTable data', rs)

  return _.filter(rs, (o) => {
    return o.currentDriver !== null
  })
})

function parseLapMs(ms: number): string {
  if (ms <= 0) {
    return '--'
  }

  return moment(ms).format('mm:ss.SSS')
}

const columns: TableColumn<CarState>[] = [
  {
    accessorKey: 'position',
    header: '#',
    cell: ({ row }) => row.index + 1,
  },
  {
    accessorKey: 'currentDriver',
    header: 'Driver',
    cell: ({ row }) => {
      const currentDriver: DriverState = row.getValue('currentDriver')
      return currentDriver?.name || '--'
    },
  },
  {
    accessorKey: 'raceNumber',
    header: 'Nr',
  },
  {
    accessorKey: 'nrLaps',
    header: 'Laps',
  },
  {
    accessorKey: 'fuel',
    header: 'Fuel',
  },
  {
    accessorKey: 'bestLapMS',
    header: 'Best Lap',
    cell: ({ row }) => parseLapMs(row.getValue('bestLapMS')),
  },
  {
    accessorKey: 'lastLapMS',
    header: 'Last Lap',
    cell: ({ row }) => parseLapMs(row.getValue('lastLapMS')),
  },
  {
    accessorKey: 'currLap.s1',
    header: 'S1',
  },
  {
    accessorKey: 'currLap.s2',
    header: 'S2',
  },
  {
    accessorKey: 'currLap.s3',
    header: 'S3',
  },
  {
    accessorKey: 'gap',
    header: 'Gap',
    cell: ({ row }) => calcGap(row.index),
  },
  {
    accessorKey: 'flags',
    header: 'Flags',
  },
]

function calcGap(idx: number): string {
  if (idx === 0) {
    return ''
  }

  const curr = data.value[idx]
  const prev = data.value[idx - 1]

  if (curr?.lastLapTimestampMS === 0) {
    return ''
  }

  if (prev?.nrLaps !== curr?.nrLaps) {
    const diff = data.value[0]!.nrLaps - curr!.nrLaps
    return `+${diff} lap${diff > 1 ? 's' : ''}`
  }

  const gap = curr!.lastLapTimestampMS - prev!.lastLapTimestampMS

  if (gap < 0) {
    return '--'
  }

  return parseLapMs(gap)
}

const meta: TableMeta<CarState> = {
  class: {
    tr: (row: Row<CarState>) => {
      if (row.original.currLap.flags > 0) {
        return 'bg-error/10'
      }

      return ''
    },
  },
}
</script>

<template>
  <UTable :data="data" :columns="columns" :meta="meta" sticky :ui="{ td: 'p-1' }">
    <template #flags-cell="{ row }">
      <div class="flex gap-1">
        <UIcon name="i-lucide-scissors" class="" v-if="row.original.currLap.hasCut" />
        <UIcon name="i-lucide-log-in" class="" v-if="row.original.currLap.inLap" />
        <UIcon name="i-lucide-log-out" class="" v-if="row.original.currLap.outLap" />
        <UIcon name="i-ph-flag-checkered-fill" class="" v-if="row.original.currLap.sessionOver" />
      </div>
    </template>
  </UTable>
</template>
