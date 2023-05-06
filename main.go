package main

import (
	"calories-counter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//Starting router and listen to 5500 port
	r := mux.NewRouter()
	r.HandleFunc("/input", handlers.InputHandler).Methods("POST")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(":5500", r)
}
