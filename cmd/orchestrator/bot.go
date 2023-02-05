package main

import (
	"fmt"
	"log"
	"sync"

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
	fmt.Printf("\nEvent %v created with id: %v", title, eventId)
	events.CreateFreeTicket(eventId, token)
	events.ValidateEvent(eventId)

}

func getEventPicture(eventType string) events.Picture {
	url := unsplash.FetchImage(eventType)

	return events.Picture{
		BucketPath: url,
		Url:        url,
	}
}

func main() {
	var wg sync.WaitGroup
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tokens := events.ConnectBotUsers()
	nbOfEvents := 5

	for i := 0; i < nbOfEvents; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			generateEvent(tokens[0])
			generateEvent(tokens[1])

		}()

	}

	wg.Wait()

}
