<template>
  <UModal :title="title" :description="content">
    <slot />

    <template #footer="{ close }">
      <div class="flex gap-2 w-full">
        <div class="flex-none"><UButton label="Confirm" @click="success(close)" /></div>
        <div class="flex-auto" />
        <div class="flex-none">
          <UButton color="neutral" label="Dismiss" @click="close()" />
        </div>
      </div>
    </template>
  </UModal>
</template>

<script setup lang="ts">
interface Props {
  title?: string
  content?: string
  exec: () => void
}

const props = defineProps<Props>()

function success(clFn: () => void) {
  props.exec()
  clFn()
}
</script>
