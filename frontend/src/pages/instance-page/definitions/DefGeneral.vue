<script setup lang="ts">
import NBCheckbox from '@/components/NBCheckbox.vue'
import { type ConfigurationJson } from '@/lib/accweb/types'

const model = defineModel({
  required: true,
  default: {} as ConfigurationJson,
})
</script>

<template>
  <div class="flex flex-col gap-3">
    <TFormField
      label="UDP Port"
      name="acc.configuration.udpPort"
      help="ACC clients will use this port to establish a connection to the server."
    >
      <UInput v-model="model.udpPort" type="number" />
    </TFormField>

    <TFormField
      label="TCP Port"
      name="acc.configuration.tcpPort"
      help="Connected clients will use this Port to stream the car positions and is used
for the ping test. In case you never see your server getting a ping value, this
indicates that the udpPort is not accessible "
    >
      <UInput v-model="model.tcpPort" type="number" />
    </TFormField>

    <TFormField
      label="Max connections"
      name="acc.configuration.maxConnections"
      help="Replaces “maxClients”. The maximum amount of connections a server will
accept at a time. If you own the hardware server, you can just set any high
number you want. If you rented a 16 or 24 slot server, your Hosting
Provider probably has set this here and doesn’t give you write-access to this
configuration file."
    >
      <UInput v-model="model.maxConnections" type="number" min="0" />
    </TFormField>

    <TFormField
      label="Public IP"
      name="acc.configuration.publicIP"
      help="Explicitly defines the public IP address this server is listening to. Useful if
the backend is connected via a different gateway (for example the AWS
Accelerator IP).
If the publicIP is used, the server has to respond to an additional handshake,
or it will immediately shutdown on backend connect. "
    >
      <UInput v-model="model.publicIP" type="number" />
    </TFormField>

    <NBCheckbox
      v-model="model.registerToLobby"
      label="Register to lobby"
      description="When unchecked, this server won’t register to the backend. Is useful for LAN
sessions. If unchecked, the server is declared “Private Multiplayer”.
See serverList.json to learn how to bypass the backend’s server list and
discover servers not listed."
    />

    <NBCheckbox
      v-model="model.lanDiscovery"
      label="LAN Discovery"
      description="Defines if the server will listen to LAN discovery requests. Can be turned
off for dedicated servers."
    />
  </div>
</template>
