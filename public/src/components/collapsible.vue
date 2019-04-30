<template>
    <div class="collapsible">
        <div class="collapsible-title" v-on:click="expanded = !expanded">
            {{title}}
            <i v-bind:class="{fas: true, 'fa-chevron-down': !expanded, 'fa-chevron-up': expanded}"></i>
        </div>
        <div class="collapsible-content" v-show="expanded">
            <slot></slot>
        </div>
    </div>
</template>

<script>
export default {
    props: ["title"],
    data() {
        return {
            expanded: false
        };
    }
}
</script>
