package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"wasatext/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setGroupPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	id := ps.ByName("group_id")

	if id == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if id != ctx.Token {
		return
	}

	err := r.ParseMultipartForm(5 * (10 ^ 6))
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Image size limit exceeded", http.StatusBadRequest)
		return
	}

	rawfile, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	image, err := io.ReadAll(rawfile)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println(http.DetectContentType(image))
	if http.DetectContentType(image) != "image/jpeg" {
		http.Error(w, "Wrong file type", http.StatusBadRequest)
		return
	}

	defer rawfile.Close()
	err = os.WriteFile("grouppics/"+id, image, 0700)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Ok", http.StatusOK)
}
