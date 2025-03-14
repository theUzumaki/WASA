package api

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"wasatext/service/api/reqcontext"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) newChat(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var id = ps.ByName("id")
	var chat structs.Chat
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

	if chat.Name != "chat" && chat.Picture == "" {
		imagePath := "./stdpics/grouppic.jpg"
		imageData, err := ioutil.ReadFile(imagePath)
		if err != nil {
			http.Error(w, "Unable to read image", http.StatusInternalServerError)
			return
		}

		chat.Picture = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imageData)
	}

	chat_id, err := rt.db.NewChat(id, chat)
	if err != nil && err.Error() == "chat already existing" {
		chatDB, err := rt.db.GetConversation(strconv.Itoa(chat_id))
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		chat = chatDB
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	} else {
		chat.Id = chat_id
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(chat)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
