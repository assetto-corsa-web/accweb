<template>
<div class="box">
    <h2>{{$t(resolveSessionType(session.sessionType))}}</h2>
    <div class="server-settings-container four-columns">
        <field type="number" :label="$t('hourofday_label')" v-model="session.hourOfDay"></field>
        <field type="number" :label="$t('dayofweekend_label')" v-model="session.dayOfWeekend"></field>
        <field type="number" :label="$t('sessiondurationminutes_label')" v-model="session.sessionDurationMinutes"></field>
        <field type="number" :label="$t('timemultiplier_label')" v-model="session.timeMultiplier"></field>
    </div>
    <selection :label="$t('type_label')" :options="types" v-model="session.sessionType"></selection>
    <v-btn small v-on:click="$emit('remove', session.index)"><i class="fas fa-ban"></i> {{$t("remove_button")}}</v-btn>
</div>
</template>

<script>
import field from "../field.vue";
import selection from "../selection.vue";

export default {
    components: {
        field,
        selection
    },
    props: ["session"],
    data() {
        return {
            types: [{
                    value: "P",
                    label: "Practice"
                },
                {
                    value: "Q",
                    label: "Qualifying"
                },
                {
                    value: "R",
                    label: "Race"
                }
            ]
        };
    },
    methods: {
        resolveSessionType(key) {
            const type = this.types.find(type => type.value === key);

            if (type === undefined) {
                return key;
            }

            return type.label;
        }
    }
}
</script>

<i18n>
{
    "en": {
        "hourofday_label": "Hour of day",
        "dayofweekend_label": "Day of weekend",
        "timemultiplier_label": "Time multiplier",
        "type_label": "Type",
        "sessiondurationminutes_label": "Session duration minutes",
        "remove_button": "Remove session",
        "Race": "Race",
        "Qualifying": "Qualifying",
        "Practice": "Practice"
    }
}
</i18n>
