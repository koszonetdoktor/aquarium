package measurements

import (
	"config"
	"fmt"
	"log"

	client "github.com/influxdata/influxdb1-client/v2"
)

type sensorResult [][]interface{}

func readWaterTemp() (sensorResult, error) {
	fmt.Println("Reading the water")

	q := client.NewQuery("SELECT * FROM water WHERE sensor='temperature'", "aquarium_db", "ms")

	res, err := config.InfluxDB.Query(q)
	if err != nil {
		log.Println("ERROR: Could make query to get temperature")
		return nil, err
	}

	return res.Results[0].Series[0].Values, nil
}
