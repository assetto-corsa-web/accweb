<template>
	<layout>
		<div class="title">
			<h1>{{$t("title")}}</h1>
			<div class="menu">
				<button class="primary" v-on:click="$router.push('/server')" v-if="is_admin"><i class="fas fa-plus"></i> {{$t("add_new")}}</button>
				<button v-bind:class="stopAllClass" v-on:click="stopAllServers" v-if="is_mod || is_admin"><i class="fas fa-stop"></i> {{$t("stop_all")}}</button>
				<button v-on:click="loadServer(true)"><i class="fas fa-sync"></i> {{$t("refresh")}}</button>
				<button class="logout-btn" v-on:click="logout" v-bind:title="$t('log_out')"><i class="fas fa-sign-out-alt"></i></button>
			</div>
		</div>

        <div class="sort">
            <label>
                {{$t('sorting')}}

                <select v-model="sorting">
                    <option v-for="(o, k) in sortingOptions" v-bind:key="k" v-bind:value="k">{{o.label}}</option>
                </select>
            </label>
        </div>

		<server v-for="s in orderedServers"
			:key="s.id"
			:server="s"
			:ro="!is_mod && !is_admin"
			v-on:copied="loadServer"
			v-on:deleted="loadServer"
			v-on:started="loadServer"
			v-on:stopped="loadServer"></server>
		<p v-if="!servers || !servers.length">No servers found.</p>
	</layout>
</template>

<script>
import axios from "axios";
import {end, layout, server} from "../components";
import _ from "lodash";

let toId = null;

export default {
    components: {layout, server, end},
    data() {
        return {
            stopAllRunning: false,
            servers: [],
            sortingOptions: {
                "name": {label: "Server Name", field: "name", order: "asc"},
                "udp_port": {label: "UDP Port", field: "udpPort", order: "asc"},
                "tcp_port": {label: "TCP Port", field: "tcpPort", order: "asc"},
                "is_running": {label: "Is Running", field: "isRunning", order: "desc"},
                "nr_clients": {label: "Driver Count", field: "nrClients", order: "desc"}
            },
            sorting: "name"
        };
    },
    mounted() {
        if (localStorage.sorting) {
            this.sorting = localStorage.sorting;
        }

        this.refreshList();
    },
    beforeDestroy() {
        if (toId !== null) {
            clearTimeout(toId);
            toId = null;
        }
    },
    watch: {
        sorting(newSorting) {
            localStorage.sorting = newSorting
        }
    },
    computed: {
        orderedServers: function () {
            const o = this.sortingOptions[this.sorting]
            return _.orderBy(this.servers, o.field, o.order)
        },
        stopAllClass: function () {
            return {
                disabled: this.stopAllRunning
            }
        }
    },
    methods: {
        logout() {
            this.$store.commit("logout");
            this.$router.push("/login");
        },
        loadServer(refresh) {
            let timeout = 0;

            if (refresh) {
                this.servers = [];
                timeout = 100;
            }

            setTimeout(() => {
                axios.get("/api/servers")
                    .then(r => {
                        this.servers = r.data;
                    })
                    .catch(e => {
                        this.$store.commit("toast", this.$t("receive_server_list_error"))
                    });
            }, timeout);
        },
        stopAllServers() {
            if (!window.confirm(this.$t("confirm_stop_all"))) {
                return;
            }

            this.stopAllRunning = true;
            axios.post("/api/servers/stop-all")
                .then(d => {
                    this.loadServer(false);
                    this.stopAllRunning = false
                })
                .catch(e => {
                    this.$store.commit("toast", this.$t("stop_all_error"))
                    this.stopAllRunning = false
                });
        },
        refreshList() {
            this.loadServer();
            toId = setTimeout(() => {
                this.refreshList();
            }, 10000);
        }
    }
}
</script>

<style scoped>
.logout-btn .fas {
    margin: 0;
}

.sort select {
    display: inline;
}
</style>

<i18n>
{
    "en": {
        "title": "Servers",
        "add_new": "Add Server",
        "stop_all": "Stop All Servers",
        "refresh": "Refresh",
        "receive_server_list_error": "Error receiving server list.",
        "stop_all_error": "Error while stopping all servers",
        "sorting": "Sorting: ",
        "confirm_stop_all": "Do you really want to stop all acc servers?",
        "log_out": "Log out"
    }
}
</i18n>
