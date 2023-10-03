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
                <div id="state"><strong>Status:</strong> {{data.live.serverState}}</div>
                <div id="track"><strong>Track:</strong> {{data.live.track}}</div>
                <div id="phase">
                    <strong>Phase:</strong> {{data.live.sessionType}} ({{data.live.sessionPhase}})
                    <span v-if="data.live.sessionRemaining > 0">[{{data.live.sessionRemaining}} min]</span>
                </div>
                <div id="nrdrivers"><strong>Drivers:</strong> {{data.live.nrClients}}</div>
                <div id="updatedat"><strong>Last Update:</strong> {{new Date(data.live.updatedAt).toLocaleString()}}</div>
            </div>

            <div class="body">
                <table id="leaderboard">
                    <tr class="tbl-header">
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
                        <th v-if="data.live.sessionType == 'Race'">Gap</th>
                        <th>Flags</th>
                    </tr>

                    <tr v-for="(car, carId) in orderedCars" :key="carId"
                        v-on:click="setShowLaps(car.carID)"
                        v-bind:class="{'tbl-row': true, active: car.carID === showLaps}"
                    >
                        <td>{{carId+1}}</td>
                        <td>{{car.currentDriver ? car.currentDriver.name : car.carID}}</td>
                        <td>{{car.raceNumber}}</td>
                        <td>{{car.carModel}}</td>
                        <td>{{car.nrLaps}}</td>
                        <td>{{car.fuel}}</td>
                        <td>{{msToTime(car.bestLapMS)}}</td>
                        <td v-bind:class="{invalid: car.currLap.flags > 0}">{{msToTime(car.lastLapMS)}}</td>
                        <td v-bind:class="{invalid: car.currLap.flags > 0}">{{car.currLap.s1}}</td>
                        <td v-bind:class="{invalid: car.currLap.flags > 0}">{{car.currLap.s2}}</td>
                        <td v-bind:class="{invalid: car.currLap.flags > 0}">{{car.currLap.s3}}</td>
                        <td v-if="data.live.sessionType == 'Race'">{{calcGap(carId)}}</td>
                        <td v-bind:class="{invalid: car.currLap.flags > 0}">
                            <i class="fas fa-cut" v-if="car.currLap.hasCut" title="Has Cut"></i>
                            <i class="fas fa-sign-in-alt" v-if="car.currLap.inLap" title="In Lap"></i>
                            <i class="fas fa-sign-out-alt" v-if="car.currLap.outLap" title="Out Lap"></i>
                            <i class="fas fa-flag-checkered" v-if="car.currLap.sessionOver" title="Session is Over"></i>
                        </td>
                    </tr>
                </table>

                <div id="laps" v-if="showLaps">
                    <h3>Car {{showLapsCar.raceNumber}} Laps</h3>

                    <table>
                        <tr class="tbl-header">
                            <th>Nr</th>
                            <th>Driver</th>
                            <th>Fuel</th>
                            <th>Lap Time</th>
                            <th>S1</th>
                            <th>S2</th>
                            <th>S3</th>
                            <th>Delta</th>
                            <th>Flags</th>
                        </tr>

                        <tr v-for="(lap, i) in showLapsCar.laps" :key="i" v-bind:class="{ 'tbl-row': true, invalid: lap.flags > 0, best: lap.lapTimeMS === showLapsCar.bestLapMS }">
                            <td>{{i+1}}</td>
                            <td>{{showLapsCar.drivers[lap.driverIndex] ? showLapsCar.drivers[lap.driverIndex].name : '--'}}</td>
                            <td>{{lap.fuel}}</td>
                            <td>{{msToTime(lap.lapTimeMS)}}</td>
                            <td>{{lap.s1}}</td>
                            <td>{{lap.s2}}</td>
                            <td>{{lap.s3}}</td>
                            <td align="right">{{calcDelta(i)}}</td>
                            <td>
                                <i class="fas fa-cut" v-if="lap.hasCut" title="Has Cut"></i>
                                <i class="fas fa-sign-in-alt" v-if="lap.inLap" title="In Lap"></i>
                                <i class="fas fa-sign-out-alt" v-if="lap.outLap" title="Out Lap"></i>
                                <i class="fas fa-flag-checkered" v-if="lap.sessionOver" title="Session is Over"></i>
                            </td>
                        </tr>
                    </table>
                </div>

                <div id="chat">
                    <h3>Chat</h3>

                    <div class="message" v-for="item in data.live.chats.slice().reverse()" :key="item.ts">
                        <div class="ts">{{new Date(item.ts).toLocaleString()}}</div>
                        <div class="name">{{item.name}}:</div> 
                        <div class="msg">{{item.message}}</div>
                    </div>
                </div>
            </div>
        </div>
    </layout>
