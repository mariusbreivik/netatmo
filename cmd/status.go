package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"

	netatmoAPI "github.com/mariusbreivik/netatmo/api/netatmo"
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

		// Indoor section (Base Station)
		fmt.Printf("  %s📍 %s (%s)%s\n", chalk.Bold, netatmoAPI.ModuleTypeDescription(device.Type), device.ModuleName, chalk.Reset)
		fmt.Println("  ─────────────────────────────────")
		fmt.Printf("  🌡️ Temperature    %s %s\n", netatmo.FormatTemperature(dashboard.Temperature), netatmo.FormatTrend(dashboard.TempTrend))
		fmt.Printf("  💧 Humidity       %s\n", netatmo.FormatHumidity(dashboard.Humidity))
		fmt.Printf("  🌫️ CO2            %s\n", netatmo.FormatCO2(dashboard.CO2))
		fmt.Printf("  🔊 Noise          %s\n", netatmo.FormatNoise(dashboard.Noise))
		fmt.Println()

		// Display all modules with appropriate sections based on type
		if len(device.Modules) > 0 {
			for _, module := range device.Modules {
				printModuleStatus(&module)
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

// getModuleIcon returns an appropriate emoji icon for the module type
func getModuleIcon(moduleType string) string {
	switch moduleType {
	case netatmoAPI.ModuleTypeOutdoor:
		return "🌳"
	case netatmoAPI.ModuleTypeWind:
		return "💨"
	case netatmoAPI.ModuleTypeRain:
		return "🌧️"
	case netatmoAPI.ModuleTypeIndoor:
		return "🏠"
	default:
		return "📡"
	}
}

// printModuleStatus prints the status section for a module based on its type
func printModuleStatus(module *netatmoAPI.Module) {
	icon := getModuleIcon(module.Type)
	typeDesc := netatmoAPI.ModuleTypeDescription(module.Type)

	fmt.Printf("  %s%s %s (%s)%s\n", chalk.Bold, icon, typeDesc, module.ModuleName, chalk.Reset)
	fmt.Println("  ─────────────────────────────────")

	switch module.Type {
	case netatmoAPI.ModuleTypeOutdoor, netatmoAPI.ModuleTypeIndoor:
		// Temperature/Humidity modules
		fmt.Printf("  🌡️ Temperature    %s %s\n", netatmo.FormatTemperature(module.DashboardData.Temperature), netatmo.FormatTrend(module.DashboardData.TempTrend))
		fmt.Printf("  💧 Humidity       %s\n", netatmo.FormatHumidity(module.DashboardData.Humidity))
	case netatmoAPI.ModuleTypeWind:
		// Wind gauge - show wind data if available
		fmt.Printf("  💨 Wind data available\n")
	case netatmoAPI.ModuleTypeRain:
		// Rain gauge - show rain data if available
		fmt.Printf("  🌧️ Rain data available\n")
	}

	fmt.Printf("  🔋 Battery        %s\n", netatmo.FormatBattery(module.BatteryPercent))
	fmt.Println()
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
