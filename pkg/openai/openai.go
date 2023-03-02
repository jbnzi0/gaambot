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

type ChatGPTConversation struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTRequest struct {
	Messages         []ChatGPTConversation `json:"messages"`
	MaxTokens        int                   `json:"max_tokens"`
	Model            string                `json:"model"`
	Temperature      float32               `json:"temperature"`
	FrequencyPenalty int                   `json:"frequency_penalty"`
}

type ChatGPTResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message      ChatGPTConversation `json:"message"`
		Index        int                 `json:"index"`
		FinishReason string              `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     string `json:"prompt_tokens"`
		CompletionTokens string `json:"completion_tokens"`
		TotalTokens      string `json:"total_tokens"`
	} `json:"usage"`
}

func Chat(text string) string {
	var (
		apiURL = os.Getenv("OPENAI_API_URL")
		apiKey = os.Getenv("OPENAI_API_KEY")
	)

	message := ChatGPTConversation{Role: "user", Content: text}

	payload := ChatGPTRequest{
		Messages:         []ChatGPTConversation{message},
		MaxTokens:        35,
		Model:            "gpt-3.5-turbo",
		Temperature:      1,
		FrequencyPenalty: 2,
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

	answer := strings.TrimLeft(result.Choices[0].Message.Content, "\r\n\"")

	return strings.ReplaceAll(answer, "\"", "")
}
