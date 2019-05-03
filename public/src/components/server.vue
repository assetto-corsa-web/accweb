<template>
    <div class="server">
        <div>
            <div class="name">
                {{server.settings.serverName}}
                <i class="fas fa-cog" v-on:click="edit" v-if="!ro"></i>
                <i class="fas fa-terminal" v-on:click="logs" v-if="!ro"></i>
                <i class="fas fa-trash" v-on:click="$emit('delete')" v-if="is_admin && !ro"></i>
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
export default {
    props: ["server", "ro"],
    methods: {
        edit() {
            this.$router.push(`/server?id=${this.server.id}`);
        },
        logs() {
            this.$router.push(`/logs?id=${this.server.id}`);
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
        "stop_server": "Stop"
    }
}
</i18n>
