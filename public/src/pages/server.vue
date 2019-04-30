<template>
    <layout>
        <div class="title">
            <h1>TODO Servername</h1>
            <div class="menu">
                <button class="primary" v-on:click="save" v-if="is_admin"><i class="fas fa-save"></i> {{$t("save")}}</button>
                <button v-on:click="$router.push('/')" v-if="is_admin"><i class="fas fa-ban"></i> {{$t("cancel")}}</button>
                <button v-on:click="deleteServer" v-if="is_admin"><i class="fas fa-trash"></i> {{$t("delete")}}</button>
                <button class="primary" v-on:click="$router.push('/')" v-if="!is_admin"><i class="fas fa-arrow-left"></i> {{$t("back")}}</button>
            </div>
        </div>
        TODO Form or Import
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
        "delete": "Delete",
        "back": "Back"
    }
}
</i18n>
