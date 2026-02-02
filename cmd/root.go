package cmd

import (
	"fmt"
	"os"

	"github.com/mariusbreivik/netatmo/internal/netatmo"
	"github.com/spf13/cobra"
)

var sharedClient *netatmo.Client

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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "netatmo",
	Short: "read data from your personal netatmo weather station\n",
	Long: "Uses the Netatmo Weatherstation API to get your indoor/outdoor\n" +
		"temperature, co2 level, pressure level, nois level, humidity, firmware data, wifi signal strength,\n" +
		"and more",
	Example: "netatmo temp --indoor",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("debug", "d", false, "debug logging")
}
