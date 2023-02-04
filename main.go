package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func create() {
	tokens := ConnectBotUsers()

	for _, token := range tokens {
		fmt.Println("Token:", token)
		// GenerateEvent(token)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// create()
	// UploadPicture("party", "testId")
}
