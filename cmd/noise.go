/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	netatmo2 "github.com/mariusbreivik/netatmo/api/netatmo"
	"github.com/mariusbreivik/netatmo/internal/netatmo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// noiseCmd represents the noise command
var noiseCmd = &cobra.Command{
	Use:     "noise",
	Short:   "read noise data from netatmo station",
	Long:    `read noise data from netatmo station`,
	Example: "netatmo noise",
	RunE: func(cmd *cobra.Command, args []string) error {
		netatmoClient, err := netatmo.NewClient(netatmo.Config{
			ClientID:     viper.GetString("netatmo.clientID"),
			ClientSecret: viper.GetString("netatmo.clientSecret"),
			Username:     viper.GetString("netatmo.username"),
			Password:     viper.GetString("netatmo.password"),
		})

		if len(args) > 0 {
			fmt.Println(cmd.UsageString())
		}
		if err != nil {
			return err
		}

		printNoiseLevel(netatmoClient.GetStationData())
		return nil
	},
}

func printNoiseLevel(stationData netatmo2.StationData) {
	fmt.Println("Station name: ", stationData.Body.Devices[0].StationName)
	fmt.Println("Noise level:", chalk.Green, stationData.Body.Devices[0].DashboardData.Noise, "dB", chalk.Reset)

}

func init() {
	rootCmd.AddCommand(noiseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// noiseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// noiseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
