package auth

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type User struct {
	Username string
	Password string
	FullName string
}

func CreateUser(req *http.Request) error {
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	var newUser User
	err := json.Unmarshal(bs, &newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println("BODY", newUser.Username)
	return nil
}