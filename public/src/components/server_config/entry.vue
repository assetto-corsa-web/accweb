<template>
    <div class="box">
        <field type="number" :label="$t('racenumber_label')" v-model="entry.raceNumber"></field>
        <selection :label="$t('forcedcarmodel_label')" :options="carModels" v-model="entry.forcedCarModel"></selection>
        <checkbox :label="$t('overridedriverinfo_label')" v-model="entry.overrideDriverInfo"></checkbox>
        <checkbox :label="$t('isserveradmin_label')" v-model="entry.isServerAdmin"></checkbox>
        <field type="text" :label="$t('customcar_label')" v-model="entry.customCar"></field>
        <checkbox :label="$t('overridecarmodelforcustomcar_label')" v-model="entry.overrideCarModelForCustomCar"></checkbox>
        <field type="number" :label="$t('ballast_label')" v-model="entry.ballastKg"></field>
        <field type="number" :label="$t('restrictor_label')" v-model="entry.restrictor"></field>
        <field type="number" :label="$t('defaultgridposition_label')" v-model="entry.defaultGridPosition"></field>
        <driver v-for="driver in drivers"
            :key="driver.index"
            :driver="driver"
            v-on:update="updateDriver"
            v-on:remove="removeDriver"></driver>
        <button v-on:click="addDriver">{{$t("add_driver_button")}}</button>
        <button v-on:click="$emit('remove', entry.index)">{{$t("remove_button")}}</button>
    </div>
</template>

<script>
import field from "../field.vue";
import driver from "./driver.vue";
import selection from "../selection.vue";
import checkbox from "../checkbox.vue";

export default {
    components: {field, driver, selection, checkbox},
    props: ["entry"],
    data() {
        return {
            driverIndex: 0,
            carModels: [
              {value: -1, label: "None"},
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
            drivers: []
        };
    },
    mounted() {
        if(this.entry && this.entry.drivers) {
            this.drivers = this.entry.drivers;
        }
    },
    methods: {
        updateDriver(driver) {
            let index = parseInt(driver.index);

            for(let i = 0; i < this.drivers.length; i++) {
                if(this.drivers[i].index === index) {
                    this.drivers[i] = driver;
                    break;
                }
            }

            this.entry.drivers = this.drivers;
        },
        addDriver() {
            this.drivers.push({
                index: this.driverIndex,
                firstName: "",
                lastName: "",
                shortName: "",
                driverCategory: 0,
                playerID: ""
            });
            this.driverIndex++;
        },
        removeDriver(index) {
            index = parseInt(index);

            for(let i = 0; i < this.drivers.length; i++) {
                if(this.drivers[i].index === index) {
                    this.drivers.splice(i, 1);
                    break;
                }
            }
        }
    }
}
</script>

<i18n>
{
    "en": {
        "racenumber_label": "Race Number",
        "forcedcarmodel_label": "Forced Car Model",
        "overridedriverinfo_label": "Override Driver Info",
        "isserveradmin_label": "Is Server Admin",
        "customcar_label": "Custom Car",
        "overridecarmodelforcustomcar_label": "Override Car Model For Custom Car",
        "ballast_label": "Ballast (0 - 100kg)",
        "restrictor_label": "Restrictor (0 - 20%)",
        "defaultgridposition_label": "Default Grid Position",
        "add_driver_button": "Add Driver",
        "remove_button": "Remove Entry"
    }
}
</i18n>
