package config

import (
	"log"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

var InfluxDB client.Client

func init() {
	var err error
	InfluxDB, err = client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		log.Println("ERROR: could not connect to the influx databse")
	}

	_, _, err = InfluxDB.Ping(1 * time.Minute)
	if err != nil {
		// log.Fatal("ERROR: Influx server is not answering!")
		log.Println("ERROR: Influx server is not answering!")
	}

	log.Println("InfluxDB is connected", InfluxDB)
}

func Insert(media string, sensor string, measuredValue interface{}) error {

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "aquarium_db",
		Precision: "m",
	})
	if err != nil {
		return err
	}

	sensorTag := map[string]string{"sensor": sensor}

	valueField := map[string]interface{}{"value": measuredValue}

	pt, err := client.NewPoint(media, sensorTag, valueField, time.Now())
	if err != nil {
		return err
	}
	bp.AddPoint(pt)

	err = InfluxDB.Write(bp)
	if err != nil {
		return err
	}

	return nil
}

type Record struct {
	Name  string
	Value float64
	Date  string
}

func MassInsertRecord(media string, records []Record) error {

	log.Println("Mass inserting to media ", media)

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "aquarium_db",
		Precision: "m",
	})
	if err != nil {
		return err
	}

	dateForm := "2006-01-02 15:04"
	//NOTE console.log(moment.utc(1566688980000).format("YYYY-MM-DD HH:MM"))

	for _, record := range records {

		t, err := time.Parse(dateForm, record.Date)
		if err != nil {
			return err
		}

		sensorTag := map[string]string{"sensor": record.Name}

		valueField := map[string]interface{}{"value": record.Value}

		pt, err := client.NewPoint(media, sensorTag, valueField, t)
		if err != nil {
			return err
		}
		bp.AddPoint(pt)
	}

	err = InfluxDB.Write(bp)
	if err != nil {
		return err
	}

	return nil
}

func CloseInfluxClient() {
	InfluxDB.Close()
}
