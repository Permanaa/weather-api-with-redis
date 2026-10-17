package handler

import "github.com/permanaa/weather-api-with-redis/internal/config"

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}
