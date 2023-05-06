package handlers

import (
	"bytes"
	"calories-counter/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var input *model.Input

func GetChatGPTResponse(prompt string) ([]byte, error) {

	//Creating the requestbody for openAPI endpoint
	requestBody, err := json.Marshal(map[string]interface{}{
		"prompt":      prompt,
		"max_tokens":  150,
		"temperature": 0.7,
	})
	if err != nil {
		return nil, err
	}

	//Creating a request object
	request, err := http.NewRequest("POST", os.Getenv("openAPI_endpoint"), bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	//Setting the headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+os.Getenv("openAPI_key"))

	//Creating the HTTP client object and performing the request with its function
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	//Getting the response.body and returning it
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func InputHandler(w http.ResponseWriter, r *http.Request) {

	//Getting the user input values
	food := r.FormValue("food")
	amount := r.FormValue("amount")

	//Creating an input object
	input = &model.Input{
		Food:   food,
		Amount: amount,
	}

	// Converting the bool value (if exists)
	if r.FormValue("cooked") != "" {
		cooked, err := strconv.ParseBool(r.FormValue("cooked"))
		if err != nil {
			// Trata o erro, se houver
			http.Error(w, "Valor inválido", http.StatusBadRequest)
			return
		}
		input.Cooked = cooked
	}

	//Process the GPT question and return the response according to the inputs
	w.Write([]byte(processGPTResponse()))
}

func processGPTResponse() string {

	godotenv.Load(".env")
	food := input.Food     // Change for the desired food
	amount := input.Amount // Change for the desired amount
	cooked := ""

	if input.Cooked {
		cooked = "cozido"
	}

	// Using chatGPT to generate the response
	prompt := fmt.Sprintf("Me diga somente o valor de quantas calorias tem %v gramas de %s %s de acordo com a Tabela de Composição de Alimentos do IBGE", amount, food, cooked)
	response, err := GetChatGPTResponse(prompt)
	if err != nil {
		log.Panic("Erro ao obter resposta do ChatGPT:", err)
	}

	// Extract the JSON response
	var chatGPTResponse model.ChatGPTResponse
	err = json.Unmarshal(response, &chatGPTResponse)
	if err != nil {
		log.Panic("Erro ao analisar resposta do ChatGPT:", err)
	}

	//Returning the message
	message := chatGPTResponse.Choices[0].Text
	return message
}
