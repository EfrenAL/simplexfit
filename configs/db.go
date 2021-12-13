package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/heroku/go-getting-started/controllers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connecting to db
func Connect() *sql.DB {

	port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }

    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
    }

	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		log.Printf("Error while creating exercise table, Reason: %v\n", err)
		os.Exit(100)
	}
	log.Printf("Connected to gormDB")

	controllers.InitiateDB(db, gormDB)
	controllers.CreateExerciseTable()

	return db
}