<template>
    <div>
        <DateInput title="Date" type="text" v-model="name" @onValidate="onValidateDate" />
    </div>
</template>
<script lang="ts">
import Vue from "vue";
import DateInput from "@/components/DateInput.vue";
import Button from "@/components/Button.vue";

export default Vue.extend({
    name: "events",
    components: { DateInput, Button },
    data() {
        return {
            name: ""
        };
    },
    methods: {
        onValidateDate(valid) {
            console.log("Valduite", valid);
        }
    }
});
</script>