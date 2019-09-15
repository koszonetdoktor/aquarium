<template>
    <div>
        <Input
            v-for="(record) in recordables"
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
        <Error v-if="savingState.error" />
        <Loading v-else-if="savingState.saving" />
        <Success v-else-if="savingState.success" />
    </div>
</template>

<script lang="ts">
import Vue from "vue";
import Input from "@/components/Input.vue";
import Button from "@/components/Button.vue";
import moment, { min } from "moment";
import axios from "../../../utils/axios";
import Error from "@/components/statusSigns/Error.vue";
import Loading from "@/components/statusSigns/Loading.vue";
import Success from "@/components/statusSigns/Success.vue";

export default Vue.extend({
    name: "measurements",
    components: { Input, Error, Loading, Success },
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
            hasDateError: false,
            savingState: {
                saving: false,
                success: false,
                error: false
            }
        };
    },
    mounted() {
        this.recordTs = moment().format("YYYY-MM-DD HH:MM");
    },
    methods: {
        onSave() {
            if (validateDateInputFormat(this.recordTs)) {
                this.hasDateError = false;
                this.savingState.saving = true;

                axios
                    .post("/record/measurements", {
                        measurements: constructBody(
                            this.recordables,
                            this.recordTs
                        )
                    })
                    .then(() => {
                        this.savingState.saving = false;
                        this.savingState.success = true;
                    })
                    .catch(() => {
                        this.savingState.saving = false;
                        this.savingState.error = true;
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

function wait() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve();
        }, 2000);
    });
}

function constructBody(
    records: { name: string; value: number | null }[],
    date: string
): Body {
    const body: Body = [];
    records.forEach(record => {
        if (record.value) {
            body.push({
                name: record.name,
                value: Number(record.value),
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
