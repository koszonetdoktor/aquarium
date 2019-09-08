package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"routes"

	"github.com/rs/cors"
	"sensors"
	"config"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	var port int

	go sensors.SamplingWaterTemp()
	
	//I am not sure, this call makes any sense
	defer config.CloseInfluxClient()

	mux := http.NewServeMux()
	routes.Use(mux)


	if appEnv == "production" {
		port = 5000

		fs := http.FileServer(http.Dir("../../ui/dist"))
		mux.Handle("/", fs)

		log.Println("Server is running in production on port: ", port)

		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
	} else {
		port = 8081

		log.Println("Server is running in development on port: ", port)

		handler := cors.Default().Handler(mux)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
	}
}
