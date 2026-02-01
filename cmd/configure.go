package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mariusbreivik/netatmo/internal/netatmo"
	"github.com/spf13/cobra"
)

var clientID string
var clientSecret string

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure Netatmo API credentials",
	Long: `Configure your Netatmo API credentials (client ID and client secret).

You can get these credentials from the Netatmo developer portal:
  1. Go to https://dev.netatmo.com/apps/
  2. Create a new app or select an existing one
  3. Copy the client ID and client secret`,
	Example: `  # Interactive mode (prompts for credentials)
  netatmo configure

  # Non-interactive mode (for scripting)
  netatmo configure --client-id YOUR_CLIENT_ID --client-secret YOUR_CLIENT_SECRET`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If credentials not provided via flags, prompt interactively
		if clientID == "" {
			fmt.Print("Enter client ID: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("failed to read client ID: %w", err)
			}
			clientID = strings.TrimSpace(input)
		}

		if clientSecret == "" {
			fmt.Print("Enter client secret: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("failed to read client secret: %w", err)
			}
			clientSecret = strings.TrimSpace(input)
		}

		// Validate credentials are not empty
		if clientID == "" {
			return fmt.Errorf("client ID cannot be empty")
		}
		if clientSecret == "" {
			return fmt.Errorf("client secret cannot be empty")
		}

		// Load existing config to preserve tokens if present
		config, err := netatmo.LoadConfig()
		if err != nil {
			// Config doesn't exist, create a new one
			config = &netatmo.AppConfig{}
		}

		// Update credentials
		config.ClientID = clientID
		config.ClientSecret = clientSecret

		// Save config
		if err := netatmo.SaveConfig(config); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		configPath, _ := netatmo.GetConfigPath()
		fmt.Println("✓ Credentials saved successfully!")
		fmt.Printf("  Config file: %s\n", configPath)

		if !config.HasTokens() {
			fmt.Println("\nNext step: Run 'netatmo login' to authenticate with your access tokens")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.Flags().StringVar(&clientID, "client-id", "", "Netatmo API client ID")
	configureCmd.Flags().StringVar(&clientSecret, "client-secret", "", "Netatmo API client secret")
}
