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

	"github.com/mariusbreivik/netatmo/api/netatmo"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

// humidityCmd represents the humidity command
var humidityCmd = &cobra.Command{
	Use:     "humidity",
	Short:   "read humidity data from netatmo station",
	Long:    `read humidity data from netatmo station`,
	Example: "netatmo humidity",
	RunE: func(cmd *cobra.Command, args []string) error {
		netatmoClient, err := getClient()

		if err != nil {
			return err
		}

		if indoor {
			printIndoorHumidity(netatmoClient.GetStationData())
		} else if outdoor {
			printOutdoorHumidity(netatmoClient.GetStationData())
		} else {
			fmt.Println(cmd.UsageString())
		}

		return nil
	},
}

func printOutdoorHumidity(stationData netatmo.StationData) {
	fmt.Println("Station name: ", stationData.Body.Devices[0].StationName)
	fmt.Println("Humidity outdoor:", chalk.Blue, stationData.Body.Devices[0].Modules[0].DashboardData.Humidity, chalk.Reset)

}

func printIndoorHumidity(stationData netatmo.StationData) {
	fmt.Println("Station name: ", stationData.Body.Devices[0].StationName)
	fmt.Println("Humidity indoor:", chalk.Red, stationData.Body.Devices[0].DashboardData.Humidity, chalk.Reset)
}

func init() {
	rootCmd.AddCommand(humidityCmd)

	humidityCmd.Flags().BoolVarP(&indoor, "indoor", "i", false, "netatmo humidity -i|--indoor")
	humidityCmd.Flags().BoolVarP(&outdoor, "outdoor", "o", false, "netatmo humidity -o|--outdoor")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// humidityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// humidityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
