<template>
    <Input
        :value="this.recordTs"
        :type="type"
        :title="title"
        @input="onDateChange"
        :hasError="wrongFormat"
    />
</template>
<script lang="ts">
import Vue from "vue";
import Input from "./Input.vue";
import moment from "moment";

export default Vue.extend({
    name: "DateInput",
    components: { Input },
    props: {
        title: String,
        type: String
    },
    data() {
        return {
            wrongFormat: false,
            recordTs: ""
        };
    },
    mounted() {
        this.recordTs = moment().format("YYYY-MM-DD HH:MM");
    },
    methods: {
        onDateChange(date) {
            const validationResult = validateDateInputFormat(date);

            this.$emit("input", date);
            this.$emit("onValidate", validationResult);
            this.wrongFormat = !validationResult;
            this.recordTs = date;
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