<template>
    <layout>
        <div class="title">
            <h1>{{servername}}</h1>
            <div class="menu">
                <button class="primary" v-on:click="save" v-if="is_admin"><i class="fas fa-save"></i> {{$t("save")}}</button>
                <button v-on:click="$router.push('/')" v-if="is_admin"><i class="fas fa-ban"></i> {{$t("cancel")}}</button>
                <button class="primary" v-on:click="$router.push('/')" v-if="!is_admin"><i class="fas fa-arrow-left"></i> {{$t("back")}}</button>
            </div>
        </div>
        <div class="tabs">
            <div v-bind:class="{tab: true, 'tab-active': activeTab === 0}" v-on:click="activeTab = 0">{{$t("create_server")}}</div>
            <div v-bind:class="{tab: true, 'tab-active': activeTab === 1}" v-on:click="activeTab = 1" v-if="is_admin">{{$t("import_server")}}</div>
        </div>
        <div v-show="activeTab === 0">
            <basic ref="basic"></basic>
            <settings ref="settings"></settings>
            <event ref="event"></event>
        </div>
        <div v-show="activeTab === 1 && is_admin">
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
import axios from "axios";
import {layout, end, basic, settings, event} from "../components";

export default {
    components: {layout, end, basic, settings, event},
    data() {
        return {
            activeTab: 0,
            id: 0,
            servername: "New Server"
        };
    },
    mounted() {
        this.id = this.$route.query.id;

        if(this.id) {
            this.loadServer();
        }
    },
    methods: {
        loadServer() {
            axios.get("/api/server", {params: {id: this.id}})
            .then(r => {
                this.servername = r.data.settings.serverName;
                this.$refs.basic.setData(r.data.basic);
                this.$refs.settings.setData(r.data.settings);
                this.$refs.event.setData(r.data.event);
            });
        },
        save() {
            let basic = this.$refs.basic.getData();
            let settings = this.$refs.settings.getData();
            let event = this.$refs.event.getData();
            let data = {
                id: parseInt(this.id),
                basic,
                settings,
                event
            };

            axios.post("/api/server", data)
            .then(() => {
                this.$router.push("/");
            });
        }
    }
}
</script>

<i18n>
{
    "en": {
        "save": "Save",
        "cancel": "Cancel",
        "back": "Back",
        "create_server": "Create server",
        "import_server": "Import server",
        "upload_hint": "You can import an existing server by uploading its configuration files. Select the files of your server configuration and press import.",
        "import_button": "Import"
    }
}
</i18n>
