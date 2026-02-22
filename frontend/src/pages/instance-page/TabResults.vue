<script lang="ts" setup>
import type { InstancePayload, InstanceResultsPayload } from '@/lib/accweb/types'
import { getInstanceResults } from '@/lib/accweb/client'
import { onMounted, ref, watch } from 'vue'

const props = defineProps<{ instance: InstancePayload }>()

const results = ref<InstanceResultsPayload | null>(null)

function onInit() {
  getInstanceResults(props.instance.id)
    .then((res: InstanceResultsPayload) => {
      results.value = res
    })
    .catch((err) => {
      results.value = { results: [] } as InstanceResultsPayload
      console.error('Error fetching instance results:', err)
    })
}

onMounted(() => {
  onInit()
})

watch(
  () => props.instance,
  (id) => {
    if (id) {
      onInit()
    }
  },
)
</script>

<template>
  <div class="p-4">
    <div v-if="results">
      <h2 class="text-xl font-bold mb-4">Race Results</h2>
      <table class="table-auto w-full border-collapse border border-gray-300">
        <thead>
          <tr>
            <th class="border border-gray-300 px-4 py-2">File</th>
          </tr>
        </thead>

        <tbody>
          <tr v-for="file in results.results" :key="file">
            <td class="border border-gray-300 px-4 py-2">
              {{ file }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div v-else>
      <p>Loading results...</p>
    </div>
  </div>
</template>
