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
        <div class="tabs">
            <div v-bind:class="{tab: true, 'tab-active': activeTab === 0}" v-on:click="activeTab = 0">{{$t("create_server")}}</div>
            <div v-bind:class="{tab: true, 'tab-active': activeTab === 1}" v-on:click="activeTab = 1">{{$t("import_server")}}</div>
        </div>
        <div v-show="activeTab === 0">
            <basic></basic>
            <settings></settings>
            <event></event>
        </div>
        <div v-show="activeTab === 1">
            <p>{{$t("upload_hint")}}</p>
            <label>configuration.json</label>
            <input type="file" name="configuration.json" />
            <label>settings.json</label>
            <input type="file" name="settings.json" />
            <label>event.json</label>
            <input type="file" name="event.json" />
            <input class="primary" type="submit" :value="$t('import_button')" />
        </div>
    </layout>
</template>

<script>
import {layout, end, basic, settings, event} from "../components";

export default {
    components: {layout, end, basic, settings, event},
    data() {
        return {
            activeTab: 0
        };
    },
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
        "back": "Back",
        "create_server": "Create server",
        "import_server": "Import server",
        "upload_hint": "You can import an existing server by uploading its configuration files. Select the files of your server configuration and press import.",
        "import_button": "Import"
    }
}
</i18n>
