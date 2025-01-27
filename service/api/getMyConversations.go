package api

import (
	"encoding/json"
	"net/http"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyConversations(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user := ps.ByName("id")
	userId := ps.ByName("id")

	if userId != ctx.Token {
		return
	}

	if user == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	chats, err := rt.db.GetMyConversations(user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(chats)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
