package routes

import (
	"fmt"
	"io"
	"net/http"
)

func authenticate(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}

	io.WriteString(res, "You got in my friend")
}

func Use() {
	fmt.Println("Use routes!")
	http.HandleFunc("/authenticate", authenticate)
}
