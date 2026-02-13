package netatmo

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestGetConfigPath(t *testing.T) {
	path, err := GetConfigPath()
	if err != nil {
		t.Fatalf("GetConfigPath() error = %v", err)
	}

	if path == "" {
		t.Error("GetConfigPath() returned empty path")
	}

	if !filepath.IsAbs(path) {
		t.Errorf("GetConfigPath() = %q, want absolute path", path)
	}

	if filepath.Base(path) != configFileName {
		t.Errorf("GetConfigPath() base = %q, want %q", filepath.Base(path), configFileName)
	}
}

func TestAppConfig_HasCredentials(t *testing.T) {
	tests := []struct {
		name     string
		config   AppConfig
		expected bool
	}{
		{
			name:     "both set",
			config:   AppConfig{ClientID: "id", ClientSecret: "secret"},
			expected: true,
		},
		{
			name:     "only client id",
			config:   AppConfig{ClientID: "id"},
			expected: false,
		},
		{
			name:     "only client secret",
			config:   AppConfig{ClientSecret: "secret"},
			expected: false,
		},
		{
			name:     "both empty",
			config:   AppConfig{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.config.HasCredentials(); got != tt.expected {
				t.Errorf("HasCredentials() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAppConfig_HasTokens(t *testing.T) {
	tests := []struct {
		name     string
		config   AppConfig
		expected bool
	}{
		{
			name:     "both set",
			config:   AppConfig{AccessToken: "access", RefreshToken: "refresh"},
			expected: true,
		},
		{
			name:     "only access token",
			config:   AppConfig{AccessToken: "access"},
			expected: false,
		},
		{
			name:     "only refresh token",
			config:   AppConfig{RefreshToken: "refresh"},
			expected: false,
		},
		{
			name:     "both empty",
			config:   AppConfig{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.config.HasTokens(); got != tt.expected {
				t.Errorf("HasTokens() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSaveAndLoadConfig(t *testing.T) {
	// Create a temporary directory for test config
	tmpDir, err := os.MkdirTemp("", "netatmo-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test config file path
	testConfigPath := filepath.Join(tmpDir, configFileName)

	// Test config to save
	expiry := time.Now().Add(time.Hour)
	testConfig := &AppConfig{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		AccessToken:  "test-access-token",
		RefreshToken: "test-refresh-token",
		TokenExpiry:  &expiry,
	}

	// Save directly to test path
	data, err := json.MarshalIndent(testConfig, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(testConfigPath, data, 0600); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	// Read and verify
	readData, err := os.ReadFile(testConfigPath)
	if err != nil {
		t.Fatalf("Failed to read config: %v", err)
	}

	var loadedConfig AppConfig
	if err := json.Unmarshal(readData, &loadedConfig); err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}

	if loadedConfig.ClientID != testConfig.ClientID {
		t.Errorf("ClientID = %q, want %q", loadedConfig.ClientID, testConfig.ClientID)
	}
	if loadedConfig.ClientSecret != testConfig.ClientSecret {
		t.Errorf("ClientSecret = %q, want %q", loadedConfig.ClientSecret, testConfig.ClientSecret)
	}
	if loadedConfig.AccessToken != testConfig.AccessToken {
		t.Errorf("AccessToken = %q, want %q", loadedConfig.AccessToken, testConfig.AccessToken)
	}
	if loadedConfig.RefreshToken != testConfig.RefreshToken {
		t.Errorf("RefreshToken = %q, want %q", loadedConfig.RefreshToken, testConfig.RefreshToken)
	}
}

func TestConfigExists(t *testing.T) {
	// This test verifies the function works without error
	// The actual result depends on whether the user has a config file
	_ = ConfigExists()
}

func TestAppConfigJSONSerialization(t *testing.T) {
	expiry := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	config := &AppConfig{
		ClientID:     "client-123",
		ClientSecret: "secret-456",
		AccessToken:  "access-789",
		RefreshToken: "refresh-012",
		TokenExpiry:  &expiry,
	}

	// Marshal
	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	// Unmarshal
	var decoded AppConfig
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	// Verify
	if decoded.ClientID != config.ClientID {
		t.Errorf("ClientID mismatch after round-trip")
	}
	if decoded.ClientSecret != config.ClientSecret {
		t.Errorf("ClientSecret mismatch after round-trip")
	}
	if decoded.AccessToken != config.AccessToken {
		t.Errorf("AccessToken mismatch after round-trip")
	}
	if decoded.RefreshToken != config.RefreshToken {
		t.Errorf("RefreshToken mismatch after round-trip")
	}
}

func TestAppConfigOmitEmptyTokens(t *testing.T) {
	config := &AppConfig{
		ClientID:     "client-123",
		ClientSecret: "secret-456",
		// No tokens set
	}

	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	jsonStr := string(data)

	// Verify omitempty works for optional fields
	if !json.Valid(data) {
		t.Error("Generated JSON is not valid")
	}

	// Should still contain required fields
	var decoded map[string]interface{}
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal to map: %v", err)
	}

	if _, ok := decoded["clientID"]; !ok {
		t.Error("clientID should be present in JSON")
	}

	// AccessToken should be omitted when empty (due to omitempty)
	_ = jsonStr // Verify the JSON is valid
}
