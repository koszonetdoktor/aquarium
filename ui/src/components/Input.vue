<template>
    <div>
        <lable>{{title}}</lable>
        <input :class="{error: hasError}" :type="type" @input="$emit('input', $event.target.value)" />
    </div>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
    name: "Input",
    props: {
        title: String,
        type: String,
        hasError: Boolean
    }
});
</script>

<style scoped>
div {
    padding: 8px 8px 8px 0;
}

label {
    display: block;
    font-size: 8px;
    text-transform: uppercase;
    color: #267dbd;
}

input {
    background: transparent;
    border: none;
    border-bottom: solid 1px #267dbd;
    outline: none;
    color: #93b7cc;
    font-size: 24px;
    font-weight: bold;
}
.error {
    border-bottom-color: rgb(199, 19, 58);
}
</style>