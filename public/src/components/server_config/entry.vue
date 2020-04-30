<template>
    <div class="box">
        <field type="number" :label="$t('racenumber_label')" v-model="entry.raceNumber"></field>
        <field type="number" :label="$t('forcedcarmodel_label')" v-model="entry.forcedCarModel"></field>
        <field type="number" :label="$t('overridedriverinfo_label')" v-model="entry.overrideDriverInfo"></field>
        <field type="number" :label="$t('isserveradmin_label')" v-model="entry.isServerAdmin"></field>
        <field type="text" :label="$t('customcar_label')" v-model="entry.customCar"></field>
        <field type="number" :label="$t('overridecarmodelforcustomcar_label')" v-model="entry.overrideCarModelForCustomCar"></field>
        <field type="number" :label="$t('ballast_label')" v-model="entry.ballastKg"></field>
        <field type="number" :label="$t('restrictor_label')" v-model="entry.restrictor"></field>
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

export default {
    components: {field, driver},
    props: ["entry"],
    data() {
        return {
            driverIndex: 0,
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
        "add_driver_button": "Add Driver",
        "remove_button": "Remove Entry"
    }
}
</i18n>
