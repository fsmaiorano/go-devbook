package controllers

import "net/http"

func Authentication(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authentication"))
}
