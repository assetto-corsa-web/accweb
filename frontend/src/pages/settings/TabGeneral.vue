<script setup lang="ts">
import { ref } from 'vue'

const data = ref(null)
async function onSubmit() {}
</script>

<template>
  <UForm :state="data" class="space-y-4" @submit="onSubmit"></UForm>
</template>
