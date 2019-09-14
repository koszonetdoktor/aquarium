package record

import (
	"config"
	"encoding/json"
	"log"
	"net/http"
)

type collection struct {
	Measurements []config.Record
}

func saveMeasurementRecords(req *http.Request) error {
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)

	var records collection

	if err := json.Unmarshal(bs, &records); err != nil {
		return err
	}

	if err := config.MassInsertRecord("water", records.Measurements); err != nil {
		log.Println("ERROR: could not save the records in Influx", err)
		return err
	}
	log.Println("Records are inserted succesfully")

	return nil
}
