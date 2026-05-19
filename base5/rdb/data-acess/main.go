package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"Arhs910/queries"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("dot env not loaded")
		return
	}
	dsn := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=postgres sslmode=disable", os.Getenv("DBUSER"), os.Getenv("DBPASS"))

	fmt.Println("DBUSER = ", os.Getenv("DBUSER"))

	db, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("DB not connected", err)
		return
	}

	fmt.Println("DB Connected")

	album, err := queries.GetAlbum(db, 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(album.Title)
}
