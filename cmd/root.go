package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "netatmo",
	Short: "read data from your personal netatmo weather station\n",
	Long: "Uses the Netatmo Weatherstation API to get your indoor/outdoor\n" +
		"temperature, co2 level, pressure level, nois level, humidity, firmware data, wifi signal strength,\n" +
		"and more",
	Example: "netatmo temp --indoor",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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
