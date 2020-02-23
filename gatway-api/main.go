package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stack-attack/the-roomy-dragonflys/auth-api"
	"github.com/stack-attack/the-roomy-dragonflys/bet-api"
	"github.com/stack-attack/the-roomy-dragonflys/challenge-api"
	"github.com/stack-attack/the-roomy-dragonflys/user-api"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.Use(JwtAuthentication)

	HandleAuthentication(r)
	HandleUserRequests(r)
	HandleChallengeRequests(r)
	HandleBetRequests(r)

	err := http.ListenAndServe(":"+getPort(), r)
	if err != nil {
		fmt.Print(err)
	}
}

func getPort() string {
	if value, ok := os.LookupEnv("PORT"); ok {
		return value
	}
	return "8080"
}

func HandleAuthentication(r *mux.Router) {
	r.HandleFunc("/auth/login", auth_api.Authenticate).Methods("POST")
	r.HandleFunc("/auth/register", auth_api.Register).Methods("POST")
}

func HandleUserRequests(r *mux.Router) {
	r.HandleFunc("/users", user_api.GetUsers).Methods("GET")
	r.HandleFunc("/users", user_api.CreateUser).Methods("POST")
	r.HandleFunc("/users/{uuid}", user_api.GetUser).Methods("GET")
	r.HandleFunc("/users/{uuid}/challenges", user_api.GetUserChallenges).Methods("GET")
	r.HandleFunc("/users/{uuid}/bets", user_api.GetUserBets).Methods("GET")
	r.HandleFunc("/users/{uuid}", user_api.DeleteUserById).Methods("DELETE")
}

func HandleChallengeRequests(r *mux.Router) {
	r.HandleFunc("/challenges", challenge_api.GetChallenges).Methods("GET")
	r.HandleFunc("/challenges", challenge_api.CreateChallenge).Methods("POST")
	r.HandleFunc("/challenges/{uuid}", challenge_api.GetChallenge).Methods("GET")
	r.HandleFunc("/challenges/{uuid}", challenge_api.UpdateChallenge).Methods("PATCH")
	r.HandleFunc("/challenges/{uuid}/bets", challenge_api.GetChallengeBets).Methods("GET")
	r.HandleFunc("/challenges/{uuid}/bets/amount", challenge_api.GetAmountByChallenges).Methods("GET")
}

func HandleBetRequests(r *mux.Router) {
	r.HandleFunc("/bets", bet_api.GetBets).Methods("GET")
	r.HandleFunc("/bets", bet_api.CreateBet).Methods("POST")
	r.HandleFunc("/bets/{uuid}", bet_api.GetBet).Methods("GET")
}
