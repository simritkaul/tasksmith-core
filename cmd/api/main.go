package main

import (
	"log"
	"os"

	"github.com/simritkaul/tasksmith-core/internal/db"
)

func main() {
    connStr := os.Getenv("DATABASE_URL")
    if connStr == "" {
        log.Fatal("DATABASE_URL not set")
    }

    database, err := db.New(connStr)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("DB connected successfully")
    _ = database
}
