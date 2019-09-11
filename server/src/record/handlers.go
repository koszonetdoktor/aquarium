package record

import (
	"fmt"
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
