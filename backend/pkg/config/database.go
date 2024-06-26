package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDatabaseConnection(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
