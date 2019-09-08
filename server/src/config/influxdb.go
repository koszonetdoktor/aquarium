package config

import (
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
)

var InfluxDB client.Client

func init(){
	var err error
	InfluxDB, err = client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		log.Println("ERROR: could not connect to the influx databse")
	}

	log.Println("InfluxDB is connected", InfluxDB)
}

func CloseInfluxClient() {
	InfluxDB.Close()
}