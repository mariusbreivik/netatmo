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

// humidityCmd represents the humidity command
var humidityCmd = &cobra.Command{
	Use:     "humidity",
	Short:   "Read humidity data from Netatmo station",
	Long:    `Read indoor or outdoor humidity data from your Netatmo weather station.`,
	Example: "netatmo humidity --indoor\nnetatmo humidity --outdoor",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !indoor && !outdoor {
			return cmd.Help()
		}

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

		if indoor {
			printIndoorHumidity(stationData)
		} else if outdoor {
			if len(stationData.Body.Devices[0].Modules) == 0 {
				return fmt.Errorf("no outdoor module found")
			}
			printOutdoorHumidity(stationData)
		}

		return nil
	},
}

func printOutdoorHumidity(stationData netatmo.StationData) {
	device := stationData.Body.Devices[0]
	module := device.Modules[0]
	fmt.Println("Station name:", device.StationName)
	fmt.Println("Humidity outdoor:", internalNetatmo.FormatHumidity(module.DashboardData.Humidity))
}

func printIndoorHumidity(stationData netatmo.StationData) {
	device := stationData.Body.Devices[0]
	fmt.Println("Station name:", device.StationName)
	fmt.Println("Humidity indoor:", internalNetatmo.FormatHumidity(device.DashboardData.Humidity))
}

func init() {
	rootCmd.AddCommand(humidityCmd)

	humidityCmd.Flags().BoolVarP(&indoor, "indoor", "i", false, "Show indoor humidity")
	humidityCmd.Flags().BoolVarP(&outdoor, "outdoor", "o", false, "Show outdoor humidity")
}
