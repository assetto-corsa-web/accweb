<script setup lang="ts">
import type { AssistRulesJson } from '@/lib/accweb/types'

const model = defineModel({
  required: true,
  default: {} as AssistRulesJson,
})
</script>

<template>
  <div class="flex flex-col gap-3">
    <TFormField
      label="Stability Control Level Max: 0 to 100 max"
      name="acc.assistRules.stabilityControlLevelMax"
      help="Set’s the maximum % of SC that can be used. In
case a client has a higher SC set than allowed by
the server, he will only run what is allowed (25%
in this example). Obviously setting this property to
0 removes all SC, including mouse and keyboard
users.
The Stability Control is an artificial driving aid
that allows the car to act out of the physics
boundaries, and highly recommended to overcome
input methods like Keyboards, Gamepads and
Mouse steering. However, there is a built-in effect
that makes the SC performance inferior, so in
theory using (and relying) on SC is already more
than enough penalty, and the way to improve
performance is to practice driving without.
Default: 100"
    >
      <UInput v-model="model.stabilityControlLevelMax" type="number" min="0" max="100" step="1" />
    </TFormField>

    <NBCheckbox
      v-model="model.disableAutosteer"
      label="Disable Auto-Steer"
      description="Disables the steering aid that is only available for
gamepad controllers. Unlike SC, this works inside
the physics and does not allow unrealistic driving
behaviour – except that this is a very strong aid
with superhuman feeling for grip and high reaction
speed. There is a built-in penalty that should
balance the driving performance in most cases, and
give an incentive to learn not to use the driving
aid.
Default: 0"
    />

    <NBCheckbox
      v-model="model.disableAutoLights"
      label=" Disable Auto-Lights"
      description="Forces the equivalent assist option to “off”"
    />

    <NBCheckbox
      v-model="model.disableAutoWiper"
      label="Disable Auto-Wiper"
      description="Forces the equivalent assist option to “off”"
    />

    <NBCheckbox
      v-model="model.disableAutoEngineStart"
      label="Disable Auto-Engine Start"
      description="Forces the equivalent assist option to “off”"
    />

    <NBCheckbox
      v-model="model.disableAutoPitLimiter"
      label="Disable Auto Pit Limiter"
      description="Forces the equivalent assist option to “off”"
    />

    <NBCheckbox
      v-model="model.disableAutoGear"
      label="Disable Auto-Gear"
      description="Forces the equivalent assist option to “off”"
    />

    <NBCheckbox
      v-model="model.disableAutoClutch"
      label="Disable Auto-Clutch"
      description="Forces the equivalent assist option to “off”"
    />

    <NBCheckbox
      v-model="model.disableIdealLine"
      label="Disable Ideal Line"
      description="Forces the equivalent assist option to “off”"
    />
  </div>
</template>
