package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stack-attack/the-roomy-dragonflys/repository"
	util "github.com/stack-attack/the-roomy-dragonflys/utils"
	"net/http"
)

func GetChallenges(w http.ResponseWriter, r *http.Request) {
	data := repository.GetAllChallenge()
	resp := util.Message("success")
	resp["data"] = data
	util.Respond(w, resp)
}

func CreateChallenge(w http.ResponseWriter, r *http.Request) {
	var newChallenge repository.Challenge

	err := json.NewDecoder(r.Body).Decode(&newChallenge)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.Respond(w, util.Message("Invalid request"))
		return
	}

	resp, status := newChallenge.CreateChallenge()
	w.WriteHeader(status)
	util.Respond(w, resp)
}

func GetChallenge(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	challengeId := vars["uuid"]
	challenge := repository.GetChallengeById(challengeId)

	if challenge != nil {
		resp := util.Message("success")
		resp["data"] = challenge
		util.Respond(w, resp)
	} else {
		resp := util.Message("Challenge not found with id {" + challengeId + "}")
		w.WriteHeader(http.StatusNotFound)
		util.Respond(w, resp)
	}
}

func GetChallengeBets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	challengeId := vars["uuid"]
	challenge := repository.GetBetsByChallengeId(challengeId)

	if challenge != nil {
		resp := util.Message("success")
		resp["data"] = challenge
		util.Respond(w, resp)
	} else {
		resp := util.Message("Challenge not found with id {" + challengeId + "}")
		w.WriteHeader(http.StatusNotFound)
		util.Respond(w, resp)
	}
}
