<template>
    <div class="server">
        <div>
            <div class="name">
                Servername
                <i class="fas fa-cog" v-on:click="edit"></i>
                <i class="fas fa-terminal" v-on:click="logs"></i>
            </div>
            <div class="info">PID TCP/UDP Track...</div>
        </div>
        <button class="start" v-on:click="start" v-if="is_mod">{{$t("start_server")}}</button>
    </div>
</template>

<script>
export default {
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
