package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	muxHandler := http.NewServeMux()

	muxHandler.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Hello, Universe!",
		})
	})

	server := &http.Server{
		Addr:    ":1804",
		Handler: muxHandler,
	}

	log.Println("Serving on port: 1804")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
