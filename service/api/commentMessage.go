package api

import (
	"encoding/json"
	"log"
	"net/http"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userid := ps.ByName("id")
	messageid := ps.ByName("message_id")
	var comment text

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil || messageid == "" || comment.Text == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	if userid != ctx.Token {
		return
	}

	err = rt.db.CommentMessage(messageid, comment.Text)
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Comment succesfully added", http.StatusOK)
}
