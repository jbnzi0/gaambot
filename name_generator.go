package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiURL = "https://api.openai.com/v1/completions"
	// apiKey = "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"
	apiKey = "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"
)

func GenerateEventName(eventType string, country string) {
	text := "Can you generate a random event name for an event " + eventType + "in " + country + " ?"

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

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	fmt.Println("Response:", result)

}
