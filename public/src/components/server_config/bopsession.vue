<template>
    <div class="box">
        <selection :label="$t('track_label')" :options="tracks" v-model="entry.track"></selection>
        <selection :label="$t('carModel_label')" :options="carModels" v-model="entry.carModel"></selection>
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
        carModels: [
          {value: 0, label: "Porsche 991 GT3"},
          {value: 1, label: "Mercedes AMG GT3"},
          {value: 2, label: "Ferrari 488 GT3"},
          {value: 3, label: "Audi R8 LMS"},
          {value: 4, label: "Lamborghini Huracan GT3"},
          {value: 5, label: "McLaren 650s GT3"},
          {value: 6, label: "Nissan GT R Nismo GT3 2018"},
          {value: 7, label: "BMW M6 GT3"},
          {value: 8, label: "Bentley Continental GT3 2018"},
          {value: 9, label: "Porsche 991.2 GT3 Cup"},
          {value: 10, label: "Nissan GT-R Nismo GT3 2017"},
          {value: 11, label: "Bentley Continental GT3 2016"},
          {value: 12, label: "Aston Martin Vantage V12 GT3"},
          {value: 13, label: "Lamborghini Gallardo R-EX"},
          {value: 14, label: "Jaguar G3"},
          {value: 15, label: "Lexus RC F GT3"},
          {value: 16, label: "Lamborghini Huracan Evo (2019)"},
          {value: 17, label: "Honda NSX GT3"},
          {value: 18, label: "Lamborghini Huracan SuperTrofeo"},
          {value: 19, label: "Audi R8 LMS Evo (2019)"},
          {value: 20, label: "AMR V8 Vantage (2019)"},
          {value: 21, label: "Honda NSX Evo (2019)"},
          {value: 22, label: "McLaren 720S GT3 (2019)"},
          {value: 23, label: "Porsche 911 II GT3 R (2019)"},
          {value: 50, label: "Alpine A1110 GT4"},
          {value: 51, label: "Aston Martin Vantage GT4"},
          {value: 52, label: "Audi R8 LMS GT4"},
          {value: 53, label: "BMW M4 GT4"},
          {value: 55, label: "Chevrolet Camaro GT4"},
          {value: 56, label: "Ginetta G55 GT4"},
          {value: 57, label: "KTM X-Bow GT4"},
          {value: 58, label: "Maserati MC GT4"},
          {value: 59, label: "McLaren 570S GT4"},
          {value: 60, label: "Mercedes AMG GT4"},
          {value: 61, label: "Porsche 718 Cayman GT4"}
        ],
        carModel: 0,
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