package main

import "fmt"

func create() {
	tokens := ConnectBotUsers()

	for _, token := range tokens {
		fmt.Println("Token:", token)
		// GenerateEvent(token)
	}
}

func main() {
	// create()
	UploadPicture("party", "testId")
}
