package routes

import (
	"auth"
	"fmt"
	"net/http"
)

//Use :This will setup the routes fro the server
func Use(mux *http.ServeMux) {
	fmt.Println("Use routes!")
	mux.HandleFunc("/authenticate", auth.Authenticate)
	mux.HandleFunc("/signup", auth.SignUp)
	mux.HandleFunc("/testBody", testBody)
}

func testBody(res http.ResponseWriter, req *http.Request) {
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	fmt.Println("BODY: ", string(bs))
}
