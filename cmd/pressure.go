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
	"math"

	netatmo2 "github.com/mariusbreivik/netatmo/api/netatmo"
	"github.com/mariusbreivik/netatmo/internal/netatmo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// pressure2Cmd represents the pressure command
var pressureCmd = &cobra.Command{
	Use:     "pressure",
	Short:   "read pressure data from netatmo station",
	Long:    `read pressure data from netatmo station`,
	Example: "netatmo pressure",
	RunE: func(cmd *cobra.Command, args []string) error {
		netatmoClient, err := netatmo.NewClient(netatmo.Config{
			ClientID:     viper.GetString("netatmo.clientID"),
			ClientSecret: viper.GetString("netatmo.clientSecret"),
			Username:     viper.GetString("netatmo.username"),
			Password:     viper.GetString("netatmo.password"),
		})

		if err != nil {
			return err
		}

		if len(args) > 0 {
			fmt.Println(cmd.UsageString())
		}

		printPressureLevel(netatmoClient.GetStationData())

		return nil
	},
}

func printPressureLevel(stationData netatmo2.StationData) {
	fmt.Println("Station name: ", stationData.Body.Devices[0].StationName)
	fmt.Println("Pressure:", chalk.Green, math.Round(stationData.Body.Devices[0].DashboardData.AbsolutePressure / 1000 * 760), "mm", chalk.Reset)

}

func init() {
	rootCmd.AddCommand(pressureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pressureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pressureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
