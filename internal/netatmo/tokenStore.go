package netatmo

import (
	"time"

	"golang.org/x/oauth2"
)

// SaveTokenFromStrings saves access and refresh tokens to the config file
func SaveTokenFromStrings(accessToken, refreshToken string) error {
	// Load existing config to preserve credentials
	config, err := LoadConfig()
	if err != nil {
		// If config doesn't exist, create a new one
		config = &AppConfig{}
	}

	// Set expiry to past so token refresh is triggered on first use
	expiry := time.Now().Add(-1 * time.Hour)

	config.AccessToken = accessToken
	config.RefreshToken = refreshToken
	config.TokenExpiry = &expiry

	return SaveConfig(config)
}

// saveToken saves the OAuth2 token to the config file
func saveToken(token *oauth2.Token) error {
	// Load existing config to preserve credentials
	config, err := LoadConfig()
	if err != nil {
		// If config doesn't exist, create a new one
		config = &AppConfig{}
	}

	config.AccessToken = token.AccessToken
	config.RefreshToken = token.RefreshToken
	config.TokenExpiry = &token.Expiry

	return SaveConfig(config)
}

// loadToken loads the OAuth2 token from the config file
func loadToken() (*oauth2.Token, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	if !config.HasTokens() {
		return nil, nil
	}

	token := &oauth2.Token{
		AccessToken:  config.AccessToken,
		RefreshToken: config.RefreshToken,
		TokenType:    "Bearer",
	}

	if config.TokenExpiry != nil {
		token.Expiry = *config.TokenExpiry
	}

	return token, nil
}

// tokenExists checks if tokens exist in the config
func tokenExists() bool {
	config, err := LoadConfig()
	if err != nil {
		return false
	}
	return config.HasTokens()
}
