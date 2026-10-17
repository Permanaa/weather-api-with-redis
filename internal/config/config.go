package config

import "os"

type Config struct {
	ApiKey  string
	BaseUrl string
}

func Load() *Config {
	return &Config{
		ApiKey:  os.Getenv("VISUAL_CROSSING_WEATHER_API_KEY"),
		BaseUrl: os.Getenv("BASE_URL"),
	}
}
