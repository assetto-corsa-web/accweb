<script setup lang="ts">
import NBCheckbox from '@/components/NBCheckbox.vue'
import { type SettingsJson } from '@/lib/accweb/types'
import type { SelectItem } from '@nuxt/ui'
import carGroups from '@/data/cargroup'
import { ref } from 'vue'

const carGroupItems = ref<SelectItem[]>(carGroups)

const formationLapItems = ref<SelectItem[]>([
  { label: 'New Position control and UI', value: 3 },
  { label: 'Old with limiter', value: 0 },
  { label: 'Free no limiter', value: 1 },
])

const model = defineModel({
  required: true,
  default: {} as SettingsJson,
})
</script>

<template>
  <div class="flex flex-col gap-3">
    <TFormField
      label="Password"
      name="acc.settings.password"
      help="Password required to enter ths server. If a password is set, the
server is declared 'Private Multiplayer'."
    >
      <UInput v-model="model.password" />
    </TFormField>

    <TFormField
      label="Admin Password"
      name="acc.settings.adminPassword"
      help="Password to elevate via 'Server admin commands'"
    >
      <UInput v-model="model.adminPassword" />
    </TFormField>

    <TFormField
      label="Spectator Password"
      name="acc.settings.spectatorPassword"
      help="Password to enter the server as spectator. Must be different to
'password' if both is set. "
    >
      <UInput v-model="model.spectatorPassword" />
    </TFormField>

    <TFormField
      label="Track Medals (0 - 3)"
      name="acc.settings.trackMedalsRequirement"
      help="Defines the amount of track medals that a user has to have for
the given track (values 0, 1, 2, 3)"
    >
      <UInput v-model="model.trackMedalsRequirement" type="number" min="0" max="3" />
    </TFormField>

    <TFormField
      label="Safity Rating requirement (-1 a 99)"
      name="acc.settings.safetyRatingRequirement"
      help="Defines the Safety Rating (SA) that a user must have to join this
server (values -1, 0, …. 99)"
    >
      <UInput v-model="model.safetyRatingRequirement" type="number" min="-1" max="99" />
    </TFormField>

    <TFormField
      label="Race Craft Rating requirement (-1 a 99)"
      name="acc.settings.racecraftRatingRequirement"
      help="Defines the Safety Rating (RC) that a user must have to join
this server (values -1, 0, …. 99)"
    >
      <UInput v-model="model.racecraftRatingRequirement" type="number" min="-1" max="99" />
    </TFormField>

    <TFormField
      label="Max Car Slots"
      name="acc.settings.maxCarSlots"
      help="Defines
the amount of car slots the server can occupy; this value is
overridden if the pit count of the track is lower, or with 30 for
public MP. The gap between maxCarSlots and maxConnections
defines how many spectators or other irregular connections (ie
entry list entries) can be on the server."
    >
      <UInput v-model="model.maxCarSlots" type="number" min="0" />
    </TFormField>

    <NBCheckbox
      v-model="model.ignorePrematureDisconnects"
      label="Ignore Premature Disconnects"
      description="Removes a (very good) fix where users can randomly lose the
connection. There is no sane reason to turn this off. UNCHECK if you are on Linux"
    />

    <NBCheckbox
      v-model="model.dumpLeaderboards"
      label="Dump Leaderboards"
      description="If checked, any session will write down the result leaderboard in
a “results” folder."
    />

    <NBCheckbox
      v-model="model.isRaceLocked"
      label="Is Race Locked"
      description="If unchecked, the server will allow joining during a race session. Is
not useful in “Public Multiplayer”, as the user-server matching
will ignore ongoing race sessions. "
    />

    <NBCheckbox
      v-model="model.randomizeTrackWhenEmpty"
      label="Randomize Track When Empty"
      description="If set to 1, the server will change to a random track when the
last drivers leaves (which causes a reset to FP1). The “track”
property will only define the default state for the first session. "
    />

    <NBCheckbox
      v-model="model.allowAutoDQ"
      label="Allow Auto DQ"
      description="If unchecked, the server won’t automatically disqualify drivers, and
instead hand out Stop&Go (30s) penalties. This way a server
admin / race director has 3 laps time to review the incident, and
either use /dq or /clear based on his judgement."
    />

    <TFormField
      label="Car Group"
      name="acc.settings.carGroup"
      help="Defines the car group for this server."
    >
      <USelect :items="carGroupItems" v-model="model.carGroup" class="w-40" />
    </TFormField>

    <TFormField
      label="Formation Lap Type"
      name="acc.settings.formationLapType"
      help="Defines the car group for this server."
    >
      <USelect :items="formationLapItems" v-model="model.formationLapType" class="w-56" />
    </TFormField>

    <NBCheckbox
      v-model="model.shortFormationLap"
      label="Shot Formation Lap"
      description="Toggles the short and long formation lap. Long formation is
only usable on private servers."
    />

    <TFormField
      label="Central Entry List Path"
      name="acc.settings.centralEntryListPath"
      help="Can override the default entryList path “cfg/entrylist.json”, so
multiple ACC servers on the machine can use the same entrylist
(and custom car files). Set a full path like
“C:/customEntryListSeriesA/”, where the entrylist is stored.
Attention: The path seperators have to be slashes (/),
backslashes (\) will not work."
    >
      <UInput v-model="model.centralEntryListPath" />
    </TFormField>

    <NBCheckbox
      v-model="model.dumpEntryList"
      label="Dump Entry List"
      description="Will save an entry list at the end of any Qualifying session. This
can be a quick way to collect a starting point to build an entry
list, and is a way to save the defaultGridPositions which can be
used to run a race without Qualifying session and predefined
grid. Also see the corresponding admin command."
    />
  </div>
</template>
