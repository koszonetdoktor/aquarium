<template>
    <div>
        <Button @click="onRecord">Record</Button>
    </div>
</template>

<script lang="ts">
import Vue from "vue";
import Button from "../components/Button.vue";
export default Vue.extend({
    name: "home",
    components: {
        Button
    },
    methods: {
        onRecord() {
            this.$router.push("record");
        }
    }
});
</script>