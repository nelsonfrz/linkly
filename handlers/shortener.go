package handlers

import (
	"fmt"
	"linkly/storage"
	"net/http"

	"github.com/gorilla/mux"
)

const BaseUrl string = "http://linkly.com:3000"

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	id := storage.SaveUrl(url)
	fmt.Fprintf(w, "%s/r/%s", BaseUrl, id)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	redirect_url := storage.GetUrl(id)
	http.Redirect(w, r, redirect_url, http.StatusFound)
}
