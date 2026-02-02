package netatmo

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/mariusbreivik/netatmo/api/netatmo"
	"golang.org/x/oauth2"
)

const (
	baseURL        = "https://api.netatmo.com/"
	tokenURL       = baseURL + "oauth2/token"
	deviceURL      = baseURL + "api/getstationsdata"
	defaultTimeout = 30 * time.Second
)

// Client use to make request to Netatmo API
type Client struct {
	oauth      *oauth2.Config
	httpClient *http.Client
	ctx        context.Context
}

// NewClient creates a handle for authentication to Netatmo API using stored config
func NewClient() (*Client, error) {
	return NewClientWithContext(context.Background())
}

// NewClientWithContext creates a client with a custom context for timeout/cancellation
func NewClientWithContext(ctx context.Context) (*Client, error) {
	Debug("initializing Netatmo client")

	// Load config
	config, err := LoadConfig()
	if err != nil {
		Debug("failed to load config", "error", err)
		return nil, err
	}

	// Validate credentials
	if !config.HasCredentials() {
		return nil, NewAuthError(
			"credentials not configured",
			"Run 'netatmo configure' to set up your API credentials",
			ErrNotConfigured,
		)
	}

	// Validate tokens
	if !config.HasTokens() {
		return nil, NewAuthError(
			"not authenticated",
			"Run 'netatmo login' to authenticate with Netatmo",
			ErrNotAuthenticated,
		)
	}

	oauthConfig := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"read_station"},
		Endpoint: oauth2.Endpoint{
			TokenURL: tokenURL,
		},
	}

	// Load token
	token, err := loadToken()
	if err != nil {
		Debug("failed to load token", "error", err)
		return nil, NewAuthError(
			"failed to load authentication token",
			"Run 'netatmo login' to re-authenticate",
			err,
		)
	}

	if token == nil {
		return nil, NewAuthError(
			"not authenticated",
			"Run 'netatmo login' to authenticate with Netatmo",
			ErrNotAuthenticated,
		)
	}

	// Token is valid, use it directly
	if token.Valid() {
		Debug("using valid token")
		return &Client{
			oauth:      oauthConfig,
			httpClient: oauthConfig.Client(ctx, token),
			ctx:        ctx,
		}, nil
	}

	// Token expired but we have a refresh token - try to refresh
	if token.RefreshToken != "" {
		Debug("token expired, attempting refresh")
		tokenSource := oauthConfig.TokenSource(ctx, token)
		newToken, err := tokenSource.Token()
		if err != nil {
			Debug("token refresh failed", "error", err)
			return nil, NewAuthError(
				"token refresh failed",
				"Run 'netatmo login' to re-authenticate",
				err,
			)
		}

		// Save the refreshed token
		if saveErr := saveToken(newToken); saveErr != nil {
			Warn("could not save refreshed token", "error", saveErr)
		} else {
			Debug("refreshed token saved successfully")
		}

		return &Client{
			oauth:      oauthConfig,
			httpClient: oauthConfig.Client(ctx, newToken),
			ctx:        ctx,
		}, nil
	}

	return nil, NewAuthError(
		"token expired and no refresh token available",
		"Run 'netatmo login' to re-authenticate",
		ErrTokenExpired,
	)
}

// doHTTP executes an HTTP request with timeout
func (c *Client) doHTTP(req *http.Request) (*http.Response, error) {
	// Add timeout context
	ctx, cancel := context.WithTimeout(c.ctx, defaultTimeout)
	defer cancel()
	req = req.WithContext(ctx)

	Debug("executing HTTP request", "method", req.Method, "url", req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		Debug("HTTP request failed", "error", err)
		return nil, NewNetworkError("failed to connect to Netatmo API", err)
	}

	Debug("HTTP response received", "status", resp.StatusCode)
	return resp, nil
}

// GetStationData returns data from netatmo api
func (c *Client) GetStationData() (netatmo.StationData, error) {
	Debug("fetching station data")

	resp, err := c.doHTTPGet(deviceURL, url.Values{"app_type": {"app_station"}})
	if err != nil {
		return netatmo.StationData{}, err
	}

	return processHTTPResponse(resp)
}

// doHTTPGet sends an HTTP GET request
func (c *Client) doHTTPGet(urlStr string, data url.Values) (*http.Response, error) {
	if data != nil {
		urlStr = urlStr + "?" + data.Encode()
	}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, NewNetworkError("failed to create request", err)
	}

	return c.doHTTP(req)
}

// processHTTPResponse parses the API response and returns station data
func processHTTPResponse(resp *http.Response) (netatmo.StationData, error) {
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		Debug("API error response", "status", resp.StatusCode, "body", string(body))
		return netatmo.StationData{}, NewAPIError(resp.StatusCode, string(body), ErrAPIError)
	}

	var stationData netatmo.StationData
	if err := json.NewDecoder(resp.Body).Decode(&stationData); err != nil {
		Debug("failed to decode response", "error", err)
		return netatmo.StationData{}, NewAPIError(resp.StatusCode, "failed to parse API response", err)
	}

	Debug("station data received", "devices", len(stationData.Body.Devices))
	return stationData, nil
}
