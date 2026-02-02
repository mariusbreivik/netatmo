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
	internalNetatmo "github.com/mariusbreivik/netatmo/internal/netatmo"
	"github.com/spf13/cobra"
)

// co2Cmd represents the co2 command
var co2Cmd = &cobra.Command{
	Use:     "co2",
	Short:   "Read CO2 data from Netatmo station",
	Long:    `Read CO2 level data from your Netatmo indoor weather station module.`,
	Example: "netatmo co2",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := getClient()
		if err != nil {
			return err
		}

		stationData, err := client.GetStationData()
		if err != nil {
			return err
		}

		if err := validateStationData(stationData); err != nil {
			return err
		}

		printCo2Level(stationData)
		return nil
	},
}

func printCo2Level(stationData netatmo.StationData) {
	device := stationData.Body.Devices[0]
	fmt.Println("Station name:", device.StationName)
	fmt.Println("CO2:", internalNetatmo.FormatCO2(device.DashboardData.CO2))
}

func init() {
	rootCmd.AddCommand(co2Cmd)
}
