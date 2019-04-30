<template>
    <collapsible :title="$t('title')">
        <field :label="$t('servername_label')" v-model="serverName"></field>
        <field type="password" :label="$t('password_label')" v-model="password"></field>
        <field type="password" :label="$t('adminpassword_label')" v-model="adminPassword"></field>
        <field type="number" :label="$t('trackmedalsrequirement_label')" v-model="trackMedalsRequirement"></field>
        <field type="number" :label="$t('safetyratingrequirement_label')" v-model="safetyRatingRequirement"></field>
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
            trackMedalsRequirement: 3,
            safetyRatingRequirement: -1
        };
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
        "safetyratingrequirement_label": "Safety rating requirement"
    }
}
</i18n>
