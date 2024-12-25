package api

import (
	"net/http"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userid := ps.ByName("id")
	messageid := ps.ByName("message_id")

	if messageid == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	if userid != ctx.Token {
		return
	}

	err := rt.db.DeleteMessage(messageid)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Message succesfully deleted", http.StatusOK)
}
