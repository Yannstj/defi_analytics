package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func healthHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Création d'un contexte avec timeout pour le ping
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()

		// Test de la connexion
		err := db.Ping(ctx)

		if err != nil {
			// Si erreur, on renvoie 503 Service Unavailable
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":   "error",
				"database": "disconnected",
			})
			return
		}

		// Si tout est OK, on renvoie 200 OK
		c.JSON(http.StatusOK, gin.H{
			"status":   "OK",
			"database": "connected",
		})
	}
}

func main() {
	router := gin.Default()
	
	dbURL := os.Getenv("DATABASE_URL")
	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Impossible de créer le pool de connexion: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World :)",
		})
	})

	router.GET("/health", healthHandler(db))

	fmt.Println("Serveur démarré sur :8080")
	router.Run(":8080")
}