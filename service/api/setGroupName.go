package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setGroupName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var id = ps.ByName("id")
	var groupid = ps.ByName("group_id")
	type groupname struct {
		Name string `json:"name"`
	}

	var name groupname
	err := json.NewDecoder(r.Body).Decode(&name)

	if id != ctx.Token {
		return
	}

	if err != nil || id == "" || groupid == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var valid = regexp.MustCompile("^.*?$").MatchString(name.Name)
	if !valid {
		http.Error(w, "Groupname not valid", http.StatusExpectationFailed)
		return
	}

	err = rt.db.SetGroupName(id, name.Name)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Name succesfully changed", http.StatusOK)
}
