# GoLLMChat 🚀
A Go package for integrating Large Language Models (LLMs) like Gemini, OpenAI, and Grok into your Go projects with ease.

## Features
- **Chat API (`Chat()`)** – Send a message to the selected LLM and receive a response synchronously.
- **Streaming Support (`StreamChat()`)** – Get real-time responses via WebSockets for interactive applications.
- **Multi-Provider Support** – Seamlessly integrate with Gemini, OpenAI, and Grok APIs.
- **Authentication Handling** – Configure provider-specific API keys effortlessly.
- **Error Handling & Logging** – Built-in structured logging for debugging and monitoring.
- **Easy Installation & Usage** – Install with a single command: `go get github.com/bniladridas/go-llm-chat`.

## Requirements
- Go 1.18 or later
- Valid API keys from OpenAI, Gemini, or xAI (Grok)

## Installation
```sh
go get github.com/bniladridas/go-llm-chat
```

## Usage
Here’s how to use `GoLLMChat` in your Go projects:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/bniladridas/go-llm-chat/chatapi"
)

func main() {
    // Initialize client with your API key and provider (e.g., "openai", "gemini", or "grok").
    // Replace "YOUR_API_KEY" with a valid key from your chosen provider's website.
    client := chatapi.NewClient("YOUR_API_KEY", "openai")

    // Synchronous chat: send a message and get a response.
    response, err := client.Chat("Hello AI!")
    if err != nil {
        log.Fatalf("Chat error: %v", err)
    }
    fmt.Println("AI Response:", response)

    // Streaming chat: receive real-time responses.
    ch := make(chan string)
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    go func() {
        err := client.StreamChat(ctx, "Tell me a story", ch)
        if err != nil {
            log.Printf("StreamChat error: %v", err)
        }
    }()

    // Print streaming responses as they arrive.
    for msg := range ch {
        fmt.Println("Streaming Response:", msg)
    }
}
```

## Supported Providers
- **OpenAI**: GPT-3.5, GPT-4
- **Gemini**: 1.0
- **Grok**: xAI, latest as of March 2025

## Troubleshooting
- Ensure your API key is valid and matches the selected provider.
- Check your internet connection for streaming issues.
- For detailed logs, enable verbose logging (if supported by the package).

## Contributing
Pull requests are welcome! To contribute:
1. Fork the repository and clone it locally.
2. Run `go test ./...` to ensure tests pass.
3. Submit a pull request with a clear description of your changes.

For major changes, please open an issue first to discuss your ideas. 

See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.

## License
[MIT](https://choosealicense.com/licenses/mit/)
