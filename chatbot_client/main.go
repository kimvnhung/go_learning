package main

import (
	"log"

	"github.com/kimvnhung/go_learning/chatbot_client/server"
)

func main() {
	log.Println("Starting Chatbot Client")
	// Create a new ChatbotClient
	apiOutPort := 1997
	apiServicePort := 1998
	chatbotServicePort := 1999
	err := server.StartServices(apiOutPort, apiServicePort, chatbotServicePort)
	if err != nil {
		log.Fatalf("Error starting services: %v", err)
	}
}