</template>

<script>
import axios from "axios";
import {layout} from "../components";
import _ from "lodash";

let toId = null;

export default {
    name: "live",
    components: {layout},
    data() {
        return {
            id: "0",
            showLaps: null,
            showLapsCar: null,
            data: {
                name: "",
                track: "",
                live: {
                    serverState: "",
                    nrClients: 0,
                    sessionType: "",
                    sessionPhase: "",
                    sessionRemaining: 0,
                    cars: {},
                    chats: []
                }
            },
        };
    },
    mounted() {
        this.id = this.$route.query.id;
        this.refreshList();
    },
    beforeDestroy() {
        if (toId !== null) {
            clearTimeout(toId);
            toId = null;
        }
    },
    computed: {
        orderedCars: function () {
            const ordered = _.orderBy(this.data.live.cars, "position");
            return _.filter(ordered, (o) => { return o.currentDriver !== null });
        },
        classObject: function(lap) {
            return {
                "tbl-row": true,
                "lap-invalid": lap.flags > 0
            }
        }
    },
    methods: {
        loadLive() {
            axios.get(`/api/instance/${this.id}/live`)
                .then(r => {
                    this.data = r.data;

                    if (this.showLaps !== null && this.data.live.cars[this.showLaps] === undefined) {
                        this.showLaps = null;
                    }

                    if (this.showLaps !== null) {
                        this.setShowLaps(this.showLaps);
                    }
                })
                .catch(e => {
                    this.$store.commit("toast", this.$t("load_live_error"))
                });
        },
        refreshList() {
            this.loadLive();
            toId = setTimeout(() => {
                this.refreshList();
            }, 10000);
        },
        setShowLaps(carID) {
            this.showLaps = carID;
            this.showLapsCar = this.data.live.cars[carID]
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

            let sign = "";

            if (ms < 0) {
                sign = "-";
                ms = Math.abs(ms);
            }

            const s = ms / 1000;
            const m = s / 60;

            return `${sign}${Math.floor(m % 60)}:${_.padStart(Math.floor(s%60).toString(), 2, '0')}.${_.padStart(Math.floor(ms%1000).toString(), 3, '0')}`;
        },
        calcDelta(idx) {
            const bestLap = this.showLapsCar.bestLapMS;
            const curLap = this.showLapsCar.laps[idx].lapTimeMS;

            return this.msToTime(curLap - bestLap);
        },
        calcGap(idx) {
            if (idx === 0) {
                return "";
            }

            const curr = this.orderedCars[idx];
            const prev = this.orderedCars[idx-1];

            if (curr.lastLapTimestampMS === 0) {
                return "";
            }

            if (prev.nrLaps !== curr.nrLaps) {
                const diff = this.orderedCars[0].nrLaps - curr.nrLaps;
                return `+${diff} lap${(diff > 1) ? "s" : ""}`;
            }

            const gap = curr.lastLapTimestampMS - prev.lastLapTimestampMS;

            if (gap < 0) {
                return "--";
            }

            return this.msToTime(gap);
        }
    }
}
</script>	

<style scoped>
.header {
    margin-bottom: 30px;
}

.header div {
    display: inline;
    margin-right: 20px;
}

#leaderboard .tbl-row:hover {
    background-color: #304363;
    cursor: pointer;
}

.active {
    background-color: #27344c !important;
}

.invalid {
    background-color: #803c3c !important;
}

.best {
    background-color: #274c29 !important;
}

th {
    background-color: #1b2838;
}

td, th {
    padding: 5px;
}

tr:nth-child(odd) {
  background-color: #1f2936;
}

#chat .message div {
    display: inline;
    margin-right: 10px;
} 

#chat .message {
    margin-bottom: 5px;
} 

.message .ts {
    color: #304363;
}

.message .name {
    font-weight: bold;
}
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
