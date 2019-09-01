package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "production" {
		fmt.Println("App is in production! Serve static site on port :5000")
		http.ListenAndServe(":5000", http.FileServer(http.Dir("../ui/dist")))
	} else {
		fmt.Println("App is running in development!")
	}
}
