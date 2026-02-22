<script setup lang="ts">
import type { InstancePayload, SessionSettings } from '@/lib/accweb/types'
import { ref } from 'vue'
import TabDefinitions from './instance-page/TabDefinitions.vue'
import { createInstance } from '@/lib/accweb/client'
import { useRouter } from 'vue-router'

const toast = useToast()
const router = useRouter()

const instance = ref<InstancePayload>({
  acc: {
    configuration: {
      lanDiscovery: 0,
      registerToLobby: 1,
      publicIP: '',
      maxConnections: 10,
      tcpPort: 9600,
      udpPort: 9600,
    },
    settings: {
      serverName: 'Server name (by accweb)',
      adminPassword: '',
      spectatorPassword: '',
      password: '',
      trackMedalsRequirement: 0,
      safetyRatingRequirement: -1,
      racecraftRatingRequirement: -1,
      maxCarSlots: 30,
      ignorePrematureDisconnects: 1,
      dumpLeaderboards: 0,
      isRaceLocked: 0,
      randomizeTrackWhenEmpty: 0,
      allowAutoDQ: 1,
      carGroup: 'FreeForAll',
      formationLapType: 3,
      shortFormationLap: 0,
      centralEntryListPath: '',
      dumpEntryList: 0,
    },
    assistRules: {
      stabilityControlLevelMax: 100,
    },
    bop: {},
    entrylist: {
      forceEntryList: 0,
    },
    event: {
      track: 'barcelona',
      ambientTemp: 26,
      trackTemp: 26,
      cloudLevel: 0.3,
      rain: 0,
      weatherRandomness: 1,
      preRaceWaitingTimeSeconds: 30,
      sessionOverTimeSeconds: 120,
      postQualySeconds: 0,
      postRaceSeconds: 0,
      metaData: '',
      simracerWeatherConditions: 0,
      isFixedConditionQualification: 0,
      sessions: [] as SessionSettings[],
    },
    eventRules: {
      qualifyStandingType: 1,
      pitWindowLengthSec: -1,
      driverStintTimeSec: -1,
      mandatoryPitstopCount: 0,
      maxTotalDrivingTime: -1,
      maxDriversCount: 1,
      tyreSetCount: 50,
      isRefuellingAllowedInRace: true,
      isMandatoryPitstopRefuellingRequired: false,
      isMandatoryPitstopSwapDriverRequired: false,
      isMandatoryPitstopTyreChangeRequired: false,
      isRefuellingTimeFixed: false,
    },
  },
  accExtraSettings: {
    adminPasswordIsEmpty: true,
    spectatorPasswordIsEmpty: true,
    passwordIsEmpty: true,
  },
  accWeb: {
    autoStart: false,
    enableAdvWindowsCfg: false,
    enableGlobalBanlist: true,
    enableGlobalEntrylist: true,
  },
} as InstancePayload)

async function onSubmit(data: InstancePayload) {
  createInstance(data)
    .then((rs) => {
      toast.add({
        title: 'Success',
        description: 'The new instance has been created.',
        color: 'success',
      })

      router.push({ path: `/instance/${rs.id}` })
    })
    .catch((err) => {
      toast.add({
        title: 'Failed to create instance',
        description: err.response.data.error,
        color: 'error',
      })
    })
}
</script>

<template>
  <TabDefinitions :instance="instance" @instance-updated="() => {}" :overrideOnSubmit="onSubmit" />
</template>
