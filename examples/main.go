package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bniladridas/go-llm-chat/chatapi"
)

func main() {
	// Load API key from environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set")
	}

	// Initialize client with API key and provider
	client := chatapi.NewClient(apiKey, "openai")

	// Example of synchronous chat
	response, err := client.Chat("Hello AI!")
	if err != nil {
		log.Fatalf("Chat error: %v", err)
	}
	fmt.Println("AI Response:", response)

	// Example of streaming chat
	ch := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	go func() {
		err := client.StreamChat(ctx, "Tell me a story", ch)
		if err != nil {
			log.Printf("StreamChat error: %v", err)
		}
	}()

	for msg := range ch {
		fmt.Println("Streaming Response:", msg)
	}
}
