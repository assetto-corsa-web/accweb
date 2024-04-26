<template>
    <div class="box">
        <selection :label="$t('track_label')" :options="tracks" v-model="entry.track"></selection>
        <selection :label="$t('carModel_label')" :options="carModels" v-model="entry.carModel"></selection>
        <field type="number" :label="$t('ballast_label')" v-model="entry.ballastKg"></field>
        <field type="number" :label="$t('restrictor_label')" v-model="entry.restrictor"></field>
        <button v-on:click="$emit('remove', entry.index)">{{$t("remove_button")}}</button>
    </div>
</template>

<script>
import field from "../field.vue";
import selection from "../selection.vue";
import tracks from "../../data/tracks";
import cars from "../../data/cars";
import _ from "lodash";

export default {
    components: {field, selection},
    props: ["entry"],
    data() {
      return {
        tracks: _.orderBy(tracks, "label", "asc"),
        track: "barcelona",
        carModels: _.sortBy(
                _.mapValues(cars, function(o) { return {value: o.id.toString(), label: o.model, brand: o.brand}; }),
                ["brand", "label"]),
        carModel: 0,
        ballastKg: 0,
        restrictor: 0
      };
    }
}
</script>

<i18n>
{
    "en": {
        "track_label": "Track",
        "carModel_label": "Car Model #",
        "ballast_label": "Ballast: 0 to 100kg max.",
        "restrictor_label": "Motor Restrictor: 0 to 20% max.",
        "remove_button": "Remove BOP"
    }
}
</i18n>
