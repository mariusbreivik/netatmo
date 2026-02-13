package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mariusbreivik/netatmo/api/netatmo"
	internalNetatmo "github.com/mariusbreivik/netatmo/internal/netatmo"
)

// noiseCmd represents the noise command
var noiseCmd = &cobra.Command{
	Use:     "noise",
	Short:   "Read noise data from Netatmo station",
	Long:    `Read noise level data from your Netatmo indoor weather station module.`,
	Example: "netatmo noise",
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

		printNoiseLevel(stationData)
		return nil
	},
}

func printNoiseLevel(stationData netatmo.StationData) {
	device := stationData.Body.Devices[0]
	fmt.Println("Station name:", device.StationName)
	fmt.Println("Noise level:", internalNetatmo.FormatNoise(device.DashboardData.Noise))
}

func init() {
	rootCmd.AddCommand(noiseCmd)
}
