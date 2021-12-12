package config

import (
	"log"
	"net/url"
	"os"

	"github.com/go-pg/pg/v9"
	"github.com/heroku/go-getting-started/controllers"
)

// Connecting to db
func Connect() *pg.DB {

	parsedUrl, err := url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil { 
		panic(err)
	}

	log.Printf("ParsedUrl: %q", parsedUrl)

	pgOptions := &pg.Options{
		User: parsedUrl.User.Username(),
		Database: parsedUrl.Path[1:],
		Addr: parsedUrl.Host,
	}

	if password, ok := parsedUrl.User.Password(); ok {
		pgOptions.Password = password
	}

	db := pg.Connect(pgOptions)

	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")

	controllers.CreateExerciseTable(db)
	controllers.InitiateDB(db)


	return db
}