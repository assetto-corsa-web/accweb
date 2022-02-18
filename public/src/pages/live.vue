<template>
    <layout>
        <div class="title">
            <h1>Live: {{data.name}}</h1>
            <div class="menu">
                <button v-on:click="loadLive"><i class="fas fa-sync"></i> {{$t("refresh")}}</button>
                <button class="primary" v-on:click="$router.push('/')"><i class="fas fa-arrow-left"></i> {{$t("back")}}</button>
            </div>
        </div>

        <div class="content">
            <div class="header">
                <div id="state">{{data.live.serverState}}</div>
                <div id="track">{{data.track}}</div>
                <div id="phase">{{data.live.sessionType}} ({{data.live.sessionPhase}})</div>
                <div id="nrdrivers">{{data.live.nrClients}}</div>
            </div>

            <div class="body">
                <table>
                    <tr>
                        <th>Pos</th>
                        <th>Driver</th>
                        <th>Nr</th>
                        <th>Model</th>
                        <th>Laps</th>
                        <th>Fuel</th>
                        <th>Best Lap</th>
                        <th>Last Lap</th>
                        <th>S1</th>
                        <th>S2</th>
                        <th>S3</th>
                        <th>Flags</th>
                        <th>Cut</th>
                        <th>InLap</th>
                        <th>OutLap</th>
                        <th>SessionOver</th>
                    </tr>

                    <tr v-for="(car, carId) in orderedCars" :key="carId">
                        <td>{{carId+1}}</td>
                        <td>{{car.currentDriver ? car.currentDriver.name : car.carID}}</td>
                        <td>{{car.raceNumber}}</td>
                        <td>{{car.carModel}}</td>
                        <td>{{car.nrLaps}}</td>
                        <td>{{car.fuel}}</td>
                        <td>{{msToTime(car.bestLapMS)}}</td>
                        <td>{{msToTime(car.lastLapMS)}}</td>
                        <td>{{lastLap(car.laps).s1}}</td>
                        <td>{{lastLap(car.laps).s2}}</td>
                        <td>{{lastLap(car.laps).s3}}</td>
                        <td>{{lastLap(car.laps).flags}}</td>
                        <td>{{lastLap(car.laps).hasCut}}</td>
                        <td>{{lastLap(car.laps).inLap}}</td>
                        <td>{{lastLap(car.laps).outLap}}</td>
                        <td>{{lastLap(car.laps).sessionOver}}</td>
                    </tr>
                </table>
            </div>
        </div>
    </layout>
</template>

<script>
import axios from "axios";
import {layout} from "../components";
import _ from "lodash";

export default {
    name: "live",
    components: {layout},
    data() {
        return {
            id: 0,
            data: {
                name: "",
                track: "",
                live: {
                    serverState: "",
                    nrClients: 0,
                    sessionType: "",
                    sessionPhase: "",
                    cars: {},
                }
            },
        };
    },
    mounted() {
        this.id = parseInt(this.$route.query.id);
        this.refreshList();
    },
    computed: {
        orderedCars: function () {
            return _.orderBy(
                _.filter(this.data.live.cars, 'currentDriver')
                , "position")
        },
    },
    methods: {
        loadLive() {
            axios.get(`/api/instance/${this.id}/live`)
                .then(r => {
                    this.data = r.data;
                })
                .catch(e => {
                    this.$store.commit("toast", this.$t("load_live_error"))
                });
        },
        refreshList() {
            this.loadLive();
            setTimeout(() => {
                this.refreshList();
            }, 10000);
        },
        lastLap(laps) {
            if (laps === undefined || laps.length === 0) {
                return {};
            }

            return laps[laps.length-1]
        },
        msToTime(ms) {
            if (ms === 0 || ms === undefined) {
                return "--";
            }

            const m = Math.floor(ms / 60000);
            const s = Math.floor((ms - (m * 60000)) / 1000)
            const c = ms - (m * 60000) - (s * 1000)

            return `${m}:${_.pad(s, 2, '0')}:${_.pad(c, 3, '0')}`;
        }
    }
}
</script>

<style scoped>

</style>

<i18n>
{
    "en": {
        "refresh": "Refresh",
        "back": "Back",
        "load_live_error": "Error loading server live data."
    }
}
</i18n>
