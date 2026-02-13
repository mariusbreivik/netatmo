package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mariusbreivik/netatmo/api/netatmo"
)

// firmwareCmd represents the firmware command
var firmwareCmd = &cobra.Command{
	Use:     "firmware",
	Short:   "Read firmware data from Netatmo station",
	Long:    `Read firmware version information from your Netatmo weather station and all connected modules.`,
	Example: "netatmo firmware",
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

		printFirmwareInfo(stationData)
		return nil
	},
}

func printFirmwareInfo(stationData netatmo.StationData) {
	device := stationData.Body.Devices[0]

	fmt.Println("Station name:", device.StationName)
	fmt.Println()
	fmt.Printf("  📟 Base station (%s)\n", device.ModuleName)
	fmt.Printf("     Firmware: %d\n", device.Firmware)
	fmt.Printf("     Reachable: %v\n", device.Reachable)
	fmt.Println()

	if len(device.Modules) > 0 {
		for _, module := range device.Modules {
			fmt.Printf("  📡 Module (%s)\n", module.ModuleName)
			fmt.Printf("     Type: %s\n", module.Type)
			fmt.Printf("     Firmware: %d\n", module.Firmware)
			fmt.Printf("     Reachable: %v\n", module.Reachable)
			fmt.Println()
		}
	}
}

func init() {
	rootCmd.AddCommand(firmwareCmd)
}
