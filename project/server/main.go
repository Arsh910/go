package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	PORT := os.Getenv("PORT")

	if PORT == "" {
		log.Fatal("PORT NOT FOUND")
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

	// mount the routes
	router.Mount("/v1", v1Router)

	fmt.Println("Server started on PORT: ", PORT)
	serv := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	err := serv.ListenAndServe()

	if err != nil {
		log.Fatal("something wrong")
	}

}
