package api

import (
	"encoding/json"
	"net/http"
	"wasatext/service/api/reqcontext"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	id := ps.ByName("id")

	if id == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if id != ctx.Token {
		return
	}

	var user structs.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = rt.db.SetMyPhoto(id, user.Picture)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Picture succesfully changed", http.StatusOK)
}
