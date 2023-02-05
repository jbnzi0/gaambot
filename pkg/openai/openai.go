package openai

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type ChatGPTRequest struct {
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Model       string  `json:"model"`
	Temperature float32 `json:"temperature"`
}

type ChatGPTResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string `json:"text"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     string `json:"prompt_tokens"`
		CompletionTokens string `json:"completion_tokens"`
		TotalTokens      string `json:"total_tokens"`
	} `json:"usage"`
}

func TextCompletion(text string) string {
	var (
		apiURL = os.Getenv("OPENAI_API_URL")
		apiKey = os.Getenv("OPENAI_API_KEY")
	)

	payload := ChatGPTRequest{
		Prompt:      text,
		MaxTokens:   7,
		Model:       "text-davinci-003",
		Temperature: 0.5,
	}

	data, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var result ChatGPTResponse
	json.Unmarshal(body, &result)

	return strings.TrimLeft(result.Choices[0].Text, "\r\n\"")
}
