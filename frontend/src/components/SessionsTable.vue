<script setup lang="ts">
import { type SessionSettings } from '@/lib/accweb/types'
import type { TableColumn, SelectItem, DropdownMenuItem } from '@nuxt/ui'
import { ref } from 'vue'

const newTpl: SessionSettings = {
  dayOfWeekend: 1,
  hourOfDay: 9,
  sessionDurationMinutes: 1,
  sessionType: 'P',
  timeMultiplier: 1,
}

const addQRTpl: SessionSettings[] = [
  {
    dayOfWeekend: 2,
    hourOfDay: 9,
    sessionDurationMinutes: 15,
    sessionType: 'Q',
    timeMultiplier: 1,
  },
  {
    dayOfWeekend: 3,
    hourOfDay: 9,
    sessionDurationMinutes: 30,
    sessionType: 'R',
    timeMultiplier: 1,
  },
]

const addPQRTpl: SessionSettings[] = [
  {
    dayOfWeekend: 1,
    hourOfDay: 9,
    sessionDurationMinutes: 7,
    sessionType: 'P',
    timeMultiplier: 1,
  },
  {
    dayOfWeekend: 2,
    hourOfDay: 9,
    sessionDurationMinutes: 15,
    sessionType: 'Q',
    timeMultiplier: 1,
  },
  {
    dayOfWeekend: 3,
    hourOfDay: 9,
    sessionDurationMinutes: 30,
    sessionType: 'R',
    timeMultiplier: 1,
  },
]

const model = defineModel<SessionSettings[]>()

const sessionTypes = ref<SelectItem[]>([
  {
    label: 'Pratice',
    value: 'P',
  },
  {
    label: 'Qualifying',
    value: 'Q',
  },
  {
    label: 'Race',
    value: 'R',
  },
])

const daysOfWeek = ref<SelectItem[]>([
  {
    label: 'Friday',
    value: 1,
  },
  {
    label: 'Saturday',
    value: 2,
  },
  {
    label: 'Sunday',
    value: 3,
  },
])

// Columns definition only for headers and keys
const columns: TableColumn<SessionSettings>[] = [
  { accessorKey: 'sessionType', header: 'Type' },
  { accessorKey: 'dayOfWeekend', header: 'Day of Weekend' },
  { accessorKey: 'hourOfDay', header: 'Hour of Day' },
  { accessorKey: 'sessionDurationMinutes', header: 'Duration' },
  { accessorKey: 'timeMultiplier', header: 'Time Multiplier' },
  { accessorKey: 'actions', header: '' },
]

function getRowItems(index: number): DropdownMenuItem[] {
  return [
    {
      type: 'label',
      label: 'Actions',
    },
    {
      label: 'Move Up',
      icon: 'i-lucide-move-up',
      disabled: index === 0,
      onSelect() {
        if (model.value && index > 0) {
          const sessions = model.value
          ;[sessions[index], sessions[index - 1]] = [sessions[index - 1]!, sessions[index]!]
        }
      },
    },
    {
      label: 'Move Down',
      icon: 'i-lucide-move-down',
      disabled: !model.value || index >= model.value.length - 1,
      onSelect() {
        if (model.value && index < model.value.length - 1) {
          const sessions = model.value
          ;[sessions[index], sessions[index + 1]] = [sessions[index + 1]!, sessions[index]!]
        }
      },
    },
    {
      label: 'Duplicate',
      icon: 'i-lucide-copy',
      onSelect() {
        if (model.value) {
          const duplicated = JSON.parse(JSON.stringify(model.value[index]))
          model.value.splice(index, 0, duplicated)
        }
      },
    },
    {
      type: 'separator',
    },
    {
      label: 'Delete',
      color: 'error',
      icon: 'i-lucide-trash',
      onSelect() {
        if (model.value) {
          model.value.splice(index, 1)
        }
      },
    },
  ]
}
</script>

<template>
  <UTable ref="table" :data="model" :columns="columns" :ui="{ td: 'p-1' }">
    <template #sessionType-cell="{ row }">
      <USelect
        v-model="model![row.index]!.sessionType"
        size="sm"
        :items="sessionTypes"
        class="w-28"
      />
    </template>

    <template #dayOfWeekend-cell="{ row }">
      <USelect
        v-model="model![row.index]!.dayOfWeekend"
        size="sm"
        :items="daysOfWeek"
        class="w-26"
      />
    </template>

    <template #hourOfDay-cell="{ row }">
      <TFormField :name="`acc.event.sessions.${row.index}.hourOfDay`">
        <UInput
          v-model="model![row.index]!.hourOfDay"
          type="number"
          min="0"
          size="sm"
          max="23"
          class="w-16"
        />
      </TFormField>
    </template>

    <template #sessionDurationMinutes-cell="{ row }">
      <TFormField :name="`acc.event.sessions.${row.index}.sessionDurationMinutes`">
        <UInput
          v-model="model![row.index]!.sessionDurationMinutes"
          type="number"
          min="0"
          size="sm"
          class="w-16"
        />
      </TFormField>
    </template>

    <template #timeMultiplier-cell="{ row }">
      <TFormField :name="`acc.event.sessions.${row.index}.timeMultiplier`">
        <UInput
          v-model="model![row.index]!.timeMultiplier"
          size="sm"
          type="number"
          min="0"
          class="w-16"
        />
      </TFormField>
    </template>

    <template #actions-cell="{ row }">
      <UDropdownMenu
        :items="getRowItems(row.index)"
        :content="{ align: 'end' }"
        aria-label="Actions dropdown"
      >
        <UButton
          icon="i-lucide-ellipsis-vertical"
          color="neutral"
          variant="ghost"
          aria-label="Actions dropdown"
        />
      </UDropdownMenu>
    </template>
  </UTable>

  <div class="flex gap-2 mt-5">
    <UButton
      label="Add"
      icon="i-lucide-plus"
      @click="model?.push(newTpl)"
      color="secondary"
      variant="subtle"
      class="flex-none"
    />

    <UButton
      label="Add (Q/R)"
      icon="i-lucide-plus"
      @click="model?.push(...addQRTpl)"
      color="secondary"
      variant="subtle"
      class="flex-none"
    />

    <UButton
      label="Add (P/Q/R)"
      icon="i-lucide-plus"
      @click="model?.push(...addPQRTpl)"
      color="secondary"
      variant="subtle"
      class="flex-none"
    />

    <ConfirmDialog
      title="Confirmation"
      content="Do you really want to clear all sessions?"
      :exec="
        () => {
          model = []
        }
      "
    >
      <UButton
        label="Clear all"
        icon="i-lucide-trash"
        color="error"
        variant="subtle"
        class="flex-none"
      />
    </ConfirmDialog>
  </div>
</template>
