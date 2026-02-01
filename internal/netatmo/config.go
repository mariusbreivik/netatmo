package netatmo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const configFileName = ".netatmo-config.json"

// AppConfig holds all configuration including credentials and tokens
type AppConfig struct {
	ClientID     string     `json:"clientID"`
	ClientSecret string     `json:"clientSecret"`
	AccessToken  string     `json:"accessToken,omitempty"`
	RefreshToken string     `json:"refreshToken,omitempty"`
	TokenExpiry  *time.Time `json:"tokenExpiry,omitempty"`
}

// GetConfigPath returns the path to the config file in the user's home directory
func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

// LoadConfig loads the configuration from the config file
func LoadConfig() (*AppConfig, error) {
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("config file not found. Please run 'netatmo configure' first")
		}
		return nil, err
	}

	var config AppConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

// SaveConfig saves the configuration to the config file
func SaveConfig(config *AppConfig) error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

// ConfigExists checks if the config file exists
func ConfigExists() bool {
	path, err := GetConfigPath()
	if err != nil {
		return false
	}
	_, err = os.Stat(path)
	return err == nil
}

// HasCredentials checks if the config has clientID and clientSecret set
func (c *AppConfig) HasCredentials() bool {
	return c.ClientID != "" && c.ClientSecret != ""
}

// HasTokens checks if the config has tokens set
func (c *AppConfig) HasTokens() bool {
	return c.AccessToken != "" && c.RefreshToken != ""
}
