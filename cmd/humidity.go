package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mariusbreivik/netatmo/api/netatmo"
	internalNetatmo "github.com/mariusbreivik/netatmo/internal/netatmo"
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
