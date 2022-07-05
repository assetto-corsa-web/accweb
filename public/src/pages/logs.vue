<template>
<layout>
    <div class="title">
        <h1>{{ servername }}</h1>
        <div class="menu">
            <v-btn small v-on:click="loadLogs"><i class="fas fa-sync"></i> {{ $t("refresh") }}</v-btn>
            <v-btn small class="primary" v-on:click="$router.push('/')"><i class="fas fa-arrow-left"></i> {{ $t("back") }}</v-btn>
        </div>
    </div>
    <label for="output">{{ $t("output_label") }}</label>
    <textarea v-model="logs" style="min-height: 300px;" ref="output" id="output"></textarea>
</layout>
</template>

<script>
import axios from "axios";
import {
    layout,
    end
} from "../components";

export default {
    components: {
        layout,
        end
    },
    data() {
        return {
            id: 0,
            servername: "",
            logs: ""
        };
    },
    /*watch: {
        logs() {
            // vue needs some time to render the textareas content
            setTimeout(() => {
                this.$refs.output.scrollTop = this.$refs.output.scrollHeight;
            }, 200);
        }
    },*/
    mounted() {
        this.id = parseInt(this.$route.query.id);
        this.loadServer();
        this.loadLogs();
    },
    methods: {
        loadServer() {
            axios.get(`/api/instance/${this.id}`).then(r => {
                this.servername = r.data.acc.settings.serverName;
            });
        },
        loadLogs() {
            axios
                .get(`/api/instance/${this.id}/logs`)
                .then(r => {
                    this.logs = r.data.logs;
                })
                .catch(e => {
                    this.$store.commit("toast", this.$t("load_logs_error"));
                });
        }
    }
};
</script>

<i18n>
{
    "en": {
        "refresh": "Refresh",
        "back": "Back",
        "output_label": "Log output",
        "load_logs_error": "Error loading logs."
    }
}
</i18n>
