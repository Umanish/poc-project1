package handler

import (
	"net/http"
)

func Timer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Welcome Tuesday!"))
}
