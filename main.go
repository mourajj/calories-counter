package main

import (
	"calories-counter/handlers"
	"calories-counter/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

	//Loading environment variables

	godotenv.Load(".env")
	food := "arroz branco" // Change for the desired food
	amount := 100.0        // Change for the desired amount

	// Using chatGPT to generate the response
	prompt := fmt.Sprintf("Me diga somente a quantidade exata de calorias (numero) que tem em %v gramas de %s", amount, food)
	response, err := handlers.GetChatGPTResponse(prompt)
	if err != nil {
		fmt.Println("Erro ao obter resposta do ChatGPT:", err)
		return
	}

	// Extract the JSON response
	var chatGPTResponse model.ChatGPTResponse
	err = json.Unmarshal(response, &chatGPTResponse)
	if err != nil {
		fmt.Println("Erro ao analisar resposta do ChatGPT:", err)
		return
	}

	message := chatGPTResponse.Choices[0].Text
	fmt.Println("Mensagem gerada pelo ChatGPT:", message)
}
