<template>
    <div>
        <h1>Dashboard</h1>
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import axios from "../../utils/axios";
import * as api from "../../api";
import { SensorMeasurement, Event } from "../../types";
export default Vue.extend({
    name: "dashboard",
    data() {
        return {
            waterTemperatures: [] as SensorMeasurement<"temperature">[],
            waterPhs: [] as SensorMeasurement<"ph">[],
            events: [] as Event[]
        };
    },
    mounted: async function() {
        //TODO
        //- get the events
        //- get the measurements
        //- put them tougether ona Graph
        // all the options of the graph should be managed here,
        // - I am not even sure that the Line chart should be used or not
        try {
            this.waterTemperatures = await api.getWaterTemperature();
            this.waterPhs = await api.getPhs();
            this.events = await api.getEvents();
            console.log(
                "Mounted: ",
                this.waterTemperatures,
                this.waterTemperatures,
                this.events
            );
        } catch (error) {
            console.error(error);
            //TODO set some error state
        }
    }
});
</script>