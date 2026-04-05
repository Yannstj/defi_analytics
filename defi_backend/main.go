package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

// 1. Structure pour le binding JSON
type TokenInput struct {
	Symbol string `json:"symbol" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

func healthHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		err := db.Ping(ctx)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "database": "disconnected"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "OK", "database": "connected"})
	}
}

// 2. Handler pour le POST
func createTokenHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input TokenInput

		// ShouldBindJSON décode le body et vérifie les règles "binding"
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides", "details": err.Error()})
			return
		}

		// Ici, tu peux insérer dans ta base de données
		query := "INSERT INTO tokens (symbole, name) VALUES ($1, $2)"
		_, err := db.Exec(c.Request.Context(), query, input.Symbol, input.Name)
		if err != nil {
			log.Printf("erreur insertion: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Échec de l'insertion en base"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Token créée avec succès",
			"data":    input,
		})
	}
}
func TokenHandler(db *pgxpool.Pool) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. On initialise un slice vide (pour éviter de renvoyer 'null' en JSON si la table est vide)
        tokens := []TokenInput{} 

        query := "SELECT symbole, name FROM tokens"
        rows, err := db.Query(c.Request.Context(), query)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la lecture", "details": err.Error()})
            return
        }
        // Toujours fermer les rows pour libérer la connexion au pool
        defer rows.Close()

        // 2. On itère sur les résultats
        for rows.Next() {
            var t TokenInput
            // Scan doit recevoir les pointeurs des champs dans l'ordre du SELECT
            err := rows.Scan(&t.Symbol, &t.Name)
            if err != nil {
                log.Printf("Erreur scan: %v", err)
                continue // On passe au suivant si une ligne est corrompue
            }
            // 3. On ajoute le token au slice
            tokens = append(tokens, t)
        }

        // 4. On vérifie si une erreur est survenue pendant l'itération
        if err = rows.Err(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur pendant l'itération"})
            return
        }

        // 5. On renvoie le tableau final
        c.JSON(http.StatusOK, tokens)
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

	// Migrations
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	// Routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World :)"})
	})
	router.GET("/health", healthHandler(db))
	router.GET("/tokens", TokenHandler(db))
    
	// 3. Ajout de la route POST
	router.POST("/tokens", createTokenHandler(db))

	fmt.Println("Serveur démarré sur :8080")
	router.Run(":8080")
}