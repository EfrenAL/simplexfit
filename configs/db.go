package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/EfrenAL/simplexfit/controllers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connecting to db
func Connect() *sql.DB {

	//Production env
	conection := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	//Local env
	//conection := "user=efrenal dbname=postgres password=secure-password host=localhost sslmode=disable"	
	//

    if port == "" {
        //log.Fatal("$PORT must be set")
		log.Printf("Running in local")
		port = "8080"
		conection = "user=efrenal dbname=postgres password=secure-password host=localhost sslmode=disable"	
    }

    db, err := sql.Open("postgres", conection)
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

	controllers.CreateTables()

	controllers.CreateExerciseTable()
	controllers.CreateWorkoutTable()

	return db
}