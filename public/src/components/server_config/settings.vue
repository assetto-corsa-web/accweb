<template>
    <collapsible :title="$t('title')">
        <field :label="$t('servername_label')" v-model="serverName"></field>
        <field type="password" :label="$t('password_label')" v-model="password"></field>
        <field type="password" :label="$t('adminpassword_label')" v-model="adminPassword"></field>
        <field type="number" :label="$t('trackmedalsrequirement_label')" v-model="trackMedalsRequirement"></field>
        <field type="number" :label="$t('safetyratingrequirement_label')" v-model="safetyRatingRequirement"></field>
        <field type="number" :label="$t('racecraftratingrequirement_label')" v-model="racecraftRatingRequirement"></field>
        <field type="password" :label="$t('spectatorpassword_label')" v-model="spectatorPassword"></field>
        <field type="number" :label="$t('dumpleaderboards_label')" v-model="dumpLeaderboards"></field>
        <field type="number" :label="$t('isracelocked_label')" v-model="isRaceLocked"></field>
        <field type="number" :label="$t('randomizetrackwhenempty_label')" v-model="randomizeTrackWhenEmpty"></field>
        <field type="number" :label="$t('maxcarslots_label')" v-model="maxCarSlots"></field>
        <field type="text" :label="$t('centralentrylistpath_label')" v-model="centralEntryListPath"></field>
        <field type="number" :label="$t('shortformationlap_label')" v-model="shortFormationLap"></field>
        <field type="number" :label="$t('allowautodq_label')" v-model="allowAutoDQ"></field>
        <field type="number" :label="$t('dumpentrylist_label')" v-model="dumpEntryList"></field>
        <selection :label="$t('formationlaptype_label')" :options="formationLapTypes" v-model="formationLapType"></selection>
        <selection :label="$t('cargroup_label')" :options="carGroups" v-model="carGroup"></selection>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import field from "../field.vue";
import selection from "../selection.vue";

export default {
    components: {collapsible, field, selection},
    data() {
        return {
            serverName: "Servername",
            password: "",
            adminPassword: "",
            trackMedalsRequirement: 0,
            safetyRatingRequirement: -1,
            racecraftRatingRequirement: 0,
            spectatorPassword: "",
            dumpLeaderboards: 0,
            isRaceLocked: 0,
            randomizeTrackWhenEmpty: 0,
            maxCarSlots: 30,
            centralEntryListPath: "",
            shortFormationLap: 0,
            allowAutoDQ: 0,
            dumpEntryList: 0,
    		formationLapTypes: [    			
    			{value: "0", label: "0 - Old with limiter"},
    			{value: "1", label: "1 - Free No Limiter"},
				{value: "3", label: "3 - New Position control and UI"},
    		],
            formationLapType: 3,
    		carGroups: [    			
    			{value: "FreeForAll", label: "Mode FreeForAll"},
    			{value: "GT3_2018", label: "Mode GT3"},
				{value: "GT4", label: "Mode GT4"},
				{value: "CUP", label: "Mode CUP"},
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
            this.trackMedalsRequirement = data.trackMedalsRequirement;
            this.safetyRatingRequirement = data.safetyRatingRequirement;
            this.racecraftRatingRequirement = data.racecraftRatingRequirement;
            this.spectatorPassword = data.spectatorPassword;
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
                trackMedalsRequirement: parseInt(this.trackMedalsRequirement),
                safetyRatingRequirement: parseInt(this.safetyRatingRequirement),
                racecraftRatingRequirement: parseInt(this.racecraftRatingRequirement),
                spectatorPassword: this.spectatorPassword,
                dumpLeaderboards: parseInt(this.dumpLeaderboards),
                isRaceLocked: parseInt(this.isRaceLocked),
                randomizeTrackWhenEmpty: parseInt(this.randomizeTrackWhenEmpty),
                maxCarSlots: parseInt(this.maxCarSlots),
                centralEntryListPath: this.centralEntryListPath,
                shortFormationLap: parseInt(this.shortFormationLap),
                allowAutoDQ: parseInt(this.allowAutoDQ),
                dumpEntryList: parseInt(this.dumpEntryList),
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
        "trackmedalsrequirement_label": "Track medals requirement",
        "safetyratingrequirement_label": "Safety rating requirement",
        "racecraftratingrequirement_label": "Racecraft Training Requirement",
        "spectatorpassword_label": "Spectator Password",
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
