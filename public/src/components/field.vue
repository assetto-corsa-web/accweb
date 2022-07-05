<template>
<div>
    <label>
        {{label}}
        <input :type="type" :placeholder="placeholder" v-bind:value="value" v-on:input="$emit('input', $event.target.value)" :step="step" :disabled="disabled" class="input" />
    </label>
    <div class="error" v-show="error">{{error}}</div>
</div>
</template>

<script>
export default {
    props: {
        value: null,
        type: {
            default: "text"
        },
        label: null,
        placeholder: null,
        step: null,
        disabled: false,
        error: null
    }
}
</script>

<style>
.input {
    background-color: transparent;
    border-style: solid;
}
</style>
