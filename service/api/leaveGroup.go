package api

import (
	"log"
	"net/http"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) leaveGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var userid = ps.ByName("id")
	var chatid = ps.ByName("group_id")

	err := rt.db.LeaveGroup(chatid, userid)
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Group succesfully left", http.StatusOK)
}
