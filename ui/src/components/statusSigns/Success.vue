<template>
    <div>
        <i class="fas fa-smile-wink fa-3x"></i>
    </div>
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({
    name: "Success"
});
</script>
<style lang="scss" scoped>
@import "../../scss/variables.scss";
i {
    color: map-get($map: $colors, $key: "success");
    padding: map-get($map: $sizes, $key: "s");
}
</style>