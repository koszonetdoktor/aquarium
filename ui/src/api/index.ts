import axios from "../utils/axios";
import { SensorMeasurement, Event } from "../types";

export async function getWaterTemperature(): Promise<SensorMeasurement<"temperature">> {
    try {
        const response = await axios.get(
            "/measurements/water/temperature",
        );
        return response.data.Measurements;
    } catch (error) {
        throw Error(error);
    }
}

export async function getPhs(): Promise<SensorMeasurement<"ph">> {
    try {
        const response = await axios.get(
            "/measurements/water/ph",
        );
        return response.data.Measurements;
    } catch (error) {
        throw Error(error);
    }
}

export async function getEvents(): Promise<Event[]> {
    try {
        const response = await axios.get("/record/events");
        return response.data;
    } catch (error) {
        throw Error(error);
    }
}

export function saveEvent(event: Event) {
    return new Promise((resolve, reject) => {
        axios.post("/record/events", event)
            .then(() => {
                resolve();
            }).catch((error) => {
                reject();
            });
    });
}
