package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"

	"github.com/mariusbreivik/netatmo/internal/netatmo"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Display a full dashboard of your weather station",
	Long:    `Display a comprehensive overview of all your Netatmo weather station data in one view.`,
	Example: "netatmo status",
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

		device := stationData.Body.Devices[0]
		dashboard := device.DashboardData

		// Header
		fmt.Println()
		fmt.Printf("%s🏠 %s%s\n", chalk.Bold, device.StationName, chalk.Reset)
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()

		// Indoor section
		fmt.Printf("  %s📍 Indoor (%s)%s\n", chalk.Bold, device.ModuleName, chalk.Reset)
		fmt.Println("  ─────────────────────────────────")
		fmt.Printf("  🌡️ Temperature    %s %s\n", netatmo.FormatTemperature(dashboard.Temperature), netatmo.FormatTrend(dashboard.TempTrend))
		fmt.Printf("  💧 Humidity       %s\n", netatmo.FormatHumidity(dashboard.Humidity))
		fmt.Printf("  🌫️ CO2            %s\n", netatmo.FormatCO2(dashboard.CO2))
		fmt.Printf("  🔊 Noise          %s\n", netatmo.FormatNoise(dashboard.Noise))
		fmt.Println()

		// Outdoor section (if modules exist)
		if len(device.Modules) > 0 {
			for _, module := range device.Modules {
				// Check if it's an outdoor module (has temperature)
				hasTemp := false
				for _, dt := range module.DataType {
					if dt == "Temperature" {
						hasTemp = true
						break
					}
				}

				if hasTemp {
					fmt.Printf("  %s🌳 Outdoor (%s)%s\n", chalk.Bold, module.ModuleName, chalk.Reset)
					fmt.Println("  ─────────────────────────────────")
					fmt.Printf("  🌡️ Temperature    %s %s\n", netatmo.FormatTemperature(module.DashboardData.Temperature), netatmo.FormatTrend(module.DashboardData.TempTrend))
					fmt.Printf("  💧 Humidity       %s\n", netatmo.FormatHumidity(module.DashboardData.Humidity))
					fmt.Printf("  🔋 Battery        %s\n", netatmo.FormatBattery(module.BatteryPercent))
					fmt.Println()
				}
			}
		}

		// System info section
		fmt.Printf("  %s📊 System%s\n", chalk.Bold, chalk.Reset)
		fmt.Println("  ─────────────────────────────────")
		fmt.Printf("  🌀 Pressure       %.1f hPa %s\n", dashboard.Pressure, netatmo.FormatTrend(dashboard.PressureTrend))
		fmt.Printf("  📶 WiFi           %s\n", netatmo.FormatWifiSignal(device.WifiStatus))
		fmt.Printf("  ⚙️ Firmware       %d\n", device.Firmware)
		fmt.Println()

		// Last updated
		fmt.Printf("  %s⏱️  Last updated: %s%s\n", chalk.Dim, netatmo.FormatRelativeTime(dashboard.TimeUtc), chalk.Reset)
		fmt.Println()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
