package chatapi

import (
	"fmt"
	"os"
	"sync"
)

// ConfigManager handles API key and settings management
type ConfigManager struct {
	apiKey string
	mu     sync.RWMutex
}

// NewConfigManager creates a new ConfigManager instance
func NewConfigManager() *ConfigManager {
	return &ConfigManager{}
}

// SetAPIKey sets the API key
func (cm *ConfigManager) SetAPIKey(key string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.apiKey = key
}

// GetAPIKey retrieves the API key
func (cm *ConfigManager) GetAPIKey() string {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.apiKey
}

// LoadAPIKeyFromEnv loads the API key from environment variables
func (cm *ConfigManager) LoadAPIKeyFromEnv(envVar string) error {
	key := os.Getenv(envVar)
	if key == "" {
		return fmt.Errorf("environment variable %s is not set", envVar)
	}
	cm.SetAPIKey(key)
	return nil
}
