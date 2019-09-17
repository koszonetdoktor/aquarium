package routes

import (
	"auth"
	"fmt"
	"measurements"
	"net/http"
	"os/user"
	"record"
)

//Use :This will setup the routes fro the server
func Use(mux *http.ServeMux) {
	fmt.Println("Use routes!")

	mux.HandleFunc("/authenticate", auth.Authenticate)
	mux.HandleFunc("/signup", auth.SignUp)

	mux.HandleFunc("/record/measurements", record.Measurements)
	mux.HandleFunc("/record/events", record.Event)

	mux.HandleFunc("/measurements/water/temperature", measurements.GetWaterTemp)
	mux.HandleFunc("/measurements/water/ph", measurements.GetPh)

	mux.HandleFunc("/testBody", testBody)
	mux.HandleFunc("/desk", testDesktop)
}

func testBody(res http.ResponseWriter, req *http.Request) {
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	fmt.Println("BODY: ", string(bs))
}

func testDesktop(res http.ResponseWriter, req *http.Request) {
	myself, err := user.Current()

	if err != nil {
		http.Error(res, "Problem with the user", 404)
	}

	fmt.Println("USER DIR: ", myself.HomeDir)
}
