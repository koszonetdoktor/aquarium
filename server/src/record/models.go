package record

import (
	"config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type collection struct {
	Measurements []config.Record
}

func saveMeasurementRecords(req *http.Request) error {
	bs := make([]byte, req.ContentLength)
	fmt.Println("BODY ", req.Body)
	req.Body.Read(bs)

	var records collection

	if err := json.Unmarshal(bs, &records); err != nil {
		return err
	}

	if err := config.MassInsertRecord("water", records.Measurements); err != nil {
		log.Println("ERROR: could not save the records in Influx")
	}

	return nil
}
