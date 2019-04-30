<template>
    <collapsible :title="$t('title')">
        <selection :label="$t('track_label')" :options="tracks" v-model="track"></selection>
        <selection :label="$t('eventtype_label')" :options="eventTypes" v-model="eventType"></selection>
        <field type="number" :label="$t('preracewaitingtimeseconds_label')" v-model="preRaceWaitingTimeSeconds"></field>
        <field type="number" :label="$t('sessionovertimeseconds_label')" v-model="sessionOverTimeSeconds"></field>
        <field type="number" :label="$t('ambienttemp_label')" v-model="ambientTemp"></field>
        <field type="number" :label="$t('tracktemp_label')" v-model="trackTemp"></field>
        <field type="number" :label="$t('cloudlevel_label')" :step="0.01" v-model="cloudLevel"></field>
        <field type="number" :label="$t('rain_label')" :step="0.01" v-model="rain"></field>
        <field type="number" :label="$t('weatherrandomness_label')" v-model="weatherRandomness"></field>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import field from "../field.vue";
import selection from "../selection.vue";

export default {
    components: {collapsible, field, selection},
    data() {
    	return {
    		tracks: [
    			"zolder",
    			"idontknowyet"
    		],
    		eventTypes: [
    			"E_3h"
    		],

    		track: "zolder",
			eventType: "E_3h",
			preRaceWaitingTimeSeconds: 15,
			sessionOverTimeSeconds: 120,
			ambientTemp: 26,
			trackTemp: 30,
			cloudLevel: 0.3,
			rain: 0.0,
			weatherRandomness: 1
    	};
    },
    methods: {
    	getData() {
    		return {
				track: this.track,
				eventType: this.eventType,
				preRaceWaitingTimeSeconds: this.preRaceWaitingTimeSeconds,
				sessionOverTimeSeconds: this.sessionOverTimeSeconds,
				ambientTemp: this.ambientTemp,
				trackTemp: this.trackTemp,
				cloudLevel: this.cloudLevel,
				rain: this.rain,
				weatherRandomness: this.weatherRandomness
    		};
    	}
    }
}
</script>

<i18n>
{
    "en": {
        "title": "Event settings",
        "track_label": "Track",
        "eventtype_label": "Event type",
        "preracewaitingtimeseconds_label": "Pre race waiting time seconds",
        "sessionovertimeseconds_label": "Session overtime seconds",
        "ambienttemp_label": "Ambient temp",
        "tracktemp_label": "Track temp",
        "cloudlevel_label": "Cloud level",
        "rain_label": "Rain",
        "weatherrandomness_label": "Weather randomness"
    }
}
</i18n>
