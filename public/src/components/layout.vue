<template>
    <div v-bind:class="{main: true, 'main-small': small}">
    	<img src="static/favicon.png" alt="" class="logo" />
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
