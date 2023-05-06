package handlers

import (
	"bytes"
	"calories-counter/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func GetChatGPTResponse(prompt string) ([]byte, error) {
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

func InputHandler(w http.ResponseWriter, r *http.Request) {
	// Recebe o valor da caixa de texto do formulário

	food := r.FormValue("food")
	amount := r.FormValue("amount")

	input := model.Input{
		Food:   food,
		Amount: amount,
	}

	// Converte o valor recebido de string para bool
	if r.FormValue("cooked") != "" {
		cooked, err := strconv.ParseBool(r.FormValue("cooked"))
		if err != nil {
			// Trata o erro, se houver
			http.Error(w, "Valor inválido", http.StatusBadRequest)
			return
		}
		input.Cooked = cooked
	}

	fmt.Println(input.Food)
	fmt.Println(input.Amount)
	fmt.Println(input.Cooked)
}
