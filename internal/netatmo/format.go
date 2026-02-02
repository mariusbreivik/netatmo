package netatmo

import (
	"fmt"
	"time"

	"github.com/ttacon/chalk"
)

var colorEnabled = true

func SetColorEnabled(enabled bool) {
	colorEnabled = enabled
}

func color(c chalk.Color) string {
	if colorEnabled {
		return c.String()
	}
	return ""
}

func style(s chalk.TextStyle) string {
	if colorEnabled {
		return s.String()
	}
	return ""
}

func reset() string {
	if colorEnabled {
		return chalk.Reset.String()
	}
	return ""
}

// Netatmo WiFi status: 86=good, 71-85=average, <71=bad (lower is better, it's signal attenuation in dB)
func FormatWifiSignal(status int) string {
	switch {
	case status >= 86:
		return fmt.Sprintf("%sPoor 📶%s", color(chalk.Red), reset())
	case status >= 71:
		return fmt.Sprintf("%sFair 📶%s", color(chalk.Yellow), reset())
	default:
		return fmt.Sprintf("%sGood 📶%s", color(chalk.Green), reset())
	}
}

func FormatBattery(percent int) string {
	switch {
	case percent > 50:
		return fmt.Sprintf("%s%d%% 🔋%s", color(chalk.Green), percent, reset())
	case percent > 20:
		return fmt.Sprintf("%s%d%% 🔋%s", color(chalk.Yellow), percent, reset())
	default:
		return fmt.Sprintf("%s%d%% 🪫%s", color(chalk.Red), percent, reset())
	}
}

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

func FormatTrend(trend string) string {
	switch trend {
	case "up":
		return fmt.Sprintf("%s↑%s", color(chalk.Red), reset())
	case "down":
		return fmt.Sprintf("%s↓%s", color(chalk.Blue), reset())
	case "stable":
		return fmt.Sprintf("%s→%s", color(chalk.Green), reset())
	default:
		return ""
	}
}

func FormatTemperature(temp float64) string {
	switch {
	case temp >= 25:
		return fmt.Sprintf("%s%.1f°C%s", color(chalk.Red), temp, reset())
	case temp >= 18:
		return fmt.Sprintf("%s%.1f°C%s", color(chalk.Green), temp, reset())
	case temp >= 10:
		return fmt.Sprintf("%s%.1f°C%s", color(chalk.Yellow), temp, reset())
	default:
		return fmt.Sprintf("%s%.1f°C%s", color(chalk.Blue), temp, reset())
	}
}

func FormatCO2(co2 int) string {
	switch {
	case co2 >= 1500:
		return fmt.Sprintf("%s%d ppm%s", color(chalk.Red), co2, reset())
	case co2 >= 1000:
		return fmt.Sprintf("%s%d ppm%s", color(chalk.Yellow), co2, reset())
	default:
		return fmt.Sprintf("%s%d ppm%s", color(chalk.Green), co2, reset())
	}
}

func FormatHumidity(humidity int) string {
	switch {
	case humidity >= 70:
		return fmt.Sprintf("%s%d%%%s", color(chalk.Blue), humidity, reset())
	case humidity >= 30:
		return fmt.Sprintf("%s%d%%%s", color(chalk.Green), humidity, reset())
	default:
		return fmt.Sprintf("%s%d%%%s", color(chalk.Yellow), humidity, reset())
	}
}

func FormatNoise(noise int) string {
	switch {
	case noise >= 70:
		return fmt.Sprintf("%s%d dB%s", color(chalk.Red), noise, reset())
	case noise >= 50:
		return fmt.Sprintf("%s%d dB%s", color(chalk.Yellow), noise, reset())
	default:
		return fmt.Sprintf("%s%d dB%s", color(chalk.Green), noise, reset())
	}
}
