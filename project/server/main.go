package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Arsh910/go/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")
	PORT := os.Getenv("PORT")
	DB_URL := os.Getenv("DB_URL")

	if PORT == "" {
		log.Fatal("PORT NOT FOUND")
	}
	if DB_URL == "" {
		log.Fatal("DB URL NOT FOUND")
	}

	// db
	conn, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal("Can't connect to DB", err)
	}

	apic := apiConfig{
		DB: database.New(conn),
	}

	// routes setup
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELTE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// routes
	v1Router := chi.NewRouter()
	v1Router.Get("/healthy", handleReadiness)
	v1Router.Get("/err", handleError)
	v1Router.Post("/user", apic.handleCreateUser)

	// mount the routes
	router.Mount("/v1", v1Router)

	fmt.Println("Server started on PORT: ", PORT)
	serv := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	err2 := serv.ListenAndServe()

	if err2 != nil {
		log.Fatal("something wrong")
	}

}
