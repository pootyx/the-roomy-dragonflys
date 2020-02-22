package controllers

import (
	repository "../repository"
	u "../utils"
	"encoding/json"
	"fmt"
	_ "github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := repository.GetUsers()
	resp :=  u.Message(true, "success")
	resp["data"] = users
	w.WriteHeader(http.StatusPermanentRedirect)
	u.Respond(w, resp)
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
	fmt.Println(w, "Get User by ID!")
	params := mux.Vars(r)
	userId := params["uuid"]

	user := repository.GetUser(userId)
	fmt.Println(user)
	resp := u.Message(true, "success")
	resp["data"] = user
	u.Respond(w, resp)
}

func GetUserChallenges(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get User Challenges!")
}

func GetUserBets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get User Bets!")
}
