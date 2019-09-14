<template>
    <div>
        <Button @click="onRecord">Record</Button>
        <Button @click="onGet">Get</Button>
        <LineChart v-if="showChart" :dataSeries="dataS" />
    </div>
</template>

<script lang="ts">
import Vue from "vue";
import Button from "@/components/Button.vue";
import axios from "../../utils/axios";
import moment from "moment-timezone";
import LineChart from "@/components/charts/LineChart.vue";
import { SensorMeasurement } from "./types";
import { getWaterTemperature } from "./api";

export default Vue.extend({
    name: "home",
    components: {
        Button,
        LineChart
    },
    data: function() {
        return {
            dataS: [
                {
                    name: "Temp",
                    data: [] as { x: string; y: number }[]
                }
            ]
        };
    },
    computed: {
        showChart() {
            return this.dataS[0].data.length > 0;
        }
    },
    methods: {
        onRecord() {
            this.$router.push("record");
        },
        async onGet() {
            try {
                const tempMeasurePoints = await getWaterTemperature();

                const ts = moment.utc(1568385332230).tz("Europe/Copenhagen");

                const data: { x: string; y: number }[] = [];
                tempMeasurePoints.forEach((point, index) => {
                    data.push({
                        x: moment(ts).format("MM-DD HH:MM"),
                        y: point[2]
                    });
                });

                this.dataS.splice(0, 1, {
                    name: "Celcius",
                    data
                });
            } catch (error) {
                console.error("Coudl not get it", error);
            }
        }
    }
});

function convertDate(unix: number): string {
    const ts = moment.utc(unix).tz("Europe/Copenhagen");
    return moment(ts).format("YYYY-MM-DD HH:MM");
}
</script>