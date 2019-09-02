<template>
    <div class="container">
        <Input title="Username" type="text" v-model="username" />
        <Input title="Password" type="password" v-model="password" />
        <Button @click="onLogin">Login</Button>
    </div>
</template>

<script lang="ts">
import Vue from "vue";
import Input from "@/components/Input.vue";
import Button from "@/components/Button.vue";

export default Vue.extend({
    name: "login",
    components: {
        Input,
        Button
    },
    data: function() {
        return {
            username: "",
            password: ""
        };
    },
    methods: {
        onLogin: function() {
            console.log("Login: ", this.username, "pass: ", this.password);
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