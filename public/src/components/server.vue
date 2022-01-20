<template>
    <div class="server">
        <div>
            <div class="name">
                {{server.name}}
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
                Track: {{server.track}}
                <span v-if="!ro">&bull; Config Dir: {{server.id}}</span>
            </div>
        </div>
        <button class="start" v-on:click="start" v-if="is_mod && !ro && !server.pid">{{$t("start_server")}}</button>
        <button class="stop" v-on:click="stop" v-if="is_mod && !ro && server.pid">{{$t("stop_server")}}</button>
        <div class="online" v-if="ro && server.pid">Running</div>
        <div class="offline" v-if="ro && !server.pid">Offline</div>
    </div>
</template>

<script>
import axios from "axios";

export default {
    props: ["server", "ro"],
    methods: {
        edit() {
            this.$router.push(`/server?id=${this.server.id}`);
        },
        logs() {
            this.$router.push(`/logs?id=${this.server.id}`);
        },
        copyConfig() {
            axios.put("/api/server", {id: this.server.id})
            .then(() => {
                this.$emit("copied");
            })
            .catch(e => {
                this.$store.commit("toast", this.$t("copy_server_error"))
            });
        },
        exportConfig() {
            // replace everything that's not a "normal" character or number so we export using a valid filename
            // in case it's empty afterwards, set a default filename
            let filename = this.server.name.replace(/[^a-z0-9]/gi, '_').toLowerCase();
            
            if(!filename.length) {
                filename = "server";
            }
            
            let link = document.createElement("a");
            link.setAttribute("type", "hidden");
            link.href = `/api/server/export/${this.server.id}_${filename}.zip?id=${this.server.id}&token=${this.$store.state.auth.token}`;
            document.body.appendChild(link);
            link.click();
            link.remove();
        },
        deleteServer() {
            axios.delete("/api/server", {params: {id: this.server.id}})
            .then(() => {
                this.$emit("deleted");
            })
            .catch(e => {
                this.$store.commit("toast", this.$t("delete_server_error"))
            });
        },
        start() {
            axios.post("/api/instance", {id: this.server.id})
            .then(() => {
                this.$emit("started");
            })
            .catch(e => {
                this.$store.commit("toast", this.$t("start_server_error"))
            });
        },
        stop() {
            axios.delete("/api/instance", {params: {id: this.server.id}})
            .then(() => {
                this.$emit("stopped");
            })
            .catch(e => {
                this.$store.commit("toast", this.$t("stop_server_error"))
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
        "copy_config": "Copy config",
        "export_config": "Export config",
        "delete_server": "Delete server",
        "copy_server_error": "Error copying server configuration.",
        "delete_server_error": "Error deleting server configuration.",
        "start_server_error": "Error starting server, please check the logs.",
        "stop_server_error": "Error stopping server."
    }
}
</i18n>
