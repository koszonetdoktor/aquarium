<template>
    <div>
        <Input
            v-for="(record) in recordables"
            @input="onRecordInput"
            v-model="record.value"
            :key="record.name"
            :title="record.name"
            type="number"
        />
        <Input
            v-model="recordTs"
            :hasError="hasDateError"
            type="text"
            title="Time"
            :value="recordTs"
        />
        <Button @click="onSave">Save</Button>
    </div>
</template>

<script lang="ts">
import Vue from "vue";
import Input from "@/components/Input.vue";
import Button from "@/components/Button.vue";
import moment, { min } from "moment";
import axios from "../utils/axios";

export default Vue.extend({
    name: "record",
    components: { Input },
    data: function() {
        return {
            recordables: [
                {
                    name: "ammonia",
                    value: null
                },
                {
                    name: "nitrate",
                    value: null
                },
                {
                    name: "nitrite",
                    value: null
                }
            ],
            recordTs: "",
            hasDateError: false
        };
    },
    mounted() {
        this.recordTs = moment().format("YYYY-MM-DD HH:MM");
    },
    methods: {
        onRecordInput() {
            console.log(this.recordables[0].value, this.recordables[1].value);
        },
        onSave() {
            if (validateDateInputFormat(this.recordTs)) {
                this.hasDateError = false;

                axios.post("/record/measurements", {
                    measurements: constructBody(this.recordables, this.recordTs)
                });
            } else {
                this.hasDateError = true;
            }
        }
    }
});

type Body = {
    name: string;
    value: number;
    date: string;
}[];

function constructBody(
    records: { name: string; value: number | null}[],
    date: string
): Body {
    const body = [];
    records.forEach(record => {
        if (record.value) {
            body.push({
                name: record.name,
                value: record.value,
                date
            });
        }
    });
    return body;
}

function validateDateInputFormat(fullDate: string): boolean {
    const [date, time] = fullDate.split(" ");
    if (date && time) {
        const [year, month, day] = date.split("-");
        const [hour, minute] = time.split(":");

        if (year && month && day && hour && minute) {
            if (
                year.length === 4 &&
                month.length === 2 &&
                day.length === 2 &&
                hour.length === 2 &&
                minute.length === 2
            ) {
                return [year, month, day, hour, minute].every(
                    part => !isNaN(Number(part))
                );
            }
        }
    }
    return false;
}
</script>
