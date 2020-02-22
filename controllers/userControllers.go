package controllers

import (
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Users!")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Users!")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get User!")
}

func GetUserChallenges(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get User Challenges!")
}

func GetUserBets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get User Bets!")
}
