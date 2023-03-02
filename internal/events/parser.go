package events

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type PartialEvent struct {
	Name        string
	Category    string
	Description string
	City        string
}

func GetRandomEventData(events []PartialEvent) PartialEvent {
	rand.Seed(time.Now().UnixNano())
	max := len(events) - 1

	return events[rand.Intn(max+1)]
}

func ReadEventsFile() []PartialEvent {
	file, err := os.Open("../../assets/events.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return createEventList(data)
}

func createEventList(data [][]string) []PartialEvent {
	var usersList []PartialEvent

	for i, line := range data {
		if i == 0 {
			continue
		}

		event := PartialEvent{
			Name:        line[0],
			Category:    strings.ToLower(line[1]),
			Description: line[2],
			City:        strings.ToLower(line[3]),
		}
		usersList = append(usersList, event)
	}

	return usersList
}
