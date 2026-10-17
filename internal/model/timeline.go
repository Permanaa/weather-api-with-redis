package model

type Timeline struct {
	QueryCost       int      `json:"queryCost"`
	Latitude        float64  `json:"latitude"`
	Longitude       float64  `json:"longitude"`
	ResolvedAddress string   `json:"resolvedAddress"`
	Address         string   `json:"address"`
	Timezone        string   `json:"timezone"`
	Tzoffset        float64  `json:"tzoffset"`
	Days            []Days   `json:"days"`
	Stations        Stations `json:"stations"`
}

type Days struct {
	Datetime   string   `json:"datetime"`
	Tempmax    float64  `json:"tempmax"`
	Tempmin    float64  `json:"tempmin"`
	Temp       float64  `json:"temp"`
	Humidity   float64  `json:"humidity"`
	Precip     float64  `json:"precip"`
	Windspeed  float64  `json:"windspeed"`
	Pressure   float64  `json:"pressure"`
	Cloudcover float64  `json:"cloudcover"`
	Sunrise    string   `json:"sunrise"`
	Sunset     string   `json:"sunset"`
	Conditions string   `json:"conditions"`
	Icon       string   `json:"icon"`
	Stations   []string `json:"stations,omitempty"`
	Hours      []Hours  `json:"hours,omitempty"`
}

type Hours struct {
	Datetime   string  `json:"datetime"`
	Temp       float64 `json:"temp"`
	Humidity   float64 `json:"humidity"`
	Dew        float64 `json:"dew,omitempty"`
	Windspeed  float64 `json:"windspeed"`
	Pressure   float64 `json:"pressure"`
	Cloudcover float64 `json:"cloudcover,omitempty"`
	Conditions string  `json:"conditions"`
	Icon       string  `json:"icon"`
}

type Stations struct {
	Kiad Kiad `json:"KIAD"`
	Kjyo Kjyo `json:"KJYO"`
}

type Kiad struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  float64 `json:"distance"`
}

type Kjyo struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  float64 `json:"distance"`
}
