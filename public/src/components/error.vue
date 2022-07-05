<template>
<div class="error-toast" v-show="message">
    <div class="error-content">
        {{message}}
    </div>
</div>
</template>

<script>
export default {
    computed: {
        message() {
            return this.$store.state.error.errorMsg;
        }
    }
}
</script>
