package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	var userName = user.Name

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var valid = regexp.MustCompile("^.*?$").MatchString(userName)
	if !valid {
		http.Error(w, "Username not valid", http.StatusExpectationFailed)
		return
	}

	user = User{
		Id:   -1,
		Name: userName,
	}

	var status string
	id, status, err := rt.db.LoginManager(user.ApiUserToDB())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	type token struct {
		UserId int `json:"userId"`
		Token  int `json:"token"`
	}
	var tok token

	tok.UserId = id
	tok.Token = id

	if status == "user exist" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(tok)
}
