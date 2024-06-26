package main

import (
	"database/sql"
	"fmt"
	"github/rowmur/insta-clone/internal/auth"
	"github/rowmur/insta-clone/internal/database"
	"github/rowmur/insta-clone/internal/graph"
	"github/rowmur/insta-clone/internal/loaders"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
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

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	}))
	router.Use(auth.Middleware(dbQueries))
	router.Use(loaders.Middleware(dbQueries))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	gqlSrv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DBQueries: dbQueries,
	}}))
	router.Handle("/playground", playground.Handler("GraphQL Playground", "/query"))
	router.Handle("/query", gqlSrv)

	port := 8080
	fmt.Printf("Serving on port :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
