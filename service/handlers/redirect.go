package handlers

import "net/http"

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Redirect URL"))
}
