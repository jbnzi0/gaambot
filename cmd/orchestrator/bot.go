package main

import (
	"fmt"
	"log"

	"github.com/jbnzi0/gaambot/internal/datagenerator"
	"github.com/jbnzi0/gaambot/internal/events"
	"github.com/jbnzi0/gaambot/pkg/openai"
	"github.com/jbnzi0/gaambot/pkg/unsplash"
	"github.com/joho/godotenv"
)

func GenerateEvent(token string) {
	category := events.GetRandomEventCategory()
	eventType := events.GetRandomEventType()
	fmt.Println(category, eventType)
	address, formattedAddress := datagenerator.GetRandomAddress()
	title := openai.TextCompletion("Can you generate a " + datagenerator.GetRandomAdjective() + " event name for an event " + eventType + " " + category + " in " + formattedAddress + " ?")
	description := openai.TextCompletion("Can you generate a short description for an event called " + title + " which is a " + eventType + " " + category + " in " + formattedAddress + " ?")
	picture := getEventPicture(eventType)

	event := events.Event{
		Category:    category,
		Address:     address,
		Picture:     picture,
		EventType:   eventType,
		Date:        datagenerator.GetRandomDate(),
		Description: description,
		Title:       title,
	}

	eventId := events.CreateEvent(event, token)
	events.CreateFreeTicket(eventId, token)
}

// func create() {
// 	tokens := ConnectBotUsers()

// 	for _, token := range tokens {
// 		fmt.Println("Token:", token)
// 		// GenerateEvent(token)
// 	}
// }

func getEventPicture(eventType string) events.Picture {
	url := unsplash.FetchImage(eventType)

	return events.Picture{
		BucketPath: url,
		Url:        url,
	}
}

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// create()
	GenerateEvent("")
	// picture := getPicture("party", "testId")
}
