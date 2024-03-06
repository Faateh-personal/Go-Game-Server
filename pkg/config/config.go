package config

import (
	"encoding/json"
	"os"
)

// Config represents the configuration for the game server.
type Config struct {
	Port        string `json:"port"`
	DatabaseURL string `json:"database_url"`
}

// LoadConfig loads the configuration from a JSON file.
func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
