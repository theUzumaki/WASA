package api

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) newChatApi(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var id = ps.ByName("id")
	var chat Chat
	var err error

	err = json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var valid = regexp.MustCompile("^.*?$").MatchString(chat.Name)
	if !valid {
		http.Error(w, "Chat name not valid", http.StatusExpectationFailed)
		return
	}

	err = rt.db.NewChat(id, chat.ApiChatToDB())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Chat succesfully created", http.StatusCreated)
}
