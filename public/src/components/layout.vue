<template>
    <div v-bind:class="{main: true, 'main-small': small}">
    	<img src="static/favicon.png" alt="" class="logo" />
        <slot></slot>
        <end></end>
        <error></error>
    </div>
</template>

<script>
import end from "./end.vue";
import error from "./error.vue";

export default {
    components: {end, error},
    props: ["small"]
}
</script>
