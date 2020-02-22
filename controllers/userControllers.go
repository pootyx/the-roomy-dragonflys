package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stack-attack/the-roomy-dragonflys/repository"
	util "github.com/stack-attack/the-roomy-dragonflys/utils"
	"io/ioutil"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := repository.GetAllUsers()
	resp :=  util.Message("success")
	resp["data"] = users
	util.Respond(w, resp)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser repository.User

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.Respond(w, util.Message("Invalid request"))
		return
	}

	json.Unmarshal(reqBody, &newUser)
	newUser.CreateUser()

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["uuid"]
	user := repository.GetUserById(userId)

	if user != nil {
		resp := util.Message("success")
		resp["data"] = user
		util.Respond(w, resp)
	} else {
		resp := util.Message("User not found with id {" + userId + "}")
		w.WriteHeader(http.StatusNotFound)
		util.Respond(w, resp)
	}
}

func GetUserChallenges(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get User Challenges!")
}

func GetUserBets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get User Bets!")
}
