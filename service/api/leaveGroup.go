package api

import (
	"net/http"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) leaveGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var userid = ps.ByName("id")
	var chatid = ps.ByName("group_id")

	if userid == "" || chatid == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	if userid != ctx.Token {
		return
	}

	err := rt.db.LeaveGroup(chatid, userid)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Group succesfully left", http.StatusOK)
}
