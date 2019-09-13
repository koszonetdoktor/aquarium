package measurements

import (
	"encoding/json"
	"log"
	"net/http"
)

type sensorDataPoints struct {
	Measurements [][]interface{}
}

func GetWaterTemp(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	temps, err := readWaterTemp()
	if err != nil {
		log.Println("ERROR", err)
	}

	log.Println(temps)

	tempSensorMs := sensorDataPoints{
		Measurements: temps,
	}

	respBytes, err := json.Marshal(tempSensorMs)
	if err != nil {
		log.Println("ERROR: could not marshal temperature sensor data")
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Write(respBytes)
	log.Println("Water temperature has been read successfully!")
}
