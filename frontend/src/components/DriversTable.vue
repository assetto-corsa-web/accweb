<script setup lang="ts">
import { type DriverSettings } from '@/lib/accweb/types'
import type { TableColumn, SelectItem, DropdownMenuItem } from '@nuxt/ui'
import { ref } from 'vue'
import nationalities from '@/data/nationalities'

const model = defineModel<DriverSettings[]>()
const props = defineProps<{ index: number; hidePrefix: boolean }>()
const driverCategories = ref<SelectItem[]>(['Bronze', 'Silver', 'Gold', 'Platinum'])
const nationalitiesItems = ref<SelectItem[]>(nationalities)

// Columns definition only for headers and keys
const columns: TableColumn<DriverSettings>[] = [
  { accessorKey: 'firstName', header: 'First Name' },
  { accessorKey: 'lastName', header: 'Last Name' },
  { accessorKey: 'shortName', header: 'Short' },
  { accessorKey: 'nationality', header: 'Nationality' },
  { accessorKey: 'driverCategory', header: 'Driver Category' },
  { accessorKey: 'playerID', header: 'Player ID' },
  { accessorKey: 'actions', header: '' },
]

const prefix = props.hidePrefix ? '' : 'acc.entrylist.'

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
  <UTable ref="table" :data="model" :columns="columns" :ui="{ td: 'p-1', th: 'p-1 py-4' }">
    <template #firstName-cell="{ row }">
      <TFormField :name="`${prefix}entries.${index}.drivers.${row.index}.firstName`">
        <UInput v-model="model![row.index]!.firstName" size="sm" class="w-18" />
      </TFormField>
    </template>

    <template #lastName-cell="{ row }">
      <TFormField :name="`${prefix}entries.${index}.drivers.${row.index}.lastName`">
        <UInput v-model="model![row.index]!.lastName" size="sm" class="w-18" />
      </TFormField>
    </template>

    <template #shortName-cell="{ row }">
      <TFormField :name="`${prefix}entries.${index}.drivers.${row.index}.shortName`">
        <UInput v-model="model![row.index]!.shortName" size="sm" class="w-16" />
      </TFormField>
    </template>

    <template #nationality-cell="{ row }">
      <USelect
        v-model="model![row.index]!.nationality"
        size="sm"
        :items="nationalitiesItems"
        class="w-26 mt-1"
      />
    </template>

    <template #driverCategory-cell="{ row }">
      <USelect
        v-model="model![row.index]!.driverCategory"
        :items="driverCategories"
        size="sm"
        class="w-26 mt-1"
      />
    </template>

    <template #playerID-cell="{ row }">
      <TFormField :name="`${prefix}entries.${index}.drivers.${row.index}.playerID`">
        <UInput v-model="model![row.index]!.playerID" size="sm" class="w-36" />
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
      @click="model?.push({ playerID: '' } as DriverSettings)"
      color="secondary"
      variant="subtle"
      class="flex-none"
    />

    <ConfirmDialog
      title="Confirmation"
      content="Do you really want to clear all drivers?"
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
