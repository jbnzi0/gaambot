package main

import (
	"log"
	"sync"

	"github.com/jbnzi0/gaambot/internal/events"
	"github.com/joho/godotenv"
)

func orchestrate() {
	var wg sync.WaitGroup
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	tokens := events.ConnectBotUsers()
	nbOfEvents := 5

	partialEvents := events.ReadEventsFile()
	for i := 0; i < nbOfEvents; i++ {
		wg.Add(1)
		firstEvent := events.GetRandomEventData(partialEvents)
		secondEvent := events.GetRandomEventData(partialEvents)

		go func() {
			defer wg.Done()
			events.GenerateEvent(tokens[0], firstEvent)
			events.GenerateEvent(tokens[1], secondEvent)
		}()

	}

	wg.Wait()
}

// TODO: Add cron to run it every week
func main() {
	orchestrate()
}
