<template>
    <collapsible :title="$t('title')">
        <field type="number" :label="$t('udp_label')" v-model="udpPort"></field>
        <field type="number" :label="$t('tcp_label')" v-model="tcpPort"></field>
        <field type="number" :label="$t('maxclients_label')" v-model="maxClients"></field>
        <field type="number" :label="$t('registertolobby_label')" v-model="registerToLobby"></field>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import field from "../field.vue";

export default {
    components: {collapsible, field},
    data() {
        return {
            udpPort: 9600,
            tcpPort: 9600,
            maxClients: 10,
            registerToLobby: 1
        };
    },
    methods: {
        setData(data) {
            this.udpPort = data.udpPort;
            this.tcpPort = data.tcpPort;
            this.maxClients = data.maxClients;
            this.registerToLobby = data.registerToLobby;
        },
        getData() {
            return {
                udpPort: parseInt(this.udpPort),
                tcpPort: parseInt(this.tcpPort),
                maxClients: parseInt(this.maxClients),
                registerToLobby: parseInt(this.registerToLobby)
            };
        }
    }
}
</script>

<i18n>
{
    "en": {
        "title": "Basic configuration",
        "udp_label": "UDP port",
        "tcp_label": "TCP port",
        "maxclients_label": "Max. clients",
        "registertolobby_label": "Register to lobby"
    }
}
</i18n>
