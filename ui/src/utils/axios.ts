import axios from "axios";

let instance;
if (process.env.NODE_ENV === "development") {
    console.log("USE development");
    instance = axios.create({
        baseURL: "http://localhost:8081",
    });
} else {
    instance = axios.create();
}
export default instance;
