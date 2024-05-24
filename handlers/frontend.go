package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("assets/index.html")
	if err != nil {
		panic(err)
	}
	html := fmt.Sprint(string(data))
	fmt.Fprintf(w, html)
}
