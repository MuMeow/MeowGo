package main

import (
	cats "MeowGo/services/cat/controller"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()

	cat := r.PathPrefix("/cats").Subrouter()
	cats.Controller(cat)

	r.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"alive":true}`))
	}).Methods("GET")

	log.Print("running on :10801")

	http.ListenAndServe(":10801", handlers.CORS(header, methods, origins, credentials)(r))
}
