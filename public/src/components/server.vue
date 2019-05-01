<template>
    <div class="server">
        <div>
            <div class="name">
                Servername
                <i class="fas fa-cog" v-on:click="edit" v-if="!ro"></i>
                <i class="fas fa-terminal" v-on:click="logs" v-if="!ro"></i>
                <i class="fas fa-trash" v-on:click="$emit('delete')" v-if="is_admin && !ro"></i>
            </div>
            <div class="info">PID TCP/UDP Track...</div>
        </div>
        <button class="start" v-on:click="start" v-if="is_mod && !ro">{{$t("start_server")}}</button>
        <div class="online" v-if="ro">Running</div>
        <div class="offline" v-if="ro">Offline</div>
    </div>
</template>

<script>
export default {
    props: ["ro"],
    methods: {
        edit() {
            // TODO
            //this.$router.push(`/server?name=${this.server.name}`);
        },
        logs() {
            // TODO
            //this.$router.push(`/logs?name=${this.server.name}`);
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
