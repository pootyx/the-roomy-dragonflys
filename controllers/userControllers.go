package controllers

import (
	repository "../repository"
	u "../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "github.com/google/uuid"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Users!")
	repository.CreateUserToDb()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST call /users")
	var newUser repository.User

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	fmt.Println(reqBody)
	json.Unmarshal(reqBody, &newUser)

	newUser.Create()
	fmt.Println(newUser)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get User!")
}

func GetUserChallenges(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get User Challenges!")
}

func GetUserBets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get User Bets!")
}
