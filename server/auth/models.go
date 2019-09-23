package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
)

type UserInfo struct {
	Username string
	Password string
	FullName string
}

func CreateUser(req *http.Request) error {
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	var newUser UserInfo

	err := json.Unmarshal(bs, &newUser)
	if err != nil {
		return err
	}

	myHomePath, err := user.Current()
	if err != nil {
		return err
	}

	userJSONFilePath := fmt.Sprintf("%s/Desktop/users.json", myHomePath.HomeDir)

	fmt.Println("Saving user")
	os.Stdout.Write(bs)

	err = ioutil.WriteFile(userJSONFilePath, bs, 0644)
	if err != nil {
		return err
	}
	fmt.Println("User ", newUser.Username, " is created")
	return nil
}

type UserCredentials struct {
	Username string
	Password string
}

func Login(req *http.Request) (bool, error) {

	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)

	var userToLogin UserCredentials
	err := json.Unmarshal(bs, &userToLogin)
	if err != nil {
		return false, err
	}

	myHomePath, err := user.Current()
	if err != nil {
		return false, err
	}

	userJSONFilePath := fmt.Sprintf("%s/Desktop/users.json", myHomePath.HomeDir)

	userFile, err := ioutil.ReadFile(userJSONFilePath)
	if err != nil {
		return false, err
	}

	var readUserJSON UserInfo
	err = json.Unmarshal(userFile, &readUserJSON)
	if err != nil {
		return false, err
	}

	matched := areCredentialsMatched(userToLogin, readUserJSON)

	return matched, nil
}

func areCredentialsMatched(loggingUser UserCredentials, savedUser UserInfo) bool {
	usernameMatch := loggingUser.Username == savedUser.Username
	passMatch := loggingUser.Password == savedUser.Password
	return usernameMatch && passMatch
}
