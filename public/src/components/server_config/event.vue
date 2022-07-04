<template>
    <collapsible :title="$t('title')" with-import="true" import-filename="event.json" @load="setData">
        <div style="margin-bottom: 1.0rem;">
            <selection :label="$t('track_label')" :options="tracks" v-model="track"></selection>
        </div>
        <div class="server-settings-container two-columns">
            <div>
                <field type="number" :label="$t('ambienttemp_label')" v-model="ambientTemp"></field>
                <field type="number" :label="$t('tracktemp_label')" v-model="trackTemp"></field>
                <field type="number" :label="$t('cloudlevel_label')" :step="0.01" v-model="cloudLevel"></field>
                <field type="number" :label="$t('rain_label')" :step="0.01" v-model="rain"></field>
                <field type="number" :label="$t('weatherrandomness_label')" v-model="weatherRandomness"></field>
                <checkbox :label="$t('simracerWeatherConditions_label')" v-model="simracerWeatherConditions"></checkbox>
                <checkbox :label="$t('isFixedConditionQualification_label')" v-model="isFixedConditionQualification"></checkbox>
            </div>
            <div>
                <field type="number" :label="$t('preracewaitingtimeseconds_label')" v-model="preRaceWaitingTimeSeconds"></field>
                <field type="number" :label="$t('sessionovertimeseconds_label')" v-model="sessionOverTimeSeconds"></field>
                <field type="number" :label="$t('postqualyseconds_label')" v-model="postQualySeconds"></field>
                <field type="number" :label="$t('postraceseconds_label')" v-model="postRaceSeconds"></field>
            </div>
        </div>
        <session v-for="session in sessions"
            :key="session.index"
            :session="session"
            v-on:remove="removeSession"></session>
        <div class="button-row">
          <button v-on:click="addSession">{{$t("add_session_button")}}</button>
          <button v-on:click="addDefaultSessions('Q/R')">{{$t("add_sessions_q_r_button")}}</button>
          <button v-on:click="addDefaultSessions('P/Q/R')">{{$t("add_sessions_p_q_r_button")}}</button>
          <button v-on:click="clearSessions">{{$t("clear_sessions_button")}}</button>
        </div>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import field from "../field.vue";
import selection from "../selection.vue";
import session from "./session.vue";
import checkbox from "../checkbox.vue";
import tracks from "../../data/tracks";

export default {
    components: {collapsible, field, selection, session, checkbox},
    data() {
    	  return {
            tracks: tracks,
            track: "barcelona",
            preRaceWaitingTimeSeconds: 15,
            sessionOverTimeSeconds: 120,
            ambientTemp: 26,
            trackTemp: 26,
            cloudLevel: 0.3,
            rain: 0.0,
            weatherRandomness: 1,
            postQualySeconds: 0,
            postRaceSeconds: 0,
            simracerWeatherConditions: false,
            isFixedConditionQualification: false,
            sessionIndex: 0,
            sessions: []
        };
    },
    methods: {
        setData(data) {
            this.track = data.track;
            this.preRaceWaitingTimeSeconds = data.preRaceWaitingTimeSeconds;
            this.sessionOverTimeSeconds = data.sessionOverTimeSeconds;
            this.ambientTemp = data.ambientTemp;
            this.trackTemp = data.trackTemp;
            this.cloudLevel = data.cloudLevel;
            this.rain = data.rain;
            this.weatherRandomness = data.weatherRandomness;
            this.postQualySeconds = data.postQualySeconds;
            this.postRaceSeconds = data.postRaceSeconds;
            this.simracerWeatherConditions = data.simracerWeatherConditions;
            this.isFixedConditionQualification = data.isFixedConditionQualification;
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
                preRaceWaitingTimeSeconds: parseInt(this.preRaceWaitingTimeSeconds),
                sessionOverTimeSeconds: parseInt(this.sessionOverTimeSeconds),
                ambientTemp: parseInt(this.ambientTemp),
                trackTemp: parseInt(this.trackTemp),
                cloudLevel: this.toFloat(this.cloudLevel),
                rain: this.toFloat(this.rain),
                weatherRandomness: parseInt(this.weatherRandomness),
                postQualySeconds: parseInt(this.postQualySeconds),
                postRaceSeconds: parseInt(this.postRaceSeconds),
                simracerWeatherConditions: this.simracerWeatherConditions ? 1 : 0,
                isFixedConditionQualification: this.isFixedConditionQualification ? 1 : 0,
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
        addDefaultSessions(type) {
            if (type === 'P/Q/R') {
                this.sessions.push({
                    index: this.sessionIndex++,
                    hourOfDay: 9,
                    dayOfWeekend: 1,
                    timeMultiplier: 1,
                    sessionType: "P",
                    sessionDurationMinutes: 7
                });
            }

            this.sessions.push({
                index: this.sessionIndex++,
                hourOfDay: 9,
                dayOfWeekend: 2,
                timeMultiplier: 1,
                sessionType: "Q",
                sessionDurationMinutes: 15
            });

            this.sessions.push({
                index: this.sessionIndex++,
                hourOfDay: 9,
                dayOfWeekend: 3,
                timeMultiplier: 1,
                sessionType: "R",
                sessionDurationMinutes: 30
            });
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
        clearSessions() {
            if (!window.confirm(this.$t("confirm_clear_sessions"))) {
                return;
            }
            while(this.sessions.length > 0) {
                this.sessions.splice(this.sessions[this.sessions.length - 1], 1);
            }
        },
        toFloat(value) {
            if(typeof value === "string") {
                return parseFloat(value.replace(",", "."), 64);
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
        "preracewaitingtimeseconds_label": "Pre race waiting time seconds",
        "sessionovertimeseconds_label": "Session overtime seconds",
        "ambienttemp_label": "Ambient temp",
        "tracktemp_label": "Track temp (empty to omit)",
        "cloudlevel_label": "Cloud level",
        "rain_label": "Rain",
        "weatherrandomness_label": "Weather randomness",
        "postqualyseconds_label": "Post Qualy Seconds",
        "postraceseconds_label": "Post Race Seconds",
        "simracerWeatherConditions_label": "Simracer Weather Conditions",
        "isFixedConditionQualification_label": "Is Fixed Weather Conditions in Qualification",
        "add_session_button": "Add session",
        "add_sessions_q_r_button": "Add sessions (Q/R)",
        "add_sessions_p_q_r_button": "Add sessions (P/Q/R)",
        "clear_sessions_button": "Clear all Sessions",
        "confirm_clear_sessions": "Do you really want to remove all sessions?"
    }
}
</i18n>
