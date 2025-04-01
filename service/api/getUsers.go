package api

import (
	"encoding/json"
	"net/http"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userId := ps.ByName("id")

	if userId != ctx.Token {
		return
	}

	username := ps.ByName("searchquery")

	if username == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	users, err := rt.db.GetUsers(username)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
