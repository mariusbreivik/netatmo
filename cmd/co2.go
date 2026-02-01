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

// co2Cmd represents the co2 command
var co2Cmd = &cobra.Command{
	Use:     "co2",
	Short:   "read co2 data from netatmo station",
	Long:    `read co2 data from netatmo station`,
	Example: "netatmo co2",
	RunE: func(cmd *cobra.Command, args []string) error {
		netatmoClient, err := netatmo.NewClient(netatmo.Config{
			ClientID:     viper.GetString("netatmo.clientID"),
			ClientSecret: viper.GetString("netatmo.clientSecret"),
		})

		if err != nil {
			return err
		}

		if len(args) > 0 {
			fmt.Println(cmd.UsageString())
		}

		printCo2Level(netatmoClient.GetStationData())

		return nil
	},
}

func printCo2Level(stationData netatmo2.StationData) {
	fmt.Println("Station name: ", stationData.Body.Devices[0].StationName)
	fmt.Println("Co2:", chalk.Green, stationData.Body.Devices[0].DashboardData.CO2, "ppm", chalk.Reset)

}

func init() {
	rootCmd.AddCommand(co2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// co2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// co2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
