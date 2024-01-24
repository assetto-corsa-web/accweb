<template>
    <div class="server">
        <div>
            <div class="name">
                {{server.name}}
                <span v-if="is_ro">
                    <i class="fas fa-tv" v-if="server.pid" v-on:click="live" :title="$t('view_live')"></i>
                </span>
                <span v-if="!ro">
                    <i class="fas fa-cog" v-on:click="edit" :title="$t('change_config')"></i>
                    <i class="fas fa-terminal" v-on:click="logs" :title="$t('view_logs')"></i>
                    <i class="fas fa-copy" v-on:click="copyConfig" v-if="is_admin" :title="$t('copy_config')"></i>
                    <i class="fas fa-file-download" v-on:click="exportConfig" :title="$t('export_config')"></i>
                    <i class="fas fa-trash" v-on:click="deleteServer" v-if="is_admin" :title="$t('delete_server')"></i>
                </span>
            </div>
            <div class="info">
                <span v-if="server.pid">PID: {{server.pid}}</span>
                UDP: {{server.udpPort}} &bull;
                TCP: {{server.tcpPort}} &bull;
                {{$t("track")}}: {{server.track}}
                <span v-if="!ro">&bull; {{$t("configuration_directory")}}: {{server.id}}</span>
            </div>
            <div class="info state" v-if="server.pid">
                <b>{{$t("state")}}: </b>{{$t(server.serverState)}} &bull;
                <b>{{$t("number_of_drivers")}}: </b>{{formattedServerClientCount}} &bull;
                <b>{{$t("session")}}: </b>
                <span v-if="server.sessionType">{{server.sessionType}} ({{server.sessionPhase}}) - {{server.sessionRemaining}} min(s)</span>
                <span v-else>{{$t('not_detected')}}</span>
            </div>
        </div>
        <button class="start" v-on:click="start" v-if="is_mod && !ro && !server.pid">{{$t("start_server")}}</button>
        <button class="stop" v-on:click="stop" v-if="is_mod && !ro && server.pid">{{$t("stop_server")}}</button>
        <div class="online" v-if="ro && server.pid">{{$t("running")}}</div>
        <div class="offline" v-if="ro && !server.pid">{{$t("offline")}}</div>
    </div>
</template>

<style>
.state {
    margin-top: 10px;
}

.state b {
    color: #505050;
}
</style>

<script>
import axios from "axios";

export default {
    props: ["server", "ro"],
    computed: {
        formattedServerClientCount: function () {
            return this.server.serverState === 'not_registered' ? '-' : this.server.nrClients;
        }
    },
    methods: {
        edit() {
            this.$router.push(`/server?id=${this.server.id}`);
        },
        logs() {
            this.$router.push(`/logs?id=${this.server.id}`);
        },
        live() {
            this.$router.push(`/live?id=${this.server.id}`);
        },
        copyConfig() {
            axios.post(`/api/instance/${this.server.id}/clone`)
            .then(() => {
                this.$emit("copied");
            })
            .catch(e => {
                this.$store.commit("toast", this.$t("copy_server_error"))
            });
        },
        exportConfig() {
            let link = document.createElement("a");
            link.setAttribute("type", "hidden");
            link.href = `/api/instance/${this.server.id}/export?token=${this.$store.state.auth.token}`;
            document.body.appendChild(link);
            link.click();
            link.remove();
        },
        deleteServer() {
            if (!window.confirm(this.$t("confirm_delete"))) {
                return;
            }
            axios.delete(`/api/instance/${this.server.id}`)
            .then(() => {
                this.$emit("deleted");
            })
            .catch(e => {
                this.$store.commit("toast", this.$t("delete_server_error"))
            });
        },
        start() {
            axios.post(`/api/instance/${this.server.id}/start`)
            .then(() => {
                this.$emit("started");
            })
            .catch(e => {
                this.$store.commit("toast", this.$t("start_server_error", {error: e.response.data.error}))
            });
        },
        stop() {
            if (!window.confirm(this.$t("confirm_stop"))) {
                return;
            }
            axios.post(`/api/instance/${this.server.id}/stop`)
            .then(() => {
                this.$emit("stopped");
            })
            .catch(e => {
                this.$store.commit("toast", this.$t("stop_server_error", {error: e.response.data.error}))
            });
        }
    }
}
</script>

<i18n>
{
    "en": {
        "start_server": "Start",
        "stop_server": "Stop",
        "change_config": "Change config",
        "view_logs": "View logs",
        "view_live": "View live",
        "copy_config": "Copy config",
        "export_config": "Export config",
        "delete_server": "Delete server",
        "copy_server_error": "Error copying server configuration.",
        "delete_server_error": "Error deleting server configuration.",
        "start_server_error": "Error starting server, please check the logs. ERROR: {error}",
        "stop_server_error": "Error stopping server. ERROR: {error}",
        "track": "Track",
        "configuration_directory": "Config dir",
        "running": "Running",

        "state": "State",
        "number_of_drivers": "Drivers",
        "session": "Session",
        "not_detected": "Not detected",

        "offline": "Offline",
        "starting": "Starting",
        "not_registered": "Waiting for events",
        "online": "Online",
        "confirm_stop": "Do you really want to stop this server?",
        "confirm_delete": "Do you really want to delete this server?"
    }
}
</i18n>
