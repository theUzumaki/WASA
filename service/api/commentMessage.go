package api

import (
	"encoding/json"
	"net/http"
	"wasatext/service/api/reqcontext"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userid := ps.ByName("id")
	messageid := ps.ByName("message_id")
	var comm structs.Text

	err := json.NewDecoder(r.Body).Decode(&comm)
	if err != nil || messageid == "" || comm.Text == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	if userid != ctx.Token {
		return
	}

	err = rt.db.CommentMessage(messageid, comm.Text, userid)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Comment succesfully added", http.StatusOK)
}
