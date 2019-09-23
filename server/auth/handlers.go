package auth

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func Authenticate(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}

	isAuthenticated, err := Login(req)
	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	if isAuthenticated {
		createSessionCookie(res, req)
	} else {
		http.Error(res, http.StatusText(401), http.StatusUnauthorized)
	}
	//NOTE example code to check if session exist and give back the user name if it does
	// c.Value is the cookie value, so the userId
	// if un, ok := dbSessions[c.Value]; ok {
	// 	u = dbUsers[un]
	// }
	// io.WriteString(res, "You got in my friend")
}

func SignUp(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Checking the sign up")
	if req.Method != http.MethodPost {
		http.Error(res, "Not found", http.StatusNotFound)
		return
	}
	fmt.Println("Sign up")
	err := CreateUser(req)
	if err != nil {
		fmt.Println("Create user failed with ", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
	}
	createSessionCookie(res, req)
	// res.WriteHeader(http.StatusOK)
}

func createSessionCookie(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println("Cookie", cookie)
}
