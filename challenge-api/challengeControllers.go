package challenge_api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stack-attack/the-roomy-dragonflys/bet-api"
	util "github.com/stack-attack/the-roomy-dragonflys/utils"
	"net/http"
)

func GetChallenges(w http.ResponseWriter, r *http.Request) {
	data := GetAllChallenge()
	resp := util.Message("success")
	resp["data"] = data
	util.Respond(w, resp)
}

func CreateChallenge(w http.ResponseWriter, r *http.Request) {
	var newChallenge Challenge

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
	challenge := GetChallengeById(challengeId)

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
	challenge := bet_api.GetBetsByChallengeId(challengeId)

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

func GetAmountByChallenges(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	challengeId := vars["uuid"]
	challenges := bet_api.GetBetsByChallengeId(challengeId)

	var amount int
	for _, challange := range challenges {
		amount += challange.Amount
	}

	resp := util.Message("success")
	resp["data"] = amount
	util.Respond(w, resp)
}

func UpdateChallenge(w http.ResponseWriter, r *http.Request) {
	var newChallenge Challenge
	err := json.NewDecoder(r.Body).Decode(&newChallenge)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.Respond(w, util.Message("Invalid request"))
		return
	}

	vars := mux.Vars(r)
	challengeId := vars["uuid"]

	resp, status := newChallenge.UpdateChallenge(challengeId)
	w.WriteHeader(status)
	util.Respond(w, resp)
}
