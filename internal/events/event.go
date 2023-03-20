package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jbnzi0/gaambot/internal/datagenerator"
	"github.com/jbnzi0/gaambot/internal/maps"
	"github.com/jbnzi0/gaambot/pkg/chatgpt"
	"github.com/jbnzi0/gaambot/pkg/unsplash"
	"github.com/kr/pretty"
)

type Event struct {
	Title        string       `json:"title"`
	Date         string       `json:"date"`
	EventType    string       `json:"type"`
	Category     string       `json:"category"`
	Address      maps.Address `json:"address"`
	Picture      Picture      `json:"picture"`
	Description  string       `json:"description"`
	RefundPolicy string       `json:"refundPolicy"`
}

type Picture struct {
	BucketPath string `json:"bucketPath"`
	Url        string `json:"url"`
}

type EventResponse struct {
	Id string `json:"_id"`
}

type TicketRequest struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

type UpdateEventRequest struct {
	// Picture      Picture `json:"picture"`
	RefundPolicy string `json:"refundPolicy"`
}

func GetRandomEventCategory() string {
	rand.Seed(time.Now().UnixNano())
	categories := []string{
		"party",
		"networking",
		"outdoor",
		"sports",
	}

	max := len(categories) - 1

	return categories[rand.Intn(max+1)]
}

func createFreeTicket(eventId string, token string) {
	url := os.Getenv("CHESS_API_URL") + "/organizer/events/" + eventId + "/tickets"

	data, err := json.Marshal(TicketRequest{
		Name:     "Free",
		Price:    0,
		Quantity: 100,
	})

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

	if response.StatusCode != 201 {
		log.Fatal(response.Status)
	}
}

func createEvent(event Event, token string) string {
	url := os.Getenv("CHESS_API_URL") + "/events"

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

	if response.StatusCode != 201 {
		log.Fatal(response.Status)
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

func UpdateEvent(eventId string, token string, payload UpdateEventRequest) string {
	url := os.Getenv("CHESS_API_URL") + "/events/" + eventId

	data, err := json.Marshal(payload)

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

	if response.StatusCode != 201 {
		log.Fatal(response.Status)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var result EventResponse
	json.Unmarshal(body, &result)

	fmt.Println("Updated eventId: " + result.Id)
	return result.Id
}

func validateEvent(eventId string) {
	url := os.Getenv("CHESS_API_URL") + "/bendo/events/" + eventId + "/review"

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
		log.Fatal(response.Status)
	}
}

func getEventPicture(title string) Picture {
	url := unsplash.Search(strings.ToLower(title))

	return Picture{
		BucketPath: url,
		Url:        url,
	}
}

func GenerateEvent(token string, data PartialEvent) {
	category := data.Category
	address, formattedAddress := maps.SearchPlace(maps.GetRandomNeighbourhood())
	title := chatgpt.Chat("Short " + datagenerator.GetRandomAdjective() + " event name for a " + category + " in" + formattedAddress)
	description := chatgpt.Chat("Short description for a " + category + " called " + title + " in " + formattedAddress)
	picture := getEventPicture(category + " " + title)

	event := Event{
		Category:     category,
		Address:      address,
		Picture:      picture,
		EventType:    "exclusive",
		Date:         datagenerator.GetRandomDate(),
		Description:  description,
		Title:        title,
		RefundPolicy: "up-to-seven-days",
	}

	eventId := createEvent(event, token)
	pretty.Printf("\nEvent \"%v\" created in %v with ID: %v\n", title, formattedAddress, eventId)

	createFreeTicket(eventId, token)
	validateEvent(eventId)

}
