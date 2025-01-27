package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) newChat(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var id = ps.ByName("id")
	var chat Chat
	var err error

	if id != ctx.Token {
		return
	}

	err = json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var valid = regexp.MustCompile("^.*?$").MatchString(chat.Name)
	if !valid {
		http.Error(w, "Chat name not valid", http.StatusExpectationFailed)
		return
	}

	chat_id, err := rt.db.NewChat(id, chat.ApiChatToDB())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	chat.Id = chat_id

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(chat)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
