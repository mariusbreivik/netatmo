package netatmo

// StationData is a struct representation of the response from netatmo weather api
type StationData struct {
	Body       StationDataBody `json:"body"`
	Status     string          `json:"status"`
	TimeExec   float64         `json:"time_exec"`
	TimeServer int64           `json:"time_server"`
}

// StationDataBody contains the main data payload
type StationDataBody struct {
	Devices []Device `json:"devices"`
	User    User     `json:"user"`
}

// Device represents a Netatmo weather station device
type Device struct {
	ID              string          `json:"_id"`
	StationName     string          `json:"station_name"`
	DateSetup       int64           `json:"date_setup"`
	LastSetup       int64           `json:"last_setup"`
	Type            string          `json:"type"`
	LastStatusStore int64           `json:"last_status_store"`
	ModuleName      string          `json:"module_name"`
	Firmware        int             `json:"firmware"`
	LastUpgrade     int64           `json:"last_upgrade"`
	WifiStatus      int             `json:"wifi_status"`
	Reachable       bool            `json:"reachable"`
	Co2Calibrating  bool            `json:"co2_calibrating"`
	DataType        []string        `json:"data_type"`
	Place           Place           `json:"place"`
	HomeID          string          `json:"home_id"`
	HomeName        string          `json:"home_name"`
	DashboardData   DeviceDashboard `json:"dashboard_data"`
	Modules         []Module        `json:"modules"`
}

// Place represents the location information for a device
type Place struct {
	Altitude int       `json:"altitude"`
	Country  string    `json:"country"`
	Timezone string    `json:"timezone"`
	Location []float64 `json:"location"`
}

// DeviceDashboard contains the main station's sensor readings
type DeviceDashboard struct {
	TimeUtc          int64   `json:"time_utc"`
	Temperature      float64 `json:"Temperature"`
	CO2              int     `json:"CO2"`
	Humidity         int     `json:"Humidity"`
	Noise            int     `json:"Noise"`
	Pressure         float64 `json:"Pressure"`
	AbsolutePressure float64 `json:"AbsolutePressure"`
	MinTemp          float64 `json:"min_temp"`
	MaxTemp          float64 `json:"max_temp"`
	DateMaxTemp      int64   `json:"date_max_temp"`
	DateMinTemp      int64   `json:"date_min_temp"`
	TempTrend        string  `json:"temp_trend"`
	PressureTrend    string  `json:"pressure_trend"`
}

// Module represents an external module (e.g., outdoor sensor)
type Module struct {
	ID             string          `json:"_id"`
	Type           string          `json:"type"`
	ModuleName     string          `json:"module_name"`
	LastSetup      int64           `json:"last_setup"`
	DataType       []string        `json:"data_type"`
	BatteryPercent int             `json:"battery_percent"`
	Reachable      bool            `json:"reachable"`
	Firmware       int             `json:"firmware"`
	LastMessage    int64           `json:"last_message"`
	LastSeen       int64           `json:"last_seen"`
	RfStatus       int             `json:"rf_status"`
	BatteryVp      int             `json:"battery_vp"`
	DashboardData  ModuleDashboard `json:"dashboard_data"`
}

// ModuleDashboard contains a module's sensor readings
type ModuleDashboard struct {
	TimeUtc     int64   `json:"time_utc"`
	Temperature float64 `json:"Temperature"`
	Humidity    int     `json:"Humidity"`
	MinTemp     float64 `json:"min_temp"`
	MaxTemp     float64 `json:"max_temp"`
	DateMaxTemp int64   `json:"date_max_temp"`
	DateMinTemp int64   `json:"date_min_temp"`
	TempTrend   string  `json:"temp_trend"`
}

// User represents the Netatmo user information
type User struct {
	Mail           string         `json:"mail"`
	Administrative Administrative `json:"administrative"`
}

// Administrative contains user preferences and regional settings
type Administrative struct {
	Country      string `json:"country"`
	RegLocale    string `json:"reg_locale"`
	Lang         string `json:"lang"`
	Unit         int    `json:"unit"`
	Windunit     int    `json:"windunit"`
	Pressureunit int    `json:"pressureunit"`
	FeelLikeAlgo int    `json:"feel_like_algo"`
}
