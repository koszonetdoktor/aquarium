package sensors

import (
	"aquarium/config"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/MichaelS11/go-ads"
)

func SamplingPh() {
	err := ads.HostInit()
	if err != nil {
		fmt.Println(err)
	}

	var ads1 *ads.ADS
	ads1, err = ads.NewADS("I2C1", 0x48, "")
	if err != nil {
		fmt.Println(err)
	}

	ads1.SetConfigGain(ads.ConfigGain1)

	for {
		var result uint16
		result, err = ads1.ReadRetry(5)
		if err != nil {
			ads1.Close()
			fmt.Println("ERROR in Ph sampling", err)
		}

		phRaw := caluclatePh(result)

		ph := math.Floor(phRaw*100) / 100

		err = config.Insert("water", "ph", ph)
		if err != nil {
			log.Println("ERROR: Saving the read ph!", err)
		}
		log.Println("Sample water ph: ", ph)
		time.Sleep(1 * time.Minute)
	}
}

func caluclatePh(rawMeasurement uint16) float64 {
	slope := 0.000488
	intersection := 1.667

	ph := slope*float64(rawMeasurement) + intersection
	return ph
}
