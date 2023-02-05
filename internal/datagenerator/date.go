package datagenerator

import (
	"math/rand"
	"time"
)

func GetRandomDate() string {
	rand.Seed(time.Now().UnixNano())

	date := time.Now()
	daysToAdd := rand.Intn(15-1) + 1

	randomDate := date.AddDate(0, 0, daysToAdd)

	return randomDate.Format(time.RFC3339)
}
