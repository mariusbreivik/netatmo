package cmd

import (
	"fmt"

	"github.com/mariusbreivik/netatmo/api/netatmo"
	internalNetatmo "github.com/mariusbreivik/netatmo/internal/netatmo"
	"github.com/spf13/cobra"
)

var indoor bool
var outdoor bool

// tempCmd is the command for retrieving temperature
var tempCmd = &cobra.Command{
	Use:     "temp",
	Short:   "Read temperature data from Netatmo station",
	Long:    `Read indoor or outdoor temperature data from your Netatmo weather station.`,
	Example: "netatmo temp --indoor\nnetatmo temp --outdoor",
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
			printIndoorTemp(stationData)
		} else if outdoor {
			if len(stationData.Body.Devices[0].Modules) == 0 {
				return fmt.Errorf("no outdoor module found")
			}
			printOutdoorTemp(stationData)
		}

		return nil
	},
}

func printOutdoorTemp(stationData netatmo.StationData) {
	device := stationData.Body.Devices[0]
	module := device.Modules[0]
	fmt.Println("Station name:", device.StationName)
	fmt.Println("Temperature outdoor:", internalNetatmo.FormatTemperature(module.DashboardData.Temperature))
}

func printIndoorTemp(stationData netatmo.StationData) {
	device := stationData.Body.Devices[0]
	fmt.Println("Station name:", device.StationName)
	fmt.Println("Temperature indoor:", internalNetatmo.FormatTemperature(device.DashboardData.Temperature))
}

func init() {
	rootCmd.AddCommand(tempCmd)

	tempCmd.Flags().BoolVarP(&indoor, "indoor", "i", false, "Show indoor temperature")
	tempCmd.Flags().BoolVarP(&outdoor, "outdoor", "o", false, "Show outdoor temperature")
}
