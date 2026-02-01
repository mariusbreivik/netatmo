package netatmo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mariusbreivik/netatmo/api/netatmo"
	"golang.org/x/oauth2"
)

const (
	baseURL   = "https://api.netatmo.com/"
	tokenURL  = baseURL + "oauth2/token"
	deviceURL = baseURL + "api/getstationsdata"
)

// Config used to create a netatmo client
type Config struct {
	ClientID     string
	ClientSecret string
}

// Client use to make request to Netatmo API
type Client struct {
	oauth        *oauth2.Config
	httpClient   *http.Client
	httpResponse *http.Response
}

// NewClient creates a handle for authentication to Netatmo API using stored OAuth2 tokens
func NewClient(config Config) (*Client, error) {
	ctx := context.Background()

	oauthConfig := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"read_station"},
		Endpoint: oauth2.Endpoint{
			TokenURL: tokenURL,
		},
	}

	// Check if token exists
	if !tokenExists() {
		return nil, fmt.Errorf("not authenticated. Please run 'netatmo login' first")
	}

	// Load existing token
	token, err := loadToken()
	if err != nil {
		return nil, fmt.Errorf("failed to load token: %w. Please run 'netatmo login' again", err)
	}

	// Token is valid, use it directly
	if token.Valid() {
		return &Client{
			oauth:      oauthConfig,
			httpClient: oauthConfig.Client(ctx, token),
		}, nil
	}

	// Token expired but we have a refresh token - try to refresh
	if token.RefreshToken != "" {
		tokenSource := oauthConfig.TokenSource(ctx, token)
		newToken, err := tokenSource.Token()
		if err != nil {
			return nil, fmt.Errorf("token refresh failed: %w. Please run 'netatmo login' again", err)
		}

		// Save the refreshed token
		if saveErr := saveToken(newToken); saveErr != nil {
			fmt.Printf("Warning: could not save refreshed token: %v\n", saveErr)
		}

		return &Client{
			oauth:      oauthConfig,
			httpClient: oauthConfig.Client(ctx, newToken),
		}, nil
	}

	return nil, fmt.Errorf("token expired and no refresh token available. Please run 'netatmo login' again")
}

// do a generic HTTP request
func (c *Client) doHTTP(req *http.Request) (*http.Response, error) {

	var err error
	c.httpResponse, err = c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return c.httpResponse, nil
}

// GetStationData returns data from netatmo api
func (c *Client) GetStationData() netatmo.StationData {
	resp, err := c.doHTTPGet(deviceURL, url.Values{"app_type": {"app_station"}})

	if err != nil {
		fmt.Println(err)
	}
	stationData := processHTTPResponse(resp, err)

	return stationData
}

// send http GET request
func (c *Client) doHTTPGet(url string, data url.Values) (*http.Response, error) {
	if data != nil {
		url = url + "?" + data.Encode()
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return c.doHTTP(req)
}

// process HTTP response
func processHTTPResponse(resp *http.Response, err error) netatmo.StationData {
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("An error occured %s", err)
	}

	// debug
	//debug, _ := httputil.DumpResponse(resp, true)
	//fmt.Printf("%s\n\n", debug)

	// check http return code
	if resp.StatusCode != 200 {
		fmt.Printf("Bad HTTP return code %d", resp.StatusCode)
	}

	var devices netatmo.StationData
	err = json.NewDecoder(resp.Body).Decode(&devices)
	if err != nil {
		fmt.Println(err)
	}

	return devices
}
