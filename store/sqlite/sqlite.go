package sqlite

import (
	"database/sql"
	"log"
)

type Store struct {
	ctx *sql.DB
}

func New(path string) Store {
	log.Printf("Opening database at %s", path)
	ctx, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("Could not open database: %v", err)
	}

	store := Store{ctx: ctx}

	return store
}
