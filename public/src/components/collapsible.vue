<template>
    <div class="collapsible">
        <div class="collapsible-title" v-on:click="expanded = !expanded">
            <div class="action-title">
                {{title}}
                <filereader v-show="withImport" @load="onLoadContent" :filename="importFilename"></filereader>
            </div>

            <i class="collapse-icon" v-bind:class="{fas: true, 'fa-chevron-down': !expanded, 'fa-chevron-up': expanded}"></i>
        </div>
        <div class="collapsible-content" v-show="expanded">
            <div v-show="loadError" class="alert">
                This is an invalid JSON file or maybe there is encoding issues.
            </div>
            <slot></slot>
        </div>
    </div>
</template>

<script>
import filereader from "./filereader.vue";

export default {
    props: ["title", "withImport", "importFilename"],
    components: { filereader },
    data() {
        return {
            expanded: false,
            loadError: ''
        };
    },
    methods: {
        onLoadContent: function (e) {
            this.expanded = true;
            this.loadError = false;
            try {
                const obj = JSON.parse(e.replaceAll('\x00', ''));
                this.$emit("load", obj);
                
            } catch (err) {
                this.loadError = true;
            }
        }
    }
}
</script>

<style>
  .action-title {
    display: grid;
    grid-template-columns: 1fr auto;
    grid-gap: 10px;
  }
</style>
