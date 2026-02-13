package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mariusbreivik/netatmo/api/netatmo"
	internalNetatmo "github.com/mariusbreivik/netatmo/internal/netatmo"
)

// pressureCmd represents the pressure command
var pressureCmd = &cobra.Command{
	Use:     "pressure",
	Short:   "Read pressure data from Netatmo station",
	Long:    `Read atmospheric pressure data from your Netatmo weather station.`,
	Example: "netatmo pressure",
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

		printPressureLevel(stationData)
		return nil
	},
}

func printPressureLevel(stationData netatmo.StationData) {
	device := stationData.Body.Devices[0]
	dashboard := device.DashboardData
	fmt.Println("Station name:", device.StationName)
	fmt.Printf("Pressure: %.1f hPa %s\n", dashboard.Pressure, internalNetatmo.FormatTrend(dashboard.PressureTrend))
}

func init() {
	rootCmd.AddCommand(pressureCmd)
}
