package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Status  string `json:"status"`
	Env     string `json:"environment"`
	Message string `json:"message"`
}

func main() {
	// Read docker environment variables or just use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

    // to tell where it is running
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := Response{
			Status:  "ok",
			Env:     appEnv,
			Message: "Go backend running inside Docker!",
		}
		json.NewEncoder(w).Encode(resp)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Printf("Server starting on port %s in %s mode...\n", port, appEnv)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
