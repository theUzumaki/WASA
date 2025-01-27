package api

import (
	"encoding/json"
	"net/http"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	chatId := ps.ByName("chat_id")
	userId := ps.ByName("id")

	if userId != ctx.Token {
		return
	}

	if chatId == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	chat, err := rt.db.GetConversation(chatId)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(chat)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
