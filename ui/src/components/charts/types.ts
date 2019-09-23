export interface Annotation {
    x: any;
    borderColor: string;
    label: {
        borderColor: string
        style: object
        text: string,
    };
}

// options = {
//     annotations: {
//         xaxis: [
//             {
//                 // in a datetime series, the x value should be a timestamp, just like it is generated below
//                 x: new Date("11/17/2017").getTime(),
//                 strokeDashArray: 0,
//                 borderColor: "#775DD0",
//                 label: {
//                     borderColor: "#775DD0",
//                     style: {
//                         color: "#fff",
//                         background: "#775DD0",
//                     },
//                     text: "X Axis Anno Vertical",
//                 },
//             },
//         ],
//         points: [
//             {
//                 x: new Date("27 Nov 2017").getTime(),
//                 y: 8500.9,
//                 marker: {
//                     size: 6,
//                     fillColor: "#fff",
//                     strokeColor: "#2698FF",
//                     radius: 2,
//                 },
//                 label: {
//                     borderColor: "#FF4560",
//                     offsetY: 0,
//                     style: {
//                         color: "#fff",
//                         background: "#FF4560",
//                     },

//                     text: "Point Annotation (XY)",
//                 },
//             },
//         ],
//     },
// };
