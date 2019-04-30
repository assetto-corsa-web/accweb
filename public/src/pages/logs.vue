<template>
    <layout>
        <div class="title">
            <h1>TODO Servername</h1>
            <div class="menu">
                <button class="primary" v-on:click="$router.push('/')"><i class="fas fa-arrow-left"></i> {{$t("back")}}</button>
            </div>
        </div>
        TODO Logs
    </layout>
</template>

<script>
import {layout, end} from "../components";

export default {
    components: {layout, end}
}
</script>

<i18n>
{
    "en": {
        "back": "Back"
    }
}
</i18n>
