<template>
    <collapsible :title="$t('title')">
        <selection :label="$t('track_label')" :options="tracks" v-model="track"></selection>
        <selection :label="$t('eventtype_label')" :options="eventTypes" v-model="eventType"></selection>
        <field type="number" :label="$t('preracewaitingtimeseconds_label')" v-model="preRaceWaitingTimeSeconds"></field>
        <field type="number" :label="$t('sessionovertimeseconds_label')" v-model="sessionOverTimeSeconds"></field>
        <field type="number" :label="$t('ambienttemp_label')" v-model="ambientTemp"></field>
        <field type="number" :label="$t('cloudlevel_label')" :step="0.01" v-model="cloudLevel"></field>
        <field type="number" :label="$t('rain_label')" :step="0.01" v-model="rain"></field>
        <field type="number" :label="$t('weatherrandomness_label')" v-model="weatherRandomness"></field>
        <field type="number" :label="$t('postqualyseconds_label')" v-model="postQualySeconds"></field>
        <field type="number" :label="$t('postraceseconds_label')" v-model="postRaceSeconds"></field>
        <session v-for="session in sessions"
            :key="session.index"
            :session="session"
            v-on:remove="removeSession"></session>
        <button v-on:click="addSession">{{$t("add_session_button")}}</button>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import field from "../field.vue";
import selection from "../selection.vue";
import session from "./session.vue";

export default {
    components: {collapsible, field, selection, session},
    data() {
    	return {
    		tracks: [
                {value: "misano", label: "Misano"},
                {value: "paul_ricard", label: "Paul Ricard"},
                {value: "nurburgring", label: "Nürburgring GP"},
                {value: "hungaroring", label: "Hungaroring"},
                {value: "zolder", label: "Zolder"},
    			{value: "monza", label: "Monza"},
                {value: "brands_hatch", label: "Brands Hatch"},
                {value: "barcelona", label: "Barcelona"},
                {value: "silverstone", label: "Silverstone"},
                {value: "spa", label: "Spa"},
                {value: "zandvoort", label: "Zandvoort"},
                {value: "monza_2019", label: "Monza (2019)"},
                {value: "zolder_2019", label: "Zolder (2019)"},
                {value: "brands_hatch_2019", label: "Brands Hatch (2019)"},
                {value: "silverstone_2019", label: "Silverstone (2019)"},
                {value: "paul_ricard_2019", label: "Paul Ricard (2019)"},
                {value: "misano_2019", label: "Misano (2019)"},
                {value: "spa_2019", label: "Spa (2019)"},
                {value: "nurburgring_2019", label: "Nürburgring (2019)"},
                {value: "barcelona_2019", label: "Barcelona (2019)"},
                {value: "hungaroring_2019", label: "Hungaroring (2019)"},
                {value: "zandvoort_2019", label: "Zandvoort (2019)"},
                {value: "kyalami_2019", label: "Kyalami (2019)"},
                {value: "mount_panorama_2019", label: "Mount Panorama (2019)"},
                {value: "suzuka_2019", label: "Suzuka (2019)"},
                {value: "laguna_seca_2019", label: "Laguna Seca (2019)"},
    		],
    		eventTypes: [
    			{value: "E_3h", label: "Endurance - 3h"},
                {value: "E_6h", label: "Endurance - 6h"},
    		],

    		track: "misano",
			eventType: "E_3h",
			preRaceWaitingTimeSeconds: 15,
			sessionOverTimeSeconds: 120,
			ambientTemp: 26,
			cloudLevel: 0.3,
			rain: 0.0,
			weatherRandomness: 1,
            postQualySeconds: 0,
            postRaceSeconds: 0,
            
            sessionIndex: 0,
            sessions: []
    	};
    },
    methods: {
        setData(data) {
            this.track = data.track;
            this.eventType = data.eventType;
            this.preRaceWaitingTimeSeconds = data.preRaceWaitingTimeSeconds;
            this.sessionOverTimeSeconds = data.sessionOverTimeSeconds;
            this.ambientTemp = data.ambientTemp;
            this.cloudLevel = data.cloudLevel;
            this.rain = data.rain;
            this.weatherRandomness = data.weatherRandomness;
            this.postQualySeconds = data.postQualySeconds;
            this.postRaceSeconds = data.postRaceSeconds;
            this.setSessionData(data.sessions);
        },
        setSessionData(data) {
            for(let i = 0; i < data.length; i++) {
                this.sessions.push({
                    index: this.sessionIndex,
                    hourOfDay: data[i].hourOfDay,
                    dayOfWeekend: data[i].dayOfWeekend,
                    timeMultiplier: data[i].timeMultiplier,
                    sessionType: data[i].sessionType,
                    sessionDurationMinutes: data[i].sessionDurationMinutes
                });
                this.sessionIndex++;
            }
        },
    	getData() {
    		return {
				track: this.track,
				eventType: this.eventType,
				preRaceWaitingTimeSeconds: parseInt(this.preRaceWaitingTimeSeconds),
				sessionOverTimeSeconds: parseInt(this.sessionOverTimeSeconds),
				ambientTemp: parseInt(this.ambientTemp),
				cloudLevel: this.toFloat(this.cloudLevel),
				rain: this.toFloat(this.rain),
				weatherRandomness: parseInt(this.weatherRandomness),
                postQualySeconds: parseInt(this.postQualySeconds),
                postRaceSeconds: parseInt(this.postRaceSeconds),
                sessions: this.getSessionData()
    		};
    	},
        getSessionData() {
            let sessions = [];

            for(let i = 0; i < this.sessions.length; i++) {
                sessions.push({
                    hourOfDay: parseInt(this.sessions[i].hourOfDay),
                    dayOfWeekend: parseInt(this.sessions[i].dayOfWeekend),
                    timeMultiplier: parseInt(this.sessions[i].timeMultiplier),
                    sessionType: this.sessions[i].sessionType,
                    sessionDurationMinutes: parseInt(this.sessions[i].sessionDurationMinutes)
                });
            }

            return sessions;
        },
        addSession() {
            this.sessions.push({
                index: this.sessionIndex,
                hourOfDay: 9,
                dayOfWeekend: 1,
                timeMultiplier: 1,
                sessionType: "P",
                sessionDurationMinutes: 7
            });
            this.sessionIndex++;
        },
        removeSession(index) {
            index = parseInt(index);

            for(let i = 0; i < this.sessions.length; i++) {
                if(this.sessions[i].index === index) {
                    this.sessions.splice(i, 1);
                    break;
                }
            }
        },
        toFloat(value) {
            if(typeof value === "string") {
                return parseFloat(value.replace(",", "."));
            }

            return value;
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
        "cloudlevel_label": "Cloud level",
        "rain_label": "Rain",
        "weatherrandomness_label": "Weather randomness",
        "add_session_button": "Add session",
        "postqualyseconds_label": "Post Qualy Seconds",
        "postraceseconds_label": "Post Race Seconds"
    }
}
</i18n>
