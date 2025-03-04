# go-llm-chat

A Go package for integrating Large Language Models (LLMs) like Gemini, OpenAI, and Grok into your Go projects.

## Features

- **Chat API (`Chat()`)** – Sends a message to the selected LLM and gets a response.
- **Streaming Support (`StreamChat()`)** – Uses WebSockets for real-time responses.
- **Multi-Provider Support** – Works with Gemini, OpenAI, and Grok APIs.
- **Authentication Handling** – Uses API keys dynamically.
- **Error Handling & Logging** – Provides structured logs for debugging.
- **Easy Installation & Usage** – Available via `go get github.com/bniladridas/go-llm-chat`.

## Installation

```sh
go get github.com/bniladridas/go-llm-chat
```

## Usage

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bniladridas/go-llm-chat/chatapi"
)

func main() {
	// Initialize client with API key and provider
	client := chatapi.NewClient("YOUR_API_KEY", "openai")

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
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
