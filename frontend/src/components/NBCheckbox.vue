<script setup lang="ts">
import { ref, watch } from 'vue'

defineProps({
  label: {
    type: String,
    required: true,
  },
  description: {
    type: String,
    default: '',
  },
})

const model = defineModel({ default: 0 })
const boolModel = ref(model.value == 1 ? true : false)

watch(boolModel, (newVal) => {
  model.value = newVal ? 1 : 0
})
</script>

<template>
  <div class="flex items-center gap-2">
    <TCheckbox v-bind="$attrs" :label="label" :description="description" v-model="boolModel" />
  </div>
</template>
