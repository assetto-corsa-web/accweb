<template>
    <div class="box">
        <selection :label="$t('track_label')" :options="tracks" v-model="entry.track"></selection>
		<field type="number" :label="$t('carModel_label')" v-model="entry.carModel"></field>
        <field type="number" :label="$t('ballast_label')" v-model="entry.ballast"></field>
        <field type="number" :label="$t('restrictor_label')" v-model="entry.restrictor"></field>
        <button v-on:click="$emit('remove', entry.index)">{{$t("remove_button")}}</button>
    </div>
</template>

<script>
import field from "../field.vue";
import selection from "../selection.vue";

export default {
    components: {field, selection},
    props: ["entry"],
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
                {value: "laguna_seca_2019", label: "Laguna Seca (2019)"}
    		],
			track: "spa",
			carModel: 99,
			ballast: 0,
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