import axios from "axios";

let instance;
if (process.env.APP_ENV === "development") {
    instance = axios.create({
        baseURL: "http://localhost:8081",
    });
} else {
    instance = axios.create();
}
export default instance;
