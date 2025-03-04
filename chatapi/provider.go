package chatapi

import (
	"fmt"
	"strings"
)

// ProviderConfig holds configuration for each provider
type ProviderConfig struct {
	BaseURL    string
	Model      string
	APIVersion string
}

// GetProviderConfig returns the configuration for the specified provider
func GetProviderConfig(provider string) (*ProviderConfig, error) {
	switch strings.ToLower(provider) {
	case "openai":
		return &ProviderConfig{
			BaseURL:    "https://api.openai.com/v1",
			Model:      "gpt-4",
			APIVersion: "2023-05-15",
		}, nil
	case "gemini":
		return &ProviderConfig{
			BaseURL:    "https://generativelanguage.googleapis.com/v1",
			Model:      "gemini-pro",
			APIVersion: "2023-08-01",
		}, nil
	case "grok":
		return &ProviderConfig{
			BaseURL:    "https://api.grok.ai/v1",
			Model:      "grok-1",
			APIVersion: "2023-10-01",
		}, nil
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}
}
