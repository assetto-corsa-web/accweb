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
import cars from "../../data/cars";
import _ from "lodash";

export default {
    components: {field, driver, selection, checkbox},
    props: ["entry"],
    data() {
        return {
            driverIndex: 0,
            carModels: _.sortBy(
                _.mapValues(cars, function(o) { return {value: o.id.toString(), label: o.model, brand: o.brand}; }),
                ["brand", "label"]),
            drivers: []
        };
    },
    mounted() {
        if(this.entry && this.entry.drivers) {
            this.drivers = this.entry.drivers;

            for(let i = 0; i < this.drivers.length; i++) {
                this.drivers[i].index = this.driverIndex = i;
            }
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
                firstName: undefined,
                lastName: undefined,
                shortName: undefined,
                driverCategory: undefined,
                playerID: "",
                nationality: undefined,
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
