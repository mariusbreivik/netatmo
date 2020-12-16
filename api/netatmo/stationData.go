package netatmo

// StationData is a struct representation of the response fra netatmo weather api
type StationData struct {
	Body struct {
		Devices []struct {
			ID              string   `json:"_id"`
			StationName     string   `json:"station_name"`
			DateSetup       int      `json:"date_setup"`
			LastSetup       int      `json:"last_setup"`
			Type            string   `json:"type"`
			LastStatusStore int      `json:"last_status_store"`
			ModuleName      string   `json:"module_name"`
			Firmware        int      `json:"firmware"`
			LastUpgrade     int      `json:"last_upgrade"`
			WifiStatus      int      `json:"wifi_status"`
			Reachable       bool     `json:"reachable"`
			Co2Calibrating  bool     `json:"co2_calibrating"`
			DataType        []string `json:"data_type"`
			Place           struct {
				Altitude int       `json:"altitude"`
				Country  string    `json:"country"`
				Timezone string    `json:"timezone"`
				Location []float64 `json:"location"`
			} `json:"place"`
			HomeID        string `json:"home_id"`
			HomeName      string `json:"home_name"`
			DashboardData struct {
				TimeUtc          int     `json:"time_utc"`
				Temperature      float64 `json:"Temperature"`
				CO2              int     `json:"CO2"`
				Humidity         int     `json:"Humidity"`
				Noise            int     `json:"Noise"`
				Pressure         float64 `json:"Pressure"`
				AbsolutePressure float64 `json:"AbsolutePressure"`
				MinTemp          float64 `json:"min_temp"`
				MaxTemp          float64 `json:"max_temp"`
				DateMaxTemp      float64 `json:"date_max_temp"`
				DateMinTemp      float64 `json:"date_min_temp"`
				TempTrend        string  `json:"temp_trend"`
				PressureTrend    string  `json:"pressure_trend"`
			} `json:"dashboard_data"`
			Modules []struct {
				ID             string   `json:"_id"`
				Type           string   `json:"type"`
				ModuleName     string   `json:"module_name"`
				LastSetup      int      `json:"last_setup"`
				DataType       []string `json:"data_type"`
				BatteryPercent int      `json:"battery_percent"`
				Reachable      bool     `json:"reachable"`
				Firmware       int      `json:"firmware"`
				LastMessage    int      `json:"last_message"`
				LastSeen       int      `json:"last_seen"`
				RfStatus       int      `json:"rf_status"`
				BatteryVp      int      `json:"battery_vp"`
				DashboardData  struct {
					TimeUtc     int     `json:"time_utc"`
					Temperature float64 `json:"Temperature"`
					Humidity    int     `json:"Humidity"`
					MinTemp     float64 `json:"min_temp"`
					MaxTemp     float64 `json:"max_temp"`
					DateMaxTemp float64 `json:"date_max_temp"`
					DateMinTemp float64 `json:"date_min_temp"`
					TempTrend   string  `json:"temp_trend"`
				} `json:"dashboard_data"`
			} `json:"modules"`
		} `json:"devices"`
		User struct {
			Mail           string `json:"mail"`
			Administrative struct {
				Country      string `json:"country"`
				RegLocale    string `json:"reg_locale"`
				Lang         string `json:"lang"`
				Unit         int    `json:"unit"`
				Windunit     int    `json:"windunit"`
				Pressureunit int    `json:"pressureunit"`
				FeelLikeAlgo int    `json:"feel_like_algo"`
			} `json:"administrative"`
		} `json:"user"`
	} `json:"body"`
}
