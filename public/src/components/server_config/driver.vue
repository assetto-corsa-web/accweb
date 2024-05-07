<template>
    <div class="box">
        <div class="server-settings-container three-columns">
            <field type="text" :label="$t('firstname_label')" v-model="firstName"></field>
            <field type="text" :label="$t('lastname_label')" v-model="lastName"></field>
            <field type="text" :label="$t('shortname_label')" v-model="shortName"></field>
        </div>
        <div class="server-settings-container three-columns">
            <selection :label="$t('nationality_label')" :options="nationalities" v-model="nationality"></selection>
            <selection :label="$t('drivercategory_label')" :options="driverCategoryTypes" v-model="driverCategory"></selection>
            <field type="text" :label="$t('playerid_label')" v-model="playerID"></field>
        </div>
        <button v-on:click="$emit('remove', driver.index)">{{$t("remove_button")}}</button> 
    </div>
</template>

<script>
import field from "../field.vue";
import selection from "../selection.vue";
import nationalities from "../../data/nationalities";
import _ from "lodash";

export default {
    components: {field, selection},
    props: ["driver"],
    watch: {
        firstName(value) {
            this.driver.firstName = value == '' ? undefined : value;
            this.$emit("update", this.driver);
        },
        lastName(value) {
            this.driver.lastName = value == '' ? undefined : value;;
            this.$emit("update", this.driver);
        },
        shortName(value) {
            this.driver.shortName = value == '' ? undefined : value;;
            this.$emit("update", this.driver);
        },
        driverCategory(value) {
            this.driver.driverCategory = parseInt(value) == 0 ? undefined : parseInt(value);
            this.$emit("update", this.driver);
        },
        playerID(value) {
            this.driver.playerID = value;
            this.$emit("update", this.driver);
        },
        nationality(value) {
            this.driver.nationality = parseInt(value) == 0 ? undefined : parseInt(value);
            this.$emit("update", this.driver);
        }
    },
    data() {
        return {
            firstName: undefined,
            lastName: undefined,
            shortName: undefined,
            driverCategoryTypes: [
              {value: 0, label: "Bronze"},
              {value: 1, label: "Silver"},
              {value: 2, label: "Gold"},
              {value: 3, label: "Platinum"}
            ],
            driverCategory: undefined,
            playerID: "",
            nationalities: _.mapValues(_.orderBy(nationalities, "country", "asc"), function(o) { return {value: o.id, label: o.country}; }),
            nationality: undefined,
        };
    },
    mounted() {
        this.firstName = this.driver.firstName;
        this.lastName = this.driver.lastName;
        this.shortName = this.driver.shortName;
        this.driverCategory = parseInt(this.driver.driverCategory);
        this.playerID = this.driver.playerID;
        this.nationality = this.driver.nationality
    }
}
</script>

<i18n>
{
    "en": {
        "firstname_label": "First Name",
        "lastname_label": "Last Name",
        "shortname_label": "Short Name",
        "drivercategory_label": "Driver Category",
        "playerid_label": "PlayerID",
        "remove_button": "Remove Driver",
        "nationality_label": "Nationality"
    }
}
</i18n>
