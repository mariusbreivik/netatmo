package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "netatmo",
	Short: "read data from your personal netatmo weather station\n",
	Long: "Uses the Netatmo Weatherstation API to get your indoor/outdoor\n" +
		"temperature, co2 level, nois level, humidity, firmware data, wifi signal strength,\n" +
		"and more",
	Example: "netatmo temp --indoor",
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.netatmo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("debug", "d", false, "debug logging")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv()
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".netatmo" (without extension).
		viper.AddConfigPath(home)

		viper.SetConfigName(".netatmo")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
