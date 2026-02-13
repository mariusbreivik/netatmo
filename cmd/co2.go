package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mariusbreivik/netatmo/api/netatmo"
	internalNetatmo "github.com/mariusbreivik/netatmo/internal/netatmo"
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
