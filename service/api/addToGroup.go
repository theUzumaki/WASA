package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var userid = ps.ByName("id")
	var chatid = ps.ByName("group_id")
	var newuser User
	err := json.NewDecoder(r.Body).Decode(&newuser)

	if err != nil || userid == "" || chatid == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if userid != ctx.Token {
		return
	}

	err = rt.db.AddToGroup(userid, chatid, strconv.Itoa(newuser.Id))
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "User succesfully added", http.StatusOK)
}
