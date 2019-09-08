package sensors

import (
	"io/ioutil"
	"strings"
	"log"
	"errors"
	"time"
	"fmt"
	"strconv"
)

func SamplingWaterTemp() {
	connectedSensors, err := getSensors() 
	if err != nil {
		log.Println("Could not get the conencted sensors!")
	}
	
	log.Println("Conencted Sensors: ", connectedSensors)

	for {
		temperature, err := readTemperature(connectedSensors[0])
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Temp: ", temperature)
		time.Sleep(1 * time.Second)
	}
}

func getSensors()([]string, error) {
	data, err := ioutil.ReadFile("/sys/bus/w1/devices/w1_bus_master1/w1_master_slaves")
	if err != nil {
		return nil, err
	} 

	sensors := strings.Split(string(data), "\n")

	if len(sensors) > 0 {
		sensors = sensors[:len(sensors)-1] 
	}
	return sensors, nil
}

func readTemperature(sensorName string) (float64, error) {
	data, err := ioutil.ReadFile("/sys/bus/w1/devices/" + sensorName + "/w1_slave")
	if err != nil {
		return 0.0, errors.New("Temperature read error")
	}

	raw := string(data)

	i := strings.LastIndex(raw, "t=")
	if i == -1 {
		return 0.0, errors.New("Temperature read error")
	}

	c, err := strconv.ParseFloat(raw[i+2:len(raw)-1], 64)
	if err != nil {
		return 0.0, errors.New("Temperature read error")
	}
	
	return c/ 1000.0, nil
}