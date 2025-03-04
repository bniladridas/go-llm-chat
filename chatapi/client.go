package chatapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Client represents the chat client
type Client struct {
	apiKey     string
	provider   string
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new chat client
func NewClient(apiKey, provider string) *Client {
	return &Client{
		apiKey:     apiKey,
		provider:   provider,
		httpClient: &http.Client{},
	}
}

// Chat sends a message to the selected LLM and returns the response
func (c *Client) Chat(message string) (string, error) {
	if c.apiKey == "" {
		return "", errors.New("API key is required")
	}

	if message == "" {
		return "", errors.New("message cannot be empty")
	}

	// Implement provider-specific API calls
	switch strings.ToLower(c.provider) {
	case "openai":
		return c.chatWithOpenAI(message)
	case "gemini":
		return c.chatWithGemini(message)
	case "grok":
		return c.chatWithGrok(message)
	default:
		return "", fmt.Errorf("unsupported provider: %s", c.provider)
	}
}

// chatWithOpenAI handles OpenAI API requests
func (c *Client) chatWithOpenAI(message string) (string, error) {
	// Prepare the request payload
	payload := map[string]interface{}{
		"model": "gpt-4",
		"messages": []map[string]string{
			{"role": "user", "content": message},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	// Send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	// Parse the response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	// Extract the response message
	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", errors.New("invalid response format")
	}

	firstChoice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", errors.New("invalid response format")
	}

	messageContent, ok := firstChoice["message"].(map[string]interface{})
	if !ok {
		return "", errors.New("invalid response format")
	}

	content, ok := messageContent["content"].(string)
	if !ok {
		return "", errors.New("invalid response format")
	}

	return content, nil
}

// chatWithGemini handles Gemini API requests
func (c *Client) chatWithGemini(message string) (string, error) {
	// TODO: Implement Gemini API integration
	return "", errors.New("Gemini integration not yet implemented")
}

// chatWithGrok handles Grok API requests
func (c *Client) chatWithGrok(message string) (string, error) {
	// TODO: Implement Grok API integration
	return "", errors.New("Grok integration not yet implemented")
}
