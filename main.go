package main

import (
	"./controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	HandleUserRequests(r)
	HandleChallengeRequests(r)
	HandleBetRequests(r)

	err := http.ListenAndServe(":"+"8080", r)
	if err != nil {
		fmt.Print(err)
	}
}

func HandleUserRequests(r *mux.Router) {
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{uuid}", controllers.GetUser).Methods("POST")
	r.HandleFunc("/users/{uuid}/challenges", controllers.GetUserChallenges).Methods("GET")
	r.HandleFunc("/users/{uuid}/bets", controllers.GetUserBets).Methods("GET")
}

func HandleChallengeRequests(r *mux.Router) {
	r.HandleFunc("/challenges", controllers.GetChallenges).Methods("GET")
	r.HandleFunc("/challenges", controllers.CreateChallenge).Methods("POST")
	r.HandleFunc("/challenges/{uuid}", controllers.GetChallenge).Methods("GET")
	r.HandleFunc("/challenges/{uuid}/bets", controllers.GetChallengeBets).Methods("GET")
}

func HandleBetRequests(r *mux.Router) {

}
