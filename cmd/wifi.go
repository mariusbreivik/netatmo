package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mariusbreivik/netatmo/api/netatmo"
	internalNetatmo "github.com/mariusbreivik/netatmo/internal/netatmo"
)

// wifiCmd represents the wifi command
var wifiCmd = &cobra.Command{
	Use:     "wifi",
	Short:   "Read WiFi signal data from Netatmo station",
	Long:    `Read WiFi signal strength and RF status from your Netatmo weather station and connected modules.`,
	Example: "netatmo wifi",
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

		printWifiInfo(stationData)
		return nil
	},
}

func printWifiInfo(stationData netatmo.StationData) {
	device := stationData.Body.Devices[0]

	fmt.Println("Station name:", device.StationName)
	fmt.Println()
	fmt.Printf("  📟 Base station (%s)\n", device.ModuleName)
	fmt.Printf("     WiFi Signal: %s (raw: %d dB)\n", internalNetatmo.FormatWifiSignal(device.WifiStatus), device.WifiStatus)
	fmt.Println()

	if len(device.Modules) > 0 {
		for _, module := range device.Modules {
			fmt.Printf("  📡 Module (%s)\n", module.ModuleName)
			fmt.Printf("     RF Signal: %s (raw: %d)\n", formatRFSignal(module.RfStatus), module.RfStatus)
			fmt.Printf("     Battery: %s\n", internalNetatmo.FormatBattery(module.BatteryPercent))
			fmt.Println()
		}
	}
}

// formatRFSignal converts RF status to human-readable string
// Netatmo RF status: 90=low, 80=medium, 70=high, 60=full (lower is better)
func formatRFSignal(status int) string {
	switch {
	case status >= 90:
		return "Low 📡"
	case status >= 80:
		return "Medium 📡"
	case status >= 70:
		return "High 📡"
	default:
		return "Full 📡"
	}
}

func init() {
	rootCmd.AddCommand(wifiCmd)
}
