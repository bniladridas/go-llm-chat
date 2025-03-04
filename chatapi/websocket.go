package chatapi

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

// StreamChat streams responses from the selected LLM in real-time
func (c *Client) StreamChat(ctx context.Context, message string, ch chan<- string) error {
	if c.apiKey == "" {
		return errors.New("API key is required")
	}

	if message == "" {
		return errors.New("message cannot be empty")
	}

	// Get provider-specific WebSocket URL
	wsURL, err := c.getWebSocketURL()
	if err != nil {
		return fmt.Errorf("failed to get WebSocket URL: %w", err)
	}

	// Connect to WebSocket
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}
	defer conn.Close()

	// Send message
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	// Read responses
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return fmt.Errorf("failed to read message: %w", err)
			}
			ch <- string(msg)
		}
	}
}

// getWebSocketURL returns the WebSocket URL for the selected provider
func (c *Client) getWebSocketURL() (string, error) {
	switch strings.ToLower(c.provider) {
	case "openai":
		return "wss://api.openai.com/v1/chat/completions", nil
	case "gemini":
		return "wss://generativelanguage.googleapis.com/v1/stream", nil
	case "grok":
		return "wss://api.grok.ai/v1/stream", nil
	default:
		return "", fmt.Errorf("unsupported provider: %s", c.provider)
	}
}
