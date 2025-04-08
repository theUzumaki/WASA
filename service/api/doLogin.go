package api

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"wasatext/service/api/reqcontext"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var user structs.User
	err := json.NewDecoder(r.Body).Decode(&user)
	var userName = user.Name

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var valid = regexp.MustCompile("^.*?$").MatchString(userName)
	if !valid {
		http.Error(w, "Username not valid", http.StatusBadRequest)
		return
	}

	imagePath := "./stdpics/profilepic.jpg"
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		http.Error(w, "Unable to read image", http.StatusInternalServerError)
		return
	}

	user.Picture = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imageData)

	user = structs.User{
		Id:      -1,
		Name:    userName,
		Picture: user.Picture,
	}

	var status string
	var newuser structs.User
	dbUser, status, err := rt.db.LoginManager(user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	newuser = structs.User{
		Id:      dbUser.Id,
		Name:    dbUser.Name,
		Picture: dbUser.Picture,
	}

	if status == "user exist" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(newuser)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
