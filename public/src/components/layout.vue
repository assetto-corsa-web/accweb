<template>
    <div v-bind:class="{main: true, 'main-small': small}">
        <router-link to="/"><img src="static/logo.png" alt="" class="logo" /></router-link>
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
