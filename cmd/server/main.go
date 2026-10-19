package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/permanaa/weather-api-with-redis/internal/config"
	"github.com/permanaa/weather-api-with-redis/internal/db"
	"github.com/permanaa/weather-api-with-redis/internal/handler"
	"github.com/permanaa/weather-api-with-redis/internal/helper"
	"github.com/permanaa/weather-api-with-redis/internal/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env")
	}

	redisClient, err := db.NewRedisClient()
	if err != nil {
		log.Fatal("Failed to connect redis", err)
	}

	cfg := config.Load()
	cache := helper.NewCache(redisClient)
	limiter := middleware.NewIPRateLimiter()

	handler := handler.NewHandler(cfg, cache)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Hello, Universe!",
		})
	})

	mux.HandleFunc("GET /timeline/{location}", handler.GetTimeline)

	applyMiddleware := limiter.Middleware(mux)

	server := &http.Server{
		Addr:    ":1804",
		Handler: applyMiddleware,
	}

	log.Println("Serving on port: 1804")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
