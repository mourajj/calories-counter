package model

type ChatGPTResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

type Input struct {
	Food   string
	Amount string
	Cooked bool
}
