package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mariusbreivik/netatmo/internal/netatmo"
	"github.com/spf13/cobra"
)

var accessToken string
var refreshToken string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Store Netatmo API tokens for authentication",
	Long: `Store your Netatmo access token and refresh token for API authentication.

You can get these tokens from the Netatmo developer portal:
  1. Go to https://dev.netatmo.com/apps/
  2. Select your app
  3. Scroll down to "Token generator"
  4. Select scope "read_station" and click "Generate Token"
  5. Copy the access token and refresh token`,
	Example: `  # Interactive mode (prompts for tokens)
  netatmo login

  # Non-interactive mode (for scripting)
  netatmo login --access-token YOUR_ACCESS_TOKEN --refresh-token YOUR_REFRESH_TOKEN`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If tokens not provided via flags, prompt interactively
		if accessToken == "" {
			fmt.Print("Enter access token: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("failed to read access token: %w", err)
			}
			accessToken = strings.TrimSpace(input)
		}

		if refreshToken == "" {
			fmt.Print("Enter refresh token: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("failed to read refresh token: %w", err)
			}
			refreshToken = strings.TrimSpace(input)
		}

		// Validate tokens are not empty
		if accessToken == "" {
			return fmt.Errorf("access token cannot be empty")
		}
		if refreshToken == "" {
			return fmt.Errorf("refresh token cannot be empty")
		}

		// Save tokens
		if err := netatmo.SaveTokenFromStrings(accessToken, refreshToken); err != nil {
			return fmt.Errorf("failed to save tokens: %w", err)
		}

		tokenPath, _ := netatmo.GetTokenPath()
		fmt.Println("✓ Tokens saved successfully!")
		fmt.Printf("  Token file: %s\n", tokenPath)
		fmt.Println("\nYou can now use netatmo commands like 'netatmo temp --indoor'")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVar(&accessToken, "access-token", "", "Netatmo API access token")
	loginCmd.Flags().StringVar(&refreshToken, "refresh-token", "", "Netatmo API refresh token")
}
