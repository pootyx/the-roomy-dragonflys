package controllers

import (
	"../repository"
	u "../utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func GetChallenges(w http.ResponseWriter, r *http.Request) {
	data := repository.GetAll()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func CreateChallenge(w http.ResponseWriter, r *http.Request) {
	var newChallenge repository.Challenge

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	fmt.Println(reqBody)
	json.Unmarshal(reqBody, &newChallenge)

	newChallenge.Create()
	fmt.Println(newChallenge)
}

func GetChallenge(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["uuid"]
	data := repository.GetById(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func GetChallengeBets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Challenge Bets!")
}
