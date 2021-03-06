package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stack-attack/the-roomy-dragonflys/repository"
	util "github.com/stack-attack/the-roomy-dragonflys/utils"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := repository.GetAllUsers()
	resp := util.Message("success")
	resp["data"] = users
	util.Respond(w, resp)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser repository.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.Respond(w, util.Message("Invalid request"))
		return
	}

	resp, status := newUser.CreateUser()
	w.WriteHeader(status)
	util.Respond(w, resp)
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
	params := mux.Vars(r)
	userId := params["uuid"]
	challenge := repository.GetChallengeByUserId(userId)

	if challenge != nil {
		resp := util.Message("success")
		resp["data"] = challenge
		util.Respond(w, resp)
	} else {
		resp := util.Message("User not found with id {" + userId + "}")
		w.WriteHeader(http.StatusNotFound)
		util.Respond(w, resp)
	}
}

func GetUserBets(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["uuid"]

	bets := repository.GetBetsByUserId(userId)

	if bets != nil {
		resp := util.Message("success")
		resp["data"] = bets
		util.Respond(w, resp)
	} else {
		resp := util.Message("User not found with id {" + userId + "}")
		w.WriteHeader(http.StatusNotFound)
		util.Respond(w, resp)
	}
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["uuid"]

	repository.DeleteUserById(userId)

	resp := util.Message("User successfully deleted!")
	w.WriteHeader(204)
	util.Respond(w, resp)
}
