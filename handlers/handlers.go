package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
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
