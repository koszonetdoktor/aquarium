import axios from "axios";

let instance;
if (process.env.APP_ENV === "production") {
    instance = axios.create();
} else {
    instance = axios.create({
        baseURL: "http://localhost:8081",
    });
}
export default instance;
