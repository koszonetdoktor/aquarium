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

export default Vue.extend({
    name: "record",
    components: { Input },
    data: function() {
        return {
            recordables: [
                {
                    name: "Ammonia",
                    value: null
                },
                {
                    name: "Nitrate",
                    value: null
                },
                {
                    name: "Nitrite",
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
            console.log("save");
            console.log("record:", this.recordTs);
            if (validateDateInputFormat(this.recordTs)) {
                this.hasDateError = false;
            } else {
                this.hasDateError = true;
            }
        }
    }
});
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
