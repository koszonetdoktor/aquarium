export type SensorMeasurement<T> = [number, T, number];

export interface Event {
    name: string;
    date: string;
    note: string;
    category: number;
}
