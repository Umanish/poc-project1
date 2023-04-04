package main

import (
	"log"
	"net/http"

	"github.com/Umanish/poc-project1/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	port := "8081"

	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()

	r.Get("/welcome", handler.Welcome)
	r.Get("/timer", handler.Timer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello World!"))
	})

	log.Fatal(http.ListenAndServe(":"+port, r))
}
