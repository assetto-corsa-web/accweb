<template>
    <div class="box">
        <h2>{{$t(resolveSessionType(session.sessionType))}}</h2>
        <field type="number" :label="$t('hourofday_label')" v-model="session.hourOfDay"></field>
        <field type="number" :label="$t('dayofweekend_label')" v-model="session.dayOfWeekend"></field>
        <field type="number" :label="$t('timemultiplier_label')" v-model="session.timeMultiplier"></field>
        <selection :label="$t('type_label')" :options="types" v-model="session.sessionType"></selection>
        <field type="number" :label="$t('sessiondurationminutes_label')" v-model="session.sessionDurationMinutes"></field>
        <button v-on:click="$emit('remove', session.index)">{{$t("remove_button")}}</button>
    </div>
</template>

<script>
import field from "../field.vue";
import selection from "../selection.vue";

export default {
    components: {field, selection},
    props: ["session"],
    data() {
    	  return {
            types: [
                {value: "P", label: "Practice"},
                {value: "Q", label: "Qualify"},
                {value: "R", label: "Race"}
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
