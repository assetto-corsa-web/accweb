<template>
<layout>
    <div class="title">
        <h1>{{ $t("title") }}</h1>
        <div class="menu">
            <v-btn small v-on:click="loadServer(true)"><i class="fas fa-sync"></i> {{ $t("refresh") }}</v-btn>
        </div>
    </div>
    <server v-for="s in server" :key="server.id" :server="s" ro="true"></server>
</layout>
</template>

<script>
import axios from "axios";
import {
    layout,
    server,
    end
} from "../components";

export default {
    components: {
        layout,
        server,
        end
    },
    data() {
        return {
            server: []
        };
    },
    mounted() {
        this.refreshList();
    },
    methods: {
        loadServer(refresh) {
            let timeout = 0;

            if (refresh) {
                this.server = [];
                timeout = 100;
            }

            setTimeout(() => {
                axios
                    .get("/api/status")
                    .then(r => {
                        this.server = r.data;
                    })
                    .catch(() => {
                        this.$store.commit("toast", this.$t("receive_server_list_error"));
                    });
            }, timeout);
        },
        refreshList() {
            this.loadServer();
            setTimeout(() => {
                this.refreshList();
            }, 5000);
        }
    }
};
</script>

<i18n>
{

"en": {

"title": "Server Status",

"refresh": "Refresh",

"receive_server_list_error": "Error receiving server list."

}
}
</i18n>
