package main

import (
	"bytes"
	"calories-counter/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//Loading environment variables
	godotenv.Load(".env")
	food := "arroz branco" // Change for the desired food
	amount := 100.0        // Change for the desired amount

	// Using chatGPT to generate the response
	prompt := fmt.Sprintf("Me diga somente a quantidade exata de calorias (numero) que tem em %v gramas de %s", amount, food)
	response, err := getChatGPTResponse(prompt)
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

func getChatGPTResponse(prompt string) ([]byte, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"prompt":      prompt,
		"max_tokens":  150,
		"temperature": 0.7,
	})
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/text-davinci-003/completions", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+os.Getenv("openAPI_key"))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
