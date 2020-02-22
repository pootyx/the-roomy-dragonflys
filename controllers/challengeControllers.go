package controllers

import (
	"fmt"
	"net/http"
)

func GetChallenges(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Challenges!")
}

func CreateChallenge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Challenge!")
}

func GetChallenge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Challenge!")
}

func GetChallengeBets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Challenge Bets!")
}
