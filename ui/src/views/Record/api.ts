import axios from "../../utils/axios";
import { Event } from "./types";

export function saveEvent(event: Event) {
    return new Promise((resolve, reject) => {
        axios.post("/record/event", event)
            .then(() => {
                resolve();
            }).catch((error) => {
                reject();
            });
    });
}
