<script setup lang="ts">
import AppHeader from '@/components/AppHeader.vue'
</script>

<template>
  <UApp>
    <AppHeader />

    <UMain>
      <UContainer>
        <div class="pt-3 pb-6">
          <RouterView />
        </div>
      </UContainer>
    </UMain>
  </UApp>
</template>
