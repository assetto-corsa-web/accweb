<template>
    <collapsible :title="$t('title')" with-import="true" import-filename="assistRules.json" @load="setData">
        <field type="number" :label="$t('stabilityControlLevelMax_label')" v-model="stabilityControlLevelMax"></field>
        <checkbox :label="$t('disableAutosteer_label')" v-model="disableAutosteer"></checkbox>
        <checkbox :label="$t('disableAutoLights_label')" v-model="disableAutoLights"></checkbox>
        <checkbox :label="$t('disableAutoWiper_label')" v-model="disableAutoWiper"></checkbox>
        <checkbox :label="$t('disableAutoEngineStart_label')" v-model="disableAutoEngineStart"></checkbox>
        <checkbox :label="$t('disableAutoPitLimiter_label')" v-model="disableAutoPitLimiter"></checkbox>
        <checkbox :label="$t('disableAutoGear_label')" v-model="disableAutoGear"></checkbox>
        <checkbox :label="$t('disableAutoClutch_label')" v-model="disableAutoClutch"></checkbox>
        <checkbox :label="$t('disableIdealLine_label')" v-model="disableIdealLine"></checkbox>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import field from "../field.vue";
import checkbox from "../checkbox.vue";

export default {
    components: {collapsible, field, checkbox},
    data() {
        return {
            stabilityControlLevelMax: 100,
            disableAutosteer: false,
            disableAutoLights: false,
            disableAutoWiper: false,
            disableAutoEngineStart: false,
            disableAutoPitLimiter: false,
            disableAutoGear: false,
            disableAutoClutch: false,
            disableIdealLine: false,
            configVersion: 1
        };
    },
    methods: {
        setData(data) {
            this.stabilityControlLevelMax = data.stabilityControlLevelMax;
            this.disableAutosteer = data.disableAutosteer;
            this.disableAutoLights = data.disableAutoLights;
            this.disableAutoWiper = data.disableAutoWiper;
            this.disableAutoEngineStart = data.disableAutoEngineStart;
            this.disableAutoPitLimiter = data.disableAutoPitLimiter;
            this.disableAutoGear = data.disableAutoGear;
            this.disableAutoClutch = data.disableAutoClutch;
            this.disableIdealLine = data.disableIdealLine;
        },
        getData() {
            return {
                stabilityControlLevelMax: parseInt(this.stabilityControlLevelMax),
                disableAutosteer: this.disableAutosteer ? 1 : 0,
                disableAutoLights: this.disableAutoLights ? 1 : 0,
                disableAutoWiper: this.disableAutoWiper ? 1 : 0,
                disableAutoEngineStart: this.disableAutoEngineStart ? 1 : 0,
                disableAutoPitLimiter: this.disableAutoPitLimiter ? 1 : 0,
                disableAutoGear: this.disableAutoGear ? 1 : 0,
                disableAutoClutch: this.disableAutoClutch ? 1 : 0,
                disableIdealLine: this.disableIdealLine ? 1 : 0,
                configVersion: 1
            };
        }
    }
}
</script>

<i18n>
{
    "en": {
        "title": "Assists rules",
        "stabilityControlLevelMax_label": "Stability Control Level Max: 0 to 100 max",
        "disableAutosteer_label": "Disable Auto-Steer",
        "disableAutoLights_label": "Disable Auto-Lights",
        "disableAutoWiper_label": "Disable Auto-Wiper",
        "disableAutoEngineStart_label": "Disable Auto-Engine Start",
        "disableAutoPitLimiter_label": "Disable Auto Pit Limiter",
        "disableAutoGear_label": "Disable Auto-Gear",
        "disableAutoClutch_label": "Disable Auto-Clutch",
        "disableIdealLine_label": "Disable Ideal Line"
    }
}
</i18n>
