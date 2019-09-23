<template>
    <div>
        <apexchart
            type="line"
            width="500"
            height="350"
            :options="chartOptions"
            :series="this.dataSeries"
        />
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import VueApexCharts from "vue-apexcharts";
import { Annotation } from "./types";
export default Vue.extend({
    name: "LineChart",
    components: {
        apexchart: VueApexCharts
    },
    props: {
        dataSeries: Array as () => { x: string; y: number }[],
        annotation: Object as () => Annotation
    },
    data: function() {
        return {
            chartOptions: {
                chart: {
                    height: 350,
                    zoom: {
                        enabled: false
                    }
                },
                title: {
                    text: "Temperature",
                    align: "left"
                },
                grid: {
                    row: {
                        colors: ["#f3f3f3", "transparent"], // takes an array which will be repeated on columns
                        opacity: 0.5
                    }
                },
                annotations: {
                    xaxis: [
                        {
                            x: "09-13 16:09",
                            borderColor: "#775DD0",
                            label: {
                                text: "Water change",
                                borderColor: "#775DD0",
                                style: {
                                    color: "#fff",
                                    background: "#775DD0"
                                }
                            }
                        } as Annotation
                    ]
                }
            }
        };
    }
});
</script>