package api

import (
	"log"
	"net/http"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userid := ps.ByName("id")
	messageid := ps.ByName("message_id")

	if messageid == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	if userid != ctx.Token {
		return
	}

	err := rt.db.UncommentMessage(messageid)
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Comment succesfully deleted", http.StatusOK)
}
