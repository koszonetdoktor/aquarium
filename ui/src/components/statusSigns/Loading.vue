<template>
    <div>
        <i class="fas fa-spinner fa-3x"></i>
    </div>
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({
    name: "Loading"
});
</script>
<style lang="scss" scoped>
@import "../../scss/variables.scss";
i {
    color: map-get($map: $colors, $key: "blue");
    padding: map-get($map: $sizes, $key: "s");
}
</style>