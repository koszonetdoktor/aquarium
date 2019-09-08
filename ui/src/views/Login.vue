<template>
    <div class="container">
        <Input
            @input="onInputChange"
            :hasError="loginError"
            title="Username"
            type="text"
            v-model="username"
        />
        <Input
            @input="onInputChange"
            :hasError="loginError"
            title="Password"
            type="password"
            v-model="password"
        />
        <Button @click="onLogin">Login</Button>
    </div>
</template>

<script lang="ts">
import Vue from "vue";
import Input from "@/components/Input.vue";
import Button from "@/components/Button.vue";
import axios from "../utils/axios";

export default Vue.extend({
    name: "login",
    components: {
        Input,
        Button
    },
    data: function() {
        return {
            username: "",
            password: "",
            loginError: false
        };
    },
    methods: {
        onLogin: function() {
            axios
                .post("/authenticate", {
                    username: this.username,
                    password: this.password
                })
                .then(() => {
                    this.$router.push("home");
                })
                .catch(() => {
                    console.error("Could not log in!");
                    this.loginError = true;
                });
        },
        onInputChange: function() {
            this.loginError = false;
        }
    }
});
</script>
<style scoped>
.container {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    vertical-align: center;
}
</style>