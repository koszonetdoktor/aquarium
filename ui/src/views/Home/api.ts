import axios from "../../utils/axios";
import { SensorMeasurement } from "./types";

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
