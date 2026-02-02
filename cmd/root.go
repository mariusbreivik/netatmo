package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	netatmoAPI "github.com/mariusbreivik/netatmo/api/netatmo"
	"github.com/mariusbreivik/netatmo/internal/netatmo"
	"github.com/spf13/cobra"
)

// Global flags
var (
	verbose    bool
	jsonOutput bool
	noColor    bool
)

// Shared client instance
var sharedClient *netatmo.Client

// getClient returns a shared Netatmo client, creating it if necessary
func getClient() (*netatmo.Client, error) {
	if sharedClient != nil {
		return sharedClient, nil
	}
	client, err := netatmo.NewClient()
	if err != nil {
		return nil, err
	}
	sharedClient = client
	return sharedClient, nil
}

// validateStationData checks if station data contains valid device information
func validateStationData(data netatmoAPI.StationData) error {
	if len(data.Body.Devices) == 0 {
		return fmt.Errorf("no devices found. Make sure your Netatmo weather station is set up correctly")
	}
	return nil
}

// handleError formats and prints errors appropriately
func handleError(err error) {
	if jsonOutput {
		errJSON, _ := json.Marshal(map[string]string{"error": err.Error()})
		fmt.Fprintln(os.Stderr, string(errJSON))
	} else {
		fmt.Fprintln(os.Stderr, "Error:", netatmo.FormatError(err))
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "netatmo",
	Short: "Read data from your personal Netatmo weather station",
	Long: `A command-line interface for the Netatmo Weather Station API.

Retrieve indoor/outdoor temperature, CO2 levels, pressure, noise levels,
humidity, firmware data, WiFi signal strength, and more from your 
Netatmo weather station.`,
	Example: `  netatmo status              # Full dashboard view
  netatmo temp --indoor       # Indoor temperature
  netatmo temp --outdoor      # Outdoor temperature  
  netatmo co2                 # CO2 level
  netatmo humidity --indoor   # Indoor humidity`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Initialize logger based on verbose flag
		netatmo.InitLogger(verbose)

		// Set color mode
		if noColor || os.Getenv("NO_COLOR") != "" {
			netatmo.SetColorEnabled(false)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handleError(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose/debug output")
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "Output in JSON format")
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "Disable colored output")
}
