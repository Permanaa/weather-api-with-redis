package handler

import (
	"github.com/permanaa/weather-api-with-redis/internal/config"
	"github.com/permanaa/weather-api-with-redis/internal/helper"
)

type Handler struct {
	cfg   *config.Config
	cache *helper.Cache
}

func NewHandler(cfg *config.Config, cache *helper.Cache) *Handler {
	return &Handler{
		cfg:   cfg,
		cache: cache,
	}
}
