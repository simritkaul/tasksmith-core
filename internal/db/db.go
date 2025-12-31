package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func New (connStr string) (*sql.DB, error) {
	db, error := sql.Open("postgres", connStr);
	if error != nil {
		return nil, error;
	}

	if err := db.Ping(); err != nil {
		return nil, error;
	}

	return db, nil;
}