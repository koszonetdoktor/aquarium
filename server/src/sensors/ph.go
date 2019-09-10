package sensors

import (
	"log"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
	"fmt"
	"time"
	"github.com/MichaelS11/go-ads"
)

func TestRead() error {
	log.Println("Test PH read")
	
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	b, err := i2creg.Open("")
	if err != nil {
		
		log.Fatal(err)
	}
	defer b.Close()
	
	// Dev is a valid conn.Conn.
	d := &i2c.Dev{Addr: 48, Bus: b}
	log.Println("BUS", b)
	log.Println("d", d)

	// Send a command 0x10 and expect a 5 bytes reply.
	// write := []byte{0x30}
	read := make([]byte, 5)
	if err := d.Tx(nil, read); err != nil {
		log.Println("ERror her")
		log.Fatal(err)
	}
	fmt.Printf("%v\n", read)

	return nil
}

func TestRead2() {
		// call HostInit once
		err := ads.HostInit()
		if err != nil {
			fmt.Println(err)
		}
	
		// create new ads with wanted busName and address. 
		var ads1 *ads.ADS
		ads1, err = ads.NewADS("I2C1", 0x48, "")
		if err != nil {
			fmt.Println(err)
		}
	
		// example changing config gain (2/3 is default, so only an example)
		ads1.SetConfigGain(ads.ConfigGain1)
	
for{		// read retry from ads chip
		var result uint16
		result, err = ads1.ReadRetry(5)
		if err != nil {
			ads1.Close()
			fmt.Println(err)
		}
	
		// close ads bus
		// elif mygain == ADS1115_REG_CONFIG_PGA_4_096V:
		// coefficient = 0.125
		//NOTE Python magic
		// # Convert the data
		// raw_adc = data[0] * 256 + data[1]
		
		// if raw_adc > 32767:
		// 	raw_adc -= 65535
		// raw_adc = int(float(raw_adc)*coefficient)
		
		// print results
		fmt.Println("raw:", result)
		voltsPyth := float64(result)
		if result > 32767 {
			voltsPyth = voltsPyth - 65535
		}
		voltsPyth = voltsPyth * 0.125
		fmt.Println("RESULT: ", voltsPyth, "mV") 
	// 	volts := (float64(result) / 32767.0 * 5.0)
	// 	fmt.Println("volts:", volts)
	// 	psi := 360.0*volts - 25.0
	// 	fmt.Println("psi:", psi)
		time.Sleep(2 * time.Second)
	}
		err = ads1.Close()
		if err != nil {
			fmt.Println(err)
		}
}