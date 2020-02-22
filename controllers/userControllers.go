package controllers

import (
	repository "../repository"
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Users!")
	repository.CreateUserToDb()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Users!")
	repository.CreateUserToDb()
	/*user := &repository.User{}

	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := user.Create()
	u.Respond(w, resp)*/

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
