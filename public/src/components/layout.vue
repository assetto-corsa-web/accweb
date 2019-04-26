<template>
    <div v-bind:class="{main: true, 'main-small': small}">
        <slot></slot>
        <end></end>
    </div>
</template>

<script>
import end from "./end.vue";

export default {
    components: {end},
    props: ["small"]
}
</script>
