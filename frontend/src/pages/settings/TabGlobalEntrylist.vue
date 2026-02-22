<script setup lang="ts">
import EntryList from '@/components/EntryList.vue'
import { accwebGlobalEntrylist } from '@/lib/accweb/schema'
import type { AccwebGlobalEntrylistJson, EntrylistJson } from '@/lib/accweb/types'
import type { FormError } from '@nuxt/ui'
import { onMounted, ref } from 'vue'
import _ from 'lodash'
import { getGlobalEntryList, saveGlobalEntryList } from '@/lib/accweb/client'

const data = ref<AccwebGlobalEntrylistJson>()
const entrylist = ref<EntrylistJson>()
const toast = useToast()

onMounted(() => {
  getGlobalEntryList().then((res) => {
    data.value = res
    entrylist.value = { entries: res.entries } as EntrylistJson
  })
})

function validate(): FormError[] {
  const errors: FormError[] = []
  const result = accwebGlobalEntrylist.safeParse(data.value)

  if (result.success) {
    return []
  }

  _.forEach(result.error.issues, (err) => {
    errors.push({ name: err.path.join('.'), message: err.message })
  })

  return errors
}

async function onSubmit() {
  if (!data.value || !entrylist.value) {
    return
  }

  data.value.entries = entrylist.value.entries

  saveGlobalEntryList(data.value)
    .then(() => {
      toast.add({ title: 'Success', description: 'The form has been submitted.', color: 'success' })
    })
    .catch((err) => {
      toast.add({
        title: 'Failed to save',
        description: err.response.data.error,
        color: 'error',
      })
    })
}
</script>

<template>
  <UForm
    :validate="validate"
    :state="data"
    class="space-y-4"
    @submit="onSubmit"
    v-if="data && entrylist"
  >
    <div class="flex flex-col gap-3">
      <p>Share a list of entries across all your servers that have this feature enabled.</p>

      <TCheckbox
        v-model="data.enabled"
        label="Enable Global Entry List"
        description="When checked, add this entrylist to all servers."
      />

      <UButton type="submit" label="Save list" icon="i-lucide-save" class="flex-none w-fit" />

      <USeparator />

      <EntryList v-model="entrylist" :hide-force-entry-list-checkbox="true" />
    </div>
  </UForm>
</template>
