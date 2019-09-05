package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"routes"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	var port int

	if appEnv == "production" {
		port = 5000

		fs := http.FileServer(http.Dir("../../ui/dist"))
		http.Handle("/", fs)

		fmt.Println("Server is running in production!")
	} else {
		port = 8081
		http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
			io.WriteString(w, "Dev Mode Test")
		})
		fmt.Println("Server is running in development!")
	}

	fmt.Println("Server is running on ", port)

	routes.Use()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
