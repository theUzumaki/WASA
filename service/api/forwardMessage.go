package api

import (
	"encoding/json"
	"log"
	"net/http"
	"wasatext/service/api/reqcontext"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userid := ps.ByName("id")
	oldchatid := ps.ByName("chat_id")
	messageid := ps.ByName("message_id")

	if userid != ctx.Token || oldchatid == "" || messageid == "" {
		return
	}

	var newchatid structs.Id

	err := json.NewDecoder(r.Body).Decode(&newchatid)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = rt.db.ForwardMessage(userid, messageid, newchatid.Id)
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Message succesfully forwarded", http.StatusCreated)
}
