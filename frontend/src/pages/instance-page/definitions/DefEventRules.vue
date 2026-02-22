<script setup lang="ts">
import type { EventRulesJson } from '@/lib/accweb/types'
import type { SelectItem } from '@nuxt/ui'
import { ref } from 'vue'

const model = defineModel({
  required: true,
  default: {} as EventRulesJson,
})

const qualifyStandingTypes = ref<SelectItem[]>([
  { label: 'Fastest Lap', value: 1 },
  { label: 'Average Lap', value: 2 },
])
</script>

<template>
  <div class="flex flex-row gap-3">
    <div class="flex-1 flex flex-col gap-3">
      <TFormField
        label="Qualify Standing Type"
        name="acc.eventRules.qualifyStandingType"
        help="1 = fastest lap, 2 = average lap (running Endurance
mode for multiple Q sessions) . Use 1, averaging
Qualy is not yet officially supported."
      >
        <USelect
          :items="qualifyStandingTypes"
          v-model="model.qualifyStandingType"
          class="w-30"
          default-value="1"
        />
      </TFormField>

      <TFormField
        label="Pit Window Length (seconds)"
        name="acc.eventRules.pitWindowLengthSec"
        help="Defines a pit window at the middle of the race.
Obviously covers the Sprint series format. -1 will
disable the pit window. Use this combined with a
mandatoryPitstopCount = 1."
      >
        <UInput v-model="model.pitWindowLengthSec" type="number" min="-1" step="1" />
      </TFormField>

      <TFormField
        label="Driver Stint Time (seconds)"
        name="acc.eventRules.driverStintTimeSec"
        help="Defines the maximum time a driver can stay out
without getting a penalty. Can be used to balance
fuel efficient cars in endurance races. The stint
time resets in the pitlane, no real stop is required.
-1 will disable the stint times. driverStintTimeSec
and maxTotalDrivingTime are interdependent features,
make sure both are set or off."
      >
        <UInput v-model="model.driverStintTimeSec" type="number" min="-1" step="1" />
      </TFormField>

      <TFormField
        label="Mandatory Pitstop Count"
        name="acc.eventRules.mandatoryPitstopCount"
        help="Defines the basic mandatory pit stops. If the value
is greater zero, any car that did not execute the
mandatory pitstops will be disqualified at the end
of the race. The necessary actions can be further
configured using the
“isMandatoryPitstopXYRequired” properties. A
value of zero disables the feature."
      >
        <UInput v-model="model.mandatoryPitstopCount" type="number" min="0" step="1" />
      </TFormField>

      <TFormField
        label="Max Total Driving Time"
        name="acc.eventRules.maxTotalDrivingTime"
        help="Restricts the maximum driving time for a single
driver. Is only useful for driver swap situations and
allows to enforce a minimum driving time for each
driver (IRL this is used to make sure mixed teams
like Pro/Am have a fair distributions of the slower
drivers). -1 disables the feature.
driverStintTimeSec and maxTotalDrivingTime are
interdependent features, make sure both are set or
off.
Will set the maximum driving time for the team
size defined by “maxDriversCount”, always make
sure both are set."
      >
        <UInput v-model="model.maxTotalDrivingTime" type="number" min="-1" step="1" />
      </TFormField>

      <TFormField
        label="Max Drivers Count"
        name="acc.eventRules.maxDriversCount"
        help="In driver swap situations, set this to the maximum
        number of drivers on a car. When an entry has fewer drivers
        than maxDriversCount, maxTotalDrivingTime is automatically
        compensated so that those 'smaller' entries are also able to
        complete the race Example: 3H race length, 65 minutes
        driverStintTimeSec and 65 minutes maxTotalDrivingTime will
        result in 65 minutes of maxTotalDrivingTime for entries of
        3 and 105 (!) minutes for entries of 2."
      >
        <UInput v-model="model.maxDriversCount" type="number" min="0" step="1" />
      </TFormField>

      <TFormField
        label="Tyre Set Count"
        name="acc.eventRules.tyreSetCount"
        help="Experimental/not supported: Can be used to reduce
the amount of tyre sets any car entry has for the
entire weekend. Please note that it is necessary to
force cars to remain in the server, or drastically
reduced tyre sets will be ineffective, as rejoining
will reset the tyre sets. "
      >
        <UInput v-model="model.tyreSetCount" type="number" min="0" step="1" />
      </TFormField>
    </div>

    <div class="flex-1 flex flex-col gap-3">
      <TCheckbox
        v-model="model.isRefuellingAllowedInRace"
        label="Is Refuelling Allowed in Race"
        description="Defines if refuelling is allowed during the race
pitstops. "
      />

      <TCheckbox
        v-model="model.isRefuellingTimeFixed"
        label="Is Refuelling Time Fixed"
        description="If turned on, any refuelling will take the same
amount of time. If turned off, refuelling will
consume time linear to the amount refuelled. Very
useful setting to balance fuel efficient cars,
especially if combined with other features."
      />

      <TCheckbox
        v-model="model.isMandatoryPitstopRefuellingRequired"
        label="Is Mandatory Pitstop Refuelling Required"
        description="Defines if a mandatory pitstop requires refuelling."
      />

      <TCheckbox
        v-model="model.isMandatoryPitstopTyreChangeRequired"
        label="Is Mandatory Pitstop Tyre Change Required"
        description="Defines if a mandatory pitstop requires changing
tyres."
      />

      <TCheckbox
        v-model="model.isMandatoryPitstopSwapDriverRequired"
        label="Is Mandatory Pitstop Swap Driver Required"
        description="Defines if a mandatory pitstop requires a driver
swap. Will only be effective for cars in driver
swap situations; even in a mixed field this will be
skipped for cars with a team size of 1 driver."
      />
    </div>
  </div>
</template>
