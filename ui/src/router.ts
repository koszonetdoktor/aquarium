import Vue from "vue";
import Router from "vue-router";
import Login from "./views/Login.vue";
import Signup from "./views/Signup.vue";
import Home from "./views/Home.vue";
import Auth from "./views/Auth.vue";
import Record from "./views/Record.vue";

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: "/",
            name: "auth",
            component: Auth,
            children: [
                {
                    path: "/",
                    component: Login,
                },
                {
                    path: "/signup",
                    component: Signup,
                },
            ],
        },
        {
            path: "/home",
            name: "home",
            component: Home,
        },
        {
            path: "/record",
            name: "record",
            component: Record,
        },
        {
            path: "/about",
            name: "about",
            // route level code-splitting
            // this generates a separate chunk (about.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () => import(/* webpackChunkName: "about" */ "./views/About.vue"),
        },
    ],
});
