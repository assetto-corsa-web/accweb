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

export default {
    components: {field, selection},
    props: ["entry"],
    data() {
      return {
        tracks: [
            {value: "barcelona", label: "Barcelona"},
            {value: "brands_hatch", label: "Brands Hatch"},
            {value: "donington", label: "Donnington"},
            {value: "hungaroring", label: "Hungaroring"},
            {value: "imola", label: "Imola"},
            {value: "kyalami", label: "Kyalami"},
            {value: "laguna_seca", label: "Laguna Seca"},
            {value: "misano", label: "Misano"},
            {value: "mount_panorama", label: "Mount Panorama"},
            {value: "nurburgring", label: "Nürburgring GP"},
            {value: "oulton_park", label: "Outlon Park"},
            {value: "paul_ricard", label: "Paul Ricard"},
            {value: "silverstone", label: "Silverstone"},
            {value: "snetterton", label: "Snetterton"},
            {value: "spa", label: "Spa"},
            {value: "suzuka", label: "Suzuka"},
            {value: "zandvoort", label: "Zandvoort"},
            {value: "zolder", label: "Zolder"},

        ],
        track: "barcelona",
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
          {value: 24, label: "Ferrari 488 GT3 Evo 2020"},
          {value: 25, label: "Mercedes-AMG GT3 2020"},
          {value: 30, label: "BMW M4 GT3"},
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
