<template>
    <collapsible :title="$t('title')">
        <field :label="$t('servername_label')" v-model="serverName"></field>
        <field type="password" :label="$t('password_label')" v-model="password"></field>
        <field type="password" :label="$t('adminpassword_label')" v-model="adminPassword"></field>
        <field type="number" :label="$t('trackmedalsrequirement_label')" v-model="trackMedalsRequirement"></field>
        <field type="number" :label="$t('safetyratingrequirement_label')" v-model="safetyRatingRequirement"></field>

        <field type="number" :label="$t('racecraftratingrequirement_label')" v-model="racecraftRatingRequirement"></field>
        <field type="number" :label="$t('spectatorslots_label')" v-model="spectatorSlots"></field>
        <field type="password" :label="$t('spectatorpassword_label')" v-model="spectatorPassword"></field>
        <field type="number" :label="$t('dumpleaderboards_label')" v-model="dumpLeaderboards"></field>
        <field type="number" :label="$t('isracelocked_label')" v-model="isRaceLocked"></field>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import field from "../field.vue";

export default {
    components: {collapsible, field},
    data() {
        return {
            serverName: "Servername",
            password: "",
            adminPassword: "",
            trackMedalsRequirement: 0,
            safetyRatingRequirement: -1,
            racecraftRatingRequirement: 0,
            spectatorSlots: 0,
            spectatorPassword: "",
            dumpLeaderboards: 0,
            isRaceLocked: 0
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
            this.spectatorSlots = data.spectatorSlots;
            this.spectatorPassword = data.spectatorPassword;
            this.dumpLeaderboards = data.dumpLeaderboards;
            this.isRaceLocked = data.isRaceLocked;
        },
        getData() {
            return {
                serverName: this.serverName,
                password: this.password,
                adminPassword: this.adminPassword,
                trackMedalsRequirement: parseInt(this.trackMedalsRequirement),
                safetyRatingRequirement: parseInt(this.safetyRatingRequirement),
                racecraftRatingRequirement: parseInt(this.racecraftRatingRequirement),
                spectatorSlots: parseInt(this.spectatorSlots),
                spectatorPassword: this.spectatorPassword,
                dumpLeaderboards: parseInt(this.dumpLeaderboards),
                isRaceLocked: parseInt(this.isRaceLocked)
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
        "spectatorslots_label": "Spectator Slots",
        "spectatorpassword_label": "Spectator Password",
        "dumpleaderboards_label": "Dump Leaderboards",
        "isracelocked_label": "Is Race Locked"
    }
}
</i18n>
