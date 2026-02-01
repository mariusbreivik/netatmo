package netatmo

import (
	"fmt"
	"time"

	"github.com/ttacon/chalk"
)

// FormatWifiSignal converts WiFi status to human-readable string with emoji
// Netatmo WiFi status: 86=good, 71-85=average, <71=bad (lower is better, it's signal attenuation in dB)
func FormatWifiSignal(status int) string {
	switch {
	case status >= 86:
		return fmt.Sprintf("%sPoor 📶%s", chalk.Red, chalk.Reset)
	case status >= 71:
		return fmt.Sprintf("%sFair 📶%s", chalk.Yellow, chalk.Reset)
	default:
		return fmt.Sprintf("%sGood 📶%s", chalk.Green, chalk.Reset)
	}
}

// FormatBattery converts battery percentage to colored string with emoji
func FormatBattery(percent int) string {
	switch {
	case percent > 50:
		return fmt.Sprintf("%s%d%% 🔋%s", chalk.Green, percent, chalk.Reset)
	case percent > 20:
		return fmt.Sprintf("%s%d%% 🔋%s", chalk.Yellow, percent, chalk.Reset)
	default:
		return fmt.Sprintf("%s%d%% 🪫%s", chalk.Red, percent, chalk.Reset)
	}
}

// FormatRelativeTime converts Unix timestamp to relative time string
func FormatRelativeTime(timestamp int) string {
	if timestamp == 0 {
		return "unknown"
	}

	t := time.Unix(int64(timestamp), 0)
	diff := time.Since(t)

	switch {
	case diff < time.Minute:
		return "just now"
	case diff < time.Hour:
		mins := int(diff.Minutes())
		if mins == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", mins)
	case diff < 24*time.Hour:
		hours := int(diff.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	default:
		days := int(diff.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	}
}

// FormatTrend converts trend string to arrow symbol
func FormatTrend(trend string) string {
	switch trend {
	case "up":
		return fmt.Sprintf("%s↑%s", chalk.Red, chalk.Reset)
	case "down":
		return fmt.Sprintf("%s↓%s", chalk.Blue, chalk.Reset)
	case "stable":
		return fmt.Sprintf("%s→%s", chalk.Green, chalk.Reset)
	default:
		return ""
	}
}

// FormatTemperature formats temperature with color based on value
func FormatTemperature(temp float64) string {
	switch {
	case temp >= 25:
		return fmt.Sprintf("%s%.1f°C%s", chalk.Red, temp, chalk.Reset)
	case temp >= 18:
		return fmt.Sprintf("%s%.1f°C%s", chalk.Green, temp, chalk.Reset)
	case temp >= 10:
		return fmt.Sprintf("%s%.1f°C%s", chalk.Yellow, temp, chalk.Reset)
	default:
		return fmt.Sprintf("%s%.1f°C%s", chalk.Blue, temp, chalk.Reset)
	}
}

// FormatCO2 formats CO2 level with color based on value
func FormatCO2(co2 int) string {
	switch {
	case co2 >= 1500:
		return fmt.Sprintf("%s%d ppm%s", chalk.Red, co2, chalk.Reset)
	case co2 >= 1000:
		return fmt.Sprintf("%s%d ppm%s", chalk.Yellow, co2, chalk.Reset)
	default:
		return fmt.Sprintf("%s%d ppm%s", chalk.Green, co2, chalk.Reset)
	}
}

// FormatHumidity formats humidity with color
func FormatHumidity(humidity int) string {
	switch {
	case humidity >= 70:
		return fmt.Sprintf("%s%d%%%s", chalk.Blue, humidity, chalk.Reset)
	case humidity >= 30:
		return fmt.Sprintf("%s%d%%%s", chalk.Green, humidity, chalk.Reset)
	default:
		return fmt.Sprintf("%s%d%%%s", chalk.Yellow, humidity, chalk.Reset)
	}
}

// FormatNoise formats noise level with color
func FormatNoise(noise int) string {
	switch {
	case noise >= 70:
		return fmt.Sprintf("%s%d dB%s", chalk.Red, noise, chalk.Reset)
	case noise >= 50:
		return fmt.Sprintf("%s%d dB%s", chalk.Yellow, noise, chalk.Reset)
	default:
		return fmt.Sprintf("%s%d dB%s", chalk.Green, noise, chalk.Reset)
	}
}
