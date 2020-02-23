package auth_api

import (
	"encoding/json"
	"github.com/stack-attack/the-roomy-dragonflys/user-api"
	u "github.com/stack-attack/the-roomy-dragonflys/utils"
	"net/http"
)

var Register = func(w http.ResponseWriter, r *http.Request) {

	user := &user_api.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message("Invalid request"))
		return
	}

	resp := user.Register() //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	user := &user_api.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message("Invalid request"))
		return
	}

	resp := user_api.Login(user.Email, user.Password)
	u.Respond(w, resp)
}
