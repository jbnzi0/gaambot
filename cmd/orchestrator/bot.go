package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jbnzi0/gaambot/internal/datagenerator"
	"github.com/jbnzi0/gaambot/internal/events"
	"github.com/jbnzi0/gaambot/pkg/openai"
	"github.com/jbnzi0/gaambot/pkg/unsplash"
	"github.com/joho/godotenv"
)

func generateEvent(token string) {
	category := events.GetRandomEventCategory()
	eventType := events.GetRandomEventType()
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
	fmt.Printf("Event %v created with id: %v", title, eventId)
	events.CreateFreeTicket(eventId, token)
	events.ValidateEvent(eventId)
}

func orchestrate() {
	tokens := events.ConnectBotUsers()

	for _, token := range tokens {
		generateEvent(token)
		os.Exit(-1)
	}
}

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

	orchestrate()
}
