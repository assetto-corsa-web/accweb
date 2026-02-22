<script setup lang="ts">
import DriversTable from '@/components/DriversTable.vue'
import cars from '@/data/cars'
import type { DriverSettings, EntrylistJson, EntrySettings } from '@/lib/accweb/types'
import type { SelectItem } from '@nuxt/ui'
import _ from 'lodash'
import { ref } from 'vue'

const model = defineModel({
  required: true,
  default: {
    entries: [] as EntrySettings[],
    forceEntryList: 0,
  } as EntrylistJson,
})

const props = defineProps<{
  hideForceEntryListCheckbox?: boolean
}>()

const carItems = ref<SelectItem[]>(
  _.map(_.sortBy(cars, 'brand'), (c) => {
    return { label: c.model, value: c.id }
  }),
)

const prefix = props.hideForceEntryListCheckbox ? '' : 'acc.entrylist.'

function addEntry() {
  if (!model.value.entries) {
    model.value.entries = []
  }

  model.value.entries.push({
    forcedCarModel: -1,
    ballastKg: 0,
    restrictor: 0,
    drivers: [{ playerID: '' } as DriverSettings] as DriverSettings[],
  } as EntrySettings)
}
</script>

<template>
  <div class="flex flex-col gap-3">
    <NBCheckbox
      v-if="!hideForceEntryListCheckbox"
      v-model="model.forceEntryList"
      label="Force Entry List"
      description="Will reject drivers that are not in the entry list. Default is
0, which allows the partial definition of entries in a
“normal” server configuration. Cannot be used on public
servers."
    />

    <div class="flex-auto gap-3">
      <div class="flex-auto" v-for="(entry, idx) in model.entries" :key="idx">
        <UCard
          variant="outline"
          class="flex flex-col gap-3 hover:bg-elevated/50"
          :ui="{ header: 'p-2 sm:px-4', body: 'p-2 sm:p-4', footer: 'p-2 sm:px-4' }"
        >
          <template #header>
            <ConfirmDialog
              title="Confirmation"
              :content="`Do you really want to delete the entry #${idx + 1}?`"
              :exec="
                () => {
                  model.entries.splice(idx, 1)
                }
              "
              class="float-right"
            >
              <UButton
                label="Delete Entry"
                icon="i-lucide-trash"
                color="error"
                variant="subtle"
                size="xs"
                class="flex-none"
              />
            </ConfirmDialog>
            #{{ idx + 1 }}
          </template>

          <div class="flex flex-row gap-3 w-full">
            <div class="flex-init flex flex-col gap-3">
              <div class="flex-none flex flex-row gap-2">
                <div class="flex-auto">
                  <TFormField
                    :name="`${prefix}entries.${idx}.raceNumber`"
                    label="Race Number"
                    help="The preferred race number if set, -1 if the driver may
decide by picking his car. Values 1 - 998"
                  >
                    <UInput
                      v-model="entry.raceNumber"
                      type="number"
                      min="-1"
                      max="998"
                      class="w-full"
                    />
                  </TFormField>
                </div>
                <div class="flex-auto pl-4">
                  <TFormField
                    :name="`${prefix}entries.${idx}.forcedCarModel`"
                    label="Forced Car Model"
                    help="If not set to -1: user cannot join with a different car, see
“Car model list” for the values "
                  >
                    <USelect :items="carItems" v-model="entry.forcedCarModel" class="w-62" />
                  </TFormField>
                </div>
              </div>

              <div class="flex-none">
                <TFormField
                  :name="`${prefix}entries.${idx}.customCar`"
                  label="Custom Car"
                  help="If set to a filename, the car, team and appearance will be
used no matter what the user chose (Exception:
overrideCarModelForCustomCar). This is useful for
leagues and events, where we want consistent car
appearance and the chosen car model for the
corresponding driver/team. The custom car file has to be
located in a “cars” folder next to the entrylist.json (also
works for centralEntryListPath).
Leave blank (“”, =default) to let the user chose the car via
car selection UI. "
                >
                  <UInput v-model="entry.customCar" class="w-full" />
                </TFormField>

                <NBCheckbox
                  v-model="entry.overrideCarModelForCustomCar"
                  label="Override Car Model for Custom Car"
                  description="If customCar is used, this setting will apply the car model
configured if the value is set to 1 (which is the default). If
set to 0, all values except the carModel are applied, so the
user is free to pick a car but while team name and
appearance will be applied."
                />
              </div>

              <div class="flex-none flex flex-row gap-2">
                <div class="flex-1">
                  <TFormField
                    :name="`${prefix}entries.${idx}.ballastKg`"
                    label="Ballast (0 - 100kg)"
                    help="Assigns ballast in kg for this car. Will be additive to
ballast for the car model (via bop.json), and can be
overridden by the admin command /ballast. Range is 0 to
100."
                  >
                    <UInput v-model="entry.ballastKg" type="number" min="0" max="100" />
                  </TFormField>
                </div>
                <div class="flex-1">
                  <TFormField
                    :name="`${prefix}entries.${idx}.restrictor`"
                    label="Restrictor (0 - 20%)"
                    help="Assigns restrictor in % for this car. Will be additive to
restrictor for the car model (via bop.json), and can be
overridden by the admin command /restrictor. Range is 0
to 20."
                  >
                    <UInput v-model="entry.restrictor" type="number" min="0" max="20" />
                  </TFormField>
                </div>
              </div>

              <div class="flex-none">
                <NBCheckbox
                  v-model="entry.overrideDriverInfo"
                  label="Override Driver Info"
                  description="If set to 1, the driver’s name and category will be
overridden by what is setup in the entry list. If set to 0, it’s
up to the client joining."
                />

                <NBCheckbox
                  v-model="entry.isServerAdmin"
                  label="Is Server Admin"
                  description="If set to 1, that user will be automatically elevated to
server admin when he joins."
                />
              </div>
            </div>

            <div class="flex-none px-2">
              <USeparator orientation="vertical" class="h-full" />
            </div>

            <div class="flex-auto place-items-center">
              <DriversTable
                v-model="entry.drivers"
                :index="idx"
                :hide-prefix="hideForceEntryListCheckbox"
              />
            </div>
          </div>
        </UCard>
      </div>

      <div class="flex-auto flex flex-row gap-2 mt-5">
        <UButton
          label="Add"
          icon="i-lucide-plus"
          @click="addEntry()"
          color="secondary"
          variant="subtle"
          class="flex-none"
        />

        <ConfirmDialog
          title="Confirmation"
          content="Do you really want to clear all entries?"
          :exec="
            () => {
              model.entries = []
            }
          "
        >
          <UButton
            label="Clear all"
            icon="i-lucide-trash"
            color="error"
            variant="subtle"
            class="flex-none"
          />
        </ConfirmDialog>
      </div>
    </div>
  </div>
</template>
