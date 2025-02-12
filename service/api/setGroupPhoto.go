package api

import (
	"encoding/json"
	"net/http"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setGroupPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	id := ps.ByName("id")
	group_id := ps.ByName("group_id")

	if id == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if id != ctx.Token {
		return
	}

	var chat Chat
	err := json.NewDecoder(r.Body).Decode(&chat)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = rt.db.SetGroupPhoto(group_id, chat.Picture)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Ok", http.StatusOK)
}
