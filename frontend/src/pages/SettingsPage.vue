<script setup lang="ts">
import { ref } from 'vue'
import type { TabsItem } from '@nuxt/ui'
import TabGeneral from './settings/TabGeneral.vue'
import TabGlobalEntrylist from './settings/TabGlobalEntrylist.vue'
import TabGlobalBanlist from './settings/TabGlobalBanlist.vue'
import TabAccessControl from './settings/TabAccessControl.vue'

const tabItems = ref<TabsItem[]>([
  {
    label: 'General',
    icon: 'i-lucide-settings',
    content: 'General Settings',
    slot: 'general',
  },
  {
    label: 'Access Control',
    icon: 'i-lucide-users',
    content: 'Access control settings',
    slot: 'access-control',
  },
  {
    label: 'Global Entry List',
    icon: 'i-lucide-notebook-pen',
    content: 'Global Entry List settings',
    slot: 'global-entry-list',
  },
  {
    label: 'Global Ban',
    icon: 'i-lucide-ban',
    content: 'Global Ban settings',
    slot: 'global-ban',
  },
])
</script>

<template>
  <UTabs :items="tabItems" variant="link" class="w-full gap-4">
    <template #general>
      <TabGeneral />
    </template>

    <template #access-control>
      <TabAccessControl />
    </template>

    <template #global-entry-list>
      <TabGlobalEntrylist />
    </template>

    <template #global-ban>
      <TabGlobalBanlist />
    </template>
  </UTabs>
</template>
