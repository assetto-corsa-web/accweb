<template>
  <!--
    O v-bind="$attrs" repassa todas as propriedades extras
    (como 'name', 'help', 'error', etc) diretamente para o UFormField
  -->
  <UFormField :label="label" :name="name">
    <!-- Repassa o slot padrão para permitir colocar o input/checkbox dentro -->
    <slot />

    <template #label>
      <div class="flex gap-2">
        <div class="flex-none">{{ label }}</div>

        <div class="flex-none">
          <UTooltip
            :text="help"
            v-if="help"
            portal
            :delay-duration="0"
            :ui="{
              content: 'max-w-100 h-auto py-2 px-3 flex-wrap',
              text: 'whitespace-normal wrap-break-word line-clamp-none',
            }"
          >
            <UIcon name="i-lucide-circle-help" class="w-4 h-4 text-primary cursor-help" />
          </UTooltip>
        </div>
      </div>
    </template>

    <!-- Repassa outros slots do UFormField se necessário -->
    <template v-for="(_, name) in $slots" #[name]="slotData">
      <slot :name="name" v-bind="slotData || {}" />
    </template>
  </UFormField>
</template>

<script setup lang="ts">
import type { FormFieldProps } from '@nuxt/ui'

// Definimos uma interface que estende as props originais do UFormField
// mas marcamos label e description como obrigatórias ou customizadas
interface Props extends /* @vue-ignore */ Omit<FormFieldProps, 'label' | 'description'> {
  label?: string
  name: string
  help?: string
}

defineProps<Props>()

// Desabilita a herança de atributos no elemento raiz (div)
// para que eles caiam apenas no UFormField via v-bind="$attrs"
defineOptions({
  inheritAttrs: false,
})
</script>
