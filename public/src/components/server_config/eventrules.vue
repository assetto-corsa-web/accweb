<template>
    <collapsible :title="$t('title')" with-import="true" import-filename="eventRules.json" @load="setData">
        <selection :label="$t('qualifyStandingType_label')" :options="qualifyStandingTypes" v-model="qualifyStandingType"></selection>
        <div class="server-settings-container three-columns" style="margin-top: 1.0rem;">
            <field type="number" :label="$t('pitWindowLengthSec_label')" v-model="pitWindowLengthSec"></field>
            <field type="number" :label="$t('driverStintTimeSec_label')" v-model="driverStintTimeSec"></field>
            <field type="number" :label="$t('mandatoryPitstopCount_label')" v-model="mandatoryPitstopCount"></field>
            <field type="number" :label="$t('maxTotalDrivingTime_label')" v-model="maxTotalDrivingTime"></field>
            <field type="number" :label="$t('maxDriversCount_label')" v-model="maxDriversCount"></field>
            <field type="number" :label="$t('tyreSetCount_label')" v-model="tyreSetCount"></field>
        </div>
        <checkbox :label="$t('isRefuellingAllowedInRace_label')" v-model="isRefuellingAllowedInRace"></checkbox>
        <checkbox :label="$t('isRefuellingTimeFixed_label')" v-model="isRefuellingTimeFixed"></checkbox>
        <checkbox :label="$t('isMandatoryPitstopRefuellingRequired_label')" v-model="isMandatoryPitstopRefuellingRequired"></checkbox>
        <checkbox :label="$t('isMandatoryPitstopTyreChangeRequired_label')" v-model="isMandatoryPitstopTyreChangeRequired"></checkbox>
        <checkbox :label="$t('isMandatoryPitstopSwapDriverRequired_label')" v-model="isMandatoryPitstopSwapDriverRequired"></checkbox>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import field from "../field.vue";
import checkbox from "../checkbox.vue";
import selection from "../selection.vue";

export default {
    components: {collapsible, field, checkbox, selection},
    data() {
      return {
        qualifyStandingTypes: [
          {value: 1, label: "Fastest Lap"},
          {value: 2, label: "Average Lap"}
        ],
        qualifyStandingType: 1,
        pitWindowLengthSec: -1,
        driverStintTimeSec: -1,
        mandatoryPitstopCount: 0,
        maxTotalDrivingTime: -1,
        maxDriversCount: 1,
        tyreSetCount: 50,
        isRefuellingAllowedInRace: true,
        isRefuellingTimeFixed: false,
        isMandatoryPitstopRefuellingRequired: false,
        isMandatoryPitstopTyreChangeRequired: false,
        isMandatoryPitstopSwapDriverRequired: false,
        configVersion: 1
      };
    },
    methods: {
        setData(data) {
            this.qualifyStandingType = data.qualifyStandingType;
            this.pitWindowLengthSec = data.pitWindowLengthSec;
            this.driverStintTimeSec = data.driverStintTimeSec;
            this.mandatoryPitstopCount = data.mandatoryPitstopCount;
            this.maxTotalDrivingTime = data.maxTotalDrivingTime;
            this.maxDriversCount = data.maxDriversCount;
            this.tyreSetCount = data.tyreSetCount;
            this.isRefuellingAllowedInRace = data.isRefuellingAllowedInRace;
            this.isRefuellingTimeFixed = data.isRefuellingTimeFixed;
            this.isMandatoryPitstopRefuellingRequired = data.isMandatoryPitstopRefuellingRequired;
            this.isMandatoryPitstopTyreChangeRequired = data.isMandatoryPitstopTyreChangeRequired;
            this.isMandatoryPitstopSwapDriverRequired = data.isMandatoryPitstopSwapDriverRequired;
        },
        getData() {
            return {
                qualifyStandingType: parseInt(this.qualifyStandingType),
                pitWindowLengthSec: parseInt(this.pitWindowLengthSec),
                driverStintTimeSec: parseInt(this.driverStintTimeSec),
                mandatoryPitstopCount: parseInt(this.mandatoryPitstopCount),
                maxTotalDrivingTime: parseInt(this.maxTotalDrivingTime),
                maxDriversCount: parseInt(this.maxDriversCount),
                tyreSetCount: parseInt(this.tyreSetCount),
                isRefuellingAllowedInRace: this.isRefuellingAllowedInRace,
                isRefuellingTimeFixed: this.isRefuellingTimeFixed,
                isMandatoryPitstopRefuellingRequired: this.isMandatoryPitstopRefuellingRequired,
                isMandatoryPitstopTyreChangeRequired: this.isMandatoryPitstopTyreChangeRequired,
                isMandatoryPitstopSwapDriverRequired: this.isMandatoryPitstopSwapDriverRequired,
                configVersion: 1
            };
        }
    }
}
</script>

<i18n>
{
    "en": {
        "title": "Event rules",
        "qualifyStandingType_label": "Qualify Standing Type",
        "pitWindowLengthSec_label": "Pit Window Length (seconds)",
        "driverStintTimeSec_label": "Driver Stint Time (seconds)",
        "mandatoryPitstopCount_label": "Mandatory Pitstop Count",
        "maxTotalDrivingTime_label": "Max Total Driving Time",
        "maxDriversCount_label": "Max Drivers Count",
        "tyreSetCount_label": "Tyre Set Count",
        "isRefuellingAllowedInRace_label": "Is Refuelling Allowed in Race",
        "isRefuellingTimeFixed_label": "Is Refuelling Time Fixed",
        "isMandatoryPitstopRefuellingRequired_label": "Is Mandatory Pitstop Refuelling Required",
        "isMandatoryPitstopTyreChangeRequired_label": "Is Mandatory Pitstop Tyre Change Required",
        "isMandatoryPitstopSwapDriverRequired_label": "Is Mandatory Pitstop Swap Driver Required"
    }
}
</i18n>
