package bet_api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	util "github.com/stack-attack/the-roomy-dragonflys/utils"
	"net/http"
)

func GetBets(w http.ResponseWriter, r *http.Request) {
	chellenges := GetAllBets()
	resp := util.Message("success")
	resp["data"] = chellenges
	util.Respond(w, resp)
}

func CreateBet(w http.ResponseWriter, r *http.Request) {
	var newBet Bet

	err := json.NewDecoder(r.Body).Decode(&newBet)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.Respond(w, util.Message("Invalid request"))
		return
	}

	resp, status := newBet.CreateBet()
	w.WriteHeader(status)
	util.Respond(w, resp)

}

func GetBet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	betId := vars["uuid"]
	challenge := GetBetById(betId)

	if challenge != nil {
		resp := util.Message("success")
		resp["data"] = challenge
		util.Respond(w, resp)
	} else {
		resp := util.Message("Challenge not found with id {" + betId + "}")
		w.WriteHeader(http.StatusNotFound)
		util.Respond(w, resp)
	}
}