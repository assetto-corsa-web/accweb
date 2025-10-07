<template>
    <layout>
        <div class="title">
            <h1>Live: {{ data.name }}</h1>
            <div class="menu">
                <select v-model="refreshRate">
                    <option v-for="o in [2, 5, 10, 15, 20]" v-bind:value="o">{{ o }} sec</option>
                </select>
                <button v-on:click="loadLive"><i class="fas fa-sync"></i> {{ $t("refresh") }}</button>
                <button class="primary" v-on:click="$router.push('/')"><i class="fas fa-arrow-left"></i>
                    {{ $t("back") }}</button>
            </div>
        </div>

        <div class="content">
            <div class="header">
                <div id="state"><strong>Status:</strong> {{ data.live.serverState }}</div>
                <div id="track"><strong>Track:</strong> {{ data.live.track }}</div>
                <div id="phase">
                    <strong>Phase:</strong> {{ data.live.sessionType }} ({{ data.live.sessionPhase }})
                    <span v-if="data.live.sessionRemaining > 0">[{{ data.live.sessionRemaining }} min]</span>
                </div>
                <div id="nrdrivers"><strong>Drivers:</strong> {{ data.live.nrClients }}</div>
                <div id="updatedat"><strong>Last Update:</strong> {{ new Date(data.live.updatedAt).toLocaleString() }}
                </div>
            </div>

            <div class="body">
                <table id="leaderboard" style="width: 100%;">
                    <tr class="tbl-header">
                        <th>Pos</th>
                        <th>Driver</th>
                        <th>Nr</th>
                        <th>Model</th>
                        <th>Laps</th>
                        <th>Fuel</th>
                        <th class="col-time">Best Lap</th>
                        <th class="col-time">Last Lap</th>
                        <th class="col-time">S1</th>
                        <th class="col-time">S2</th>
                        <th class="col-time">S3</th>
                        <th class="col-time" v-if="data.live.sessionType == 'Race'">Gap</th>
                        <th>Flags</th>
                    </tr>

                    <tr v-for="(car, carId) in orderedCars" :key="carId" v-on:click="setShowLaps(car.carID)"
                        v-bind:class="{ 'tbl-row': true, active: car.carID === showLaps }">
                        <td>{{ carId + 1 }}</td>
                        <td>{{ car.currentDriver ? car.currentDriver.name.toUpperCase() : car.carID }}</td>
                        <td>{{ car.raceNumber }}</td>
                        <td>{{ car.carModel }}</td>
                        <td>{{ car.nrLaps }}</td>
                        <td>{{ car.fuel }}</td>
                        <td>{{ msToTime(car.bestLapMS) }}</td>
                        <td v-bind:class="{ invalid: car.currLap.flags > 0 }">{{ msToTime(car.lastLapMS) }}</td>
                        <td v-bind:class="{ invalid: car.currLap.flags > 0 }">{{ car.currLap.s1 }}</td>
                        <td v-bind:class="{ invalid: car.currLap.flags > 0 }">{{ car.currLap.s2 }}</td>
                        <td v-bind:class="{ invalid: car.currLap.flags > 0 }">{{ car.currLap.s3 }}</td>
                        <td v-if="data.live.sessionType == 'Race'">{{ calcGap(carId) }}</td>
                        <td v-bind:class="{ invalid: car.currLap.flags > 0 }">
                            <i class="fas fa-cut" v-if="car.currLap.hasCut" title="Has Cut"></i>
                            <i class="fas fa-sign-in-alt" v-if="car.currLap.inLap" title="In Lap"></i>
                            <i class="fas fa-sign-out-alt" v-if="car.currLap.outLap" title="Out Lap"></i>
                            <i class="fas fa-flag-checkered" v-if="car.currLap.sessionOver" title="Session is Over"></i>
                        </td>
                    </tr>
                </table>

                <div id="laps" v-if="showLaps">
                    <h3>Car {{ showLapsCar.raceNumber }} Laps</h3>

                    <div class="server-settings-container one-three-columns">
                        <div>
                            <div  class="box" v-for="(d, i) in showLapsCar.drivers" :key="i" style="margin-bottom: 20px;">
                                <h4 style="margin: 0 0 5px 0">{{ d.name }}</h4>
                                ( {{ d.playerID }} )
                            </div>
                        </div>

                        <div>
                            <table style="width: 100%;">
                                <tr class="tbl-header">
                                    <th>Lap</th>
                                    <th>Driver</th>
                                    <th>Pos</th>
                                    <th>Fuel</th>
                                    <th class="col-time">Lap Time</th>
                                    <th class="col-time">S1</th>
                                    <th class="col-time">S2</th>
                                    <th class="col-time">S3</th>
                                    <th class="col-time">Delta</th>
                                    <th>Flags</th>
                                </tr>

                                <tr v-for="(lap, i) in showLapsCar.laps" :key="i"
                                    v-bind:class="{ 'tbl-row': true, invalid: lap.flags > 0, best: lap.lapTimeMS === showLapsCar.bestLapMS }">
                                    <td>{{ i + 1 }}</td>
                                    <td>{{ showLapsCar.drivers[lap.driverIndex] ? showLapsCar.drivers[lap.driverIndex].name :
                                        '--'
                                        }}
                                    </td>
                                    <td>{{ lap.position }}</td>
                                    <td>{{ lap.fuel }}</td>
                                    <td>{{ msToTime(lap.lapTimeMS) }}</td>
                                    <td>{{ lap.s1 }}</td>
                                    <td>{{ lap.s2 }}</td>
                                    <td>{{ lap.s3 }}</td>
                                    <td align="right">{{ calcDelta(i) }}</td>
                                    <td>
                                        <i class="fas fa-cut" v-if="lap.hasCut" title="Has Cut"></i>
                                        <i class="fas fa-sign-in-alt" v-if="lap.inLap" title="In Lap"></i>
                                        <i class="fas fa-sign-out-alt" v-if="lap.outLap" title="Out Lap"></i>
                                        <i class="fas fa-flag-checkered" v-if="lap.sessionOver" title="Session is Over"></i>
                                    </td>
                                </tr>
                            </table>
                        </div>
                    </div>
                </div>

                <div id="chat">
                    <h3>Events</h3>

                    <div class=" box server-settings-container four-columns">
                        <div>
                            <checkbox label="Hide damage report" v-model="hideDamage" />
                        </div>
                        <div>
                            <checkbox label="Hide session changes" v-model="hideSession" />
                        </div>
                        <div>
                            <checkbox label="Hide connection changes" v-model="hideConnections" />
                        </div>
                        <div>
                            <selection label="Show last events" :options="eventsOptions" :inline="true"
                                v-model="nrEvents" />
                        </div>
                    </div>

                    <div v-for="item in eventHistories" :key="item.id">
                        <div class="message" v-if="item.type == 'chat'">
                            <div class="ts">{{ timeSince(new Date(item.ts)) }}</div>
                            <div class="name">{{ item.data.name }}:</div>
                            <div class="msg">{{ item.data.message }}</div>
                        </div>

                        <div class="message session" v-if="item.type == 'session'">
                            <div class="ts">{{ timeSince(new Date(item.ts)) }}</div>
                            <div class="name">Session:</div>
                            <div class="msg">{{ item.data.sessionType }} - {{ item.data.sessionPhase }} ({{
                                item.data.sessionRemaining }} min)</div>
                        </div>

                        <div class="message damage" v-if="item.type == 'damage'">
                            <div class="ts">{{ timeSince(new Date(item.ts)) }}</div>
                            <div class="name">{{ item.data.name.toUpperCase() }}:</div>
                            <div class="msg">Damage Report #{{ item.data.raceNumber }}</div>
                        </div>

                        <div class="message new_connection" v-if="item.type == 'new_connection'">
                            <div class="ts">{{ timeSince(new Date(item.ts)) }}</div>
                            <div class="name">Player joined:</div>
                            <div class="msg">{{ item.data.name.toUpperCase() }}</div>
                        </div>

                        <div class="message remove_connection" v-if="item.type == 'remove_connection'">
                            <div class="ts">{{ timeSince(new Date(item.ts)) }}</div>
                            <div class="name">Player left:</div>
                            <div class="msg">{{ item.data.name.toUpperCase() }}</div>
                        </div>
                    </div>

                </div>
            </div>
        </div>
    </layout>
