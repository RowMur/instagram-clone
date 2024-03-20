package main

import (
	"context"
	"database/sql"
	"github/rowmur/insta-clone/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}

	dbURL := os.Getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening database: %s", err.Error())
	}

	dbQueries := database.New(db)

	version, err := dbQueries.GetVersion(context.Background())
	if err != nil {
		log.Fatalf("error querying database: %s", err.Error())
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(version))
	})

	http.ListenAndServe(":8080", router)
}