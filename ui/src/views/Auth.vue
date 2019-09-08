<template>
    <div id="nav">
        <router-link to="/">Login</router-link>|
        <router-link to="/signup">Signup</router-link>
        <router-view />
    </div>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
    name: "auth"
});
</script>

<style>
#nav {
    padding: 30px;
}

#nav a {
    font-weight: bold;
    color: #2c3e50;
}

#nav a.router-link-exact-active {
    color: #42b983;
}
</style>