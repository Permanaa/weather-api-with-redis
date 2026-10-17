package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/permanaa/weather-api-with-redis/internal/config"
	"github.com/permanaa/weather-api-with-redis/internal/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env")
	}

	cfg := config.Load()
	handler := handler.NewHandler(cfg)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Hello, Universe!",
		})
	})

	mux.HandleFunc("GET /timeline/{location}", handler.GetTimeline)

	server := &http.Server{
		Addr:    ":1804",
		Handler: mux,
	}

	log.Println("Serving on port: 1804")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
