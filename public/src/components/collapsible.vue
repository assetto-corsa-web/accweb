<template>
    <div class="collapsible">
        <div class="collapsible-title" v-on:click="expanded = !expanded">
            {{title}}
            <filereader v-show="withImport" @load="onLoadContent"></filereader>
            <i v-bind:class="{fas: true, 'fa-chevron-down': !expanded, 'fa-chevron-up': expanded}"></i>
        </div>
        <div class="collapsible-content" v-show="expanded">
            <slot></slot>
        </div>
    </div>
</template>

<script>
import filereader from "./filereader.vue";

export default {
    props: ["title", "withImport"],
    components: { filereader },
    data() {
        return {
            expanded: false
        };
    },
    methods: {
        onLoadContent: function (e) {
            this.$emit("load", JSON.parse(e));
            this.expanded = true
        }
    }
}
</script>
