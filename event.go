package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	chessURL = "https://chess-services.heroku.app/v1/"
)

type Event struct {
	Title           string `json:"title"`
	Date            string `json:"date"`
	Price           int64  `json:"price"`
	MaxParticipants int64  `json:"maxParticipants"`
	EventType       string `json:"type"`
	Category        string `json:"category"`
	Address         struct {
		AddressComponents []struct {
			LongName         string   `json:"long_name"`
			ShortName        string   `json:"short_name"`
			Types            []string `json:"types"`
			FormattedAddress string   `json:"formattedAddress"`
			Latitude         string   `json:"latitude"`
			Longitude        string   `json:"longitude"`
		} `json:"addressComponents"`
	} `json:"address"`
	Picture struct {
		BucketPath string `json:"bucketPath"`
		Url        string `json:"url"`
	} `json:"picture"`
	Description string `json:"description"`
}

type EventResponse struct {
	Id string `json:"_id"`
}

// func GenerateEvent(token string) {
// 	category := GetRandomEventCategory()
// 	eventType := GetRandomEventType()
// 	address, formattedAddress := GetRandomAddress()
// 	title := GenerateEventName(eventType, category, formattedAddress)
// 	// picture := UploadPicture(title)

// 	event := Event{
// 		Category:    category,
// 		Address:     address,
// 		Picture:     picture,
// 		EventType:   eventType,
// 		Date:        GetRandomDate(),
// 		Price:       0,
// 		Description: " ",
// 		Title:       title,
// 	}

// 	eventId := CreateEvent(event, token)

// 	fmt.Println("Event created with id : " + eventId)
// }

func GetRandomEventType() string {
	rand.Seed(time.Now().UnixNano())
	types := []string{"public", "exclusive"}
	min := 0
	max := len(types)
	return types[rand.Intn(max-min+1)+min]
}

func GetRandomEventCategory() string {
	rand.Seed(time.Now().UnixNano())
	categories := []string{
		"party",
		"networking",
		"outdoor",
		"sports",
	}

	min := 0
	max := len(categories)

	return categories[rand.Intn(max-min+1)+min]
}

func CreateEvent(event Event, token string) string {
	url := chessURL + "/bendo/events"

	data, err := json.Marshal(event)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Println(response.Status)
		log.Fatal()
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var result EventResponse
	json.Unmarshal(body, &result)

	return result.Id
}

func ValidateEvent(eventId string) {
	url := chessURL + "/bendo/events/" + eventId

	data, err := json.Marshal(map[string]string{
		"status": "accepted",
	})

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Println(response.Status)
		log.Fatal()
	}
}