</template>

<script>
import axios from "axios";
import { layout, checkbox, selection } from "../components";
import _ from "lodash";
import moment from "moment";

let toId = null;

export default {
    name: "live",
    components: { layout, checkbox, selection },
    data() {
        return {
            id: "",
            showLaps: null,
            showLapsCar: null,
            refreshRate: 10,
            hideDamage: false,
            hideSession: false,
            hideConnections: false,
            nrEvents: 100,
            eventsOptions: [
                { label: "30", value: 30 },
                { label: "50", value: 50 },
                { label: "100", value: 100 },
                { label: "150", value: 150 },
                { label: "200", value: 200 }
            ],
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
                    history: []
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
        classObject: function (lap) {
            return {
                "tbl-row": true,
                "lap-invalid": lap.flags > 0
            }
        },
        eventHistories: function () {
            const hd = this.hideDamage;
            const hs = this.hideSession;
            const hc = this.hideConnections;
            const data = _.filter(this.data.live.history, function (o) {
                if (hd && o.type == "damage") {
                    return false
                }

                if (hs && o.type == "session") {
                    return false
                }

                if (hc && (o.type == "new_connection" || o.type == "remove_connection")) {
                    return false
                }

                return true;
            })

            return _.slice(_.reverse(data), 0, this.nrEvents);
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
                        this.setShowLaps(this.showLaps, true);
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
            }, this.refreshRate * 1000);
        },
        setShowLaps(carID, refresh) {
            if (refresh === undefined && this.showLaps === carID) {
                this.showLaps = null;
                this.showLapsCar = null;
                return;
            }

            this.showLaps = carID;
            this.showLapsCar = this.data.live.cars[carID];
        },
        lastLap(laps) {
            if (laps === undefined || laps.length === 0) {
                return {};
            }

            return laps[laps.length - 1]
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

            return `${sign}${Math.floor(m % 60)}:${_.padStart(Math.floor(s % 60).toString(), 2, '0')}.${_.padStart(Math.floor(ms % 1000).toString(), 3, '0')}`;
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
            const prev = this.orderedCars[idx - 1];

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
        },
        timeSince(date) {
            return moment(date).format("LTS");
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

#laps {
    margin-top: 40px;
}

.col-time {
    width: 60px;
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

td,
th {
    padding: 5px;
}

tr:nth-child(odd) {
    background-color: #1f2936;
}

#chat {
    margin-top: 40px;
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

.message.damage {
    color: #888888;
}

.message.session {
    color: #803c3c;
}

.message.new_connection {
    color: #3c803c;
}

.message.remove_connection {
    color: #80803c;
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
