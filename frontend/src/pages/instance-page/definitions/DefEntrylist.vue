<script setup lang="ts">
import EntryList from '@/components/EntryList.vue'
import type { EntrylistJson, EntrySettings } from '@/lib/accweb/types'

const model = defineModel({
  required: true,
  default: {
    entries: [] as EntrySettings[],
    forceEntryList: 0,
  } as EntrylistJson,
})
</script>

<template>
  <EntryList v-model="model" />
</template>
