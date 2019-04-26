<template>
    <div class="server">
        <div>
            <div class="name">
                Servername
                <i class="fas fa-cog"></i>
                <i class="fas fa-terminal"></i>
            </div>
            <div class="info">PID TCP/UDP Track...</div>
        </div>
        <button class="start">{{$t("start_server")}}</button>
    </div>
</template>

<script>
export default {}
</script>

<i18n>
{
    "en": {
        "start_server": "Start",
        "stop_server": "Stop"
    }
}
</i18n>
