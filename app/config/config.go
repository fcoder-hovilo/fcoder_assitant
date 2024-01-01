// config/config.go

package config

import (
	"encoding/json"
	"os"
)

// Config represents the bot configuration
type Config struct {
	Prefix string `json:"Prefix"`
	Token  string `json:"Token"`
}

// Load loads the configuration from the specified file
func Load() (Config, error) {
	var config Config
	configFile, err := os.ReadFile("config/config.json")
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
