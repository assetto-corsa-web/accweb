<template>
    <div>
        <label>{{label}}</label>
        <div>
            <select v-bind:value="value"
                v-on:input="$emit('input', $event.target.value)">
                <option v-for="o in options">{{o}}</option>
            </select>
        </div>
        <div class="error" v-show="error">{{error}}</div>
    </div>
</template>

<script>
export default {
    props: {
        value: null,
        label: null,
        options: null,
        error: null
    }
}
</script>
