package netatmo

import (
	"strings"
	"testing"
	"time"
)

func TestFormatWifiSignal(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		wantText string
	}{
		{"good signal (low attenuation)", 50, "Good"},
		{"good signal (boundary)", 70, "Good"},
		{"fair signal (boundary)", 71, "Fair"},
		{"fair signal (mid)", 80, "Fair"},
		{"fair signal (upper boundary)", 85, "Fair"},
		{"poor signal (boundary)", 86, "Poor"},
		{"poor signal (high attenuation)", 100, "Poor"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatWifiSignal(tt.status)
			if !strings.Contains(result, tt.wantText) {
				t.Errorf("FormatWifiSignal(%d) = %q, want to contain %q", tt.status, result, tt.wantText)
			}
			if !strings.Contains(result, "📶") {
				t.Errorf("FormatWifiSignal(%d) = %q, want to contain emoji 📶", tt.status, result)
			}
		})
	}
}

func TestFormatBattery(t *testing.T) {
	tests := []struct {
		name        string
		percent     int
		wantPercent string
		wantEmoji   string
	}{
		{"full battery", 100, "100%", "🔋"},
		{"high battery", 75, "75%", "🔋"},
		{"medium battery (boundary)", 51, "51%", "🔋"},
		{"medium battery", 35, "35%", "🔋"},
		{"low battery (boundary)", 21, "21%", "🔋"},
		{"low battery", 15, "15%", "🪫"},
		{"critical battery", 5, "5%", "🪫"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatBattery(tt.percent)
			if !strings.Contains(result, tt.wantPercent) {
				t.Errorf("FormatBattery(%d) = %q, want to contain %q", tt.percent, result, tt.wantPercent)
			}
			if !strings.Contains(result, tt.wantEmoji) {
				t.Errorf("FormatBattery(%d) = %q, want to contain emoji %q", tt.percent, result, tt.wantEmoji)
			}
		})
	}
}

func TestFormatTemperature(t *testing.T) {
	tests := []struct {
		name string
		temp float64
		want string
	}{
		{"freezing", -5.0, "-5.0°C"},
		{"cold", 5.0, "5.0°C"},
		{"cool", 15.0, "15.0°C"},
		{"comfortable", 20.0, "20.0°C"},
		{"warm", 25.0, "25.0°C"},
		{"hot", 35.0, "35.0°C"},
		{"decimal precision", 22.5, "22.5°C"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatTemperature(tt.temp)
			if !strings.Contains(result, tt.want) {
				t.Errorf("FormatTemperature(%v) = %q, want to contain %q", tt.temp, result, tt.want)
			}
		})
	}
}

func TestFormatCO2(t *testing.T) {
	tests := []struct {
		name string
		co2  int
		want string
	}{
		{"excellent air quality", 400, "400 ppm"},
		{"good air quality", 800, "800 ppm"},
		{"moderate (boundary)", 1000, "1000 ppm"},
		{"poor air quality", 1200, "1200 ppm"},
		{"bad (boundary)", 1500, "1500 ppm"},
		{"very poor air quality", 2000, "2000 ppm"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatCO2(tt.co2)
			if !strings.Contains(result, tt.want) {
				t.Errorf("FormatCO2(%d) = %q, want to contain %q", tt.co2, result, tt.want)
			}
		})
	}
}

func TestFormatHumidity(t *testing.T) {
	tests := []struct {
		name     string
		humidity int
		want     string
	}{
		{"very dry", 20, "20%"},
		{"dry (boundary)", 30, "30%"},
		{"comfortable low", 40, "40%"},
		{"comfortable high", 60, "60%"},
		{"humid (boundary)", 70, "70%"},
		{"very humid", 85, "85%"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatHumidity(tt.humidity)
			if !strings.Contains(result, tt.want) {
				t.Errorf("FormatHumidity(%d) = %q, want to contain %q", tt.humidity, result, tt.want)
			}
		})
	}
}

func TestFormatNoise(t *testing.T) {
	tests := []struct {
		name  string
		noise int
		want  string
	}{
		{"quiet", 30, "30 dB"},
		{"normal", 45, "45 dB"},
		{"moderate (boundary)", 50, "50 dB"},
		{"loud", 60, "60 dB"},
		{"very loud (boundary)", 70, "70 dB"},
		{"extremely loud", 85, "85 dB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatNoise(tt.noise)
			if !strings.Contains(result, tt.want) {
				t.Errorf("FormatNoise(%d) = %q, want to contain %q", tt.noise, result, tt.want)
			}
		})
	}
}

func TestFormatTrend(t *testing.T) {
	tests := []struct {
		name  string
		trend string
		want  string
	}{
		{"up trend", "up", "↑"},
		{"down trend", "down", "↓"},
		{"stable trend", "stable", "→"},
		{"unknown trend", "unknown", ""},
		{"empty trend", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatTrend(tt.trend)
			if !strings.Contains(result, tt.want) {
				t.Errorf("FormatTrend(%q) = %q, want to contain %q", tt.trend, result, tt.want)
			}
		})
	}
}

func TestFormatRelativeTime(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name      string
		timestamp int64
		want      string
	}{
		{"zero timestamp", 0, "unknown"},
		{"just now", now.Unix(), "just now"},
		{"1 minute ago", now.Add(-1 * time.Minute).Unix(), "1 minute ago"},
		{"5 minutes ago", now.Add(-5 * time.Minute).Unix(), "5 minutes ago"},
		{"1 hour ago", now.Add(-1 * time.Hour).Unix(), "1 hour ago"},
		{"3 hours ago", now.Add(-3 * time.Hour).Unix(), "3 hours ago"},
		{"1 day ago", now.Add(-24 * time.Hour).Unix(), "1 day ago"},
		{"3 days ago", now.Add(-72 * time.Hour).Unix(), "3 days ago"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatRelativeTime(tt.timestamp)
			if result != tt.want {
				t.Errorf("FormatRelativeTime(%d) = %q, want %q", tt.timestamp, result, tt.want)
			}
		})
	}
}

func TestSetColorEnabled(t *testing.T) {
	// Test that colors can be disabled
	SetColorEnabled(false)
	result := FormatTemperature(25.0)
	// When colors are disabled, result should not contain ANSI escape codes
	if strings.Contains(result, "\033[") {
		t.Error("FormatTemperature should not contain ANSI codes when colors are disabled")
	}

	// Re-enable colors for other tests
	SetColorEnabled(true)
}
