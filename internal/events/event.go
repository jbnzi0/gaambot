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
	"time"

	"github.com/jbnzi0/gaambot/internal/datagenerator"
)

var (
	chessURL = os.Getenv("CHESS_API_URL")
)

type Event struct {
	Title           string                `json:"title"`
	Date            string                `json:"date"`
	Price           int64                 `json:"price"`
	MaxParticipants int64                 `json:"maxParticipants"`
	EventType       string                `json:"type"`
	Category        string                `json:"category"`
	Address         datagenerator.Address `json:"address"`
	Picture         Picture               `json:"picture"`
	Description     string                `json:"description"`
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

func GetRandomEventType() string {
	rand.Seed(time.Now().UnixNano())
	types := []string{"public", "exclusive"}
	min := 0
	max := len(types) - 1
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
	max := len(categories) - 1

	return categories[rand.Intn(max-min+1)+min]
}

func CreateFreeTicket(eventId string, token string) {
	url := chessURL + "/organizer/events/" + eventId + "/tickets"

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
		fmt.Println(response.Status)
		log.Fatal()
	}
}

func CreateEvent(event Event, token string) string {
	url := chessURL + "/events"

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
