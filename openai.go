package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	apiURL = os.Getenv("OPENAI_API_URL")
	apiKey = os.Getenv("OPENAI_API_KEY")
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

func GenerateEventName(eventType string, category string, country string) string {
	text := "Can you generate a " + getRandomAdjective() + " event name for an event " + eventType + " " + category + " in " + country + " ?"

	fmt.Println(text)
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

func GenerateDescription(title string, eventType string, category string, country string) string {
	text := "Can you generate a short description for an event called " + title + " which is a " + eventType + " " + category + " in " + country + " ?"

	data, err := json.Marshal(map[string]string{
		"prompt":      text,
		"max_tokens":  "20",
		"model":       "text-davinci-003",
		"temperature": "0.5",
	})

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
