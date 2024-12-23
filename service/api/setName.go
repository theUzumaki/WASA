package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var id = ps.ByName("id")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if id != ctx.Token {
		return
	}

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var valid = regexp.MustCompile("^.*?$").MatchString(user.Name)
	if !valid {
		http.Error(w, "Username not valid", http.StatusExpectationFailed)
		return
	}

	err = rt.db.SetName(id, user.Name)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Name succesfully changed", http.StatusOK)
}
