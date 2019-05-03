<template>
    <div class="server">
        <div>
            <div class="name">
                {{server.settings.serverName}}
                <i class="fas fa-cog" v-on:click="edit" v-if="!ro" :title="$t('change_config')"></i>
                <i class="fas fa-terminal" v-on:click="logs" v-if="!ro" :title="$t('view_logs')"></i>
                <i class="fas fa-copy" v-on:click="copyConfig" v-if="is_admin && !ro" :title="$t('copy_config')"></i>
                <i class="fas fa-file-download" v-on:click="exportConfig" v-if="!ro" :title="$t('export_config')"></i>
                <i class="fas fa-trash" v-on:click="deleteServer" v-if="is_admin && !ro" :title="$t('delete_server')"></i>
            </div>
            <div class="info">
                <span v-if="server.pid">PID: {{server.pid}}</span>
                UDP: {{server.basic.udpPort}}
                TCP: {{server.basic.tcpPort}}
                Track: {{server.event.track}}
            </div>
        </div>
        <button class="start" v-on:click="start" v-if="is_mod && !ro && !server.pid">{{$t("start_server")}}</button>
        <button class="stop" v-on:click="start" v-if="is_mod && !ro && server.pid">{{$t("stop_server")}}</button>
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
            });
        },
        exportConfig() {
            let link = document.createElement("a");
            link.href = `/api/server/export/${this.server.id}_${this.server.settings.serverName}.zip?id=${this.server.id}&token=${this.$store.state.auth.token}`;
            link.click();
        },
        deleteServer() {
            axios.delete("/api/server", {params: {id: this.server.id}})
            .then(() => {
                this.$emit("deleted");
            });
        },
        start() {
            // TODO
        },
        stop() {
            // TODO
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
        "delete_server": "Delete server"
    }
}
</i18n>
