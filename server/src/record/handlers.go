package record

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Measurements handling the saving of the measured record
func Measurements(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}
	if err := saveMeasurementRecords(req); err != nil {
		fmt.Println("ERROR: ", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
	}
	res.WriteHeader(http.StatusOK)
}

func Event(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		postEvent(res, req)
	case http.MethodGet:
		getEvents(res)
	default:
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}
}

func postEvent(res http.ResponseWriter, req *http.Request) {
	log.Println("Saving event in MongoDB")

	err := saveEvent(req)
	if err != nil {
		log.Println("ERROR: could not save event to MongoDB", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
	}
	log.Println("Event is successfully saved in MongoDB")
}

func getEvents(res http.ResponseWriter) {
	log.Println("GEt events")
	snaps, err := getEventSnapshots()
	if err != nil {
		log.Println("ERROR: get snapshots", err)
		return
	}
	results, err := json.Marshal(snaps)

	if err != nil {
		log.Println("ERROR: Marshal event issue", err)
	}

	res.Write(results)
}
