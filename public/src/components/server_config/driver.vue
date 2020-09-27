<template>
    <div class="box">
        <field type="text" :label="$t('firstname_label')" v-model="firstName"></field>
        <field type="text" :label="$t('lastname_label')" v-model="lastName"></field>
        <field type="text" :label="$t('shortname_label')" v-model="shortName"></field>
        <selection :label="$t('drivercategory_label')" :options="driverCategoryTypes" v-model="driverCategory"></selection>
        <field type="text" :label="$t('playerid_label')" v-model="playerID"></field>
        <field type="number" :label="$t('defaultgridposition_label')" v-model="defaultGridPosition"></field>
        <button v-on:click="$emit('remove', driver.index)">{{$t("remove_button")}}</button>
    </div>
</template>

<script>
import field from "../field.vue";

export default {
    components: {field},
    props: ["driver"],
    watch: {
        firstName(value) {
            this.driver.firstName = value;
            this.$emit("update", this.driver);
        },
        lastName(value) {
            this.driver.lastName = value;
            this.$emit("update", this.driver);
        },
        shortName(value) {
            this.driver.shortName = value;
            this.$emit("update", this.driver);
        },
        driverCategory(value) {
            this.driver.driverCategory = parseInt(value);
            this.$emit("update", this.driver);
        },
        playerID(value) {
            this.driver.playerID = value;
            this.$emit("update", this.driver);
        },
        defaultGridPosition(value) {
            this.driver.defaultGridPosition = parseInt(value);
            this.$emit("update", this.driver);
        }
    },
    data() {
        return {
            firstName: "",
            lastName: "",
            shortName: "",
            driverCategoryTypes: [
              {value: 0, label: "Bronze"},
              {value: 1, label: "Silver"},
              {value: 2, label: "Gold"},
              {value: 3, label: "Platinum"}
            ],
            driverCategory: 0,
            playerID: "",
            defaultGridPosition: 0
        };
    },
    mounted() {
        this.firstName = this.driver.firstName;
        this.lastName = this.driver.lastName;
        this.shortName = this.driver.shortName;
        this.driverCategory = this.driver.driverCategory;
        this.playerID = this.driver.playerID;
        this.defaultGridPosition = this.driver.defaultGridPosition;
    }
}
</script>

<i18n>
{
    "en": {
        "firstname_label": "Firstname",
        "lastname_label": "Lastname",
        "shortname_label": "Short Name",
        "drivercategory_label": "Driver Category",
        "playerid_label": "PlayerID",
        "defaultgridposition_label": "Default Grid Position",
        "remove_button": "Remove Driver"
    }
}
</i18n>
