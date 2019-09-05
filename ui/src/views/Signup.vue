<template>
    <div>
        <Input :hasError="hasError.username" title="Username" type="text" v-model="username" />
        <Input :hasError="hasError.fullName" title="Full name" type="text" v-model="fullName" />
        <Input :hasError="hasError.password" title="Password" type="password" v-model="password" />
        <Input
            :hasError="hasError.confirmPassword"
            title="Password again"
            type="password"
            v-model="confirmPassword"
        />
        <Button @click="onSignUp">Sign up</Button>
    </div>
</template>

<script lang="ts">
import Vue from "vue";
import Input from "../components/Input.vue";

export default Vue.extend({
    name: "signup",
    components: {
        Input
    },
    data: function() {
        return {
            username: "",
            password: "",
            fullName: "",
            confirmPassword: "",
            hasError: {
                username: false,
                password: false,
                confirmPassword: false,
                fullName: false
            }
        };
    },
    methods: {
        onSignUp: function() {
            this.hasError.username = this.username === "";

            this.hasError.password = this.password === "";

            this.hasError.confirmPassword =
                this.confirmPassword !== this.password;

            this.hasError.fullName = this.fullName === "";

            const existsError = Object.keys(this.hasError).find(
                eKey => this.hasError[eKey]
            );

            if (!existsError) {
                //TODO send signup request to the server

                console.log(
                    "Signup",
                    this.username,
                    this.password,
                    this.confirmPassword,
                    this.fullName
                );
            }
        }
    }
});
</script>
<style scoped>
</style>