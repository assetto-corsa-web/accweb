<script setup lang="ts">
import cars from '@/data/cars'
import tracks from '@/data/tracks'
import type { BopJson, BopSettings } from '@/lib/accweb/types'
import type { DropdownMenuItem, SelectItem, TableColumn } from '@nuxt/ui'
import { ref } from 'vue'
import _ from 'lodash'

const model = defineModel({
  required: true,
  default: {} as BopJson,
})

const props = defineProps<{
  currentTrack?: string
}>()

const tracksItems = ref<SelectItem[]>(tracks)
const carModelItems = ref<SelectItem[]>(
  _.map(_.sortBy(cars, 'brand'), (c) => {
    return { label: c.model, value: c.id }
  }),
)

// Columns definition only for headers and keys
const columns: TableColumn<BopSettings>[] = [
  { accessorKey: 'track', header: 'Track' },
  { accessorKey: 'carModel', header: 'Car Model' },
  { accessorKey: 'ballastKg', header: 'Ballast (kg)' },
  { accessorKey: 'restrictor', header: 'Restrictor (%)' },
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
          const sessions = model.value.entries
          ;[sessions[index], sessions[index - 1]] = [sessions[index - 1]!, sessions[index]!]
        }
      },
    },
    {
      label: 'Move Down',
      icon: 'i-lucide-move-down',
      disabled: !model.value || index >= model.value.entries.length - 1,
      onSelect() {
        if (model.value && index < model.value.entries.length - 1) {
          const sessions = model.value.entries
          ;[sessions[index], sessions[index + 1]] = [sessions[index + 1]!, sessions[index]!]
        }
      },
    },
    {
      label: 'Duplicate',
      icon: 'i-lucide-copy',
      onSelect() {
        if (model.value) {
          const duplicated = JSON.parse(JSON.stringify(model.value.entries[index]))
          model.value.entries.splice(index, 0, duplicated)
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
          model.value.entries.splice(index, 1)
        }
      },
    },
  ]
}

function addEntry() {
  if (!model.value.entries) {
    model.value.entries = []
  }

  model.value.entries.push({
    track: props.currentTrack,
    carModel: -1,
    ballastKg: 0,
    restrictor: 0,
  } as BopSettings)
}
</script>

<template>
  <UTable :data="model.entries" :columns="columns" :ui="{ root: 'w-fit' }">
    <template #track-cell="{ row }">
      <TFormField :name="`acc.bop.entries.${row.index}.track`">
        <USelect :items="tracksItems" v-model="model.entries[row.index]!.track" class="w-62" />
      </TFormField>
    </template>

    <template #carModel-cell="{ row }">
      <TFormField :name="`acc.bop.entries.${row.index}.carModel`">
        <USelect :items="carModelItems" v-model="model.entries[row.index]!.carModel" class="w-62" />
      </TFormField>
    </template>

    <template #ballastKg-cell="{ row }">
      <TFormField :name="`acc.bop.entries.${row.index}.ballastKg`">
        <UInput v-model="model.entries[row.index]!.ballastKg" type="number" min="0" max="100" />
      </TFormField>
    </template>

    <template #restrictor-cell="{ row }">
      <TFormField :name="`acc.bop.entries.${row.index}.restrictor`">
        <UInput v-model="model.entries[row.index]!.restrictor" type="number" min="0" max="20" />
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

  <div class="flex-auto flex flex-row gap-2 mt-5">
    <UButton
      label="Add"
      icon="i-lucide-plus"
      @click="addEntry()"
      color="secondary"
      variant="subtle"
      class="flex-none"
    />

    <ConfirmDialog
      title="Confirmation"
      content="Do you really want to clear all entries?"
      :exec="
        () => {
          model.entries = []
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
