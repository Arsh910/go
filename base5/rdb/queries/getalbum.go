package queries

import (
	"database/sql"
	"fmt"
)

type Album struct {
	Id     int
	Title  string
	Artist string
	Price  float64
}

func GetAlbum(db *sql.DB, id int) (*Album, error) {
	var album Album

	query := "SELECT title FROM album WHERE id = $1;"
	err := db.QueryRow(query, id).Scan(&album.title)

	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}

	return &album, nil
}
