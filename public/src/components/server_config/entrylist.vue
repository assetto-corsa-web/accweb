<template>
    <collapsible :title="$t('title')">
        <entry v-for="entry in entries"
            :key="entry.index"
            :entry="entry"
            v-on:remove="removeEntry"></entry>
        <field type="number" :label="$t('forceentrylist_label')" v-model="forceEntryList"></field>
        <button v-on:click="addEntry">{{$t("add_entry_button")}}</button>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import entry from "./entry.vue";
import field from "../field.vue";

export default {
    components: {collapsible, entry, field},
    data() {
        return {
            entryIndex: 0,
            entries: [],
            forceEntryList: 0
        };
    },
    methods: {
        setData(data) {
            let entries = data.entries;

            for(let i = 0; i < entries.length; i++) {
                this.entries.push({
                    index: this.entryIndex,
                    drivers: entries[i].drivers,
                    raceNumber: entries[i].raceNumber,
                    forcedCarModel: entries[i].forcedCarModel,
                    overrideDriverInfo: entries[i].overrideDriverInfo,
                    isServerAdmin: entries[i].isServerAdmin,
                    customCar: entries[i].customCar,
                    overrideCarModelForCustomCar: entries[i].overrideCarModelForCustomCar,
                    ballastKg: entries[i].ballastKg,
                    restrictor: entries[i].restrictor
                });
                this.entryIndex++;
            }

            this.forceEntryList = data.forceEntryList;
        },
        getData() {
            return {
                entries: this.getEntryData(),
                forceEntryList: parseInt(this.forceEntryList)
            };
        },
        getEntryData() {
            let entries = [];

            for(let i = 0; i < this.entries.length; i++) {
                entries.push({
                    drivers: this.entries[i].drivers,
                    raceNumber: parseInt(this.entries[i].raceNumber),
                    forcedCarModel: parseInt(this.entries[i].forcedCarModel),
                    overrideDriverInfo: parseInt(this.entries[i].overrideDriverInfo),
                    isServerAdmin: parseInt(this.entries[i].isServerAdmin),
                    customCar: this.entries[i].customCar,
                    overrideCarModelForCustomCar: parseInt(this.entries[i].overrideCarModelForCustomCar),
                    ballastKg: parseInt(this.entries[i].ballastKg),
                    restrictor: parseInt(this.entries[i].restrictor)
                });
            }

            return entries;
        },
        addEntry() {
            this.entries.push({
                index: this.entryIndex,
                drivers: [],
                raceNumber: 0,
                carModels: [
                  {value: -1, label: "None"},
                  {value: 0, label: "Porsche 991 GT3"},
                  {value: 1, label: "Mercedes AMG GT3"},
                  {value: 2, label: "Ferrari 488 GT3"},
                  {value: 3, label: "Audi R8 LMS"},
                  {value: 4, label: "Lamborghini Huracan GT3"},
                  {value: 5, label: "McLaren 650s GT3"},
                  {value: 6, label: "Nissan GT R Nismo GT3 2018"},
                  {value: 7, label: "BMW M6 GT3"},
                  {value: 8, label: "Bentley Continental GT3 2018"},
                  {value: 9, label: "Porsche 991.2 GT3 Cup"},
                  {value: 10, label: "Nissan GT-R Nismo GT3 2017"},
                  {value: 11, label: "Bentley Continental GT3 2016"},
                  {value: 12, label: "Aston Martin Vantage V12 GT3"},
                  {value: 13, label: "Lamborghini Gallardo R-EX"},
                  {value: 14, label: "Jaguar G3"},
                  {value: 15, label: "Lexus RC F GT3"},
                  {value: 16, label: "Lamborghini Huracan Evo (2019)"},
                  {value: 17, label: "Honda NSX GT3"},
                  {value: 18, label: "Lamborghini Huracan SuperTrofeo"},
                  {value: 19, label: "Audi R8 LMS Evo (2019)"},
                  {value: 20, label: "AMR V8 Vantage (2019)"},
                  {value: 21, label: "Honda NSX Evo (2019)"},
                  {value: 22, label: "McLaren 720S GT3 (2019)"},
                  {value: 23, label: "Porsche 911 II GT3 R (2019)"},
                  {value: 50, label: "Alpine A1110 GT4"},
                  {value: 51, label: "Aston Martin Vantage GT4"},
                  {value: 52, label: "Audi R8 LMS GT4"},
                  {value: 53, label: "BMW M4 GT4"},
                  {value: 55, label: "Chevrolet Camaro GT4"},
                  {value: 56, label: "Ginetta G55 GT4"},
                  {value: 57, label: "KTM X-Bow GT4"},
                  {value: 58, label: "Maserati MC GT4"},
                  {value: 59, label: "McLaren 570S GT4"},
                  {value: 60, label: "Mercedes AMG GT4"},
                  {value: 61, label: "Porsche 718 Cayman GT4"}
                ],
                forcedCarModel: -1,
                overrideDriverInfo: 0,
                isServerAdmin: 0,
                customCar: "",
                overrideCarModelForCustomCar: 0,
                ballastKg: 0,
                restrictor: 0
            });
            this.entryIndex++;
        },
        removeEntry(index) {
            index = parseInt(index);

            for(let i = 0; i < this.entries.length; i++) {
                if(this.entries[i].index === index) {
                    this.entries.splice(i, 1);
                    break;
                }
            }
        }
    }
}
</script>

<i18n>
{
    "en": {
        "title": "Entry List",
        "forceentrylist_label": "Force Entry List",
        "add_entry_button": "Add Entry"
    }
}
</i18n>
