<template>
    <div class="event-container">
        <Input v-model="name" type="text" title="Name" />
        <DateInput title="Date" type="text" v-model="date" @onValidate="onValidateDate" />
        <textarea title="Notes" v-model="note" />
        <Button @click="onSave">Save</Button>
        <Error v-if="savingState.error" />
        <Loading v-else-if="savingState.saving" />
        <Success v-else-if="savingState.success" />
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import DateInput from "@/components/DateInput.vue";
import Button from "@/components/Button.vue";
import Input from "@/components/Input.vue";
import axios from "../../../utils/axios";
import Error from "@/components/statusSigns/Error.vue";
import Loading from "@/components/statusSigns/Loading.vue";
import Success from "@/components/statusSigns/Success.vue";
import { saveEvent } from "../api";
import moment from "moment";

export default Vue.extend({
    name: "events",
    components: { DateInput, Button, Input, Error, Loading, Success },
    data() {
        return {
            name: "",
            date: "",
            note: "",
            category: 0, //TODO this can be expanded in the futurre, maybe witha drop down
            isDateValid: true,
            savingState: {
                saving: false,
                success: false,
                error: false
            }
        };
    },
    methods: {
        onValidateDate(isValid: boolean) {
            this.isDateValid = isValid;
        },
        async onSave() {
            if (this.isDateValid && this.name) {
                this.savingState.saving = true;
                try {
                    await saveEvent({
                        name: this.name,
                        date: moment(this.date).valueOf(),
                        category: this.category,
                        note: this.note
                    });
                    this.savingState.saving = false;
                    this.savingState.success = true;
                } catch {
                    this.savingState.saving = false;
                    this.savingState.error = true;
                    console.error("Could not save event");
                }
            }
        }
    }
});
</script>
<style lang="scss" scoped>
@import "../../../scss/variables.scss";

.event-container {
    display: flex;
    flex-direction: column;
}
textarea {
    border: solid 1px map-get($map: $colors, $key: 3);
    min-height: 100px;
    color: map-get($map: $colors, $key: 4);
}
</style>