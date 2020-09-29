<template>
    <collapsible :title="$t('title')">
        <field :label="$t('servername_label')" v-model="serverName"></field>
        <field type="password" :label="$t('password_label')" v-model="password"></field>
        <field type="password" :label="$t('adminpassword_label')" v-model="adminPassword"></field>
        <field type="password" :label="$t('spectatorpassword_label')" v-model="spectatorPassword"></field>
        <field type="number" :label="$t('trackmedalsrequirement_label')" v-model="trackMedalsRequirement"></field>
        <field type="number" :label="$t('safetyratingrequirement_label')" v-model="safetyRatingRequirement"></field>
        <field type="number" :label="$t('racecraftratingrequirement_label')" v-model="racecraftRatingRequirement"></field>
        <checkbox :label="$t('dumpleaderboards_label')" v-model="dumpLeaderboards"></checkbox>
        <checkbox :label="$t('isracelocked_label')" v-model="isRaceLocked"></checkbox>
        <checkbox :label="$t('randomizetrackwhenempty_label')" v-model="randomizeTrackWhenEmpty"></checkbox>
        <field type="number" :label="$t('maxcarslots_label')" v-model="maxCarSlots"></field>
        <field type="text" :label="$t('centralentrylistpath_label')" v-model="centralEntryListPath"></field>
        <checkbox :label="$t('shortformationlap_label')" v-model="shortFormationLap"></checkbox>
        <checkbox :label="$t('allowautodq_label')" v-model="allowAutoDQ"></checkbox>
        <checkbox :label="$t('dumpentrylist_label')" v-model="dumpEntryList"></checkbox>
        <selection :label="$t('formationlaptype_label')" :options="formationLapTypes" v-model="formationLapType"></selection>
        <selection :label="$t('cargroup_label')" :options="carGroups" v-model="carGroup"></selection>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import field from "../field.vue";
import selection from "../selection.vue";
import checkbox from "../checkbox.vue";

export default {
    components: {collapsible, field, selection, checkbox},
    data() {
      return {
        serverName: "Servername",
        password: "",
        adminPassword: "",
        spectatorPassword: "",
        trackMedalsRequirement: -1,
        safetyRatingRequirement: -1,
        racecraftRatingRequirement: -1,
        dumpLeaderboards: false,
        isRaceLocked: false,
        randomizeTrackWhenEmpty: false,
        maxCarSlots: 30,
        centralEntryListPath: "",
        shortFormationLap: false,
        allowAutoDQ: false,
        dumpEntryList: false,
        formationLapTypes: [
          {value: 0, label: "Old with limiter"},
          {value: 1, label: "Free No Limiter"},
          {value: 3, label: "New Position control and UI"},
        ],
        formationLapType: 3,
        carGroups: [
          {value: "FreeForAll", label: "Mode FreeForAll"},
          {value: "GT3", label: "Mode GT3"},
          {value: "GT4", label: "Mode GT4"},
          {value: "Cup", label: "Mode CUP"},
          {value: "ST", label: "Mode ST"}
        ],
        carGroup: "FreeForAll"
      };
    },
    methods: {
        setData(data) {
          this.serverName = data.serverName;
          this.password = data.password;
          this.adminPassword = data.adminPassword;
          this.spectatorPassword = data.spectatorPassword;
          this.trackMedalsRequirement = data.trackMedalsRequirement;
          this.safetyRatingRequirement = data.safetyRatingRequirement;
          this.racecraftRatingRequirement = data.racecraftRatingRequirement;
          this.dumpLeaderboards = data.dumpLeaderboards;
          this.isRaceLocked = data.isRaceLocked;
          this.randomizeTrackWhenEmpty = data.randomizeTrackWhenEmpty;
          this.maxCarSlots = data.maxCarSlots;
          this.centralEntryListPath = data.centralEntryListPath;
          this.shortFormationLap = data.shortFormationLap;
          this.allowAutoDQ = data.allowAutoDQ;
          this.dumpEntryList = data.dumpEntryList;
          this.formationLapType = data.formationLapType;
          this.carGroup = data.carGroup;
        },
        getData() {
          return {
            serverName: this.serverName,
            password: this.password,
            adminPassword: this.adminPassword,
            spectatorPassword: this.spectatorPassword,
            trackMedalsRequirement: parseInt(this.trackMedalsRequirement),
            safetyRatingRequirement: parseInt(this.safetyRatingRequirement),
            racecraftRatingRequirement: parseInt(this.racecraftRatingRequirement),
            dumpLeaderboards: this.dumpLeaderboards ? 1 : 0,
            isRaceLocked: this.isRaceLocked ? 1 : 0,
            randomizeTrackWhenEmpty: this.randomizeTrackWhenEmpty ? 1 : 0,
            maxCarSlots: parseInt(this.maxCarSlots),
            centralEntryListPath: this.centralEntryListPath,
            shortFormationLap: this.shortFormationLap ? 1 : 0,
            allowAutoDQ: this.allowAutoDQ ? 1 : 0,
            dumpEntryList: this.dumpEntryList ? 1 : 0,
            formationLapType: parseInt(this.formationLapType),
            carGroup: this.carGroup
          };
        }
    }
}
</script>

<i18n>
{
    "en": {
        "title": "Server settings",
        "servername_label": "Servername",
        "password_label": "Password",
        "adminpassword_label": "Admin password",
        "spectatorpassword_label": "Spectator Password",
        "trackmedalsrequirement_label": "Track medals requirement (-1 - 3)",
        "safetyratingrequirement_label": "Safety rating requirement (-1 - 99)",
        "racecraftratingrequirement_label": "Racecraft Training Requirement (-1 - 99)",
        "dumpleaderboards_label": "Dump Leaderboards",
        "isracelocked_label": "Is Race Locked",
        "randomizetrackwhenempty_label": "Randomize Track When Empty",
        "maxcarslots_label": "Max Car Slots",
        "centralentrylistpath_label": "Path to the central entrylist.json file - Do not put entrylist.json at the end of the path !",
        "shortformationlap_label": "Short Formation Lap",
        "allowautodq_label": "Allow Auto DQ",
        "dumpentrylist_label": "Dump Entry List",
        "formationlaptype_label": "Formation Lap Type",
		    "cargroup_label": "Car Group"
    }
}
</i18n>
