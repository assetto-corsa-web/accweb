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
            <div v-bind:class="{tab: true, 'tab-active': activeTab === 0}" v-on:click="activeTab = 0">{{$t("server_config")}}</div>
            <div v-bind:class="{tab: true, 'tab-active': activeTab === 1}" v-on:click="activeTab = 1" v-if="is_admin && !id">{{$t("import_server")}}</div>
        </div>
        <div v-show="activeTab === 0">
            <accweb ref="accweb"></accweb>
            <basic ref="basic"></basic>
            <settings ref="settings"></settings>
            <event ref="event"></event>
            <eventrules ref="eventrules"></eventrules>
            <entrylist ref="entrylist"></entrylist>
			<bop ref="bop"></bop>
			<assistrules ref="assistrules"></assistrules>
        </div>
        <div v-show="activeTab === 1">
            <p>{{$t("upload_hint")}}</p>
            <form v-on:submit.prevent="importServer">
                <label>configuration.json</label>
                <input type="file" name="configuration.json" v-on:change="configurationJsonListener" />
                <label>settings.json</label>
                <input type="file" name="settings.json" v-on:change="settingsJsonListener" />
                <label>event.json</label>
                <input type="file" name="event.json" v-on:change="eventJsonListener" />
                <label>eventRules.json</label>
                <input type="file" name="eventRules.json" v-on:change="eventRulesJsonListener" />
                <label>entrylist.json</label>
                <input type="file" name="entrylist.json" v-on:change="entrylistJsonListener" />
				        <label>bop.json</label>
                <input type="file" name="bop.json" v-on:change="bopJsonListener" />
				        <label>assistRules.json</label>
                <input type="file" name="assistRules.json" v-on:change="assistRulesJsonListener" />				
                <input class="primary" type="submit" :value="$t('import_button')" />
            </form>
        </div>
    </layout>
</template>

<script>
import axios from "axios";
import {layout, end, accweb, basic, settings, event, eventrules, entrylist, bop, assistrules} from "../components";

export default {
    components: {layout, end, accweb, basic, settings, event, eventrules, entrylist, bop, assistrules},
    data() {
        return {
            activeTab: 0,
            id: 0,
            servername: "New Server",
            configurationJson: null,
            settingsJson: null,
            eventJson: null,
            eventRulesJson: null,
            entrylistJson: null,
            bopJson: null,
            assistRulesJson: null
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
            axios.get("/api/instance/"+this.id)
            .then(r => {
                this.servername = r.data.acc.settings.serverName;
                this.$refs.accweb.setData(r.data.accWeb)
                this.$refs.basic.setData(r.data.acc.configuration);
                this.$refs.settings.setData(r.data.acc.settings);
                this.$refs.event.setData(r.data.acc.event);
                this.$refs.eventrules.setData(r.data.acc.eventRules);
                this.$refs.entrylist.setData(r.data.acc.entrylist);
                this.$refs.bop.setData(r.data.acc.bop);
                this.$refs.assistrules.setData(r.data.acc.assistRules);
            });
        },
        save() {
            let accWeb = this.$refs.accweb.getData();
            let configuration = this.$refs.basic.getData();
            let settings = this.$refs.settings.getData();
            let event = this.$refs.event.getData();
            let eventRules = this.$refs.eventrules.getData();
            let entrylist = this.$refs.entrylist.getData();
			      let bop = this.$refs.bop.getData();
			      let assistrules = this.$refs.assistrules.getData();
            let data = {
                accWeb,
                acc: {
                    configuration,
                    settings,
                    event,
                    eventRules,
                    entrylist,
                    bop,
                    assistrules
                }
            };

            let url = "/api/instance"

            if (this.id) {
                url += "/" + this.id
            }

            axios.post(url, data)
            .then(() => {
                this.$router.push("/");
            })
            .catch(e => {
                this.$store.commit("toast", this.$t("save_error"))
            });
        },
        configurationJsonListener(e) {
            this.configurationJson = e.target.files[0];
        },
        settingsJsonListener(e) {
            this.settingsJson = e.target.files[0];
        },
        eventJsonListener(e) {
            this.eventJson = e.target.files[0];
        },
        eventRulesJsonListener(e) {
            this.eventRulesJson = e.target.files[0];
        },
        entrylistJsonListener(e) {
            this.entrylistJson = e.target.files[0];
        },
        bopJsonListener(e) {
            this.bopJson = e.target.files[0];
        },
        assistRulesJsonListener(e) {
            this.assistRulesJson = e.target.files[0];
        },
		    importServer() {
            let data = new FormData();
            data.append("configuration", this.configurationJson);
            data.append("settings", this.settingsJson);
            data.append("event", this.eventJson);
            data.append("eventRules", this.eventRulesJson);
            data.append("entrylist", this.entrylistJson);
			      data.append("bop", this.bopJson);
		      	data.append("assistRules", this.assistRulesJson);

            let headers = {headers: {"Content-Type": "multipart/form-data"}};

            axios.post("/api/import-instance", data, headers)
            .then(() => {
                this.$router.push("/");
            })
            .catch(e => {
                this.$store.commit("toast", this.$t("import_error"))
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
        "server_config": "Configure server",
        "import_server": "Import server",
        "upload_hint": "You can import an existing server by uploading its configuration files. Select the files of your server configuration and press import.",
        "import_button": "Import",
        "save_error": "Error saving configuration, please check your input.",
        "import_error": "Error importing configuration files, please check your input."
    }
}
</i18n>
