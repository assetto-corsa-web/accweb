<script setup lang="ts">
import { saveInstance } from '@/lib/accweb/client'
import type { InstancePayload } from '@/lib/accweb/types'
import type { FormError, TabsItem } from '@nuxt/ui'
import { ref, watch, computed } from 'vue'
import { schema } from '@/lib/accweb/schema'
import DefAccweb from './definitions/DefAccweb.vue'
import DefSettings from './definitions/DefSettings.vue'
import DefGeneral from './definitions/DefGeneral.vue'
import DefEvent from './definitions/DefEvent.vue'
import DefEventRules from './definitions/DefEventRules.vue'
import DefAssists from './definitions/DefAssists.vue'
import DefEntrylist from './definitions/DefEntrylist.vue'
import DefBop from './definitions/DefBop.vue'
import _ from 'lodash'

const toast = useToast()

const props = defineProps<{
  instance: InstancePayload
  overrideOnSubmit?: (data: InstancePayload) => void
}>()
const data = ref<InstancePayload>(_.cloneDeep(props.instance))
const snapshot = ref<InstancePayload>(_.cloneDeep(props.instance))

const emit = defineEmits(['instanceUpdated'])

const isDirty = computed(() => !_.isEqual(data.value, snapshot.value))

watch(
  () => props.instance,
  (newValue) => {
    if (newValue.id != data.value.id) {
      data.value = _.cloneDeep(newValue)
      snapshot.value = _.cloneDeep(newValue)
    }
  },
)

const tabItems = ref<TabsItem[]>([
  {
    label: 'ACCWeb',
    content: 'This is the accweb content.',
    slot: 'accweb',
  },
  {
    label: 'General',
    content: 'This is the general content.',
    slot: 'general',
  },
  {
    label: 'Settings',
    content: 'This is the general content.',
    slot: 'settings',
  },
  {
    label: 'Event',
    content: 'This is the statistics content.',
    slot: 'event',
  },
  {
    label: 'Rules',
    content: 'This is the general content.',
    slot: 'rules',
  },
  {
    label: 'Entry List',
    content: 'This is the general content.',
    slot: 'entrylist',
  },
  {
    label: 'BoP',
    content: 'This is the bop content.',
    slot: 'bop',
  },
  {
    label: 'Assists',
    content: 'This is the general content.',
    slot: 'assists',
  },
])

function addBadge(slot: string) {
  const item = _.find(tabItems.value, { slot: slot })
  if (item) {
    item.badge = { color: 'error', variant: 'solid' }
  }
}

function clearBadges() {
  _.forEach(tabItems.value, (item) => {
    item.badge = undefined
  })
}

function validate(): FormError[] {
  const errors: FormError[] = []
  const result = schema.safeParse(data.value)

  clearBadges()

  if (result.success) {
    return []
  }

  _.forEach(result.error.issues, (err) => {
    addBadge(err.path[1] as string)
    errors.push({ name: err.path.join('.'), message: err.message })
  })

  return errors
}

async function defaultOnSubmit() {
  saveInstance(data.value)
    .then(() => {
      snapshot.value = _.cloneDeep(data.value)
      toast.add({ title: 'Success', description: 'The form has been submitted.', color: 'success' })
      emit('instanceUpdated')
    })
    .catch((err) => {
      toast.add({
        title: 'Failed to save',
        description: err.response.data.error,
        color: 'error',
      })
    })
}

async function onSubmit() {
  if (props.overrideOnSubmit) {
    props.overrideOnSubmit(data.value)
  } else {
    defaultOnSubmit()
  }
}
</script>

<template>
  <div v-if="data">
    <UForm :validate="validate" :state="data" class="space-y-4" @submit="onSubmit">
      <div class="mb-4 flex flex-row gap-4">
        <div class="flex-auto">
          <UFormField label="Server name" name="acc.settings.serverName">
            <UInput v-model="data.acc.settings.serverName" class="w-full" />
          </UFormField>
        </div>

        <div class="flex-none">
          <UButton
            :active="!instance.is_running && isDirty"
            active-color="primary"
            color="neutral"
            type="submit"
            class="mt-6"
            :disabled="instance.is_running || !isDirty"
            >Save Changes</UButton
          >
        </div>
      </div>

      <UTabs
        :items="tabItems"
        class="w-full gap-4"
        variant="link"
        color="warning"
        :unmount-on-hide="false"
        size="xl"
        :ui="{}"
      >
        <template #accweb>
          <DefAccweb v-model="data.accWeb" />
        </template>

        <template #general>
          <DefGeneral v-model="data.acc.configuration" />
        </template>

        <template #settings>
          <DefSettings v-model="data.acc.settings" />
        </template>

        <template #event>
          <DefEvent v-model="data.acc.event" />
        </template>

        <template #rules>
          <DefEventRules v-model="data.acc.eventRules" />
        </template>

        <template #entrylist>
          <DefEntrylist v-model="data.acc.entrylist" />
        </template>

        <template #bop>
          <DefBop v-model="data.acc.bop" :current-track="data.acc.event.track" />
        </template>

        <template #assists>
          <DefAssists v-model="data.acc.assistRules" />
        </template>
      </UTabs>
    </UForm>
  </div>

  <div v-else>Loading...</div>
</template>
