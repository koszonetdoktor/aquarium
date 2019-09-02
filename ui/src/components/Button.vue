<template>
    <button class="form-button" @click="$emit('click')">
        <slot></slot>
    </button>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
    name: "Button"
});
</script>
<style scoped>
.form-button {
    background: transparent;
    color: #267dbd;
    border: none;
    text-transform: uppercase;
    font-weight: bold;
}
</style>