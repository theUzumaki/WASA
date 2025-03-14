package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"wasatext/service/api/reqcontext"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var id = ps.ByName("id")
	var chatid = ps.ByName("chat_id")
	var message structs.Message

	if id != ctx.Token || id == "" || chatid == "" {
		return
	}

	err := r.ParseMultipartForm(4 << 20)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Retrieve the form values
	message.ChatId, err = strconv.Atoi(r.FormValue("chat_id"))
	if err != nil {
		http.Error(w, "Invalid chat id", http.StatusBadRequest)
		return
	}

	sender_id, err := strconv.Atoi(r.FormValue("sender_id"))
	if err != nil {
		http.Error(w, "Invalid sender", http.StatusBadRequest)
		return
	}
	message.Sender = structs.User{
		Id:      sender_id,
		Name:    r.FormValue("sender_name"),
		Picture: r.FormValue("sender_pic"),
	}
	dateStr := r.FormValue("date")
	message.Date, err = time.Parse(time.RFC3339, dateStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}
	message.Content = r.FormValue("content")

	chat, err := rt.db.SendMessage(message)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(chat)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
