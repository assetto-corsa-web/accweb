<template>
    <layout>
        <div class="title">
            <h1>TODO Servername</h1>
            <div class="menu">
                <button class="primary" v-on:click="save"><i class="fas fa-save"></i> {{$t("save")}}</button>
                <button v-on:click="$router.push('/')"><i class="fas fa-ban"></i> {{$t("cancel")}}</button>
                <button v-on:click="deleteServer"><i class="fas fa-trash"></i> {{$t("delete")}}</button>
            </div>
        </div>
        TODO Form
    </layout>
</template>

<script>
import {layout, end} from "../components";

export default {
    components: {layout, end},
    methods: {
        save() {

        },
        deleteServer() {

        }
    }
}
</script>

<i18n>
{
    "en": {
        "save": "Save",
        "cancel": "Cancel",
        "delete": "Delete"
    }
}
</i18n>
