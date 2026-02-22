<script setup lang="ts">
import { type EventJson } from '@/lib/accweb/types'
import tracks from '@/data/tracks'
import { ref } from 'vue'
import type { SelectItem } from '@nuxt/ui'

const model = defineModel({
  required: true,
  default: {} as EventJson,
})

const tracksItems = ref<SelectItem[]>(tracks)
</script>

<template>
  <div class="flex flex-row gap-2">
    <div class="flex-1 flex flex-col gap-2">
      <TFormField label="Track" name="acc.event.track" help="The track we run.">
        <USelect :items="tracksItems" v-model="model.track" class="w-62" />
      </TFormField>

      <TFormField
        label="Pre race waiting time seconds"
        name="acc.event.preRaceWaitingTimeSeconds"
        help="Preparation time before a race. Cannot be less than 30s."
      >
        <UInput v-model="model.preRaceWaitingTimeSeconds" type="number" min="30" />
      </TFormField>

      <TFormField
        label="Session overtime seconds"
        name="acc.event.sessionOverTimeSeconds"
        help="Time after that a session is forcibly closing after the timer
reached 0:00. Something like 107% of the expected laptime is
recommended (careful: default 2 minutes does not properly
cover tracks like Spa or Silverstone). "
      >
        <UInput v-model="model.sessionOverTimeSeconds" type="number" min="0" />
      </TFormField>

      <TFormField
        label="Post Qualy Seconds"
        name="acc.event.postQualySeconds"
        help="The time after the last driver is finished (or the
sessionOverTimeSeconds passed) in Q sessions and the race
start. Should not be set to 0, otherwise grid spawning is not secure."
      >
        <UInput v-model="model.postQualySeconds" type="number" />
      </TFormField>

      <TFormField
        label="Post Race Seconds"
        name="acc.event.postRaceSeconds"
        help="Additional time after the race ended for everyone, before the
next race weekend starts."
      >
        <UInput v-model="model.postRaceSeconds" type="number" />
      </TFormField>

      <TFormField
        label="Metadata"
        name="acc.event.metaData"
        help="A user defined string that will be transferred to the result
outputs."
      >
        <UInput v-model="model.metaData" />
      </TFormField>

      <TFormField
        label="Ambient temperature"
        name="acc.event.ambientTemp"
        help="Sets the baseline ambient temperature in °C, see  “Race
weekend simulation”"
      >
        <UInput v-model="model.ambientTemp" type="number" />
      </TFormField>

      <TFormField
        label="Cloud Level"
        name="acc.event.cloudLevel"
        help="Sets the baseline cloud level, see “Race weekend simulation”.
Has large impact on the cloud levels and rain chances."
      >
        <UInput v-model="model.cloudLevel" type="number" min="0" max="1" step="0.1" />
      </TFormField>

      <TFormField
        label="Rain"
        name="acc.event.rain"
        help="If weather randomness is off, defines the static rain level. With
dynamic weather, it defines the expected rain level, dependent
on weatherRandomness."
      >
        <UInput v-model="model.rain" type="number" min="0" max="1" step="0.1" />
      </TFormField>

      <TFormField
        label="Weather Randomness"
        name="acc.event.weatherRandomness"
        help="Sets the dynamic weather level, see  “Race weekend
simulation”. 0 = static weather; 1-4 fairly realistic weather; 5-7 more
sensational"
      >
        <UInput v-model="model.weatherRandomness" type="number" min="0" max="7" />
      </TFormField>

      <NBCheckbox
        v-model="model.simracerWeatherConditions"
        label="Simracer Weather Conditions"
        description="Experimental/not supported: if set to 1, this will limit the
maximum rain/wetness to roughly 2/3 of the maximum values,
translating to something between medium and heavy rain.
It may be useful if you feel forced to run very low cloudLevel
and weatherRandomness values just to avoid thunderstorm;
however high levels (0.4+ clouds combined with 5+
randomness) will still result in quite serious conditions."
      />

      <NBCheckbox
        v-model="model.isFixedConditionQualification"
        label="Fixed Weather Condition in Qualification"
        description="Experimental/not supported: if set to 1, the server will take the
rain, cloud, temperature, rain levels literally and make sure
whatever is set up never changes. Daytime transitions still
happen visually, but do not affect the temperatures or road
wetness. Also rubber/grip is always the same. This is intended
to be used for private league qualification servers only.
weatherRandomness has to be set to 0, otherwise
isFixedConditionQualification will be completely disabled. "
      />
    </div>

    <div class="flex-1">
      <SessionsTable v-model="model.sessions" />
    </div>
  </div>
</template>
