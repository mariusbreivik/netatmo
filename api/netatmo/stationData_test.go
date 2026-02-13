package netatmo

import (
	"encoding/json"
	"testing"
)

func TestStationDataJSONParsing(t *testing.T) {
	// Sample JSON response from Netatmo API
	jsonData := `{
		"body": {
			"devices": [
				{
					"_id": "70:ee:50:00:00:01",
					"station_name": "My Weather Station",
					"date_setup": 1234567890,
					"last_setup": 1234567890,
					"type": "NAMain",
					"last_status_store": 1234567890,
					"module_name": "Indoor",
					"firmware": 181,
					"last_upgrade": 1234567890,
					"wifi_status": 55,
					"reachable": true,
					"co2_calibrating": false,
					"data_type": ["Temperature", "CO2", "Humidity", "Noise", "Pressure"],
					"place": {
						"altitude": 100,
						"country": "NO",
						"timezone": "Europe/Oslo",
						"location": [10.0, 60.0]
					},
					"home_id": "home123",
					"home_name": "Home",
					"dashboard_data": {
						"time_utc": 1234567890,
						"Temperature": 22.5,
						"CO2": 850,
						"Humidity": 45,
						"Noise": 38,
						"Pressure": 1013.2,
						"AbsolutePressure": 1001.5,
						"min_temp": 20.0,
						"max_temp": 24.0,
						"date_max_temp": 1234567890,
						"date_min_temp": 1234567800,
						"temp_trend": "stable",
						"pressure_trend": "up"
					},
					"modules": [
						{
							"_id": "02:00:00:00:00:01",
							"type": "NAModule1",
							"module_name": "Outdoor",
							"last_setup": 1234567890,
							"data_type": ["Temperature", "Humidity"],
							"battery_percent": 87,
							"reachable": true,
							"firmware": 50,
							"last_message": 1234567890,
							"last_seen": 1234567890,
							"rf_status": 60,
							"battery_vp": 5500,
							"dashboard_data": {
								"time_utc": 1234567890,
								"Temperature": 8.3,
								"Humidity": 67,
								"min_temp": 5.0,
								"max_temp": 12.0,
								"date_max_temp": 1234567890,
								"date_min_temp": 1234567800,
								"temp_trend": "down"
							}
						}
					]
				}
			],
			"user": {
				"mail": "user@example.com",
				"administrative": {
					"country": "NO",
					"reg_locale": "en-US",
					"lang": "en",
					"unit": 0,
					"windunit": 0,
					"pressureunit": 0,
					"feel_like_algo": 0
				}
			}
		},
		"status": "ok",
		"time_exec": 0.05,
		"time_server": 1234567890
	}`

	var stationData StationData
	err := json.Unmarshal([]byte(jsonData), &stationData)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify top-level fields
	t.Run("top-level fields", func(t *testing.T) {
		if stationData.Status != "ok" {
			t.Errorf("Status = %q, want %q", stationData.Status, "ok")
		}
		if stationData.TimeExec != 0.05 {
			t.Errorf("TimeExec = %v, want %v", stationData.TimeExec, 0.05)
		}
		if stationData.TimeServer != 1234567890 {
			t.Errorf("TimeServer = %v, want %v", stationData.TimeServer, 1234567890)
		}
	})

	// Verify device data
	t.Run("device data", func(t *testing.T) {
		if len(stationData.Body.Devices) != 1 {
			t.Fatalf("Expected 1 device, got %d", len(stationData.Body.Devices))
		}

		device := stationData.Body.Devices[0]
		if device.ID != "70:ee:50:00:00:01" {
			t.Errorf("Device ID = %q, want %q", device.ID, "70:ee:50:00:00:01")
		}
		if device.StationName != "My Weather Station" {
			t.Errorf("StationName = %q, want %q", device.StationName, "My Weather Station")
		}
		if device.Type != "NAMain" {
			t.Errorf("Type = %q, want %q", device.Type, "NAMain")
		}
		if device.ModuleName != "Indoor" {
			t.Errorf("ModuleName = %q, want %q", device.ModuleName, "Indoor")
		}
		if device.Firmware != 181 {
			t.Errorf("Firmware = %d, want %d", device.Firmware, 181)
		}
		if device.WifiStatus != 55 {
			t.Errorf("WifiStatus = %d, want %d", device.WifiStatus, 55)
		}
		if !device.Reachable {
			t.Error("Reachable = false, want true")
		}
	})

	// Verify dashboard data
	t.Run("dashboard data", func(t *testing.T) {
		dashboard := stationData.Body.Devices[0].DashboardData
		if dashboard.Temperature != 22.5 {
			t.Errorf("Temperature = %v, want %v", dashboard.Temperature, 22.5)
		}
		if dashboard.CO2 != 850 {
			t.Errorf("CO2 = %d, want %d", dashboard.CO2, 850)
		}
		if dashboard.Humidity != 45 {
			t.Errorf("Humidity = %d, want %d", dashboard.Humidity, 45)
		}
		if dashboard.Noise != 38 {
			t.Errorf("Noise = %d, want %d", dashboard.Noise, 38)
		}
		if dashboard.Pressure != 1013.2 {
			t.Errorf("Pressure = %v, want %v", dashboard.Pressure, 1013.2)
		}
		if dashboard.TempTrend != "stable" {
			t.Errorf("TempTrend = %q, want %q", dashboard.TempTrend, "stable")
		}
		if dashboard.PressureTrend != "up" {
			t.Errorf("PressureTrend = %q, want %q", dashboard.PressureTrend, "up")
		}
	})

	// Verify module data
	t.Run("module data", func(t *testing.T) {
		if len(stationData.Body.Devices[0].Modules) != 1 {
			t.Fatalf("Expected 1 module, got %d", len(stationData.Body.Devices[0].Modules))
		}

		module := stationData.Body.Devices[0].Modules[0]
		if module.ID != "02:00:00:00:00:01" {
			t.Errorf("Module ID = %q, want %q", module.ID, "02:00:00:00:00:01")
		}
		if module.Type != "NAModule1" {
			t.Errorf("Module Type = %q, want %q", module.Type, "NAModule1")
		}
		if module.ModuleName != "Outdoor" {
			t.Errorf("ModuleName = %q, want %q", module.ModuleName, "Outdoor")
		}
		if module.BatteryPercent != 87 {
			t.Errorf("BatteryPercent = %d, want %d", module.BatteryPercent, 87)
		}
		if !module.Reachable {
			t.Error("Module Reachable = false, want true")
		}
	})

	// Verify module dashboard data
	t.Run("module dashboard data", func(t *testing.T) {
		moduleDashboard := stationData.Body.Devices[0].Modules[0].DashboardData
		if moduleDashboard.Temperature != 8.3 {
			t.Errorf("Module Temperature = %v, want %v", moduleDashboard.Temperature, 8.3)
		}
		if moduleDashboard.Humidity != 67 {
			t.Errorf("Module Humidity = %d, want %d", moduleDashboard.Humidity, 67)
		}
		if moduleDashboard.TempTrend != "down" {
			t.Errorf("Module TempTrend = %q, want %q", moduleDashboard.TempTrend, "down")
		}
	})

	// Verify place data
	t.Run("place data", func(t *testing.T) {
		place := stationData.Body.Devices[0].Place
		if place.Altitude != 100 {
			t.Errorf("Altitude = %d, want %d", place.Altitude, 100)
		}
		if place.Country != "NO" {
			t.Errorf("Country = %q, want %q", place.Country, "NO")
		}
		if place.Timezone != "Europe/Oslo" {
			t.Errorf("Timezone = %q, want %q", place.Timezone, "Europe/Oslo")
		}
		if len(place.Location) != 2 {
			t.Fatalf("Expected 2 location coords, got %d", len(place.Location))
		}
	})

	// Verify user data
	t.Run("user data", func(t *testing.T) {
		user := stationData.Body.User
		if user.Mail != "user@example.com" {
			t.Errorf("Mail = %q, want %q", user.Mail, "user@example.com")
		}
		if user.Administrative.Country != "NO" {
			t.Errorf("Administrative.Country = %q, want %q", user.Administrative.Country, "NO")
		}
		if user.Administrative.Lang != "en" {
			t.Errorf("Administrative.Lang = %q, want %q", user.Administrative.Lang, "en")
		}
	})
}

