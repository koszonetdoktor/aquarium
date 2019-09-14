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
export default Vue.extend({
    name: "LineChart",
    components: {
        apexchart: VueApexCharts
    },
    props: {
        dataSeries: Array as () => { x: string; y: number }[]
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
                dataLabels: {
                    enabled: false
                },
                stroke: {
                    curve: "straight"
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
                }
            }
        };
    }
});
</script>