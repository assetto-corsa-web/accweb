<template>
    <div :title="$t('title')">
        <checkbox :label="$t('autostart_label')" v-model="autoStart"></checkbox>
    </div>
</template>

<script>
import collapsible from "../collapsible.vue";
import checkbox from "../checkbox.vue";

export default {
    components: {collapsible, checkbox},
    data() {
        return {
            autoStart: false
        };
    },
    methods: {
        setData(data) {
            this.autoStart = data.autoStart;
        },
        getData() {
            return {
              autoStart: this.autoStart
            };
        }
    }
}
</script>

<i18n>
{
    "en": {
        "title": "ACC Web configuration",
        "autostart_label": "Server instance auto start."
    }
}
</i18n>