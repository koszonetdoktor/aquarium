<template>
    <div>
        <label>{{title}}</label>
        <input
            :class="{error: hasError}"
            :type="type"
            @input="$emit('input', $event.target.value)"
            :value="value"
        />
    </div>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
    name: "Input",
    props: {
        title: String,
        type: String,
        hasError: Boolean,
        value: String
    }
});
</script>

<style lang="scss" scoped>
@import "../scss/variables.scss";

div {
    padding: 8px 8px 8px 0;
}

label {
    display: block;
    font-size: 8px;
    text-transform: uppercase;
    color: map-get($map: $colors, $key: 3);
}

input {
    background: transparent;
    border: none;
    border-bottom: solid 1px map-get($map: $colors, $key: 3);
    outline: none;
    color: map-get($map: $colors, $key: 4);
    font-size: 24px;
    font-weight: bold;
}
.error {
    border-bottom-color: rgb(199, 19, 58);
}
</style>