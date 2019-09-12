<template>
    <div>
        <i class="fas fa-exclamation-triangle fa-3x"></i>
    </div>
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({
    name: "Error"
});
</script>
<style lang="scss" scoped>
@import "../../scss/variables.scss";
i {
    color: map-get($map: $colors, $key: "error");
    padding: map-get($map: $sizes, $key: "s");
}
</style>