func TestStationDataEmptyDevices(t *testing.T) {
	jsonData := `{
		"body": {
			"devices": [],
			"user": {
				"mail": "user@example.com",
				"administrative": {}
			}
		},
		"status": "ok",
		"time_exec": 0.01,
		"time_server": 1234567890
	}`

	var stationData StationData
	err := json.Unmarshal([]byte(jsonData), &stationData)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if len(stationData.Body.Devices) != 0 {
		t.Errorf("Expected 0 devices, got %d", len(stationData.Body.Devices))
	}
}

func TestDeviceDataTypes(t *testing.T) {
	jsonData := `{
		"_id": "70:ee:50:00:00:01",
		"station_name": "Test",
		"type": "NAMain",
		"data_type": ["Temperature", "CO2", "Humidity", "Noise", "Pressure"],
		"modules": []
	}`

	var device Device
	err := json.Unmarshal([]byte(jsonData), &device)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	expectedDataTypes := []string{"Temperature", "CO2", "Humidity", "Noise", "Pressure"}
	if len(device.DataType) != len(expectedDataTypes) {
		t.Fatalf("Expected %d data types, got %d", len(expectedDataTypes), len(device.DataType))
	}

	for i, dt := range expectedDataTypes {
		if device.DataType[i] != dt {
			t.Errorf("DataType[%d] = %q, want %q", i, device.DataType[i], dt)
		}
	}
}

func TestModuleTypes(t *testing.T) {
	// Test different Netatmo module types
	moduleTypes := []struct {
		typeCode    string
		description string
	}{
		{"NAMain", "Base station (indoor)"},
		{"NAModule1", "Outdoor module"},
		{"NAModule2", "Wind gauge"},
		{"NAModule3", "Rain gauge"},
		{"NAModule4", "Additional indoor module"},
	}

	for _, mt := range moduleTypes {
		t.Run(mt.description, func(t *testing.T) {
			jsonData := `{"type": "` + mt.typeCode + `"}`
			var module Module
			err := json.Unmarshal([]byte(jsonData), &module)
			if err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}
			if module.Type != mt.typeCode {
				t.Errorf("Type = %q, want %q", module.Type, mt.typeCode)
			}
		})
	}
}
