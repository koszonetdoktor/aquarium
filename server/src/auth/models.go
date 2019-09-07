package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
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
		return err
	}

	myHomePath, err := user.Current()
	if err != nil {
		return err
	}

	userJsonFilePath := fmt.Sprintf("%s/Desktop/users.json", myHomePath.HomeDir)

	fmt.Println("Saving user")
	os.Stdout.Write(bs)

	err = ioutil.WriteFile(userJsonFilePath, bs, 0644)
	if err != nil {
		return err
	}
	fmt.Println("User ", newUser.Username, " is saved")
	return nil
}
