<script setup lang="ts">
import type { AccWebAdvWindowsSettingsJson, AccWebSettingsJson } from '@/lib/accweb/types'

const model = defineModel<AccWebSettingsJson>({
  required: true,
  default: {
    advWindowsCfg: {} as AccWebAdvWindowsSettingsJson,
    autoStart: false,
    enableGlobalEntrylist: true,
    enableGlobalBanlist: true,
    enableAdvWindowsCfg: false,
  } as AccWebSettingsJson,
})
</script>

<template>
  <div class="flex flex-col gap-3">
    <TCheckbox
      v-model="model.autoStart"
      label="Server instance auto start."
      description="Server instance will start automatically when accweb starts."
    />

    <TCheckbox
      v-model="model.enableGlobalEntrylist"
      label="Enable global entry list."
      description="Enable Global Entry list on this server instance."
    />

    <TCheckbox
      v-model="model.enableGlobalBanlist"
      label="Enable global ban list."
      description="Enable Global Ban list on this server instance."
    />
  </div>
</template>
