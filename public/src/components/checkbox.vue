<template>
<div class="checkbox">
    <label>
        <input v-bind:value="value" v-on:input="$emit('input', $event.target.checked)" v-bind:checked="value" type="checkbox" />
        {{label}}
    </label>
    <div class="error" v-show="error">{{error}}</div>
</div>
</template>

<script>
export default {
    props: {
        value: false,
        label: null,
        error: null
    }
}
</script>
