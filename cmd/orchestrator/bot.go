package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jbnzi0/gaambot/internal/events"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func orchestrate() {
	var wg sync.WaitGroup

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

func loadDevEnvironment() {
	env := os.Getenv("ENVIRONMENT")

	if env == "production" {
		return
	}

	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadDevEnvironment()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	scheduler := cron.New()

	scheduler.AddFunc("0 6 * * *", func() {
		fmt.Println("Running CRON...")
		orchestrate()
		fmt.Println("Creation of events done!")
	})

	scheduler.Start()
	wg.Wait()
}
