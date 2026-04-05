package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}

// On définit un handler qui capture le pool de connexions
func healthHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		// Vérification réelle avec Ping
		err := db.Ping(ctx)
		
		status := "OK"
		dbStatus := "connected"
		code := http.StatusOK

		if err != nil {
			status = "error"
			dbStatus = "disconnected"
			code = http.StatusServiceUnavailable
		}

		response := HealthResponse{
			Status:   status,
			Database: dbStatus,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	// 1. Initialisation du pool de connexions
	dbURL := os.Getenv("DATABASE_URL")
	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// 2. Routage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello defi")
	})
	
	// Injection du pool dans le handler
	http.HandleFunc("/health", healthHandler(db))

	fmt.Println("Serveur démarré sur :8080")
	http.ListenAndServe(":8080", nil)
}