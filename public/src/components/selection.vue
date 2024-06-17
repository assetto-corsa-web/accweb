<template>
    <div>
        <label>
            {{ label }}

            <select v-bind:value="value" v-on:input="$emit('input', $event.target.value)"
                v-bind:class="{ inline: inline }">
                <option v-for="o in options" v-bind:value="o.value">{{ o.label }}</option>
            </select>
        </label>
        <div class="error" v-show="error">{{ error }}</div>
    </div>
</template>

<script>
export default {
    props: {
        value: null,
        label: null,
        options: null,
        inline: false,
        error: null
    }
}
</script>

<style>
.inline {
    display: inline;
    margin-left: 5px;
}
</style>
