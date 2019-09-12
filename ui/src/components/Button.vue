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
<style lang="scss" scoped>
@import "../scss/variables.scss";
.form-button {
    background: transparent;
    color: map-get($map: $colors, $key: 3);
    border: none;
    text-transform: uppercase;
    font-weight: bold;
}
</style>