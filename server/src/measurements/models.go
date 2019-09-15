package measurements

import (
	"config"
	"log"

	client "github.com/influxdata/influxdb1-client/v2"
)

type sensorResult [][]interface{}

func readWaterTemp() (sensorResult, error) {
	log.Println("Reading the water from InfluxDb")

	q := client.NewQuery("SELECT * FROM water WHERE sensor='temperature'", "aquarium_db", "ms")

	res, err := config.InfluxDB.Query(q)
	if err != nil {
		log.Println("ERROR: Could make query to get temperature")
		return nil, err
	}

	return res.Results[0].Series[0].Values, nil
}

func readPh() (sensorResult, error) {
	log.Println("Reading the ph from InfluxDb")

	q := client.NewQuery("SELECT * FROM water WHERE sensor='ph'", "aquarium_db", "ms")

	res, err := config.InfluxDB.Query(q)
	if err != nil {
		log.Println("ERROR: Could ot make query to get Ph")
		return nil, err
	}

	return res.Results[0].Series[0].Values, nil
}
