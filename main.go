package main

import (
	"calories-counter/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	//Starting router
	r := mux.NewRouter()

	//Defining the routes
	r.HandleFunc("/input", handlers.InputHandler).Methods("POST")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.ServeFile(w, r, "index.html")
	})

	//Setting the HTTP server to listen to the port 5500
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
