package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello defi")})
	http.ListenAndServe(":8080", nil);

}